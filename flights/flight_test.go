package flights

import (
	"context"
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
		context.Background(),
		Args{
			date,
			returnDate,
			[]string{"Los Angeles"},
			[]string{"SFO"},
			[]string{"London"},
			[]string{"CDG"},
			Options{Travelers{Adults: 2}, currency.PLN, Stop1, PremiumEconomy, OneWay, language.English},
		},
	)

	if err != nil {
		t.Fatal(err)
	}

	offersUSD, _, err := session.GetOffers(
		context.Background(),
		Args{
			date,
			returnDate,
			[]string{"Los Angeles"},
			[]string{"SFO"},
			[]string{"London"},
			[]string{"CDG"},
			Options{Travelers{Adults: 2}, currency.USD, Stop1, PremiumEconomy, OneWay, language.English},
		},
	)
	if err != nil {
		t.Fatal(err)
	}

	elemsNumber := min(len(offersPLN), len(offersUSD))

	if elemsNumber < 2 {
		t.Fatalf("not enough elements: %d", elemsNumber)
	}

	comparedOffers := 0

	for _, usd := range offersUSD {
		for _, pln := range offersPLN {
			if diff := deep.Equal(usd.Flight, pln.Flight); diff == nil {
				comparedOffers += 1
				if pln.Price < usd.Price || pln.Price < usd.Price*3 {
					t.Fatalf("wrong flight price: PLN %f USD %f", pln.Price, usd.Price)
				}
			}
		}
	}

	if comparedOffers < elemsNumber {
		t.Fatalf("not enought compared offers: expected %d compared %d", elemsNumber, comparedOffers)
	}
}

func compareWithThreshold(lv, rv float64) bool {
	return math.Abs(lv-rv) < lv*0.01
}

func testGetOffersTravelers(t *testing.T, session *Session, rootPrice float64, args Args, multiplier float64) {
	percentageDiff := 20.0

	offers, _, err := session.GetOffers(context.Background(), args)
	if err != nil {
		t.Fatal(err)
	}
	if len(offers) < 1 {
		t.Fatalf("not enough offers (%d) for the following Travelers: %+v", len(offers), args.Travelers)
	}
	lowerThreshold := rootPrice * multiplier * (100 - percentageDiff) / 100
	upperThreshold := rootPrice * multiplier * (100 + percentageDiff) / 100
	if offers[0].Price < lowerThreshold || offers[0].Price > upperThreshold {
		t.Errorf("The received price %f: should be between threshold: %f %f", offers[0].Price, lowerThreshold, upperThreshold)
	}
}

func TestGetOffersTravelers(t *testing.T) {
	session, err := New()
	if err != nil {
		t.Fatal(err)
	}

	date := time.Now().AddDate(0, 6, 0)
	returnDate := time.Now().AddDate(0, 7, 0)

	args := Args{
		date,
		returnDate,
		[]string{"Los Angeles"},
		[]string{"SFO"},
		[]string{"London"},
		[]string{"CDG"},
		Options{Travelers{Adults: 1}, currency.USD, Stop1, PremiumEconomy, OneWay, language.English},
	}

	offers, _, err := session.GetOffers(context.Background(), args)
	if err != nil {
		t.Fatal(err)
	}

	if len(offers) < 1 {
		t.Fatalf("not enough offers: %d", len(offers))
	}

	rootPrice := offers[0].Price

	args.Travelers = Travelers{Adults: 2}
	testGetOffersTravelers(t, session, rootPrice, args, 2)

	args.Travelers = Travelers{Adults: 2, Children: 1}
	testGetOffersTravelers(t, session, rootPrice, args, 3)

	args.Travelers = Travelers{Adults: 2, Children: 1, InfantInSeat: 1}
	testGetOffersTravelers(t, session, rootPrice, args, 4)
}

func removeUnknowns(offers []FullOffer) {
	for i := range offers {
		for j := range offers[i].Flight {
			offers[i].Flight[j].Unknown = nil
		}
	}
}

func TestGetOffersMock(t *testing.T) {
	dateTimeTimeZone := time.DateTime + " -0700 MST"

	dummyValue := 0

	t1, _ := time.Parse(dateTimeTimeZone, "2024-01-22 17:00:00 +0100 CET")
	t2, _ := time.Parse(dateTimeTimeZone, "2024-01-22 18:35:00 +0100 CET")
	d1, _ := time.ParseDuration("1h35m0s")

	t3, _ := time.Parse(dateTimeTimeZone, "2024-01-22 21:25:00 +0100 CET")
	t4, _ := time.Parse(dateTimeTimeZone, "2024-01-23 00:50:00 +0200 EET")
	d2, _ := time.ParseDuration("2h25m0s")

	// override time.Now() for testing
	timeNow = func() time.Time {
		t, _ := time.Parse(time.RFC3339, "2024-01-15T00:00:00Z")
		return t
	}
	defer func() { timeNow = time.Now }()

	date, _ := time.Parse(time.RFC3339, "2024-01-20T00:00:00Z")
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
			DepCity:        "Warsaw",
			ArrAirportName: "Munich International Airport",
			ArrAirportCode: "MUC",
			ArrCity:        "Munich",
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
			DepCity:        "Munich",
			ArrAirportName: `Athens International Airport "Eleftherios Venizelos"`,
			ArrAirportCode: "ATH",
			ArrCity:        "Athens",
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
		SrcCity:        "Warsaw",
		DstCity:        "Athens",
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
		context.Background(),
		Args{
			date,
			returnDate,
			[]string{"Warsaw"},
			[]string{},
			[]string{"Athens"},
			[]string{},
			Options{
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
		context.Background(),
		Args{
			date,
			returnDate,
			[]string{"Los Angeles"},
			[]string{"SFO"},
			[]string{"London"},
			[]string{"CDG"},
			Options{Travelers{Adults: 1}, currency.Unit{}, AnyStops, Economy, RoundTrip, language.English},
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
		context.Background(),
		Args{
			date,
			returnDate,
			[]string{"Los Angeles"},
			[]string{"SFO"},
			[]string{"London"},
			[]string{"CDG"},
			Options{Travelers{Adults: 2}, currency.Unit{}, Stop2, Business, OneWay, language.English},
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
