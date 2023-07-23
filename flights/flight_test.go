package flights

import (
	"net/url"
	"testing"
	"time"

	"github.com/go-test/deep"
	"golang.org/x/text/currency"
	"golang.org/x/text/language"
)

func TestGetOffersUSDPLN(t *testing.T) {
	session := New()

	date := time.Now().AddDate(0, 6, 0)
	returnDate := time.Now().AddDate(0, 7, 0)

	offersPLN, _, err := session.GetOffers(
		date,
		returnDate,
		[]string{"Los Angeles"},
		[]string{"SFO"},
		[]string{"London"},
		[]string{"CDG"},
		2,
		currency.PLN,
		Stop1,
		PremiumEconomy,
		OneWay,
		language.English,
	)
	if err != nil {
		t.Fatalf(err.Error())
	}

	offersUSD, _, err := session.GetOffers(
		date,
		returnDate,
		[]string{"Los Angeles"},
		[]string{"SFO"},
		[]string{"London"},
		[]string{"CDG"},
		2,
		currency.USD,
		Stop1,
		PremiumEconomy,
		OneWay,
		language.English,
	)
	if err != nil {
		t.Fatalf(err.Error())
	}

	if len(offersPLN) != len(offersUSD) {
		t.Fatalf("received offers array has different length: %d %d", len(offersPLN), len(offersUSD))
	}

	less := func(lv, rv FullOffer) bool {
		return lv.Price < rv.Price
	}

	sortSlice(offersPLN, less)
	sortSlice(offersUSD, less)

	for i := range offersPLN {
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
	dummyTime := time.Now()
	dummyValue := 0

	t1, _ := time.Parse(time.RFC3339, "2024-01-22T17:00:00+01:00")
	t2, _ := time.Parse(time.RFC3339, "2024-01-22T18:35:00+01:00")
	d1, _ := time.ParseDuration("1h35m0s")

	t3, _ := time.Parse(time.RFC3339, "2024-01-22T21:25:00+01:00")
	t4, _ := time.Parse(time.RFC3339, "2024-01-23T00:50:00+01:00")
	d2, _ := time.ParseDuration("2h25m0s")

	d3, _ := time.ParseDuration("4h25m0s")

	t5, _ := time.Parse(time.RFC3339, "2024-01-22T17:00:00+01:00")

	expectedOffer := FullOffer{
		Offer: Offer{
			StartDate: t5,
			Price:     1315,
		},
		Flight: []flight{{
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
		ReturnFlight:    []flight{},
		SrcAirportCode:  "",
		DstAirportCode:  "",
		OriginCity:      "",
		DestinationCity: "",
		Duration:        d3,
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
		dummyTime,
		dummyTime,
		[]string{"Warsaw"},
		[]string{},
		[]string{"Athens"},
		[]string{},
		dummyValue,
		currency.Unit{},
		Stops(dummyValue),
		Class(dummyValue),
		TripType(dummyValue),
		language.Tag{},
	)
	if err != nil {
		t.Fatal(err)
	}

	// Do not compare the unknown field
	removeUnknowns(offers)

	if diff := deep.Equal(expectedOffer, offers[0]); diff != nil {
		t.Fatalf("Offers are not equal: %v", diff)
	}

	if priceRange != expectedPriceRange {
		t.Fatalf("Wrong price range, received: %v, expected: %v", priceRange, expectedPriceRange)
	}
}

func TestFlightRawData(t *testing.T) {
	session := New()

	expectedUnescapedQuery1 := `[null,"[[],[null,null,1,null,[],1,[1,0,0,0],null,null,null,null,null,null,[[[[[\"SFO\",0],[\"/m/030qb3t\",5]]],[[[\"CDG\",0],[\"/m/04jpl\",5]]],null,0,[],[],\"2024-01-01\",null,[],[],[],null,null,[],3],[[[[\"CDG\",0],[\"/m/04jpl\",5]]],[[[\"SFO\",0],[\"/m/030qb3t\",5]]],null,0,[],[],\"2024-01-31\",null,[],[],[],null,null,[],3]],null,null,null,1,null,null,null,null,null,[]],1,0,0]"]`
	expectedUnescapedQuery2 := `[null,"[[],[null,null,2,null,[],3,[2,0,0,0],null,null,null,null,null,null,[[[[[\"SFO\",0],[\"/m/030qb3t\",5]]],[[[\"CDG\",0],[\"/m/04jpl\",5]]],null,3,[],[],\"2024-01-01\",null,[],[],[],null,null,[],3]],null,null,null,1,null,null,null,null,null,[]],1,0,0]"]`

	date, err := time.Parse("2006-01-02", "2024-01-01")
	if err != nil {
		t.Fatalf("Error while creating date: %v", err)
	}
	returnDate, err := time.Parse("2006-01-02", "2024-01-31")
	if err != nil {
		t.Fatalf("Error while creating return date: %v", err)
	}

	rawData1, err := session.getFlightRawData(
		date,
		returnDate,
		[]string{"Los Angeles"},
		[]string{"SFO"},
		[]string{"London"},
		[]string{"CDG"},
		1,
		AnyStops,
		Economy,
		RoundTrip,
		language.English)
	if err != nil {
		t.Fatal(err)
	}

	unescapedQuery1, err := url.QueryUnescape(rawData1)
	if err != nil {
		t.Fatal(err)
	}

	if unescapedQuery1 != expectedUnescapedQuery1 {
		t.Fatalf("wrong unescaped query, expected: %s received: %s", expectedUnescapedQuery1, unescapedQuery1)
	}

	rawData2, err := session.getFlightRawData(
		date,
		returnDate,
		[]string{"Los Angeles"},
		[]string{"SFO"},
		[]string{"London"},
		[]string{"CDG"},
		2,
		Stop2,
		Buisness,
		OneWay,
		language.English)
	if err != nil {
		t.Fatal(err)
	}

	unescapedQuery2, err := url.QueryUnescape(rawData2)
	if err != nil {
		t.Fatal(err)
	}

	if unescapedQuery2 != expectedUnescapedQuery2 {
		t.Fatalf("wrong unescaped query, expected: %s received: %s", expectedUnescapedQuery2, unescapedQuery2)
	}
}
