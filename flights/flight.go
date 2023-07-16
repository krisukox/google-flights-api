package flights

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"golang.org/x/net/html"
	"golang.org/x/text/currency"
)

func readLine(body *bufio.Reader) ([]byte, error) {
	bytesToDecode, isPrefix, err := body.ReadLine()
	if err != nil {
		return nil, err
	}
	if !isPrefix {
		return bytesToDecode, nil
	}
	bytesToDecodeFinal := make([]byte, len(bytesToDecode))
	copy(bytesToDecodeFinal, bytesToDecode)
	for isPrefix {
		bytesToDecode, isPrefix, err = body.ReadLine()
		if err != nil {
			return bytesToDecode, err
		}
		bytesToDecodeFinal = append(bytesToDecodeFinal, bytesToDecode...)
	}
	return bytesToDecodeFinal, nil
}

func contains(s []currency.Unit, str currency.Unit) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

var supportedCurrency = []currency.Unit{currency.USD, currency.PLN}

func parse(htmlToParse string) [][]string {
	tokenizer := html.NewTokenizer(strings.NewReader(htmlToParse))

	val := []string{}
	vals := [][]string{}
	isLi := false

	for {
		tt := tokenizer.Next()
		if tt == html.ErrorToken {
			if tokenizer.Err() == io.EOF {
				return vals
			}
			fmt.Printf("Error: %v", tokenizer.Err())
			return vals
		}
		tag, hasAttr := tokenizer.TagName()
		if tt == html.StartTagToken && string(tag) == "li" {
			isLi = true
			val = []string{}
		}
		if tt == html.EndTagToken && string(tag) == "li" {
			isLi = false
			vals = append(vals, val)
		}
		if hasAttr {
			for {
				_, attrValue, moreAttr := tokenizer.TagAttr()

				if isLi {
					val = append(val, string(attrValue))
				}

				if !moreAttr {
					break
				}
			}
		}
	}
}

func isSupportedCurrency(unit currency.Unit) bool {
	fmt.Println(supportedCurrency)
	return contains(supportedCurrency, unit)
}

func getPriceSuffix(unit currency.Unit) string {
	if unit == currency.USD {
		return "US dollars"
	}
	if unit == currency.PLN {
		return "Polish zlotys"
	}
	return ""
}

func GetFlights(
	date time.Time,
	returnDate time.Time,
	srcCity string,
	dstCity string,
	unit currency.Unit,
) ([]flight, error) {
	if !isSupportedCurrency(unit) {
		return nil, fmt.Errorf(unit.String() + " is not supproted yet")
	}

	url, err := SerializeUrl(
		date,
		returnDate,
		[]string{srcCity},
		[]string{},
		[]string{dstCity},
		[]string{},
		1,
		unit,
		AnyStops,
		Economy,
		RoundTrip,
	)

	if err != nil {
		return nil, err
	}

	fmt.Println(url)

	url = url + "&curr=" + unit.String()
	url = url + "&hl=en-US"

	flights := []flight{}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("accept", "*/*")
	req.Header.Set("accept-language", "en-US,en;q=0.9")
	req.Header.Set("cache-control", "no-cache")
	req.Header.Set("content-type", "application/x-www-form-urlencoded;charset=UTF-8")
	req.Header.Set("cookie", `CONSENT=PENDING+672`)
	req.Header.Set("origin", "https://www.google.com")
	req.Header.Set("pragma", "no-cache")
	req.Header.Set("user-agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/113.0.0.0 Safari/537.36")

	client := http.Client{
		Timeout: 30 * time.Second,
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	data := parse(string(body))

	priceSuffix := getPriceSuffix(unit)

	for _, i := range data {
		departure := ""
		arrival := ""
		price := ""
		for _, j := range i {
			if strings.Contains(j, "Departure") && departure == "" {
				departure = j
			}
			if strings.Contains(j, "Arrival") && arrival == "" {
				arrival = j
			}
			if strings.HasSuffix(j, priceSuffix) && price == "" {
				price = j
			}
		}
		if departure != "" && arrival != "" && price != "" {
			flights = append(flights, flight{departure, arrival, price})
		}
	}
	return flights, nil
}