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

func getPrice(tripObj []interface{}) (float64, error) {
	priceObj1, ok := getElement[[]interface{}](tripObj, 1)
	if !ok {
		return 0, fmt.Errorf("wrong price format stage 1: %v", priceObj1)
	}
	priceObj2, ok := getElement[[]interface{}](priceObj1, 0)
	if !ok {
		return 0, fmt.Errorf("wrong price format stage 2: %v", priceObj2)
	}
	price, ok := getElement[float64](priceObj2, 1)
	if !ok {
		return 0, fmt.Errorf("wrong price format stage 3: %v", priceObj2)
	}
	return price, nil
}

func getTime(flightTime interface{}, flightDate interface{}) (time.Time, error) {
	flightTime1, ok := flightTime.([]interface{})
	if !ok {
		return time.Time{}, fmt.Errorf("wrong time format: %v", flightTime)
	}

	var hours float64
	var minutes float64
	if len(flightTime1) == 2 { // time format: [6, 45] or [6]
		minutes, ok = flightTime1[1].(float64)
		if !ok {
			return time.Time{}, fmt.Errorf("wrong minutes format: %v", flightTime1[1])
		}
	} else if len(flightTime1) != 1 {
		return time.Time{}, fmt.Errorf("wrong time format: %v", flightTime)
	}

	if flightTime1[0] == nil {
		hours = 0
	} else if hours, ok = flightTime1[0].(float64); !ok {
		return time.Time{}, fmt.Errorf("wrong hours format: %v", flightTime1[0])
	}
	flightDate1, ok := flightDate.([]interface{})
	if !ok || len(flightDate1) != 3 {
		return time.Time{}, fmt.Errorf("wrong date format: %v", flightDate)
	}
	year, ok := flightDate1[0].(float64)
	if !ok {
		return time.Time{}, fmt.Errorf("wrong year format: %v", flightDate1[0])
	}
	month, ok := flightDate1[1].(float64)
	if !ok {
		return time.Time{}, fmt.Errorf("wrong month format: %v", flightDate1[1])
	}
	day, ok := flightDate1[2].(float64)
	if !ok {
		return time.Time{}, fmt.Errorf("wrong day format: %v", flightDate1[2])
	}
	location, _ := time.LoadLocation("Poland") // FIXME

	return time.Date(
		int(year),
		time.Month(month),
		int(day),
		int(hours),
		int(minutes),
		0,
		0,
		location,
	), nil
}

func getDuration(duration []interface{}) (time.Duration, error) {
	duration1, _ := getElement[float64](duration, 11)
	return time.ParseDuration(fmt.Sprintf("%dm", int(duration1)))
}

func getFlightNumberAirline(data interface{}) (string, interface{}, string, error) {
	data1, ok := data.([]interface{})
	if !ok || len(data1) != 4 {
		return "", "", "", fmt.Errorf("wrong flight number of airline type: %v", data)
	}
	flightNumberPart1, _ := getElement[string](data1, 0)
	// if !ok {
	//  return "", "", "", fmt.Errorf("wrong flight number part 1 type: %v", data1[0])
	// }
	flightNumberPart2, _ := getElement[string](data1, 1)
	// if !ok {
	//  return "", "", "", fmt.Errorf("wrong flight number part 2 type: %v", data1[1])
	// }
	airline, _ := getElement[string](data1, 3)
	// if !ok {
	//  return "", "", "", fmt.Errorf("wrong airline name type: %v", data1[3])
	// }
	return flightNumberPart1 + " " + flightNumberPart2, data1[2], airline, nil
}

func getUnknowns(flightObj1 []interface{}) []interface{} {
	unknowns := []interface{}{}
	unknowns = append(unknowns, getRawElement(flightObj1, 1))
	unknowns = append(unknowns, getRawElement(flightObj1, 2))
	unknowns = append(unknowns, getRawElement(flightObj1, 7))
	unknowns = append(unknowns, getRawElement(flightObj1, 9))
	unknowns = append(unknowns, getRawElement(flightObj1, 12))
	unknowns = append(unknowns, getRawElement(flightObj1, 13))
	unknowns = append(unknowns, getRawElement(flightObj1, 14))
	unknowns = append(unknowns, getRawElement(flightObj1, 15))
	unknowns = append(unknowns, getRawElement(flightObj1, 16))
	unknowns = append(unknowns, getRawElement(flightObj1, 18))
	unknowns = append(unknowns, getRawElement(flightObj1, 19))
	unknowns = append(unknowns, getRawElement(flightObj1, 23))
	unknowns = append(unknowns, getRawElement(flightObj1, 24))
	unknowns = append(unknowns, getRawElement(flightObj1, 25))
	unknowns = append(unknowns, getRawElement(flightObj1, 26))
	unknowns = append(unknowns, getRawElement(flightObj1, 27))
	unknowns = append(unknowns, getRawElement(flightObj1, 28))
	unknowns = append(unknowns, getRawElement(flightObj1, 29))
	unknowns = append(unknowns, getRawElement(flightObj1, 31))
	return unknowns
}

