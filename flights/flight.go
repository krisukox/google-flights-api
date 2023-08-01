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
	"golang.org/x/text/language"
)

const (
	flightAirportConst rune = '0'
	flightCityConst    rune = '5'
)

func (s *Session) getRawData(args OffersArgs) (string, error) {
	serSrcs, err := s.serializeFlightLocations(args.SrcCities, args.SrcAirports, args.Lang)
	if err != nil {
		return "", err
	}
	serDsts, err := s.serializeFlightLocations(args.DstCities, args.DstAirports, args.Lang)
	if err != nil {
		return "", err
	}

	serDate := args.Date.Format("2006-01-02")
	serReturnDate := args.ReturnDate.Format("2006-01-02")

	serAdults := serializeFlightAdults(args.Adults)
	serStops := serializeFlightStop(args.Stops)

	serClass := serializeFlightClass(args.Class)
	serTripType := serializeTripType(args.TripType)

	rawData := ""

	rawData += fmt.Sprintf(`[null,null,%d,null,[],%s,%s,null,null,null,null,null,null,[`,
		serTripType, serClass, serAdults)

	rawData += fmt.Sprintf(`[[[%s]],[[%s]],null,%s,[],[],\"%s\",null,[],[],[],null,null,[],3]`,
		serSrcs, serDsts, serStops, serDate)

	if args.TripType == RoundTrip {
		rawData += fmt.Sprintf(`,[[[%s]],[[%s]],null,%s,[],[],\"%s\",null,[],[],[],null,null,[],3]`,
			serDsts, serSrcs, serStops, serReturnDate)
	}

	return rawData, nil
}

func (s *Session) serializeFlightLocations(cities []string, airports []string, Lang language.Tag) (string, error) {
	abbrCities, err := s.abbrCities(cities, Lang)
	if err != nil {
		return "", nil
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

func serializeFlightAdults(Adults int) string {
	return fmt.Sprintf(`[%d,0,0,0]`, Adults)
}

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

func serializeFlightClass(Class Class) string {
	switch Class {
	case Economy:
		return "1"
	case PremiumEconomy:
		return "2"
	case Buisness:
		return "3"
	}
	return "4"
}

func (s *Session) getFlightReqData(args OffersArgs) (string, error) {
	rawData, err := s.getRawData(args)
	if err != nil {
		return "", nil
	}

	prefix := `[null,"[[],`
	suffix := `],null,null,null,1,null,null,null,null,null,[]],1,0,0]"]`

	reqData := prefix
	reqData += rawData
	reqData += suffix

	return url.QueryEscape(reqData), nil
}

func (s *Session) doRequestFlights(args OffersArgs) (*http.Response, error) {
	url := "https://www.google.com/_/TravelFrontendUi/data/travel.frontend.flights.FlightsFrontendService/GetShoppingResults?f.sid=-1300922759171628473&bl=boq_travel-frontend-ui_20230627.02_p1&hl=en&soc-app=162&soc-platform=1&soc-device=1&_reqid=52717&rt=c"

	reqDate, err := s.getFlightReqData(args)
	if err != nil {
		return nil, err
	}

	jsonBody := []byte(
		`f.req=` + reqDate +
			`&at=AAuQa1qjMakasqKYcQeoFJjN7RZ3%3A1687955915303&`) // Add Current unix timestamp instead of 1687955915303

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
	req.Header.Set("user-agent", `Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/113.0.0.0 Safari/537.36`)
	req.Header.Set("x-goog-ext-259736195-jspb",
		fmt.Sprintf(`["en-US","PL","%s",1,null,[-120],null,[[48676280,48710756,47907128,48764689,48627726,48480739,48593234,48707380]],1,[]]`, args.Currency)) // language, location, Currency

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

		flight.DepTime = time.Date(int(depYear), time.Month(depMonth), int(depDay), int(depHours), int(depMinutes), 0, 0, time.UTC)
		flight.ArrTime = time.Date(int(arrYear), time.Month(arrMonth), int(arrDay), int(arrHours), int(arrMinutes), 0, 0, time.UTC)
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

// GetOffers gets offers from the Google Flight search. The offers are returned in a slice of [FullOffer].
//
// GetOffers returns also [*PriceRange] which contains low and high price of the search. The values are taken from the
// "View price history" subsection of the search. If the search doesn't have the "View price history" subsection then
// GetOffers returns nil.
//
// Requirements are described by the [OffersArgs.Validate] function.
func (s *Session) GetOffers(args OffersArgs) ([]FullOffer, *PriceRange, error) {
	if err := args.Validate(); err != nil {
		return nil, nil, err
	}

	finalOffers := []FullOffer{}
	var finalPriceRange *PriceRange

	resp, err := s.doRequestFlights(args)
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
