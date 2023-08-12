package flights

import (
	"fmt"
	"net/url"
	"testing"
	"time"

	"github.com/go-test/deep"
	"golang.org/x/text/currency"
	"golang.org/x/text/language"
)

func TestGetPriceGraphReal(t *testing.T) {
	session, err := New()
	if err != nil {
		t.Fatal(err)
	}

	daysDiff1 := 60
	daysDiff2 := 90
	expectedOffers := daysDiff2 - daysDiff1 + 1

	offers, err := session.GetPriceGraph(
		PriceGraphArgs{
			time.Now().AddDate(0, 0, daysDiff1),
			time.Now().AddDate(0, 0, daysDiff2),
			7,
			[]string{"London"}, []string{}, []string{"Paris"}, []string{},
			Args{Travelers{Adults: 1}, currency.PLN, AnyStops, Economy, RoundTrip, language.English},
		})

	if err != nil {
		t.Fatalf(err.Error())
	}

	if len(offers) != expectedOffers {
		fmt.Println(offers)
		t.Fatalf("received offers array length is different than %d: %d", expectedOffers, len(offers))
	}
}

func TestGetPriceGraph(t *testing.T) {
	expectedPrices := []float64{
		922, 562, 648, 654, 660, 714, 891, 594, 594, 539, 648, 715, 654, 654, 594,
		594, 594, 680, 743, 654, 699, 654, 594, 594, 654, 654, 654, 806, 755, 617,
		747, 714, 680, 617, 654, 594, 539, 539, 508, 763, 739, 625, 508, 508, 508,
		508, 508, 739, 659, 508, 508, 508, 508, 508, 562, 628, 508, 508, 508}

	d := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	rd := time.Date(2024, 1, 8, 0, 0, 0, 0, time.UTC)

	expectedOffers := []Offer{}

	for _, p := range expectedPrices {
		expectedOffers = append(expectedOffers, Offer{d, rd, p})
		d = d.AddDate(0, 0, 1)
		rd = rd.AddDate(0, 0, 1)
	}

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
		PriceGraphArgs{
			time.Now().AddDate(0, 0, 2),
			time.Now().AddDate(0, 0, 5),
			0,
			[]string{"Athens"}, []string{}, []string{"Warsaw"}, []string{},
			Args{Travelers{}, currency.PLN, AnyStops, Economy, RoundTrip, language.English},
		},
	)
	if err != nil {
		t.Fatalf(err.Error())
	}

	sortSlice(offers, func(lv, rv Offer) bool {
		return lv.StartDate.Before(rv.StartDate)
	})

	if len(expectedPrices) != len(offers) {
		t.Fatalf("wrong offers length, expected: %d, received: %d", len(expectedPrices), len(offers))
	}
	if diff := deep.Equal(offers, expectedOffers); diff != nil {
		t.Fatal(diff)
	}
}

func TestPriceGraphReqData(t *testing.T) {
	session, err := New()
	if err != nil {
		t.Fatal(err)
	}

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
		PriceGraphArgs{
			date,
			returnDate,
			2,
			[]string{"Los Angeles"},
			[]string{"SFO"},
			[]string{"London"},
			[]string{"CDG"},
			Args{Travelers{Adults: 1}, currency.USD, AnyStops, Economy, RoundTrip, language.English},
		})
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
		PriceGraphArgs{
			date,
			returnDate,
			2,
			[]string{"Los Angeles"},
			[]string{"SFO"},
			[]string{"London"},
			[]string{"CDG"},
			Args{Travelers{Adults: 2}, currency.USD, Stop2, Business, OneWay, language.English},
		})
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