func getFlight(flightObj interface{}) (flight, error) {
	var unknowns []interface{}
	flightObj1, ok := flightObj.([]interface{})
	if !ok {
		return flight{}, fmt.Errorf("wrong flight format: %v", flightObj)
	}

	departureAirportCode, _ := getElement[string](flightObj1, 3)
	// if !ok {
	//  return flightV2{}, fmt.Errorf("wrong departure airport code type: %v", object1[3])
	// }
	departureAirportName, _ := getElement[string](flightObj1, 4)
	// if !ok {
	//  return flightV2{}, fmt.Errorf("wrong departure airport name type: %v", object1[4])
	// }
	arrivalAirportName, _ := getElement[string](flightObj1, 5)
	// if !ok {
	//  return flightV2{}, fmt.Errorf("wrong arrival airport name type: %v", object1[5])
	// }
	arrivalAirportCode, _ := getElement[string](flightObj1, 6)
	// if !ok {
	//  return flightV2{}, fmt.Errorf("wrong arrival airport code type: %v", object1[6])
	// }
	departureTime, _ := getTime(flightObj1[8], flightObj1[20])
	// if err != nil {
	//  return flightV2{}, fmt.Errorf("departure: %w", err)
	// }
	arrivalTime, _ := getTime(flightObj1[10], flightObj1[21])
	// if err != nil {
	//  return flightV2{}, fmt.Errorf("arrival: %w", err)
	// }
	duration, _ := getDuration(flightObj1)
	// if err != nil {
	//  return flightV2{}, err
	// }
	flightNumber, u, airlineName, _ := getFlightNumberAirline(flightObj1[22])
	unknowns = append(unknowns, u)
	// if err != nil {
	//  return flightV2{}, err
	// }
	airplane, _ := getElement[string](flightObj1, 17)
	// if !ok {
	//  return flightV2{}, fmt.Errorf("wrong airplane format: %v", object1[17])
	// }
	legroom, _ := getElement[string](flightObj1, 30)
	us := getUnknowns(flightObj1)
	unknowns = append(unknowns, us...)
	return flight{
		departureAirportCode,
		departureAirportName,
		arrivalAirportName,
		arrivalAirportCode,
		departureTime,
		arrivalTime,
		duration,
		airplane,
		flightNumber,
		unknowns,
		airlineName,
		legroom,
	}, nil
}

func getFlights(tripObj []interface{}) ([]flight, error) {
	flights := []flight{}

	flightObj1, ok := tripObj[0].([]interface{})
	if !ok {
		return nil, fmt.Errorf("wrong flights format stage 1: %v", tripObj[0])
	}
	flightObjs, ok := flightObj1[2].([]interface{})
	if !ok {
		return nil, fmt.Errorf("wrong flights format stage 2: %v", flightObj1[2])
	}
	for _, f := range flightObjs {
		finalFlight, err := getFlight(f)
		if err != nil {
			return nil, fmt.Errorf("cannot get flight: %w", err)
		}
		flights = append(flights, finalFlight)
	}

	return flights, nil
}

func getTripDuration(flights []flight) time.Duration {
	return flights[len(flights)-1].DepTime.Sub(flights[0].DepTime)
}

func getTrip(tripObj []interface{}) (FullOffer, error) {
	price, err := getPrice(tripObj)
	if err != nil {
		return FullOffer{}, err
	}
	flights, err := getFlights(tripObj)
	if err != nil {
		return FullOffer{}, err
	}
	return FullOffer{Offer{flights[0].DepTime, time.Time{}, price}, flights, []flight{}, "", "", "", "", getTripDuration(flights)}, nil
}

