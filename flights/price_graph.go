package flights

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/hashicorp/go-retryablehttp"
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
	if err != nil {
		return nil, err
	}

	jsonBody := []byte(
		`f.req=` + reqDate +
			`&at=AAuQa1oq5qIkgkQ2nG9vQZFTgSME%3A1688396662350&`) // Add current unix timestamp instead of 1687955915303

	req, err := retryablehttp.NewRequest(http.MethodPost, url, bytes.NewReader(jsonBody))
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

func priceGraphSchema(startDate, returnDate *string, price *float64) *[]interface{} {
	// [startDate,returnDate,[[null,price],""],1]
	return &[]interface{}{startDate, returnDate, &[]interface{}{&[]interface{}{nil, price}}}
}

func getPriceGraphSection(bytesToDecode []byte) ([]Offer, error) {
	offers := []Offer{}

	var err error

	rawOffers := []json.RawMessage{}

	if err = json.Unmarshal([]byte(bytesToDecode), &[]interface{}{nil, &rawOffers}); err != nil {
		return nil, err
	}

	for _, o := range rawOffers {
		finalOffer := Offer{}

		startDate := ""
		returnDate := ""

		if err = json.Unmarshal(o, priceGraphSchema(&startDate, &returnDate, &finalOffer.Price)); err != nil {
			continue
		}

		if finalOffer.StartDate, err = time.Parse("2006-01-02", startDate); err != nil {
			continue
		}
		if finalOffer.ReturnDate, err = time.Parse("2006-01-02", returnDate); err != nil {
			continue
		}

		offers = append(offers, finalOffer)
	}

	return offers, nil
}

func checkRangeDate(rangeStartDate time.Time, rangeEndDate time.Time) error {
	now := time.Now().Truncate(time.Hour * 24)

	days := int(rangeEndDate.Sub(rangeStartDate).Hours() / 24)
	if days > 161 {
		return fmt.Errorf("number of days between dates is larger than 161, %d", days)
	}
	if rangeEndDate.Equal(rangeStartDate) {
		return fmt.Errorf("rangeEndDate is the same as rangeStartDate")
	}
	if rangeEndDate.Before(rangeStartDate) {
		return fmt.Errorf("rangeEndDate is before rangeStartDate")
	}
	if rangeStartDate.Before(now) {
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
	rangeStartDate = rangeStartDate.Truncate(24 * time.Hour)
	rangeEndDate = rangeEndDate.Truncate(24 * time.Hour)

	if err := checkRangeDate(rangeStartDate, rangeEndDate); err != nil {
		return nil, err
	}

	if err := checkLocations(srcCities, srcAirports, dstCities, dstAirports); err != nil {
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

	body := bufio.NewReader(resp.Body)
	skipPrefix(body)
	for {
		readLine(body) // skip line
		bytesToDecode, err := getInnerBytes(body)
		if err != nil {
			return offers, nil
		}

		offers_, _ := getPriceGraphSection(bytesToDecode)
		offers = append(offers, offers_...)
	}
}
