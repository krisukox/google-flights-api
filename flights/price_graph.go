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
)

func (s *Session) getPriceGraphRawData(args PriceGraphArgs) (string, error) {
	return s.getRawData(
		OffersArgs{
			Date:        args.RangeStartDate,
			ReturnDate:  args.RangeStartDate.AddDate(0, 0, args.TripLength),
			SrcCities:   args.SrcCities,
			SrcAirports: args.SrcAirports,
			DstCities:   args.DstCities,
			DstAirports: args.DstAirports,
			Args:        args.Args,
		},
	)
}

func (s *Session) getPriceGraphReqData(args PriceGraphArgs) (string, error) {
	serializedRangeStartDate := args.RangeStartDate.Format("2006-01-02")
	serializedRangeEndDate := args.RangeEndDate.Format("2006-01-02")

	rawData, err := s.getPriceGraphRawData(args)
	if err != nil {
		return "", nil
	}

	prefix := `[null,"[null,`
	suffix := fmt.Sprintf(`],null,null,null,1,null,null,null,null,null,[]],[\"%s\",\"%s\"],null,[%d,%d]]"]`,
		serializedRangeStartDate, serializedRangeEndDate, args.TripLength, args.TripLength)

	reqData := prefix
	reqData += rawData
	reqData += suffix

	return url.QueryEscape(reqData), nil
}

func (s *Session) doRequestPriceGraph(args PriceGraphArgs) (*http.Response, error) {
	url := "https://www.google.com/_/TravelFrontendUi/data/travel.frontend.flights.FlightsFrontendService/GetCalendarGraph?f.sid=-8920707734915550076&bl=boq_travel-frontend-ui_20230627.07_p1&hl=en&soc-app=162&soc-platform=1&soc-device=1&_reqid=261464&rt=c"

	reqDate, err := s.getPriceGraphReqData(args)
	if err != nil {
		return nil, err
	}

	jsonBody := []byte(
		`f.req=` + reqDate +
			`&at=AAuQa1oq5qIkgkQ2nG9vQZFTgSME%3A1688396662350&`) // Add Current unix timestamp instead of 1687955915303

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
	req.Header.Set("x-goog-ext-259736195-jspb",
		fmt.Sprintf(`["en-US","PL","%s",1,null,[-120],null,[[48764689,47907128,48676280,48710756,48627726,48480739,48593234,48707380]],1,[]]`, args.Currency))

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

// GetPriceGraph gets offers (date range) from the "Price graph" section of Google Flight search.
// The offers are returned in a slice of [Offer].
//
// Requirements are described by the [PriceGraphArgs.Validate] function.
func (s *Session) GetPriceGraph(args PriceGraphArgs) ([]Offer, error) {
	if err := args.Validate(); err != nil {
		return nil, err
	}

	offers := []Offer{}

	resp, err := s.doRequestPriceGraph(args)
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
