package flights

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"golang.org/x/text/currency"
	"golang.org/x/text/language"
)

func (s *Session) getPriceGraphReqData(
	rangeStartDate, rangeEndDate time.Time,
	srcCities, srcAirports, dstCities, dstAirports []string,
	adults int,
	stops Stops,
	class Class,
	tripType TripType,
	lang language.Tag,
	tripLength int,
) (string, error) {
	additionalDate := rangeStartDate.AddDate(0, 0, tripLength)
	serializedRangeStartDate := rangeStartDate.Format("2006-01-02")
	serializedRangeEndDate := rangeEndDate.Format("2006-01-02")

	rawData, err := s.getRawData(
		rangeStartDate, additionalDate,
		srcCities, srcAirports, dstCities, dstAirports,
		adults, stops, class, tripType, lang)
	if err != nil {
		return "", nil
	}

	prefix := `[null,"[null,`
	suffix := fmt.Sprintf(`],null,null,null,1,null,null,null,null,null,[]],[\"%s\",\"%s\"],null,[%d,%d]]"]`,
		serializedRangeStartDate, serializedRangeEndDate, tripLength, tripLength)

	reqData := prefix
	reqData += rawData
	reqData += suffix

	return url.QueryEscape(reqData), nil
}

func (s *Session) doRequestPriceGraph(
	rangeStartDate, rangeEndDate time.Time,
	srcCities, srcAirports, dstCities, dstAirports []string,
	adults int,
	curr currency.Unit,
	stops Stops,
	class Class,
	tripType TripType,
	lang language.Tag,
	tripLength int,
) (*http.Response, error) {
	url := "https://www.google.com/_/TravelFrontendUi/data/travel.frontend.flights.FlightsFrontendService/GetCalendarGraph?f.sid=-8920707734915550076&bl=boq_travel-frontend-ui_20230627.07_p1&hl=en&soc-app=162&soc-platform=1&soc-device=1&_reqid=261464&rt=c"

	reqDate, err := s.getPriceGraphReqData(
		rangeStartDate, rangeEndDate,
		srcCities, srcAirports, dstCities, dstAirports,
		adults, stops, class, tripType, lang, tripLength)

	jsonBody := []byte(`f.req=` + reqDate +
		`&at=AAuQa1oq5qIkgkQ2nG9vQZFTgSME%3A1688396662350&`) // Add current unix timestamp instead of 1687955915303

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(jsonBody))
	if err != nil {
		return nil, err
	}
	req.Header.Set("accept", `*/*`)
	req.Header.Set("accept-language", `en-US,en;q=0.9`)
	req.Header.Set("cache-control", `no-cache`)
	req.Header.Set("content-type", `application/x-www-form-urlencoded;charset=UTF-8`)
	req.Header.Set("cookie", `CONSENT=PENDING+672`)
	req.Header.Set("pragma", `no-cache`)
	req.Header.Set("user-agent", `Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36`)
	req.Header.Set("x-goog-ext-259736195-jspb", `["en-US","PL","PLN",1,null,[-120],null,[[48764689,47907128,48676280,48710756,48627726,48480739,48593234,48707380]],1,[]]`)

	return s.client.Do(req)
}

func getPriceGraphSection(bytesToDecode []byte) ([]Offer, error) {
	offers := []Offer{}
	var outerObject [][]interface{}
	err := json.NewDecoder(bytes.NewReader(bytesToDecode)).Decode(&outerObject)
	if err != nil {
		return nil, err
	}

	if len(outerObject[0]) < 3 {
		return offers, nil
	}

	toDecode, ok := outerObject[0][2].(string)

	if !ok {
		return nil, fmt.Errorf("unexpected object format")
	}
	var innerObject []interface{}
	err = json.NewDecoder(bytes.NewReader([]byte(toDecode))).Decode(&innerObject)
	if err != nil {
		return nil, err
	}

	if len(innerObject) < 2 {
		return offers, nil
	}

	for _, o := range innerObject[1].([]interface{}) {
		object := o.([]interface{})
		startDate := object[0].(string)
		startDateParsed, err := time.Parse("2006-01-02", startDate)
		if err != nil {
			return nil, err
		}
		returnDate := object[1].(string)
		returnDateParsed, err := time.Parse("2006-01-02", returnDate)
		if err != nil {
			return nil, err
		}
		price := object[2].([]interface{})[0].([]interface{})[1].(float64)
		offers = append(offers, Offer{StartDate: startDateParsed, ReturnDate: returnDateParsed, Price: price})
	}

	return offers, nil
}

func checkRangeDate(rangeStartDate time.Time, rangeEndDate time.Time) error {
	days := int(rangeEndDate.Sub(rangeStartDate).Hours() / 24)
	if days > 161 {
		return fmt.Errorf("number of days between dates is larger than 161, %d", days)
	}
	if rangeEndDate.Before(rangeStartDate) {
		return fmt.Errorf("rangeEndDate is before rangeStartDate")
	}
	if rangeStartDate.Before(time.Now()) {
		return fmt.Errorf("rangeStartDate is before today's date")
	}
	return nil
}

func (s *Session) GetPriceGraph(
	rangeStartDate, rangeEndDate time.Time,
	srcCities, srcAirports, dstCities, dstAirports []string,
	adults int,
	curr currency.Unit,
	stops Stops,
	class Class,
	tripType TripType,
	lang language.Tag,
	tripLength int,
) ([]Offer, error) {
	if err := checkRangeDate(rangeStartDate, rangeEndDate); err != nil {
		return nil, err
	}
	offers := []Offer{}

	resp, err := s.doRequestPriceGraph(
		rangeStartDate, rangeEndDate,
		srcCities, srcAirports, dstCities, dstAirports,
		adults, curr, stops, class, tripType, lang,
		tripLength)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	bodyBytes, err := io.ReadAll(resp.Body)

	body := bufio.NewReader(bytes.NewReader(bodyBytes))
	skipPrefix(body)
	for true {
		readLine(body)
		bytesToDecode, err := readLine(body)
		if err != nil {
			return offers, nil
		}

		offers_, err := getPriceGraphSection(bytesToDecode)
		offers = append(offers, offers_...)
	}

	return offers, nil
}
