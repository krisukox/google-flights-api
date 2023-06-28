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

	// toDecode2, ok := outerObject[1].(string)
	// if !ok {
	// 	return nil, fmt.Errorf("unexpected object format")
	// }

	// fmt.Println(toDecode2)

	return []flight{}, nil
}

func getFlight(object []interface{}) (flight, error) {
	object1, ok := object[1].([]interface{})
	if !ok {
		return flight{}, fmt.Errorf("unexpected object format")
	}
	object2, ok := object1[0].([]interface{})
	if !ok {
		return flight{}, fmt.Errorf("unexpected object format")
	}
	price := object2[1].(float64)
	return flight{"", "", fmt.Sprintf("%f", price)}, nil
}

func getFlightsFromSection(section []interface{}) ([]flight, error) {
	flights := []flight{}

	object, ok := section[0].([]interface{})
	if !ok {
		return nil, fmt.Errorf("unexpected object format")
	}
	ok = true
	var object1 []interface{}
	for _, o := range object {
		object1, ok = o.([]interface{})
		if !ok {
			break
		}
		flight, err := getFlight(object1)
		if err != nil {
			break
		}
		flights = append(flights, flight)
	}
	// for i := 0; i < len(object); i++ {
	// 	object1, ok = object[i].([]interface{})
	// 	if !ok {
	// 		break
	// 	}
	// 	flight, err := getFlight(object1)
	// 	if err != nil {
	// 		break
	// 	}
	// 	flights = append(flights, flight)
	// }
	// object1, ok := object[0].([]interface{}) // 0, 1, 2 kolejne loty
	// getFlight(object1)
	// if !ok {
	// 	return nil, fmt.Errorf("unexpected object format")
	// }
	// object3, ok := object2[1].([]interface{})
	// if !ok {
	// 	return nil, fmt.Errorf("unexpected object format")
	// }
	// object4, ok := object3[0].([]interface{})
	// if !ok {
	// 	return nil, fmt.Errorf("unexpected object format")
	// }
	// fmt.Println(object4)
	return flights, nil
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
	req.Header.Set("x-goog-ext-259736195-jspb", `["en-US","PL","PLN",1,null,[-120],null,[[48676280,48710756,47907128,48764689,48627726,48480739,48593234,48707380]],1,[]]`)
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
	allFlights := []flight{}
	section, ok := innerObject[2].([]interface{})
	if !ok {
		return nil, fmt.Errorf("unexpected object format 2")
	}
	flights, err := getFlightsFromSection(section)
	if err != nil {
		return nil, err
	}
	allFlights = append(allFlights, flights...)
	section, ok = innerObject[3].([]interface{})
	if !ok {
		return nil, fmt.Errorf("unexpected object format 2")
	}
	flights, err = getFlightsFromSection(section)
	if err != nil {
		return nil, err
	}
	allFlights = append(allFlights, flights...)
	fmt.Println(allFlights)
	// getFlightsFromSection(section)
	// // fmt.Println(anotherObject)
	// // fmt.Println(anotherObject[1])
	// anotherObject1, ok := anotherObject[0].([]interface{})

	// ///////////////////////

	// anotherObject2, ok := anotherObject1[2].([]interface{}) // 0, 1, 2 kolejne loty
	// anotherObject3, ok := anotherObject2[1].([]interface{})
	// anotherObject4, ok := anotherObject3[0].([]interface{})
	// fmt.Println(anotherObject4[1])

	///////////////////////

	return []flight{}, nil
}
