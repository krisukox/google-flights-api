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

const (
	flightAirportConst rune = '0'
	flightCityConst    rune = '5'
)

func (s *Session) serializeFlightLocations(cities []string, airports []string, lang language.Tag) (string, error) {
	abbrCities, err := s.AbbrCities(cities, lang)
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

func serializeFlightAdults(adults int) string {
	return fmt.Sprintf(`[%d,0,0,0]`, adults)
}

func serializeFlightStop(stops Stops) string {
	switch stops {
	case Nonstop:
		return "1"
	case Stop1:
		return "2"
	case Stop2:
		return "3"
	}
	return "0"
}

func serializeFlightClass(class Class) string {
	switch class {
	case Economy:
		return "1"
	case PremiumEconomy:
		return "2"
	case Buisness:
		return "3"
	}
	return "4"
}

func (s *Session) getRawData(
	date, returnDate time.Time,
	srcCities, srcAirports, dstCities, dstAirports []string,
	adults int,
	stops Stops,
	class Class,
	tripType TripType,
	lang language.Tag,
) (string, error) {
	serSrcs, err := s.serializeFlightLocations(srcCities, srcAirports, lang)
	if err != nil {
		return "", err
	}
	serDsts, err := s.serializeFlightLocations(dstCities, dstAirports, lang)
	if err != nil {
		return "", err
	}

	serDate := date.Format("2006-01-02")
	serReturnDate := returnDate.Format("2006-01-02")

	serAdults := serializeFlightAdults(adults)
	serStops := serializeFlightStop(stops)

	serClass := serializeFlightClass(class)
	serTripType := serializeTripType(tripType)

	rawData := ""

	rawData += fmt.Sprintf(`[null,null,%d,null,[],%s,%s,null,null,null,null,null,null,[`,
		serTripType, serClass, serAdults)

	rawData += fmt.Sprintf(`[[[%s]],[[%s]],null,%s,[],[],\"%s\",null,[],[],[],null,null,[],3]`,
		serSrcs, serDsts, serStops, serDate)

	if tripType == RoundTrip {
		rawData += fmt.Sprintf(`,[[[%s]],[[%s]],null,%s,[],[],\"%s\",null,[],[],[],null,null,[],3]`,
			serDsts, serSrcs, serStops, serReturnDate)
	}

	return rawData, nil
}

func (s *Session) getFlightReqData(
	date, returnDate time.Time,
	srcCities, srcAirports, dstCities, dstAirports []string,
	adults int,
	stops Stops,
	class Class,
	tripType TripType,
	lang language.Tag,
) (string, error) {
	rawData, err := s.getRawData(
		date, returnDate,
		srcCities, srcAirports, dstCities, dstAirports,
		adults, stops, class, tripType, lang)
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

func (s *Session) doRequestFlights(
	date, returnDate time.Time,
	srcCities, srcAirports, dstCities, dstAirports []string,
	adults int,
	curr currency.Unit,
	stops Stops,
	class Class,
	tripType TripType,
	lang language.Tag,
) (*http.Response, error) {
	url := "https://www.google.com/_/TravelFrontendUi/data/travel.frontend.flights.FlightsFrontendService/GetShoppingResults?f.sid=-1300922759171628473&bl=boq_travel-frontend-ui_20230627.02_p1&hl=en&soc-app=162&soc-platform=1&soc-device=1&_reqid=52717&rt=c"

	reqDate, err := s.getFlightReqData(
		date, returnDate,
		srcCities, srcAirports, dstCities, dstAirports,
		adults, stops, class, tripType, lang)
	if err != nil {
		return nil, err
	}

	jsonBody := []byte(
		`f.req=` + reqDate +
			`&at=AAuQa1qjMakasqKYcQeoFJjN7RZ3%3A1687955915303&`) // Add current unix timestamp instead of 1687955915303

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
		fmt.Sprintf(`["en-US","PL","%s",1,null,[-120],null,[[48676280,48710756,47907128,48764689,48627726,48480739,48593234,48707380]],1,[]]`, curr)) // language, location, currency

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
	return &[]interface{}{&[]interface{}{&[]interface{}{nil, nil, rawFlights}, &[]interface{}{&[]interface{}{nil, price}}}}
}

func getSubsectionOffers(rawOffers []json.RawMessage) ([]FullOffer, error) {
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
		offer.FlightDuration = getFlightsDuration(flights)

		offers = append(offers, offer)
	}
	return offers, nil
}

func sectionOffersSchema(rawOffers1, rawOffers2 *[]json.RawMessage, priceRange *PriceRange) *[]interface{} {
	return &[]interface{}{nil, nil, rawOffers1, rawOffers2, nil, &[]interface{}{nil, nil, nil, nil,
		&[]interface{}{nil, &priceRange.Low}, &[]interface{}{nil, &priceRange.High}}}
}

func getSectionOffers(bytesToDecode []byte) ([]FullOffer, *PriceRange, error) {
	rawOffers1 := []json.RawMessage{}
	rawOffers2 := []json.RawMessage{}

	priceRange := PriceRange{}

	if err := json.Unmarshal(bytesToDecode, sectionOffersSchema(&rawOffers1, &rawOffers2, &priceRange)); err != nil {
		return nil, nil, err
	}

	allOffers := []FullOffer{}
	offers1, err := getSubsectionOffers(rawOffers1)
	if err != nil {
		return nil, nil, err
	}
	allOffers = append(allOffers, offers1...)

	offers2, err := getSubsectionOffers(rawOffers2)
	if err != nil {
		return nil, nil, err
	}
	allOffers = append(allOffers, offers2...)

	return allOffers, &priceRange, nil
}

func checkDate(date, returnDate time.Time) error {
	now := time.Now().Truncate(time.Hour * 24)

	if returnDate.Before(date) {
		return fmt.Errorf("returnDate is before date")
	}
	if date.Before(now) {
		return fmt.Errorf("date is before today's date")
	}
	return nil
}

func (s *Session) GetOffers(
	date, returnDate time.Time,
	srcCities, srcAirports, dstCities, dstAirports []string,
	adults int,
	curr currency.Unit,
	stops Stops,
	class Class,
	tripType TripType,
	lang language.Tag,
) ([]FullOffer, PriceRange, error) {
	date = date.Truncate(24 * time.Hour)
	returnDate = returnDate.Truncate(24 * time.Hour)

	if err := checkDate(date, returnDate); err != nil {
		return nil, PriceRange{}, err
	}

	if err := checkLocations(srcCities, srcAirports, dstCities, dstAirports); err != nil {
		return nil, PriceRange{}, err
	}

	finalOffers := []FullOffer{}
	finalPriceRange := PriceRange{}
	resp, err := s.doRequestFlights(
		date, returnDate,
		srcCities, srcAirports, dstCities, dstAirports,
		adults, curr, stops, class, tripType, lang)

	if err != nil {
		return nil, PriceRange{}, err
	}
	defer resp.Body.Close()

	body := bufio.NewReader(resp.Body)

	skipPrefix(body)

	for {
		readLine(body)
		bytesToDecode, err := getInnerBytes(body)
		if err != nil {
			return finalOffers, finalPriceRange, nil
		}
		offers, priceRange, _ := getSectionOffers(bytesToDecode)
		if offers != nil {
			finalOffers = append(finalOffers, offers...)
		}
		if priceRange != nil {
			finalPriceRange = *priceRange
		}
	}
}
