package flights

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"golang.org/x/text/language"
)

func (s *Session) doRequestCity(city string, lang language.Tag) (*http.Response, error) {
	requestURL := "https://www.google.com/_/TravelFrontendUi/data/batchexecute?rpcids=H028ib&source-path=%2Ftravel%2Fflights%2Fsearch&f.sid=-8421128425468344897&bl=boq_travel-frontend-ui_20230613.06_p0" +
		"&hl=" + lang.String() +
		"&soc-app=162&soc-platform=1&soc-device=1&_reqid=444052&rt=c"

	jsonBody := []byte(`f.req=%5B%5B%5B%22H028ib%22%2C%22%5B%5C%22` + city + `%5C%22%2C%5B1%2C2%2C3%2C5%2C4%5D%2Cnull%2C%5B1%2C1%2C1%5D%2C1%5D%22%2Cnull%2C%22generic%22%5D%5D%5D&at=AAuQa1qJpLKW2Hl-i40OwJyzmo22%3A1687083247610&`)
	bodyReader := bytes.NewReader(jsonBody)

	req, err := http.NewRequest(http.MethodPost, requestURL, bodyReader)
	if err != nil {
		return nil, err
	}
	req.Header.Set("accept", "*/*")
	req.Header.Set("cache-control", "no-cache")
	req.Header.Set("content-type", "application/x-www-form-urlencoded;charset=UTF-8")
	req.Header.Set("cookie", `CONSENT=PENDING+672`)
	req.Header.Set("pragma", "no-cache")
	req.Header.Set("user-agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/113.0.0.0 Safari/537.36")

	return s.client.Do(req)
}

func decodeInnerObject(body *bufio.Reader) ([][][][]interface{}, error) {
	bytesToDecode, isPrefix, err := body.ReadLine()
	if err != nil || isPrefix {
		return nil, err
	}

	var outerObject [][]interface{}
	err = json.NewDecoder(bytes.NewReader(bytesToDecode)).Decode(&outerObject)
	if err != nil {
		return nil, err
	}
	toDecode, ok := outerObject[0][2].(string)
	if !ok {
		return nil, fmt.Errorf("unexpected object format")
	}

	var innerObject [][][][]interface{}
	err = json.NewDecoder(bytes.NewReader([]byte(toDecode))).Decode(&innerObject)
	if err != nil {
		return nil, err
	}
	return innerObject, nil
}

func (s *Session) AbbrCity(city string, lang language.Tag) (string, error) {
	resp, err := s.doRequestCity(city, lang)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body1, err := io.ReadAll(resp.Body)

	body := bufio.NewReader(bytes.NewReader(body1))
	skipPrefix(body)
	readLine(body)

	innerObject, err := decodeInnerObject(body)
	if err != nil {
		return "", fmt.Errorf("couldn't decode inner object: %w", err)
	}

	if len(innerObject) <= 0 ||
		len(innerObject[0]) <= 0 ||
		len(innerObject[0][0]) <= 0 ||
		len(innerObject[0][0][0]) <= 4 {
		return "", fmt.Errorf("wrong inner object: %v", innerObject)
	}
	resCities := innerObject[0][0]

	foundCity, ok := resCities[0][2].(string)
	if !ok {
		return "", fmt.Errorf("couldn't get city name")
	}
	serializedCity, ok := resCities[0][4].(string)
	if !ok {
		return "", fmt.Errorf("couldn't get serialized city name")
	}
	if city != foundCity {
		return "", fmt.Errorf("the requested city name didn't match the found. requested: %s found: %s", city, foundCity)
	}

	return serializedCity, nil
}

func (s *Session) AbbrCities(cities []string, lang language.Tag) ([]string, error) {
	abbrCities := []string{}
	for _, c := range cities {
		sc, err := s.AbbrCity(c, lang)
		if err != nil {
			return nil, err
		}
		abbrCities = append(abbrCities, sc)
	}
	return abbrCities, nil
}
