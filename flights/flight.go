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
	"strings"
	"sync"
	"time"

	_ "time/tzdata"

	"github.com/hashicorp/go-retryablehttp"
	"github.com/krisukox/google-flights-api/iata"
	"golang.org/x/text/language"
)

const (
	flightAirportConst = 0
	flightCityConst    = 5
)

// Google Flight API requests need different enum values than Google Flight URLs
// ________|Flights	Url
// Nonstop	1		0
// Stop1	2		1
// Stop2	3		2
// AnyStops 0		3
func serializeFlightStop(Stops Stops) int {
	switch Stops {
	case Nonstop:
		return 1
	case Stop1:
		return 2
	case Stop2:
		return 3
	}
	return 0
}

func (s *Session) serializeFlightLocations(ctx context.Context, cities []string, airports []string, Lang language.Tag) ([]interface{}, error) {
	abbrCities, err := s.abbrCities(ctx, cities, Lang)
	if err != nil {
		return nil, err
	}

	data := []interface{}{}

	for _, l := range airports {
		data = append(data, []interface{}{l, flightAirportConst})
	}
	for _, l := range abbrCities {
		data = append(data, []interface{}{l, flightCityConst})
	}

	return []interface{}{data}, nil
}

func flightsSerializeTravelers(args Args) []interface{} {
	return []interface{}{
		args.Travelers.Adults,
		args.Travelers.Children,
		args.Travelers.InfantOnLap,
		args.Travelers.InfantInSeat,
	}
}

func (s *Session) flightSerializeFlights(ctx context.Context, args Args) ([]interface{}, error) {
	srcLoc, err := s.serializeFlightLocations(ctx, args.SrcCities, args.SrcAirports, args.Lang)
	if err != nil {
		return nil, err
	}

	dstLoc, err := s.serializeFlightLocations(ctx, args.DstCities, args.DstAirports, args.Lang)
	if err != nil {
		return nil, err
	}

	data := []interface{}{
		[]interface{}{
			srcLoc,
			dstLoc,
			nil,
			serializeFlightStop(args.Stops),
			[]interface{}{},
			[]interface{}{},
			args.Date.Format(time.DateOnly),
			nil,
			[]interface{}{},
			[]interface{}{},
			[]interface{}{},
			nil,
			nil,
			[]interface{}{},
			3,
		},
	}

	if args.TripType == RoundTrip {
		data = append(data, []interface{}{
			dstLoc,
			srcLoc,
			nil,
			serializeFlightStop(args.Stops),
			[]interface{}{},
			[]interface{}{},
			args.ReturnDate.Format(time.DateOnly),
			nil,
			[]interface{}{},
			[]interface{}{},
			[]interface{}{},
			nil,
			nil,
			[]interface{}{},
			3,
		})
	}

	return data, nil
}

func serializeReqFlights(flight []Flight) []interface{} {
	data := []interface{}{}

	for _, f := range flight {
		flightNumber := strings.Split(f.FlightNumber, " ")
		data = append(data, []interface{}{
			f.DepAirportCode, f.DepTime.Format(time.DateOnly), f.ArrAirportCode, nil, flightNumber[0], flightNumber[1],
		})
	}

	return data
}

func (s *Session) serializeReturnFlight(ctx context.Context, args Args, flight []Flight) ([]interface{}, error) {
	srcLoc, err := s.serializeFlightLocations(ctx, args.SrcCities, args.SrcAirports, args.Lang)
	if err != nil {
		return nil, err
	}

	dstLoc, err := s.serializeFlightLocations(ctx, args.DstCities, args.DstAirports, args.Lang)
	if err != nil {
		return nil, err
	}

	data := []interface{}{
		[]interface{}{
			srcLoc,
			dstLoc,
			nil,
			serializeFlightStop(args.Stops),
			[]interface{}{},
			[]interface{}{},
			args.Date.Format(time.DateOnly),
			nil,
			serializeReqFlights(flight),
			[]interface{}{},
			[]interface{}{},
			nil,
			nil,
			[]interface{}{},
		},
		[]interface{}{
			dstLoc,
			srcLoc,
			nil,
			serializeFlightStop(args.Stops),
			[]interface{}{},
			[]interface{}{},
			args.ReturnDate.Format(time.DateOnly),
			nil,
			[]interface{}{},
			[]interface{}{},
			[]interface{}{},
			nil,
			nil,
			[]interface{}{},
		},
	}

	return data, nil
}

