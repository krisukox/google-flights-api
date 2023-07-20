package flights

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"testing"
	"time"

	"golang.org/x/text/currency"
	"golang.org/x/text/language"
)

func TestGetPriceGraphReal(t *testing.T) {
	session := New()

	date, err := time.Parse("2006-01-02", "2024-01-01")
	if err != nil {
		t.Fatalf("Error while creating departure date")
	}
	returnDate, err := time.Parse("2006-01-02", "2024-01-31")
	if err != nil {
		t.Fatalf("Error while creating return date")
	}

	offers, err := session.GetPriceGraph(date, returnDate, 7, "Berlin", "Rome", currency.PLN, language.English)
	if err != nil {
		t.Fatalf(err.Error())
	}

	if len(offers) != 31 {
		fmt.Println(offers)
		t.Fatalf("received offers array length is different than 31: %d", len(offers))
	}
}

func TestGetPriceGraph(t *testing.T) {
	expected_prices := []float64{
		922, 562, 648, 654, 660, 714, 891, 594, 594, 539, 648, 715, 654, 654, 594,
		594, 594, 680, 743, 654, 699, 654, 594, 594, 654, 654, 654, 806, 755, 617,
		747, 714, 680, 617, 654, 594, 539, 539, 508, 628, 508, 508, 763, 625, 508,
		659, 739, 508, 508, 508, 508, 508, 562, 508, 508, 508, 508, 739, 508}

	respFile, err := os.Open("testdata/city_athens.resp")
	if err != nil {
		t.Fatalf(err.Error())
	}

	byteValue, err := ioutil.ReadAll(respFile)
	if err != nil {
		t.Fatalf(err.Error())
	}

	bodyCityAthens := io.NopCloser(bytes.NewReader(byteValue))

	respFile, err = os.Open("testdata/city_warsaw.resp")
	if err != nil {
		t.Fatalf(err.Error())
	}

	byteValue, err = ioutil.ReadAll(respFile)
	if err != nil {
		t.Fatalf(err.Error())
	}

	bodyCityWarsaw := io.NopCloser(bytes.NewReader(byteValue))

	respFile, err = os.Open("testdata/price_graph.resp")
	if err != nil {
		t.Fatalf(err.Error())
	}

	byteValue, err = ioutil.ReadAll(respFile)
	if err != nil {
		t.Fatalf(err.Error())
	}

	bodyPriceGraph := io.NopCloser(bytes.NewReader(byteValue))

	session := &Session{
		&HttpClientMock{
			[]func() (*http.Response, error){func() (*http.Response, error) {
				return &http.Response{
					Body: bodyCityAthens,
				}, nil
			}, func() (*http.Response, error) {
				return &http.Response{
					Body: bodyCityWarsaw,
				}, nil
			}, func() (*http.Response, error) {
				return &http.Response{
					Body: bodyPriceGraph,
				}, nil
			}},
			t,
		},
	}

	offers, err := session.GetPriceGraph(time.Now().AddDate(0, 0, 2), time.Now().AddDate(0, 0, 5), 0, "Athens", "Warsaw", currency.PLN, language.English)
	if err != nil {
		t.Fatalf(err.Error())
	}

	if len(expected_prices) != len(offers) {
		t.Fatalf("wrong offers length, expected: %d, received: %d", len(expected_prices), len(offers))
	}
	for i := range expected_prices {
		if offers[i].Price != expected_prices[i] {
			t.Fatalf("wrong offer price, expected: %f, received: %f", expected_prices[i], offers[i].Price)
		}
	}
}

func _testGetPriceGraphDateLimit(t *testing.T, session *Session, start time.Time, end time.Time, errorValue string) {
	_, err := session.GetPriceGraph(
		start,
		end,
		0,
		"",
		"",
		currency.PLN,
		language.English)

	if err == nil {
		t.Fatalf("GetPriceGraph call for the following dates start %s end %s, should result in error", start, end)
	} else if err.Error() != errorValue {
		t.Fatalf(`Wrong error "%s" for GetPriceGraph call with the following dates start %s end %s`, err.Error(), start, end)
	}
}

func TestGetPriceGraphDateLimit(t *testing.T) {
	session := New()

	_testGetPriceGraphDateLimit(t, session, time.Now().AddDate(0, 0, 5), time.Now().AddDate(0, 0, 170), "number of days between dates is larger than 161, 165")
	_testGetPriceGraphDateLimit(t, session, time.Now().AddDate(0, 0, 5), time.Now().AddDate(0, 0, 2), "rangeEndDate is before rangeStartDate")
	_testGetPriceGraphDateLimit(t, session, time.Now().AddDate(0, 0, -1), time.Now().AddDate(0, 0, 2), "rangeStartDate is before today's date")
}
