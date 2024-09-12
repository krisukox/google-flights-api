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

	_ "time/tzdata"

	"github.com/hashicorp/go-retryablehttp"
	"github.com/krisukox/google-flights-api/iata"
	"golang.org/x/text/language"
)

const (
	flightAirportConst rune = '0'
	flightCityConst    rune = '5'
)

// Google Flight API requests need different enum values than Google Flight URLs
// ________|Flights	Url
// Nonstop	1		0
// Stop1	2		1
// Stop2	3		2
// AnyStops 0		3
func serializeFlightStop(Stops Stops) string {
	switch Stops {
	case Nonstop:
		return "1"
	case Stop1:
		return "2"
	case Stop2:
		return "3"
	}
	return "0"
}

func (s *Session) serializeFlightLocations(ctx context.Context, cities []string, airports []string, Lang language.Tag) (string, error) {
	abbrCities, err := s.abbrCities(ctx, cities, Lang)
	if err != nil {
		return "", fmt.Errorf("could not get abbrivated city names: %v", err)
	}

	serialized := ""

	for _, l := range airports {
		serialized += fmt.Sprintf(`[\"%s\",%c],`, l, flightAirportConst)
	}
	for _, l := range abbrCities {
		serialized += fmt.Sprintf(`[\"%s\",%c],`, l, flightCityConst)
	}

	return serialized[:len(serialized)-1], nil
}

func serializeFlightTravelers(args Args) string {
	return fmt.Sprintf(
		`[%d,%d,%d,%d]`,
		args.Travelers.Adults,
		args.Travelers.Children,
		args.Travelers.InfantOnLap,
		args.Travelers.InfantInSeat,
	)
}

func (s *Session) getRawData(ctx context.Context, args Args) (string, error) {
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

	rawData := ""

	rawData += fmt.Sprintf(`[null,null,%d,null,[],%d,%s,null,null,null,null,null,null,[`,
		args.TripType, args.Class, serAdults)

	rawData += fmt.Sprintf(`[[[%s]],[[%s]],null,%s,[],[],\"%s\",null,[],[],[],null,null,[],3]`,
		serSrcs, serDsts, serStops, serDate)

	if args.TripType == RoundTrip {
		rawData += fmt.Sprintf(`,[[[%s]],[[%s]],null,%s,[],[],\"%s\",null,[],[],[],null,null,[],3]`,
			serDsts, serSrcs, serStops, serReturnDate)
	}

	return rawData, nil
}

func (s *Session) getFlightReqData(ctx context.Context, args Args) (string, error) {
	rawData, err := s.getRawData(ctx, args)
	if err != nil {
		return "", err
	}

	prefix := `[null,"[[],`
	suffix := `],null,null,null,1,null,null,null,null,null,[]],1,0,0]"]`

	reqData := prefix
	reqData += rawData
	reqData += suffix

	return url.QueryEscape(reqData), nil
}