func (s *Session) getRawData1(ctx context.Context, args Args) ([]interface{}, error) {
	flights, err := s.flightSerializeFlights(ctx, args)
	if err != nil {
		return nil, err
	}

	data :=
		[]interface{}{
			nil,
			nil,
			args.TripType,
			nil,
			[]interface{}{},
			args.Class,
			flightsSerializeTravelers(args),
			nil,
			nil,
			nil,
			nil,
			nil,
			nil,
			flights,
			nil,
			nil,
			nil,
			1,
			nil,
			nil,
			nil,
			nil,
			nil,
			[]interface{}{},
		}

	return data, nil
}

func (s *Session) getRawDataReturnFlight(ctx context.Context, args Args, flight []Flight) ([]interface{}, error) {
	flights, err := s.serializeReturnFlight(ctx, args, flight)
	if err != nil {
		return nil, err
	}

	data :=
		[]interface{}{
			nil,
			nil,
			args.TripType,
			nil,
			[]interface{}{},
			args.Class,
			flightsSerializeTravelers(args),
			nil,
			nil,
			nil,
			nil,
			nil,
			nil,
			flights,
			nil,
			nil,
			nil,
			1,
			nil,
			nil,
			nil,
			nil,
			nil,
			[]interface{}{},
		}

	return data, nil
}

func serializeReqData(innerData string) (string, error) {
	data, err := json.Marshal([]interface{}{nil, innerData})
	if err != nil {
		return "", err
	}

	return url.QueryEscape(string(data)), nil
}

func (s *Session) getFlightReqData1(ctx context.Context, args Args) (string, error) {
	rawData, err := s.getRawData1(ctx, args)
	if err != nil {
		return "", err
	}

	data := []interface{}{[]interface{}{}, rawData, 1, 0, 0}

	innerData, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	return serializeReqData(string(innerData))
}

func (s *Session) getFlightReqDataReturnFlight(ctx context.Context, args Args, flight []Flight) (string, error) {
	rawData, err := s.getRawDataReturnFlight(ctx, args, flight)
	if err != nil {
		return "", err
	}

	data := []interface{}{[]interface{}{}, rawData, 1, 0, 0}

	innerData, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	return serializeReqData(string(innerData))
}

func (s *Session) doRequestFlights(ctx context.Context, args Args) (*http.Response, error) {
	url := "https://www.google.com/_/TravelFrontendUi/data/travel.frontend.flights.FlightsFrontendService/GetShoppingResults?f.sid=-1300922759171628473&bl=boq_travel-frontend-ui_20230627.02_p1&hl=en&soc-app=162&soc-platform=1&soc-device=1&_reqid=52717&rt=c"

	reqDate, err := s.getFlightReqData1(ctx, args)
	if err != nil {
		return nil, err
	}

	jsonBody := []byte(
		`f.req=` + reqDate +
			`&at=AAuQa1qjMakasqKYcQeoFJjN7RZ3%3A` + strconv.FormatInt(time.Now().Unix(), 10) + `&`)

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
	req.Header.Set("user-agent", `Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/113.0.0.0 Safari/537.36`)
	req.Header.Set("x-goog-ext-259736195-jspb",
		fmt.Sprintf(`["en-US","US","%s",1,null,[-120],null,[[48676280,48710756,47907128,48764689,48627726,48480739,48593234,48707380]],1,[]]`, args.Currency)) // language, location, Currency

	return s.client.Do(req)
}

func (s *Session) doRequestReturnFlights(ctx context.Context, args Args, flight []Flight) (*http.Response, error) {
	url := "https://www.google.com/_/TravelFrontendUi/data/travel.frontend.flights.FlightsFrontendService/GetShoppingResults?f.sid=-516424537229108034&bl=boq_travel-frontend-ui_20230816.02_p3&hl=en&soc-app=162&soc-platform=1&soc-device=1&_reqid=673689&rt=c"

	reqDate, err := s.getFlightReqDataReturnFlight(ctx, args, flight)
	if err != nil {
		return nil, err
	}

	jsonBody := []byte(
		`f.req=` + reqDate +
			`&at=AAuQa1qjMakasqKYcQeoFJjN7RZ3%3A` + strconv.FormatInt(time.Now().Unix(), 10) + `&`)

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

func (s *Session) getReturnFlight(ctx context.Context, args Args, flight []Flight) []Flight {
	resp, err := s.doRequestReturnFlights(ctx, args, flight)
	if err != nil {
		return nil
	}
	defer resp.Body.Close()

	body := bufio.NewReader(resp.Body)
	skipPrefix(body)

	readLine(body) // skip line
	bytesToDecode, err := getInnerBytes(body)
	if err != nil {
		return nil
	}

	return s.getSectionReturnFlight(ctx, bytesToDecode)
}

func getFlights(rawFlights []json.RawMessage) []Flight {
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
			return nil
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
	return flights
}

func offerSchema(rawFlights *[]json.RawMessage, price *float64) *[]interface{} {
	return &[]interface{}{&[]interface{}{nil, nil, rawFlights}, &[]interface{}{&[]interface{}{nil, price}}}
}

func (s *Session) getSubsectionOffers(ctx context.Context, args Args, rawOffers []json.RawMessage) []FullOffer {
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

		flights := getFlights(rawFlights)
		if len(flights) == 0 {
			continue
		}

		offer.Flight = flights
		offer.StartDate = flights[0].DepTime
		offer.FlightDuration = getFlightsDuration(flights)

		offer.SrcAirportCode = flights[0].DepAirportCode
		offer.DstAirportCode = flights[len(flights)-1].ArrAirportCode

		offer.SrcCity = flights[0].DepCity
		offer.DstCity = flights[len(flights)-1].ArrCity

		offers = append(offers, offer)
	}
	return offers
}

