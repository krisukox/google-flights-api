package api

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"

	"golang.org/x/text/currency"
)

type flightV2 struct {
	departureAirportCode string
	departureAirportName string
	arrivalAirportName   string
	arrivalAirportCode   string
	departureTime        time.Time
	arrivalTime          time.Time
	duration             time.Duration
	airplane             string
	flightNumber         string
	unknown              interface{}
	airlineName          string
	legroom              string
}

func (f flightV2) String() string {
	out := ""
	out += fmt.Sprintf("departureAirportCode: %s ", f.departureAirportCode)
	out += fmt.Sprintf("departureAirportName: %s ", f.departureAirportName)
	out += fmt.Sprintf("arrivalAirportName: %s ", f.arrivalAirportName)
	out += fmt.Sprintf("arrivalAirportCode: %s ", f.arrivalAirportCode)
	out += fmt.Sprintf("departureTime: %s ", f.departureTime)
	out += fmt.Sprintf("arrivalTime: %s ", f.arrivalTime)
	out += fmt.Sprintf("duration: %s ", f.duration)
	out += fmt.Sprintf("airplane: %s ", f.airplane)
	out += fmt.Sprintf("flightNumber: %s ", f.flightNumber)
	out += fmt.Sprintf("unknown: %v ", f.unknown)
	out += fmt.Sprintf("airlineName: %s ", f.airlineName)
	out += fmt.Sprintf("legroom: %s ", f.legroom)
	return out
}

type trip struct {
	flight                 []flightV2
	returnFlight           []flightV2 // Not implemented yet
	originAirportCode      string
	destinationAirportCode string
	duration               time.Duration
	price                  float64
}

func (t trip) String() string {
	out := ""
	out += fmt.Sprintf("flight: %s \n", t.flight)
	out += fmt.Sprintf("returnFlight: %s \n", t.returnFlight)
	out += fmt.Sprintf("originAirportCode: %s \n", t.originAirportCode)
	out += fmt.Sprintf("destinationAirportCode: %s \n", t.destinationAirportCode)
	out += fmt.Sprintf("duration: %s \n", t.duration)
	out += fmt.Sprintf("price: %f \n", t.price)
	return out
}

// flightDate
// returnFlightDate
// originCity
// targetCity

// decodedMy := `[null,"[[null,null,null,\"HJRBIeHroQnsACCh4QBG--------lmbhp15AAAAAGScKcsGaWrAA\"],[null,null,1,null,[],1,[1,0,0,0],null,null,null,null,null,null,[` +
// `[[[[\"/m/0845b\",4]]],[[[\"/m/056_y\",5]]],null,0,[],[],\"2023-12-01\",null,[],[],[],null,null,[],3],` +
// `[[[[\"/m/056_y\",5]]],[[[\"/m/0845b\",4]]],null,0,[],[],\"2023-12-08\",null,[],[],[],null,null,[],3]],` +
// `null,null,null,1,null,null,null,null,null,[]],1,0,0]"]`

// decodedMy := `[null,"[[null,null,null,\"HJRBIeHroQnsACCh4QBG--------lmbhp15AAAAAGScKcsGaWrAA\"],[null,null,1,null,[],1,[1,0,0,0],null,null,null,null,null,null,[[[[[\"/m/0845b\",4]]],[[[\"/m/056_y\",5]]],null,0,[],[],\"2023-12-01\",null,[],[],[],null,null,[],3],[[[[\"/m/056_y\",5]]],[[[\"/m/0845b\",4]]],null,0,[],[],\"2023-12-08\",null,[],[],[],null,null,[],3]],null,null,null,1,null,null,null,null,null,[]],1,0,0]"]`

