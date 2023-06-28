package api

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

type flight struct {
	Departure string
	Arrival   string
	Price     string
}

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
		bytesToDecodeFinal = append(bytesToDecodeFinal, bytesToDecode...)
	}
	// fmt.Println(string(bytesToDecodeFinal))
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

	// var vals []string

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
			// if hasAttr && tt == html.StartTagToken && string(tag) == "li" {
			// fmt.Printf("Tag: %v\n", string(tag))
			for {
				_, attrValue, moreAttr := tokenizer.TagAttr()
				// if string(attrKey) == "" {
				//     break
				// }
				// fmt.Printf("Attr: %v\n", string(attrKey))
				// fmt.Printf("Attr: %v\n", string(attrValue))
				// fmt.Printf("Attr: %v\n", moreAttr)

				if isLi {
					val = append(val, string(attrValue))
				}

				if !moreAttr {
					break
				}
			}
		}
		// }
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

// func GetBestFlightsDateRange

func GetFlights(
	departureDate time.Time,
	returnDate time.Time,
	departureCity string,
	arivalCity string,
	unit currency.Unit,
) ([]flight, error) {
	if !isSupportedCurrency(unit) {
		return nil, fmt.Errorf(unit.String() + " is not supproted yet")
	}

	url, err := CreateFullURL(departureDate, returnDate, departureCity, arivalCity)
	if err != nil {
		return nil, err
	}

	fmt.Println(url)

	url = url + "&curr=" + unit.String()
	url = url + "&hl=en-US"

	flights := []flight{}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	req.Header.Set("authority", "www.google.com")
	req.Header.Set("accept", "*/*")
	req.Header.Set("accept-language", "en-US,en;q=0.9")
	req.Header.Set("cache-control", "no-cache")
	req.Header.Set("content-type", "application/x-www-form-urlencoded;charset=UTF-8")
	req.Header.Set("cookie", "CONSENT=YES+srp.gws-20211208-0-RC2.pl+FX+371")
	req.Header.Set("origin", "https://www.google.com")
	req.Header.Set("pragma", "no-cache")
	req.Header.Set("sec-ch-ua", "\"Google Chrome\";v=\"113\", \"Chromium\";v=\"113\", \"Not-A.Brand\";v=\"24\"")
	req.Header.Set("sec-ch-ua-arch", "\"x86\"")
	req.Header.Set("sec-ch-ua-bitness", "\"64\"")
	req.Header.Set("sec-ch-ua-full-version", "\"113.0.5672.92\"")
	req.Header.Set("sec-ch-ua-full-version-list", "\"Google Chrome\";v=\"113.0.5672.92\", \"Chromium\";v=\"113.0.5672.92\", \"Not-A.Brand\";v=\"24.0.0.0\"")
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-model", "")
	req.Header.Set("sec-ch-ua-platform", "Linux")
	req.Header.Set("sec-ch-ua-platform-version", "5.19.0")
	req.Header.Set("sec-ch-ua-wow64", "?0")
	req.Header.Set("sec-fetch-dest", "empty")
	req.Header.Set("sec-fetch-mode", "cors")
	req.Header.Set("sec-fetch-site", "same-origin")
	req.Header.Set("user-agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/113.0.0.0 Safari/537.36")
	// req.Header.Set("x-goog-ext-190139975-jspb", "[\"PL\",\"ZZ\",\"Xkt6eg==\"]")
	// req.Header.Set("x-goog-ext-259736195-jspb", "[\"pl\",\"PL\",\"PLN\",1,null,[-120],null,[[48764690,48627726,47907128,48480739,48710756,48676280,48593234,48707378]],1,[]]")
	req.Header.Set("x-same-domain", "1")
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

	// fmt.Println(string(body))

	data := parse(string(body))

	// fmt.Println(data)

	priceSuffix := getPriceSuffix(unit)

	for _, i := range data {
		departure := ""
		arrival := ""
		price := ""
		// fmt.Println(i)
		for _, j := range i {
			// fmt.Println(j)
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
