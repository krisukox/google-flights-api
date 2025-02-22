package flights

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/hashicorp/go-retryablehttp"
)

func (s *Session) getPriceGraphRawData(ctx context.Context, args PriceGraphArgs) (string, error) {
	return s.getRawData(ctx, args.Convert())
}

func (s *Session) getPriceGraphReqData(ctx context.Context, args PriceGraphArgs) (string, error) {
	serializedRangeStartDate := args.RangeStartDate.Format("2006-01-02")
	serializedRangeEndDate := args.RangeEndDate.Format("2006-01-02")

	rawData, err := s.getPriceGraphRawData(ctx, args)
	if err != nil {
		return "", err
	}

	prefix := `[null,"[null,`
	suffix := fmt.Sprintf(`],null,null,null,1,null,null,null,null,null,[]],[\"%s\",\"%s\"],null,[%d,%d]]"]`,
		serializedRangeStartDate, serializedRangeEndDate, args.TripLength, args.TripLength)

	reqData := prefix
	reqData += rawData
	reqData += suffix

	return url.QueryEscape(reqData), nil
}

func (s *Session) doRequestPriceGraph(ctx context.Context, args PriceGraphArgs) (*http.Response, error) {
	url := "https://www.google.com/_/FlightsFrontendUi/data/travel.frontend.flights.FlightsFrontendService/GetCalendarGraph?f.sid=-8920707734915550076&bl=boq_travel-frontend-ui_20230627.07_p1&hl=en&soc-app=162&soc-platform=1&soc-device=1&_reqid=261464&rt=c"

	reqDate, err := s.getPriceGraphReqData(ctx, args)
	if err != nil {
		return nil, err
	}

	jsonBody := []byte(
		`f.req=` + reqDate +
			`&at=AAuQa1oq5qIkgkQ2nG9vQZFTgSME%3A` + strconv.FormatInt(time.Now().Unix(), 10) + `&`)

	req, err := retryablehttp.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewReader(jsonBody))
	if err != nil {
		return nil, err
	}
	req.Header.Set("accept", `*/*`)
	req.Header.Set("accept-language", `en-US,en;q=0.9`)
	req.Header.Set("cache-control", `no-cache`)
	req.Header.Set("content-type", `application/x-www-form-urlencoded;charset=UTF-8`)
	req.Header["cookie"] = s.cookies
	req.Header.Set("pragma", `no-cache`)
	req.Header.Set("user-agent", `Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36`)
	req.Header.Set("x-goog-ext-259736195-jspb",
		fmt.Sprintf(`["en-US","US","%s",1,null,[-120],null,[[48764689,47907128,48676280,48710756,48627726,48480739,48593234,48707380]],1,[]]`, args.Currency))

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

// GetPriceGraph retrieves offers (date range) from the "Price graph" section of Google Flight search.
// The city names should be provided in the language described by args.Lang. The offers are returned
// in a slice of [Offer].
//
// GetPriceGraph returns an error if any of the requests fail or if any of the city names are misspelled.
//
// Requirements are described by the [PriceGraphArgs.Validate] function.
func (s *Session) GetPriceGraph(ctx context.Context, args PriceGraphArgs) ([]Offer, error) {
	if err := args.Validate(); err != nil {
		return nil, err
	}

	offers := []Offer{}

	resp, err := s.doRequestPriceGraph(ctx, args)
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
			sortSlice(offers, func(lv, rv Offer) bool {
				return lv.StartDate.Before(rv.StartDate)
			})
			return offers, nil
		}

		offers_, _ := getPriceGraphSection(bytesToDecode)
		offers = append(offers, offers_...)
	}
}
