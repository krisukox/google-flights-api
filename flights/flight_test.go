package flights

import (
	"math"
	"net/url"
	"testing"
	"time"

	_ "time/tzdata"

	"github.com/go-test/deep"
	"golang.org/x/text/currency"
	"golang.org/x/text/language"
)

func min(x, y int) int {
	return int(math.Min(float64(x), float64(y)))
}

func TestGetOffersUSDPLN(t *testing.T) {
	session, err := New()
	if err != nil {
		t.Fatal(err)
	}

	date := time.Now().AddDate(0, 6, 0)
	returnDate := time.Now().AddDate(0, 7, 0)

	offersPLN, _, err := session.GetOffers(
		OffersArgs{
			date,
			returnDate,
			[]string{"Los Angeles"},
			[]string{"SFO"},
			[]string{"London"},
			[]string{"CDG"},
			Args{Travelers{Adults: 2}, currency.PLN, Stop1, PremiumEconomy, OneWay, language.English},
		},
	)

	if err != nil {
		t.Fatalf(err.Error())
	}

	offersUSD, _, err := session.GetOffers(
		OffersArgs{
			date,
			returnDate,
			[]string{"Los Angeles"},
			[]string{"SFO"},
			[]string{"London"},
			[]string{"CDG"},
			Args{Travelers{Adults: 2}, currency.USD, Stop1, PremiumEconomy, OneWay, language.English},
		},
	)
	if err != nil {
		t.Fatalf(err.Error())
	}

	elemsNumber := min(len(offersPLN), len(offersUSD))

	less := func(lv, rv FullOffer) bool {
		return lv.Price < rv.Price
	}

	sortSlice(offersPLN, less)
	sortSlice(offersUSD, less)

	for i := 0; i < elemsNumber; i++ {
		if offersPLN[i].Price < offersUSD[i].Price || offersPLN[i].Price < offersUSD[i].Price*3 {
			t.Fatalf("wrong flight price: PLN %f USD %f", offersPLN[i].Price, offersUSD[i].Price)
		}
	}
}

func removeUnknowns(offers []FullOffer) {
	for i := range offers {
		for j := range offers[i].Flight {
			offers[i].Flight[j].Unknown = nil
		}
	}
}

func TestGetOffers(t *testing.T) {
	dateTimeTimeZone := time.DateTime + " -0700 MST"

	dummyTime := time.Now()
	dummyValue := 0

	t1, _ := time.Parse(dateTimeTimeZone, "2024-01-22 17:00:00 +0100 CET")
	t2, _ := time.Parse(dateTimeTimeZone, "2024-01-22 18:35:00 +0100 CET")
	d1, _ := time.ParseDuration("1h35m0s")

	t3, _ := time.Parse(dateTimeTimeZone, "2024-01-22 21:25:00 +0100 CET")
	t4, _ := time.Parse(dateTimeTimeZone, "2024-01-23 00:50:00 +0200 EET")
	d2, _ := time.ParseDuration("2h25m0s")

	returnDate, _ := time.Parse(time.RFC3339, "2024-01-25T00:00:00Z")

	d3, _ := time.ParseDuration("6h50m0s")

	expectedOffer := FullOffer{
		Offer: Offer{
			StartDate:  t1,
			ReturnDate: returnDate,
			Price:      1315,
		},
		Flight: []Flight{{
			DepAirportCode: "WAW",
			DepAirportName: "Warsaw Chopin Airport",
			ArrAirportName: "Munich International Airport",
			ArrAirportCode: "MUC",
			DepTime:        t1,
			ArrTime:        t2,
			Duration:       d1,
			Airplane:       "Airbus A320neo",
			FlightNumber:   "LH 1615",
			AirlineName:    "Lufthansa",
			Legroom:        "29 inches",
		}, {
			DepAirportCode: "MUC",
			DepAirportName: "Munich International Airport",
			ArrAirportName: `Athens International Airport "Eleftherios Venizelos"`,
			ArrAirportCode: "ATH",
			DepTime:        t3,
			ArrTime:        t4,
			Duration:       d2,
			Airplane:       "Airbus A321neo",
			FlightNumber:   "LH 1756",
			AirlineName:    "Lufthansa",
			Legroom:        "29 inches",
		}},
		ReturnFlight:   []Flight{},
		SrcAirportCode: "WAW",
		DstAirportCode: "ATH",
		SrcCity:        "",
		DstCity:        "",
		FlightDuration: d3,
	}
	expectedPriceRange := PriceRange{1300, 2300}

	httpClientMock, err := newHttpClientMock(
		t,
		"testdata/city_warsaw.resp",
		"testdata/city_athens.resp",
		"testdata/flight.resp",
	)
	if err != nil {
		t.Fatal(err)
	}

	session := &Session{
		client: httpClientMock,
	}

	offers, priceRange, err := session.GetOffers(
		OffersArgs{
			dummyTime,
			returnDate,
			[]string{"Warsaw"},
			[]string{},
			[]string{"Athens"},
			[]string{},
			Args{
				Travelers{},
				currency.Unit{},
				Stops(dummyValue),
				Class(dummyValue),
				TripType(dummyValue),
				language.Tag{},
			},
		},
	)
	if err != nil {
		t.Fatal(err)
	}

	if len(offers) != 21 {
		t.Fatalf("Not all offers parsed. Expected number of offers: 21, parsed: %d", len(offers))
	}

	// Do not compare the unknown field
	removeUnknowns(offers)

	if diff := deep.Equal(expectedOffer, offers[0]); diff != nil {
		t.Fatalf("Offers are not equal: %v", diff)
	}

	if priceRange == nil {
		t.Fatalf("Missing price range")
	}

	if *priceRange != expectedPriceRange {
		t.Fatalf("Wrong price range, received: %v, expected: %v", priceRange, expectedPriceRange)
	}
}