// f.req=%5Bnull%2C%22%5B%5Bnull%2Cnull%2Cnull%2C%5C%22HJRBIeHroQnsACCh4QBG--------lmbhp15AAAAAGScKcsGaWrAA%5C%22%5D%2C%5Bnull%2Cnull%2C1%2Cnull%2C%5B%5D%2C1%2C%5B1%2C0%2C0%2C0%5D%2Cnull%2Cnull%2Cnull%2Cnull%2Cnull%2Cnull%2C%5B%5B%5B%5B%5B%5C%22%2Fm%2F0845b%5C%22%2C4%5D%5D%5D%2C%5B%5B%5B%5C%22%2Fm%2F056_y%5C%22%2C5%5D%5D%5D%2Cnull%2C0%2C%5B%5D%2C%5B%5D%2C%5C%222023-12-01%5C%22%2Cnull%2C%5B%5D%2C%5B%5D%2C%5B%5D%2Cnull%2Cnull%2C%5B%5D%2C3%5D%2C%5B%5B%5B%5B%5C%22%2Fm%2F056_y%5C%22%2C5%5D%5D%5D%2C%5B%5B%5B%5C%22%2Fm%2F0845b%5C%22%2C4%5D%5D%5D%2Cnull%2C0%2C%5B%5D%2C%5B%5D%2C%5C%222023-12-08%5C%22%2Cnull%2C%5B%5D%2C%5B%5D%2C%5B%5D%2Cnull%2Cnull%2C%5B%5D%2C3%5D%5D%2Cnull%2Cnull%2Cnull%2C1%2Cnull%2Cnull%2Cnull%2Cnull%2Cnull%2C%5B%5D%5D%2C1%2C0%2C0%5D%22%5D&at=AAuQa1qjMakasqKYcQeoFJjN7RZ3%3A1687955915303&

func GetRawData(date, returnDate time.Time, originCity, targetCity string) string {
	// encodedValue := "%5Bnull%2C%22%5B%5Bnull%2Cnull%2Cnull%2C%5C%22HJRBIeHroQnsACCh4QBG--------lmbhp15AAAAAGScKcsGaWrAA%5C%22%5D%2C%5Bnull%2Cnull%2C1%2Cnull%2C%5B%5D%2C1%2C%5B1%2C0%2C0%2C0%5D%2Cnull%2Cnull%2Cnull%2Cnull%2Cnull%2Cnull%2C%5B%5B%5B%5B%5B%5C%22%2Fm%2F0845b%5C%22%2C4%5D%5D%5D%2C%5B%5B%5B%5C%22%2Fm%2F056_y%5C%22%2C5%5D%5D%5D%2Cnull%2C0%2C%5B%5D%2C%5B%5D%2C%5C%222023-12-01%5C%22%2Cnull%2C%5B%5D%2C%5B%5D%2C%5B%5D%2Cnull%2Cnull%2C%5B%5D%2C3%5D%2C%5B%5B%5B%5B%5C%22%2Fm%2F056_y%5C%22%2C5%5D%5D%5D%2C%5B%5B%5B%5C%22%2Fm%2F0845b%5C%22%2C4%5D%5D%5D%2Cnull%2C0%2C%5B%5D%2C%5B%5D%2C%5C%222023-12-08%5C%22%2Cnull%2C%5B%5D%2C%5B%5D%2C%5B%5D%2Cnull%2Cnull%2C%5B%5D%2C3%5D%5D%2Cnull%2Cnull%2Cnull%2C1%2Cnull%2Cnull%2Cnull%2Cnull%2Cnull%2C%5B%5D%5D%2C1%2C0%2C0%5D%22%5D"
	// decodedValue, err := url.QueryUnescape(encodedValue)
	// if err != nil {
	// 	log.Fatalf(err.Error())
	// }
	// fmt.Println(decodedValue)

	serializedOriginCity, err := GetSerializedCityName(originCity)
	if err != nil {
		log.Fatalf(err.Error())
	}
	serializedTargetCity, err := GetSerializedCityName(targetCity)
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

	// decodedMy := `[null,"[[null,null,null,\"HJRBIeHroQnsACCh4QBG--------lmbhp15AAAAAGScKcsGaWrAA\"],[null,null,1,null,[],1,[1,0,0,0],null,null,null,null,null,null,[[[[[\"/m/0845b\",4]]],[[[\"/m/056_y\",5]]],null,0,[],[],\"2023-12-01\",null,[],[],[],null,null,[],3],[[[[\"/m/056_y\",5]]],[[[\"/m/0845b\",4]]],null,0,[],[],\"2023-12-08\",null,[],[],[],null,null,[],3]],null,null,null,1,null,null,null,null,null,[]],1,0,0]"]`

	// fmt.Println(decodedMy)

	// fmt.Println(url.QueryEscape(decodedMy))

	return url.QueryEscape(decodedMy)

	// encodedValue = "AAuQa1qjMakasqKYcQeoFJjN7RZ3%3A1687955915303"
	// decodedValue, err = url.QueryUnescape(encodedValue)
	// if err != nil {
	// 	log.Fatalf(err.Error())
	// }
	// fmt.Println(decodedValue)
}

