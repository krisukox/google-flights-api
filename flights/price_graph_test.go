package flights

import (
	"fmt"
	"net/url"
	"testing"
	"time"

	"golang.org/x/text/currency"
	"golang.org/x/text/language"
)

func TestGetPriceGraphReal(t *testing.T) {
	session := New()

	daysDiff1 := 60
	daysDiff2 := 90
	expectedOffers := daysDiff2 - daysDiff1 + 1

	offers, err := session.GetPriceGraph(
		time.Now().AddDate(0, 0, daysDiff1), time.Now().AddDate(0, 0, daysDiff2),
		[]string{"London"}, []string{}, []string{"Paris"}, []string{},
		1, currency.PLN, AnyStops, Economy, RoundTrip,
		language.English, 7)

	if err != nil {
		t.Fatalf(err.Error())
	}

	if len(offers) != expectedOffers {
		fmt.Println(offers)
		t.Fatalf("received offers array length is different than %d: %d", expectedOffers, len(offers))
	}
}

func TestGetPriceGraph(t *testing.T) {
	expected_prices := []float64{
		922, 562, 648, 654, 660, 714, 891, 594, 594, 539, 648, 715, 654, 654, 594,
		594, 594, 680, 743, 654, 699, 654, 594, 594, 654, 654, 654, 806, 755, 617,
		747, 714, 680, 617, 654, 594, 539, 539, 508, 628, 508, 508, 763, 625, 508,
		659, 739, 508, 508, 508, 508, 508, 562, 508, 508, 508, 508, 739, 508}

	httpClientMock, err := newHttpClientMock(
		t,
		"testdata/city_athens.resp",
		"testdata/city_warsaw.resp",
		"testdata/price_graph.resp",
	)

	session := &Session{
		client: httpClientMock,
	}

	offers, err := session.GetPriceGraph(
		time.Now().AddDate(0, 0, 2),
		time.Now().AddDate(0, 0, 5),
		[]string{"Athens"}, []string{}, []string{"Warsaw"}, []string{},
		0, currency.PLN, AnyStops, Economy, RoundTrip,
		language.English, 0)
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
		start, end,
		[]string{}, []string{}, []string{}, []string{},
		0, currency.PLN, AnyStops, Economy, RoundTrip, language.English,
		0)

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

func TestPriceGraphReqData(t *testing.T) {
	session := New()

	expectedReqData1 := `[null,"[null,[null,null,1,null,[],1,[1,0,0,0],null,null,null,null,null,null,[[[[[\"SFO\",0],[\"/m/030qb3t\",5]]],[[[\"CDG\",0],[\"/m/04jpl\",5]]],null,0,[],[],\"2024-01-01\",null,[],[],[],null,null,[],3],[[[[\"CDG\",0],[\"/m/04jpl\",5]]],[[[\"SFO\",0],[\"/m/030qb3t\",5]]],null,0,[],[],\"2024-01-03\",null,[],[],[],null,null,[],3]],null,null,null,1,null,null,null,null,null,[]],[\"2024-01-01\",\"2024-01-31\"],null,[2,2]]"]`
	expectedReqData2 := `[null,"[null,[null,null,2,null,[],3,[2,0,0,0],null,null,null,null,null,null,[[[[[\"SFO\",0],[\"/m/030qb3t\",5]]],[[[\"CDG\",0],[\"/m/04jpl\",5]]],null,3,[],[],\"2024-01-01\",null,[],[],[],null,null,[],3]],null,null,null,1,null,null,null,null,null,[]],[\"2024-01-01\",\"2024-01-31\"],null,[2,2]]"]`

	date, err := time.Parse("2006-01-02", "2024-01-01")
	if err != nil {
		t.Fatalf("Error while creating date: %v", err)
	}
	returnDate, err := time.Parse("2006-01-02", "2024-01-31")
	if err != nil {
		t.Fatalf("Error while creating return date: %v", err)
	}

	_reqData1, err := session.getPriceGraphReqData(
		date, returnDate,
		[]string{"Los Angeles"},
		[]string{"SFO"},
		[]string{"London"},
		[]string{"CDG"},
		1, AnyStops, Economy, RoundTrip,
		language.English, 2)
	if err != nil {
		t.Fatal(err)
	}

	reqData1, err := url.QueryUnescape(_reqData1)
	if err != nil {
		t.Fatal(err)
	}

	if reqData1 != expectedReqData1 {
		t.Fatalf("wrong unescaped query, expected: %s received: %s", expectedReqData1, reqData1)
	}

	_reqData2, err := session.getPriceGraphReqData(
		date,
		returnDate,
		[]string{"Los Angeles"},
		[]string{"SFO"},
		[]string{"London"},
		[]string{"CDG"},
		2, Stop2, Buisness, OneWay,
		language.English, 2)
	if err != nil {
		t.Fatal(err)
	}

	reqData2, err := url.QueryUnescape(_reqData2)
	if err != nil {
		t.Fatal(err)
	}

	if reqData2 != expectedReqData2 {
		t.Fatalf("wrong unescaped query, expected: %s received: %s", expectedReqData2, reqData2)
	}
}