func TestFlightReqData(t *testing.T) {
	session, err := New()
	if err != nil {
		t.Fatal(err)
	}

	expectedReqData1 := `[null,"[[],[null,null,1,null,[],1,[1,0,0,0],null,null,null,null,null,null,[[[[[\"SFO\",0],[\"/m/030qb3t\",5]]],[[[\"CDG\",0],[\"/m/04jpl\",5]]],null,0,[],[],\"2024-01-01\",null,[],[],[],null,null,[],3],[[[[\"CDG\",0],[\"/m/04jpl\",5]]],[[[\"SFO\",0],[\"/m/030qb3t\",5]]],null,0,[],[],\"2024-01-31\",null,[],[],[],null,null,[],3]],null,null,null,1,null,null,null,null,null,[]],1,0,0]"]`
	expectedReqData2 := `[null,"[[],[null,null,2,null,[],3,[2,0,0,0],null,null,null,null,null,null,[[[[[\"SFO\",0],[\"/m/030qb3t\",5]]],[[[\"CDG\",0],[\"/m/04jpl\",5]]],null,3,[],[],\"2024-01-01\",null,[],[],[],null,null,[],3]],null,null,null,1,null,null,null,null,null,[]],1,0,0]"]`

	date, err := time.Parse("2006-01-02", "2024-01-01")
	if err != nil {
		t.Fatalf("Error while creating date: %v", err)
	}
	returnDate, err := time.Parse("2006-01-02", "2024-01-31")
	if err != nil {
		t.Fatalf("Error while creating return date: %v", err)
	}

	_reqData1, err := session.getFlightReqData(
		OffersArgs{
			date,
			returnDate,
			[]string{"Los Angeles"},
			[]string{"SFO"},
			[]string{"London"},
			[]string{"CDG"},
			Args{Travelers{Adults: 1}, currency.Unit{}, AnyStops, Economy, RoundTrip, language.English},
		},
	)
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

	_reqData2, err := session.getFlightReqData(
		OffersArgs{
			date,
			returnDate,
			[]string{"Los Angeles"},
			[]string{"SFO"},
			[]string{"London"},
			[]string{"CDG"},
			Args{Travelers{Adults: 2}, currency.Unit{}, Stop2, Business, OneWay, language.English},
		},
	)
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