func decode(toDecode string) ([]flight, error) {
	var outerObject []interface{}
	err := json.NewDecoder(bytes.NewReader([]byte(toDecode))).Decode(&outerObject)
	if err != nil {
		return nil, err
	}

	return []flight{}, nil
}

func getPrice(object []interface{}) (float64, error) {
	object1, ok := object[1].([]interface{})
	if !ok {
		return 0, fmt.Errorf("wrong price format stage 1: %v", object[1])
	}
	object2, ok := object1[0].([]interface{})
	if !ok {
		return 0, fmt.Errorf("wrong price format stage 2: %v", object1[0])
	}
	price, ok := object2[1].(float64)
	if !ok {
		return 0, fmt.Errorf("wrong price format stage 3: %v", object2[1])
	}
	return price, nil
}

func getAirlineName(object []interface{}) (string, error) {
	object1, ok := object[1].([]interface{})
	if !ok {
		return "", fmt.Errorf("wrong airline name format stage 1: %v", object[1])
	}
	object2, ok := object1[0].(string)
	if !ok {
		return "", fmt.Errorf("wrong airline name format stage 2: %v", object1[0])
	}
	return object2, nil
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

func getDuration(duration interface{}) (time.Duration, error) {
	duration1, ok := duration.(float64)
	if !ok {
		return time.Duration(0), fmt.Errorf("wrong duration format: %v", duration)
	}
	return time.ParseDuration(fmt.Sprintf("%dm", int(duration1)))
}

func getFlightNumberAirline(data interface{}) (string, interface{}, string, error) {
	data1, ok := data.([]interface{})
	if !ok || len(data1) != 4 {
		return "", "", "", fmt.Errorf("wrong flight number of airline type: %v", data)
	}
	flightNumberPart1, ok := data1[0].(string)
	if !ok {
		return "", "", "", fmt.Errorf("wrong flight number part 1 type: %v", data1[0])
	}
	flightNumberPart2, ok := data1[1].(string)
	if !ok {
		return "", "", "", fmt.Errorf("wrong flight number part 2 type: %v", data1[1])
	}
	airline, ok := data1[3].(string)
	if !ok {
		return "", "", "", fmt.Errorf("wrong airline name type: %v", data1[3])
	}
	return flightNumberPart1 + " " + flightNumberPart2, data1[2], airline, nil
}

func getFlight(object interface{}) (flightV2, error) {
	object1, ok := object.([]interface{})
	if !ok {
		return flightV2{}, fmt.Errorf("wrong flight format: %v", object)
	}
	if len(object1) < 32 {
		return flightV2{}, fmt.Errorf("incorrect amount of data for type: %v", object)
	}
	departureAirportCode, ok := object1[3].(string)
	if !ok {
		return flightV2{}, fmt.Errorf("wrong departure airport code type: %v", object1[3])
	}
	departureAirportName, ok := object1[4].(string)
	if !ok {
		return flightV2{}, fmt.Errorf("wrong departure airport name type: %v", object1[4])
	}
	arrivalAirportName, ok := object1[5].(string)
	if !ok {
		return flightV2{}, fmt.Errorf("wrong arrival airport name type: %v", object1[5])
	}
	arrivalAirportCode, ok := object1[6].(string)
	if !ok {
		return flightV2{}, fmt.Errorf("wrong arrival airport code type: %v", object1[6])
	}
	departureTime, err := getTime(object1[8], object1[20])
	if err != nil {
		return flightV2{}, fmt.Errorf("departure: %w", err)
	}
	arrivalTime, err := getTime(object1[10], object1[21])
	if err != nil {
		return flightV2{}, fmt.Errorf("arrival: %w", err)
	}
	duration, err := getDuration(object1[11])
	if err != nil {
		return flightV2{}, err
	}
	flightNumber, unknown, airlineName, err := getFlightNumberAirline(object1[22])
	if err != nil {
		return flightV2{}, err
	}
	airplane, ok := object1[17].(string)
	if !ok {
		return flightV2{}, fmt.Errorf("wrong airplane format: %v", object1[17])
	}
	legroom, _ := object1[30].(string)
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
		unknown,
		airlineName,
		legroom,
	}, nil
}