func getOffersFromSection(section []interface{}) ([]FullOffer, error) {
	trips := []FullOffer{}

	object, ok := section[0].([]interface{})
	if !ok {
		return nil, fmt.Errorf("unexpected object format")
	}
	for _, o := range object {
		object1, ok := o.([]interface{})
		if !ok {
			break
		}

		// fmt.Println(object1)

		trip, err := getTrip(object1)
		if err != nil {
			return trips, err
		}

		trips = append(trips, trip)
	}
	return trips, nil
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

	jsonBody := []byte(`f.req=` + reqDate +
		`&at=AAuQa1qjMakasqKYcQeoFJjN7RZ3%3A1687955915303&`) // Add current unix timestamp instead of 1687955915303
	bodyReader := bytes.NewReader(jsonBody)

	req, err := http.NewRequest(http.MethodPost, url, bodyReader)
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

func getPrice1(object interface{}) (float64, error) {
	object1, ok := object.([]interface{})
	if !ok {
		return 0, fmt.Errorf("wrong price range start format %v", object1)
	}
	if len(object1) < 2 {
		return 0, fmt.Errorf("wrong price range start format %v", object1)
	}

	price, ok := object1[1].(float64)
	if !ok {
		return 0, fmt.Errorf("wrong price range start format %v", object1[1])
	}

	return price, nil
}

func getPriceRange(object interface{}) (PriceRange, error) {
	object1, ok := object.([]interface{})
	if !ok {
		return PriceRange{}, fmt.Errorf("wrong price range format %v", object1)
	}
	if len(object1) < 6 {
		return PriceRange{}, fmt.Errorf("wrong price range format %v", object1)
	}
	startPrice, err := getPrice1(object1[4])
	if err != nil {
		return PriceRange{}, fmt.Errorf("wrong price range start format %w", err)
	}
	endPrice, err := getPrice1(object1[5])
	if !ok {
		return PriceRange{}, fmt.Errorf("wrong price range start format %w", err)
	}
	return PriceRange{startPrice, endPrice}, nil
}

func __getOffers(bytesToDecode []byte) ([]FullOffer, *PriceRange, error) {
	var outerObject [][]interface{}
	err := json.NewDecoder(bytes.NewReader(bytesToDecode)).Decode(&outerObject)
	if err != nil {
		return nil, nil, err
	}

	if len(outerObject[0]) < 3 {
		return nil, nil, fmt.Errorf("unexpected object format")
	}

	toDecode, ok := outerObject[0][2].(string)
	if !ok {
		return nil, nil, fmt.Errorf("unexpected object format 1")
	}
	var innerObject []interface{}
	err = json.NewDecoder(bytes.NewReader([]byte(toDecode))).Decode(&innerObject)
	if err != nil {
		return nil, nil, err
	}

	allOffers := []FullOffer{}

	section, ok := innerObject[2].([]interface{})
	if !ok {
		return allOffers, nil, fmt.Errorf("unexpected object format 2")
	}
	offers, err := getOffersFromSection(section)
	allOffers = append(allOffers, offers...)
	if err != nil {
		return allOffers, nil, err
	}

	section, ok = innerObject[3].([]interface{})
	if !ok {
		return allOffers, nil, fmt.Errorf("unexpected object format 3")
	}
	offers, err = getOffersFromSection(section)
	allOffers = append(allOffers, offers...)
	if err != nil {
		return allOffers, nil, err
	}

	priceRange, err := getPriceRange(innerObject[5])
	if err != nil {
		return allOffers, &priceRange, err
	}

	return allOffers, &priceRange, nil
}

func (s *Session) _getOffers(
	date, returnDate time.Time,
	srcCities, srcAirports, dstCities, dstAirports []string,
	adults int,
	curr currency.Unit,
	stops Stops,
	class Class,
	tripType TripType,
	lang language.Tag,
) ([]FullOffer, PriceRange, error) {
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

	buf, _ := io.ReadAll(resp.Body)

	// fmt.Println(string(buf))

	body := bufio.NewReader(bytes.NewReader(buf))

	skipPrefix(body)

	for {
		readLine(body)
		bytesToDecode, err := readLine(body)
		if err != nil {
			return finalOffers, finalPriceRange, nil
		}
		offers, priceRange, err := __getOffers(bytesToDecode)
		if offers != nil {
			finalOffers = append(finalOffers, offers...)
		}
		if priceRange != nil {
			finalPriceRange = *priceRange
		}
	}
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
	// TODO: Add date limit
	var allOffers []FullOffer
	var priceRange PriceRange
	var err error

	for i := 0; i < 1; i++ {
		allOffers, priceRange, err = s._getOffers(
			date, returnDate,
			srcCities, srcAirports, dstCities, dstAirports,
			adults, curr, stops, class, tripType, lang)
		if err == nil {
			return allOffers, priceRange, nil
		}
		fmt.Println("RETRY!!!!!!!!!!!!!!!!!")
	}

	return allOffers, priceRange, fmt.Errorf("number of retries 4 exceeded: %w", err)
}