func (s *Session) getSubsectionReturnFlight(ctx context.Context, rawOffers []json.RawMessage) []Flight {
	if len(rawOffers) == 0 {
		return nil
	}
	if string(rawOffers[0]) == "null" {
		return nil
	}

	rawFlights := []json.RawMessage{}

	if err := json.Unmarshal(rawOffers[0], offerSchema(&rawFlights, nil)); err != nil {
		return nil
	}

	return getFlights(rawFlights)
}

func sectionOffersSchema(rawOffers1, rawOffers2 *[]json.RawMessage, priceRange *PriceRange) *[]interface{} {
	return &[]interface{}{nil, nil, &[]interface{}{rawOffers1}, &[]interface{}{rawOffers2}, nil, &[]interface{}{nil, nil, nil, nil,
		&[]interface{}{nil, &priceRange.Low}, &[]interface{}{nil, &priceRange.High}}}
}

func (s *Session) getSectionOffers(ctx context.Context, args Args, bytesToDecode []byte) ([]FullOffer, *PriceRange, error) {
	rawOffers1 := []json.RawMessage{}
	rawOffers2 := []json.RawMessage{}

	priceRange := PriceRange{}

	if err := json.Unmarshal(bytesToDecode, sectionOffersSchema(&rawOffers1, &rawOffers2, &priceRange)); err != nil {
		return nil, nil, err
	}

	allOffers := []FullOffer{}

	offers1 := s.getSubsectionOffers(ctx, args, rawOffers1)
	allOffers = append(allOffers, offers1...)

	offers2 := s.getSubsectionOffers(ctx, args, rawOffers2)
	allOffers = append(allOffers, offers2...)

	return allOffers, &priceRange, nil
}

func (s *Session) getSectionReturnFlight(ctx context.Context, bytesToDecode []byte) []Flight {
	rawOffers1 := []json.RawMessage{}
	rawOffers2 := []json.RawMessage{}

	if err := json.Unmarshal(bytesToDecode, sectionOffersSchema(&rawOffers1, &rawOffers2, &PriceRange{})); err != nil {
		return nil
	}

	if flights := s.getSubsectionReturnFlight(ctx, rawOffers1); flights != nil {
		return flights
	}

	return s.getSubsectionReturnFlight(ctx, rawOffers2)
}

func (s *Session) GetReturnFlights(ctx context.Context, args Args, offers []FullOffer) {
	if args.TripType == OneWay {
		return
	}

	var wg sync.WaitGroup
	wg.Add(len(offers))

	for i := range offers {
		go func(ctx context.Context, args Args, offer *FullOffer) {
			defer wg.Done()
			offer.ReturnFlight = s.getReturnFlight(ctx, args, offer.Flight)
			if len(offer.ReturnFlight) == 0 {
				return
			}

			offer.ReturnDate = offer.ReturnFlight[0].DepTime
			offer.ReturnFlightDuration = getFlightsDuration(offer.ReturnFlight)
		}(ctx, args, &offers[i])
	}

	wg.Wait()
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
			s.GetReturnFlights(ctx, args, finalOffers)

			return finalOffers, finalPriceRange, nil
		}

		offers, priceRange, _ := s.getSectionOffers(ctx, args, bytesToDecode)
		if offers != nil {
			finalOffers = append(finalOffers, offers...)
		}
		if priceRange != nil {
			finalPriceRange = priceRange
		}
	}
}
