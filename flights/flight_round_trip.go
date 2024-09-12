package flights

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/hashicorp/go-retryablehttp"
)

func (s *Session) getReturnFlightRawData(ctx context.Context, args Args, fullOffer FullOffer) (string, error) {
	serSrcs, err := s.serializeFlightLocations(ctx, args.SrcCities, args.SrcAirports, args.Lang)
	if err != nil {
		return "", fmt.Errorf("could not serialize src flight src locations: %v", err)
	}
	serDsts, err := s.serializeFlightLocations(ctx, args.DstCities, args.DstAirports, args.Lang)
	if err != nil {
		return "", fmt.Errorf("could not serialize src flight dst locations: %v", err)
	}

	serDate := args.Date.Format("2006-01-02")
	serReturnDate := args.ReturnDate.Format("2006-01-02")

	serAdults := serializeFlightTravelers(args)
	serStops := serializeFlightStop(args.Stops)

	_, _, _, _, _, _ = serAdults, serDsts, serDate, serReturnDate, serSrcs, serStops

	rawData := ""

	rawData += fmt.Sprintf(`[null,null,%d,null,[],%d,%s,null,null,null,null,null,null,[`,
		args.TripType, args.Class, serAdults)

	rawData += fmt.Sprintf(`[[[%s]],[[%s]],null,%s,null,null,\"%s\",null,`,
		serSrcs, serDsts, serStops, serDate)

	rawData += "["
	for _, offer := range fullOffer.Flight {
		flightNumbers := strings.Split(offer.FlightNumber, " ")
		rawData += fmt.Sprintf(`[\"%s\",\"%s\",\"%s\",null,\"%s\",\"%s\"],`,
			offer.DepAirportCode, offer.DepTime.Format("2006-01-02"), offer.ArrAirportCode, flightNumbers[0], flightNumbers[1],
		)
	}
	rawData = rawData[:len(rawData)-1]
	rawData += "]"

	rawData += fmt.Sprintf(`],[[[%s]],[[%s]],null,%s,null,null,\"%s\"]]`,
		serDsts, serSrcs, serStops, serReturnDate)

	return rawData, nil
}

func (s *Session) getReturnFlightReqData(ctx context.Context, args Args, fullOffer FullOffer) (string, error) {
	rawData, err := s.getReturnFlightRawData(ctx, args, fullOffer)
	if err != nil {
		return "", err
	}

	prefix := `[null,"[[],`
	suffix := `,null,null,null,1],0,0,0,1]"]`

	reqData := prefix
	reqData += rawData
	reqData += suffix

	return url.QueryEscape(reqData), nil
}

func (s *Session) doRequestReturnFlights(ctx context.Context, args Args, fullOffer FullOffer) (*http.Response, error) {
	baseURL := "https://www.google.com/_/FlightsFrontendUi/data/travel.frontend.flights.FlightsFrontendService/GetShoppingResults"

	u, err := url.Parse(baseURL)
	if err != nil {
		return nil, fmt.Errorf("failed to parse url: %v", err)
	}
	q := u.Query()
	q.Set("f.sid", "-1300922759171628473")
	q.Set("bl", "boq_travel-frontend-ui_20230627.02_p1")
	q.Set("hl", args.Lang.String())
	q.Set("soc-app", "162")
	q.Set("soc-platform", "1")
	q.Set("soc-device", "1")
	q.Set("_reqid", "52717")
	q.Set("rt", "c")
	u.RawQuery = q.Encode()

	reqData, err := s.getReturnFlightReqData(ctx, args, fullOffer)
	if err != nil {
		return nil, fmt.Errorf("could not get request data: %v", err)
	}

	jsonBody := []byte(
		`f.req=` + reqData +
			`&at=AAuQa1qjMakasqKYcQeoFJjN7RZ3%3A` + strconv.FormatInt(time.Now().Unix(), 10) + `&`)

	req, err := retryablehttp.NewRequestWithContext(ctx, http.MethodPost, u.String(), bytes.NewReader(jsonBody))
	if err != nil {
		return nil, fmt.Errorf("failed to do GetShoppingResults request: %v", err)
	}
	req.Header.Set("accept", `*/*`)
	req.Header.Set("accept-language", `en-US,en;q=0.9`)
	req.Header.Set("cache-control", `no-cache`)
	req.Header.Set("content-type", `application/x-www-form-urlencoded;charset=UTF-8`)
	req.Header["cookie"] = s.cookies
	req.Header.Set("pragma", `no-cache`)
	req.Header.Set("user-agent", `Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/113.0.0.0 Safari/537.36`)
	req.Header.Set("x-goog-ext-259736195-jspb",
		fmt.Sprintf(`["en-US","US","%s",1,null,[-120],null,[[48676280,48710756,47907128,48764689,48627726,48480739,48593234,48707380]],1,[]]`, args.Currency)) // language, location, Currency

	return s.client.Do(req)
}

func (s *Session) GetReturnOffers(ctx context.Context, args Args, fullOffer FullOffer) ([]FullOffer, *PriceRange, error) {
	if err := args.ValidateOffersArgs(); err != nil {
		return nil, nil, err
	}

	finalOffers := []FullOffer{}
	var finalPriceRange *PriceRange

	resp, err := s.doRequestReturnFlights(ctx, args, fullOffer)
	if err != nil {
		return nil, nil, err
	}
	defer resp.Body.Close()

	body := bufio.NewReader(resp.Body)
	skipPrefix(body)

	for {
		readLine(body) // skip line
		bytesToDecode, err := getInnerBytes(body)
		if err != nil {
			return finalOffers, finalPriceRange, nil
		}

		offers, priceRange, _ := getSectionOffers(bytesToDecode, args.ReturnDate)
		if offers != nil {
			finalOffers = append(finalOffers, offers...)
		}
		if priceRange != nil {
			finalPriceRange = priceRange
		}
	}
}

func (s *Session) offersToRoundTripOffers(ctx context.Context, args Args, offers []FullOffer) ([]FullOffer, error) {
	roundTripOffers := []FullOffer{}
	for _, offer := range offers {
		returnOffers, _, err := s.GetReturnOffers(ctx, args, offer)
		if err != nil {
			return nil, err
		}
		for _, returnOffer := range returnOffers {
			roundTripOffers = append(roundTripOffers, FullOffer{
				Offer: Offer{
					StartDate:  offer.Offer.StartDate,
					ReturnDate: returnOffer.Offer.StartDate,
					Price:      returnOffer.Offer.Price,
				},
				Flight:               offer.Flight,
				ReturnFlight:         returnOffer.Flight,
				SrcAirportCode:       offer.SrcAirportCode,
				DstAirportCode:       offer.DstAirportCode,
				SrcCity:              offer.SrcCity,
				DstCity:              offer.DstCity,
				FlightDuration:       offer.FlightDuration,
				ReturnFlightDuration: returnOffer.FlightDuration,
			})
		}
	}
	return roundTripOffers, nil
}