func getFlights(object []interface{}) ([]flightV2, error) {
	flights := []flightV2{}

	object1, ok := object[0].([]interface{})
	if !ok {
		return nil, fmt.Errorf("wrong flights format stage 1: %v", object[0])
	}
	flightObjs, ok := object1[2].([]interface{})
	if !ok {
		return nil, fmt.Errorf("wrong flights format stage 2: %v", object1[2])
	}
	for _, f := range flightObjs {
		fl, err := getFlight(f)
		if err != nil {
			return nil, fmt.Errorf("cannot get flight: %w", err)
		}
		flights = append(flights, fl)
	}
	// fmt.Println(flights[0])
	// fmt.Println(flights[1])

	return flights, nil
}

func getTripDuration(flights []flightV2) time.Duration {
	return flights[len(flights)-1].departureTime.Sub(flights[0].departureTime)
}

func getTrip(object []interface{}) (trip, error) {
	price, err := getPrice(object)
	if err != nil {
		return trip{}, err
	}
	flights, err := getFlights(object)
	if err != nil {
		return trip{}, err
	}
	return trip{flights, []flightV2{}, "", "", getTripDuration(flights), price}, nil
}

func getFlightsFromSection(section []interface{}) ([]trip, error) {

	trips := []trip{}

	object, ok := section[0].([]interface{})
	if !ok {
		return nil, fmt.Errorf("unexpected object format")
	}
	ok = true
	var object1 []interface{}
	for _, o := range object {
		fmt.Println(o)
		// fmt.Println("getFlightsFromSection")
		object1, ok = o.([]interface{})
		if !ok {
			break
		}
		trip, err := getTrip(object1)
		if err != nil {
			fmt.Println(">>>>>>>>>>>>>>>>>ERROR %v", err)
		}
		// fmt.Println(trip)
		if err != nil {
			break
		}
		trips = append(trips, trip)
	}
	// // for _, o := range object {
	// object1, _ = object[0].([]interface{})

	// // if !ok {
	// // 	break
	// // }
	// trip, _ := getTrip(object1)

	// // fmt.Println(trip)

	// // if err != nil {
	// // 	break
	// // }
	// trips = append(trips, trip)
	// }
	return trips, nil
}

