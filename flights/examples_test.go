package flights_test

import (
	"fmt"
	"log"
	"time"

	"github.com/krisukox/google-flights-api/flights"
	"golang.org/x/text/currency"
	"golang.org/x/text/language"
)

func ExampleSession_GetPriceGraph() {
	session := flights.New()

	offers, err := session.GetPriceGraph(
		flights.PriceGraphArgs{
			RangeStartDate: time.Now().AddDate(0, 0, 30),
			RangeEndDate:   time.Now().AddDate(0, 0, 60),
			TripLength:     7,
			SrcCities:      []string{"San Francisco"},
			DstCities:      []string{"New York"},
			Args:           flights.ArgsDefault(),
		},
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(offers)
}

func ExampleSession_SerializeURL() {
	session := flights.New()

	url, err := session.SerializeURL(
		flights.URLArgs{
			Date:        time.Now().AddDate(0, 0, 30),
			ReturnDate:  time.Now().AddDate(0, 0, 37),
			SrcCities:   []string{"San Diego"},
			SrcAirports: []string{"LAX"},
			DstCities:   []string{"New York", "Philadelphia"},
			Args:        flights.ArgsDefault(),
		},
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(url)
}

func ExampleSession_GetOffers() {
	session := flights.New()

	offers, priceRange, err := session.GetOffers(
		flights.OffersArgs{
			Date:       time.Now().AddDate(0, 0, 30),
			ReturnDate: time.Now().AddDate(0, 0, 37),
			SrcCities:  []string{"Madrid"},
			DstCities:  []string{"Estocolmo"},
			Args: flights.Args{
				Adults:   2,
				Currency: currency.EUR,
				Stops:    flights.Stop1,
				Class:    flights.Economy,
				TripType: flights.RoundTrip,
				Lang:     language.Spanish,
			},
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	if priceRange != nil {
		fmt.Printf("High price %d\n", int(priceRange.High))
		fmt.Printf("Low price %d\n", int(priceRange.Low))
	}
	fmt.Println(offers)
}

func ExampleSession_AbbrCity() {
	session := flights.New()

	city, err := session.AbbrCity("New York", language.English)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(city)
}