func (s *Session) doRequestFlights(ctx context.Context, args Args) (*http.Response, error) {
	baseURL := "https://www.google.com/_/TravelFrontendUi/data/travel.frontend.flights.FlightsFrontendService/GetShoppingResults"

	u, err := url.Parse(baseURL)
	if err != nil {
		return nil, fmt.Errorf("failed to parse url: %v", err)
	}
	q := u.Query()
	q.Set("f.sid", "-1300922759171628473")
	q.Set("bl", "boq_travel-frontend-ui_20230627.02_p1")
	q.Set("hl", "en")
	q.Set("soc-app", "162")
	q.Set("soc-platform", "1")
	q.Set("soc-device", "1")
	q.Set("_reqid", "52717")
	q.Set("rt", "c")
	u.RawQuery = q.Encode()

	reqDate, err := s.getFlightReqData(ctx, args)
	if err != nil {
		return nil, fmt.Errorf("could not get request data: %v", err)
	}

	jsonBody := []byte(
		`f.req=` + reqDate +
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

func getFlightsDuration(flights []Flight) time.Duration {
	return flights[len(flights)-1].ArrTime.Sub(flights[0].DepTime)
}

func flightSchema(
	flight *Flight,
	depYear, depMonth, depDay, depHours, depMinutes,
	arrYear, arrMonth, arrDay, arrHours, arrMinutes,
	duration *float64,
	flightNoPart1, flightNoPart2 *string,
) *[]interface{} {
	return &[]interface{}{
		&flight.Unknown[0],
		&flight.Unknown[1],
		&flight.Unknown[2],
		&flight.DepAirportCode,
		&flight.DepAirportName,
		&flight.ArrAirportName,
		&flight.ArrAirportCode,
		&flight.Unknown[3],
		&[]interface{}{&depHours, &depMinutes},
		&flight.Unknown[4],
		&[]interface{}{&arrHours, &arrMinutes},
		&duration,
		&flight.Unknown[5],
		&flight.Unknown[6],
		&flight.Unknown[7],
		&flight.Unknown[8],
		&flight.Unknown[9],
		&flight.Airplane,
		&flight.Unknown[10],
		&flight.Unknown[11],
		&[]interface{}{&depYear, &depMonth, &depDay},
		&[]interface{}{&arrYear, &arrMonth, &arrDay},
		&[]interface{}{&flightNoPart1, &flightNoPart2, nil, &flight.AirlineName},
		&flight.Unknown[12],
		&flight.Unknown[13],
		&flight.Unknown[14],
		&flight.Unknown[15],
		&flight.Unknown[16],
		&flight.Unknown[17],
		&flight.Unknown[18],
		&flight.Legroom,
		&flight.Unknown[19],
	}
}

func iataLocation(iataCode string) (string, *time.Location) {
	iataLocation := iata.IATATimeZone(iataCode)
	location, err := time.LoadLocation(iataLocation.Tz)
	if err != nil {
		return iataLocation.City, time.UTC
	}
	return iataLocation.City, location
}

func getFlights(rawFlights []json.RawMessage) ([]Flight, error) {
	flights := []Flight{}
	for _, rawFlight := range rawFlights {
		flight := Flight{}
		flight.Unknown = make([]interface{}, 20)
		var depHours, depMinutes, arrHours, arrMinutes, duration,
			depYear, depMonth, depDay, arrYear, arrMonth, arrDay float64
		var flightNoPart1, flightNoPart2 string
		if err := json.Unmarshal(rawFlight, flightSchema(
			&flight,
			&depYear, &depMonth, &depDay, &depHours, &depMinutes,
			&arrYear, &arrMonth, &arrDay, &arrHours, &arrMinutes,
			&duration,
			&flightNoPart1, &flightNoPart2,
		)); err != nil {
			return nil, err
		}
		depCity, depLocation := iataLocation(flight.DepAirportCode)
		arrCity, arrLocation := iataLocation(flight.ArrAirportCode)
		flight.DepCity = depCity
		flight.DepTime = time.Date(int(depYear), time.Month(depMonth), int(depDay), int(depHours), int(depMinutes), 0, 0, depLocation)
		flight.ArrCity = arrCity
		flight.ArrTime = time.Date(int(arrYear), time.Month(arrMonth), int(arrDay), int(arrHours), int(arrMinutes), 0, 0, arrLocation)
		parsedDuration, _ := time.ParseDuration(fmt.Sprintf("%dm", int(duration)))
		flight.Duration = parsedDuration
		flight.FlightNumber = flightNoPart1 + " " + flightNoPart2
		flights = append(flights, flight)
	}
	return flights, nil
}

func offerSchema(rawFlights *[]json.RawMessage, price *float64) *[]interface{} {
	return &[]interface{}{&[]interface{}{nil, nil, rawFlights}, &[]interface{}{&[]interface{}{nil, price}}}
}

func getSubsectionOffers(rawOffers []json.RawMessage, returnDate time.Time) ([]FullOffer, error) {
	offers := []FullOffer{}
	for _, rawOffer := range rawOffers {
		offer := FullOffer{}
		rawFlights := []json.RawMessage{}

		if string(rawOffer) == "null" {
			continue
		}

		if err := json.Unmarshal(rawOffer, offerSchema(&rawFlights, &offer.Price)); err != nil {
			continue
		}

		flights, err := getFlights(rawFlights)
		if err != nil || len(flights) == 0 {
			continue
		}

		offer.Flight = flights
		offer.ReturnFlight = []Flight{}

		offer.StartDate = flights[0].DepTime
		offer.ReturnDate = returnDate.UTC()

		offer.FlightDuration = getFlightsDuration(flights)

		offer.SrcAirportCode = flights[0].DepAirportCode
		offer.DstAirportCode = flights[len(flights)-1].ArrAirportCode

		offer.SrcCity = flights[0].DepCity
		offer.DstCity = flights[len(flights)-1].ArrCity

		offers = append(offers, offer)
	}
	return offers, nil
}

func sectionOffersSchema(rawOffers1, rawOffers2 *[]json.RawMessage, priceRange *PriceRange) *[]interface{} {
	return &[]interface{}{nil, nil, &[]interface{}{rawOffers1}, &[]interface{}{rawOffers2}, nil, &[]interface{}{nil, nil, nil, nil,
		&[]interface{}{nil, &priceRange.Low}, &[]interface{}{nil, &priceRange.High}}}
}

func getSectionOffers(bytesToDecode []byte, returnDate time.Time) ([]FullOffer, *PriceRange, error) {
	rawOffers1 := []json.RawMessage{}
	rawOffers2 := []json.RawMessage{}

	priceRange := PriceRange{}

	if err := json.Unmarshal(bytesToDecode, sectionOffersSchema(&rawOffers1, &rawOffers2, &priceRange)); err != nil {
		return nil, nil, err
	}

	allOffers := []FullOffer{}
	offers1, err := getSubsectionOffers(rawOffers1, returnDate)
	if err != nil {
		return nil, nil, err
	}
	allOffers = append(allOffers, offers1...)
	offers2, err := getSubsectionOffers(rawOffers2, returnDate)
	if err != nil {
		return nil, nil, err
	}
	allOffers = append(allOffers, offers2...)

	return allOffers, &priceRange, nil
}

// GetOffers retrieves offers from the Google Flight search. The city names should be provided in the language
// described by args.Lang. The offers are returned in a slice of [FullOffer].
//
// GetOffers also returns [*PriceRange], which contains the low and high prices of the search. The values are
// taken from the "View price history" subsection of the search. If the search doesn't have the "View
// price history" subsection, then GetOffers returns nil.
//
// GetPriceGraph returns an error if any of the requests fail or if any of the city names are misspelled.
//
// Requirements are described by the [Args.ValidateOffersArgs] function.
func (s *Session) GetOffers(ctx context.Context, args Args) ([]FullOffer, *PriceRange, error) {
	if err := args.ValidateOffersArgs(); err != nil {
		return nil, nil, err
	}

	finalOffers := []FullOffer{}
	var finalPriceRange *PriceRange

	resp, err := s.doRequestFlights(ctx, args)
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