func GetFlightsV2(
	date time.Time,
	returnDate time.Time,
	originCity string,
	targetCity string,
	unit currency.Unit,
) ([]flight, error) {
	fmt.Println("GetFlightsV2")
	url := "https://www.google.com/_/TravelFrontendUi/data/travel.frontend.flights.FlightsFrontendService/GetShoppingResults?f.sid=-1300922759171628473&bl=boq_travel-frontend-ui_20230627.02_p1&hl=en&soc-app=162&soc-platform=1&soc-device=1&_reqid=52717&rt=c"

	jsonBody := []byte(`f.req=` + GetRawData(date, returnDate, originCity, targetCity) + `&at=AAuQa1qjMakasqKYcQeoFJjN7RZ3%3A1687955915303&`)
	// jsonBody := []byte(`f.req=%5Bnull%2C%22%5B%5Bnull%2Cnull%2Cnull%2C%5C%22HJRBIeHroQnsACCh4QBG--------lmbhp15AAAAAGScKcsGaWrAA%5C%22%5D%2C%5Bnull%2Cnull%2C1%2Cnull%2C%5B%5D%2C1%2C%5B1%2C0%2C0%2C0%5D%2Cnull%2Cnull%2Cnull%2Cnull%2Cnull%2Cnull%2C%5B%5B%5B%5B%5B%5C%22%2Fm%2F0845b%5C%22%2C4%5D%5D%5D%2C%5B%5B%5B%5C%22%2Fm%2F056_y%5C%22%2C5%5D%5D%5D%2Cnull%2C0%2C%5B%5D%2C%5B%5D%2C%5C%222023-12-01%5C%22%2Cnull%2C%5B%5D%2C%5B%5D%2C%5B%5D%2Cnull%2Cnull%2C%5B%5D%2C3%5D%2C%5B%5B%5B%5B%5C%22%2Fm%2F056_y%5C%22%2C5%5D%5D%5D%2C%5B%5B%5B%5C%22%2Fm%2F0845b%5C%22%2C4%5D%5D%5D%2Cnull%2C0%2C%5B%5D%2C%5B%5D%2C%5C%222023-12-08%5C%22%2Cnull%2C%5B%5D%2C%5B%5D%2C%5B%5D%2Cnull%2Cnull%2C%5B%5D%2C3%5D%5D%2Cnull%2Cnull%2Cnull%2C1%2Cnull%2Cnull%2Cnull%2Cnull%2Cnull%2C%5B%5D%5D%2C1%2C0%2C0%5D%22%5D&at=AAuQa1qjMakasqKYcQeoFJjN7RZ3%3A1687955915303&`)
	bodyReader := bytes.NewReader(jsonBody)

	req, err := http.NewRequest(http.MethodPost, url, bodyReader)
	req.Header.Set("authority", `www.google.com`)
	req.Header.Set("accept", `*/*`)
	req.Header.Set("accept-language", `en-US,en;q=0.9`)
	req.Header.Set("cache-control", `no-cache`)
	req.Header.Set("content-type", `application/x-www-form-urlencoded;charset=UTF-8`)
	req.Header.Set("cookie", `CONSENT=YES+srp.gws-20211208-0-RC2.en+FX+830; SEARCH_SAMESITE=CgQI5ZcB; OGP=-19031986:; OGPC=19031986-2:; OTZ=7066515_48_52_123900_48_436380; SID=YAhP_vXgn_DuaJ_MQKEanHWfq_kbeUng2kd4gnv5ZJpHpzyL9xo8pwryk4MhjzbfT2WGgg.; __Secure-1PSID=YAhP_vXgn_DuaJ_MQKEanHWfq_kbeUng2kd4gnv5ZJpHpzyL1XDedYjapIfRYZEL22i1LQ.; __Secure-3PSID=YAhP_vXgn_DuaJ_MQKEanHWfq_kbeUng2kd4gnv5ZJpHpzyLIjEbyOLGtI0dotafz6XZkQ.; HSID=AuHUufc7rt6o8QAWj; SSID=ABhM7gpFBIR6q20cX; APISID=gLl92_gy6y2lL6nl/AL5fEw8A1bT0-ZcVc; SAPISID=Xb0LnsmdCHatABmU/AUPKah2QtC9FT3rCF; __Secure-1PAPISID=Xb0LnsmdCHatABmU/AUPKah2QtC9FT3rCF; __Secure-3PAPISID=Xb0LnsmdCHatABmU/AUPKah2QtC9FT3rCF; AEC=AUEFqZcNUdbExn0DgShHLjuyEnKmgxq-_9_UkOzckjNdEiXNESdjNi4tAA; 1P_JAR=2023-06-28-12; NID=511=XwP-sF1jLScb0X4IQw410cNyVKj7JwLrCKEjKTa2yVu_oncmegISuFAUZ64i8gZfDZjmaMTcyn1Eddh5VDXDGxEF28hV1wrqIGeIThh0O5uR_MWaKOr5mN8blmhjyRFuN0GD0NmKewaZMxNWx6gL8hfNptyIsQsDyoexWgHSPddx9_PYD576fNBIse4Z_B9L3ZRYtsH92klx-2kQUS4WODXBk1im4OYevp3blem5ZRmB5_o2LzJnCp1lCuBuxR4k2qZgYtfqOXLkHfUlSjxJx_0nsXOHkTgHJx3lIVs84_hM_G69bQC2sdEfVLiPLPT-K5Bl9u8NBBAVidwY6IKaS1SUAlfRib8mhDpDVItSHaSBebwoY2z7PVFGhk3WsnWKiUQ_WWS_JWjMFQQktQKgrLZ5u7uHfz7ncZsBDdEtMd35lFcN; SIDCC=AP8dLtxjIANhYZzqRrqMT4tVMLTaF0G-uXY74s-3OZIs0RzjNbhajDZmjz1_dGLaZeh7xCvV4Rc; __Secure-1PSIDCC=AP8dLtxDdiAzXb3-weuMb0RSJ-GopCdBVLy8lynPvRWp6XO7vOQArdHB0J2alE3Ka0wo9YRge7o; __Secure-3PSIDCC=AP8dLtwqo57oD-LKpBzN-LQp7r1_7-tK8unYWlTImkBoyZEtdrc9uCzaSoytqaS4RTqkAw-fSQ5c`)
	req.Header.Set("origin", `https://www.google.com`)
	req.Header.Set("pragma", `no-cache`)
	req.Header.Set("sec-ch-ua", `"Google Chrome";v="113", "Chromium";v="113", "Not-A.Brand";v="24"`)
	req.Header.Set("sec-ch-ua-arch", `"x86"`)
	req.Header.Set("sec-ch-ua-bitness", `"64"`)
	req.Header.Set("sec-ch-ua-full-version", `"113.0.5672.92"`)
	req.Header.Set("sec-ch-ua-full-version-list", `"Google Chrome";v="113.0.5672.92", "Chromium";v="113.0.5672.92", "Not-A.Brand";v="24.0.0.0"`)
	req.Header.Set("sec-ch-ua-mobile", `?0`)
	req.Header.Set("sec-ch-ua-model", `""`)
	req.Header.Set("sec-ch-ua-platform", `"Linux"`)
	req.Header.Set("sec-ch-ua-platform-version", `"5.19.0"`)
	req.Header.Set("sec-ch-ua-wow64", `?0`)
	req.Header.Set("sec-fetch-dest", `empty`)
	req.Header.Set("sec-fetch-mode", `cors`)
	req.Header.Set("sec-fetch-site", `same-origin`)
	req.Header.Set("user-agent", `Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/113.0.0.0 Safari/537.36`)
	req.Header.Set("x-goog-ext-259736195-jspb", `["en-US","PL","PLN",1,null,[-120],null,[[48676280,48710756,47907128,48764689,48627726,48480739,48593234,48707380]],1,[]]`) // language, location, currency
	req.Header.Set("x-same-domain", `1`)
	client := http.Client{
		Timeout: 30 * time.Second,
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body := bufio.NewReader(resp.Body)
	skipPrefix(body)
	fmt.Println(">>>>>>>>>abc")
	bytesToDecode, err := readLine(body)
	// bytesToDecode, isPrefix, err := body.ReadLine()
	// if isPrefix {
	// 	return nil, fmt.Errorf("Too long structure to decode")
	// }
	// fmt.Println(isPrefix)
	if err != nil {
		return nil, err
	}
	// fmt.Println(string(bytesToDecode))
	// if isPrefix {
	// 	return nil, fmt.Errorf("Too long structure to decode")
	// }
	// fmt.Println(">>>>>>>>>abc")
	var outerObject [][]interface{}
	err = json.NewDecoder(bytes.NewReader(bytesToDecode)).Decode(&outerObject)
	if err != nil {
		return nil, err
	}
	toDecode, ok := outerObject[0][2].(string)
	if !ok {
		return nil, fmt.Errorf("unexpected object format")
	}
	var innerObject []interface{}
	err = json.NewDecoder(bytes.NewReader([]byte(toDecode))).Decode(&innerObject)
	if err != nil {
		return nil, err
	}
	// // // fmt.Println(">>>>>>>>>abc")
	// fmt.Println(innerObject[2])

	allTrips := []trip{}
	section, ok := innerObject[2].([]interface{})
	if !ok {
		return nil, fmt.Errorf("unexpected object format 2")
	}
	trips, err := getFlightsFromSection(section)
	if err != nil {
		return nil, err
	}

	allTrips = append(allTrips, trips...)
	section, ok = innerObject[3].([]interface{})
	if !ok {
		return nil, fmt.Errorf("unexpected object format 2")
	}
	trips, err = getFlightsFromSection(section)
	if err != nil {
		return nil, err
	}
	allTrips = append(allTrips, trips...)
	fmt.Println(allTrips)

	return []flight{}, nil
}
