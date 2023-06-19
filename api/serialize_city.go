package api

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

type WholeBody struct {
	Body []WholeBody2
}

type WholeBody2 struct {
	Body []Body
}

type Body struct {
	Unkn1    string
	Unkn2    string
	ToDecode string
	Unkn3    int
	Unkn4    int
	Unkn5    int
	Unkn6    string
}

func skipPrefix(body *bufio.Reader) error {
	_, isPrefix, err := body.ReadLine()
	if err != nil || isPrefix {
		return fmt.Errorf("error when reading response with the serialized city names: %w", err)
	}
	_, isPrefix, err = body.ReadLine()
	if err != nil || isPrefix {
		return fmt.Errorf("error when reading response with the serialized city names: %w", err)
	}
	return nil
}

func getNumber(body *bufio.Reader) (int, error) { // Use ReadLine
	var number []byte
	number, isPrefix, err := body.ReadLine()
	if err != nil || isPrefix {
		return 0, fmt.Errorf("error when reading response with the serialized city names: %w", err)
	}
	return strconv.Atoi(string(number))
}

func GetSerializedCityName(cityName string) (string, error) {
	requestURL := "https://www.google.com/_/TravelFrontendUi/data/batchexecute?rpcids=H028ib&source-path=%2Ftravel%2Fflights%2Fsearch&f.sid=-8421128425468344897&bl=boq_travel-frontend-ui_20230613.06_p0&hl=pl&soc-app=162&soc-platform=1&soc-device=1&_reqid=444052&rt=c"

	jsonBody := []byte(`f.req=%5B%5B%5B%22H028ib%22%2C%22%5B%5C%22` + cityName + `%5C%22%2C%5B1%2C2%2C3%2C5%2C4%5D%2Cnull%2C%5B1%2C1%2C1%5D%2C1%5D%22%2Cnull%2C%22generic%22%5D%5D%5D&at=AAuQa1qJpLKW2Hl-i40OwJyzmo22%3A1687083247610&`)
	bodyReader := bytes.NewReader(jsonBody)

	req, err := http.NewRequest(http.MethodPost, requestURL, bodyReader)
	if err != nil {
		log.Fatalf("Couldn't create a request")
	}
	req.Header.Set("authority", "www.google.com")
	req.Header.Set("accept", "*/*")
	req.Header.Set("accept-language", "pl-PL,pl;q=0.9,en-US;q=0.8,en;q=0.7,hr;q=0.6")
	req.Header.Set("cache-control", "no-cache")
	req.Header.Set("content-type", "application/x-www-form-urlencoded;charset=UTF-8")
	req.Header.Set("cookie", "CONSENT=YES+srp.gws-20211208-0-RC2.pl+FX+371")
	req.Header.Set("origin", "https://www.google.com")
	req.Header.Set("pragma", "no-cache")
	req.Header.Set("referer", "https://www.google.com/travel/flights/search?tfs=CBwQAhonEgoyMDIzLTA3LTA0agwIAhIIL20vMDg0NWJyCwgDEgcvbS8wbjJ6GicSCjIwMjMtMDctMDhqCwgDEgcvbS8wbjJ6cgwIAhIIL20vMDg0NWJAAUgBcAGCAQsI____________AZgBAQ")
	req.Header.Set("sec-ch-ua", "\"Google Chrome\";v=\"113\", \"Chromium\";v=\"113\", \"Not-A.Brand\";v=\"24\"")
	req.Header.Set("sec-ch-ua-arch", "\"x86\"")
	req.Header.Set("sec-ch-ua-bitness", "\"64\"")
	req.Header.Set("sec-ch-ua-full-version", "\"113.0.5672.92\"")
	req.Header.Set("sec-ch-ua-full-version-list", "\"Google Chrome\";v=\"113.0.5672.92\", \"Chromium\";v=\"113.0.5672.92\", \"Not-A.Brand\";v=\"24.0.0.0\"")
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-model", "")
	req.Header.Set("sec-ch-ua-platform", "Linux")
	req.Header.Set("sec-ch-ua-platform-version", "5.19.0")
	req.Header.Set("sec-ch-ua-wow64", "?0")
	req.Header.Set("sec-fetch-dest", "empty")
	req.Header.Set("sec-fetch-mode", "cors")
	req.Header.Set("sec-fetch-site", "same-origin")
	req.Header.Set("user-agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/113.0.0.0 Safari/537.36")
	req.Header.Set("x-client-data", "CKS1yQEIj7bJAQimtskBCKmdygEIivHKAQiSocsBCIegzQEI8abNAQj9qs0BCKCtzQEItrLNAQicuc0BCNm5zQEI3L3NAQ==")
	req.Header.Set("x-goog-ext-190139975-jspb", "[\"PL\",\"ZZ\",\"Xkt6eg==\"]")
	req.Header.Set("x-goog-ext-259736195-jspb", "[\"pl\",\"PL\",\"PLN\",1,null,[-120],null,[[48764690,48627726,47907128,48480739,48710756,48676280,48593234,48707378]],1,[]]")
	req.Header.Set("x-same-domain", "1")

	client := http.Client{
		Timeout: 30 * time.Second,
	}

	resp, err := client.Do(req)

	if err != nil {
		log.Fatalf("Couldn't send a request")
	}
	defer resp.Body.Close()

	body := bufio.NewReader(resp.Body)
	skipPrefix(body)
	_, err = getNumber(body)

	if err != nil {
		log.Fatalf("Couldn't get number")
	}

	bytesToDecode, _, err := body.ReadLine()

	if err != nil {
		log.Fatalf("Couldn't read line")
	}

	var decoded [][]interface{}
	err = json.NewDecoder(bytes.NewReader(bytesToDecode)).Decode(&decoded)
	if err != nil {
		log.Fatalf("Couldn't decode")
	}
	toDecode, ok := decoded[0][2].(string)
	if !ok {
		log.Fatalf("Couldn't decode")
	}

	var decoded2 [][][][]interface{}
	err = json.NewDecoder(bytes.NewReader([]byte(toDecode))).Decode(&decoded2)
	if err != nil {
		log.Fatalf("Couldn't decode")
	}

	city, ok := decoded2[0][0][0][2].(string)
	if !ok {
		return "", fmt.Errorf("couldn't get city name")
	}
	serializedCity, ok := decoded2[0][0][0][4].(string)
	if !ok {
		return "", fmt.Errorf("couldn't get serialized city name")
	}
	if city == cityName {
		return serializedCity, nil
	}
	return "", fmt.Errorf("the requested city name didn't match the found")
}
