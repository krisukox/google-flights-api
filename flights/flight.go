package flights

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"time"

	"golang.org/x/text/currency"
	"golang.org/x/text/language"
)

func (s *Session) GetRawData(date, returnDate time.Time, originCity, targetCity string, lang language.Tag) string {
	serializedOriginCity, err := s.AbbrCity(originCity, lang)
	if err != nil {
		log.Fatalf(err.Error())
	}
	serializedTargetCity, err := s.AbbrCity(targetCity, lang)
	if err != nil {
		log.Fatalf(err.Error())
	}
	serializedDate := date.Format("2006-01-02")
	serializedTargetDate := returnDate.Format("2006-01-02")

	decodedMy := `[null,"[[null,null,null,\"HJRBIeHroQnsACCh4QBG--------lmbhp15AAAAAGScKcsGaWrAA\"],[null,null,1,null,[],1,[1,0,0,0],null,null,null,null,null,null,[` +
		fmt.Sprintf(`[[[[\"%s\",4]]],[[[\"%s\",5]]],null,0,[],[],\"%s\",null,[],[],[],null,null,[],3],`,
			serializedOriginCity, serializedTargetCity, serializedDate) +
		fmt.Sprintf(`[[[[\"%s\",5]]],[[[\"%s\",4]]],null,0,[],[],\"%s\",null,[],[],[],null,null,[],3]],`,
			serializedTargetCity, serializedOriginCity, serializedTargetDate) +
		`null,null,null,1,null,null,null,null,null,[]],1,0,0]"]`

	return url.QueryEscape(decodedMy)
}

func getPrice(tripObj []interface{}) (float64, error) {
	priceObj1, ok := tripObj[1].([]interface{})
	if !ok {
		return 0, fmt.Errorf("wrong price format stage 1: %v", priceObj1[1])
	}
	priceObj2, ok := priceObj1[0].([]interface{})
	if !ok {
		return 0, fmt.Errorf("wrong price format stage 2: %v", priceObj2[0])
	}
	price, ok := priceObj2[1].(float64)
	if !ok {
		return 0, fmt.Errorf("wrong price format stage 3: %v", priceObj2[1])
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

	hours, ok = flightTime1[0].(float64)
	if !ok {
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
	duration1 := getElement[float64](duration, 11)
	return time.ParseDuration(fmt.Sprintf("%dm", int(duration1)))
}

func getFlightNumberAirline(data interface{}) (string, interface{}, string, error) {
	data1, ok := data.([]interface{})
	if !ok || len(data1) != 4 {
		return "", "", "", fmt.Errorf("wrong flight number of airline type: %v", data)
	}
	flightNumberPart1 := getElement[string](data1, 0)
	// if !ok {
	//  return "", "", "", fmt.Errorf("wrong flight number part 1 type: %v", data1[0])
	// }
	flightNumberPart2 := getElement[string](data1, 1)
	// if !ok {
	//  return "", "", "", fmt.Errorf("wrong flight number part 2 type: %v", data1[1])
	// }
	airline := getElement[string](data1, 3)
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

func getFlight(flightObj interface{}) (flightV2, error) {
	var unknowns []interface{}
	flightObj1, ok := flightObj.([]interface{})
	if !ok {
		return flightV2{}, fmt.Errorf("wrong flight format: %v", flightObj)
	}

	departureAirportCode := getElement[string](flightObj1, 3)
	// if !ok {
	//  return flightV2{}, fmt.Errorf("wrong departure airport code type: %v", object1[3])
	// }
	departureAirportName := getElement[string](flightObj1, 4)
	// if !ok {
	//  return flightV2{}, fmt.Errorf("wrong departure airport name type: %v", object1[4])
	// }
	arrivalAirportName := getElement[string](flightObj1, 5)
	// if !ok {
	//  return flightV2{}, fmt.Errorf("wrong arrival airport name type: %v", object1[5])
	// }
	arrivalAirportCode := getElement[string](flightObj1, 6)
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
	airplane := getElement[string](flightObj1, 17)
	// if !ok {
	//  return flightV2{}, fmt.Errorf("wrong airplane format: %v", object1[17])
	// }
	legroom := getElement[string](flightObj1, 30)
	us := getUnknowns(flightObj1)
	unknowns = append(unknowns, us...)
	return flightV2{
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

func getFlights(tripObj []interface{}) ([]flightV2, error) {
	flights := []flightV2{}

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

func getTripDuration(flights []flightV2) time.Duration {
	return flights[len(flights)-1].DepartureTime.Sub(flights[0].DepartureTime)
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
	return FullOffer{Offer{time.Time{}, time.Time{}, price}, flights, []flightV2{}, "", "", "", "", getTripDuration(flights)}, nil
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
	originCity, targetCity string,
	currencyUnit currency.Unit,
	lang language.Tag,
) (*http.Response, error) {
	url := "https://www.google.com/_/TravelFrontendUi/data/travel.frontend.flights.FlightsFrontendService/GetShoppingResults?f.sid=-1300922759171628473&bl=boq_travel-frontend-ui_20230627.02_p1&hl=en&soc-app=162&soc-platform=1&soc-device=1&_reqid=52717&rt=c"

	jsonBody := []byte(`f.req=` + s.GetRawData(date, returnDate, originCity, targetCity, lang) + `&at=AAuQa1qjMakasqKYcQeoFJjN7RZ3%3A1687955915303&`) // Add current unix timestamp instead of 1687955915303
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
	req.Header.Set("x-goog-ext-259736195-jspb", fmt.Sprintf(`["en-US","PL","%s",1,null,[-120],null,[[48676280,48710756,47907128,48764689,48627726,48480739,48593234,48707380]],1,[]]`, currencyUnit)) // language, location, currency

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
	date time.Time,
	returnDate time.Time,
	originCity string,
	targetCity string,
	currencyUnit currency.Unit,
	lang language.Tag,
) ([]FullOffer, PriceRange, error) {
	finalOffers := []FullOffer{}
	finalPriceRange := PriceRange{}
	resp, err := s.doRequestFlights(date, returnDate, originCity, targetCity, currencyUnit, lang)

	if err != nil {
		return nil, PriceRange{}, err
	}
	defer resp.Body.Close()

	buf, _ := io.ReadAll(resp.Body)
	reader := bytes.NewReader(buf)
	body := bufio.NewReader(reader)

	skipPrefix(body)

	for {
		readLine(body)
		bytesToDecode, err := readLine(body)
		if err != nil {
			return nil, PriceRange{}, err
		}
		offers, priceRange, err := __getOffers(bytesToDecode)
		if offers != nil {
			finalOffers = append(finalOffers, offers...)
		}
		if priceRange != nil {
			finalPriceRange = *priceRange
		}

		if err != nil && err.Error() != "unexpected object format 2" {
			break
		}
	}

	return finalOffers, finalPriceRange, err
}

func (s *Session) GetOffers(
	date time.Time,
	returnDate time.Time,
	originCity string,
	targetCity string,
	currencyUnit currency.Unit,
	lang language.Tag,
) ([]FullOffer, PriceRange, error) {
	// TODO: Add date limit
	var allOffers []FullOffer
	var priceRange PriceRange
	var err error

	for i := 0; i < 1; i++ {
		allOffers, priceRange, err = s._getOffers(date, returnDate, originCity, targetCity, currencyUnit, lang)
		if err == nil {
			return allOffers, priceRange, nil
		}
		fmt.Println("RETRY!!!!!!!!!!!!!!!!!")
	}

	return allOffers, priceRange, fmt.Errorf("number of retries 4 exceeded: %w", err)
}
