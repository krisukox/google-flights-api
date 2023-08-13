package flights_test

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/krisukox/google-flights-api/flights"
	"golang.org/x/text/currency"
	"golang.org/x/text/language"
)

func ExampleSession_GetPriceGraph() {
	session, err := flights.New()
	if err != nil {
		log.Fatal(err)
	}

	offers, err := session.GetPriceGraph(
		context.Background(),
		flights.PriceGraphArgs{
			RangeStartDate: time.Now().AddDate(0, 0, 30),
			RangeEndDate:   time.Now().AddDate(0, 0, 60),
			TripLength:     7,
			SrcCities:      []string{"San Francisco"},
			DstCities:      []string{"New York"},
			Options:        flights.OptionsDefault(),
		},
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(offers)
}

func ExampleSession_SerializeURL() {
	session, err := flights.New()
	if err != nil {
		log.Fatal(err)
	}

	url, err := session.SerializeURL(
		context.Background(),
		flights.Args{
			Date:        time.Now().AddDate(0, 0, 30),
			ReturnDate:  time.Now().AddDate(0, 0, 37),
			SrcCities:   []string{"San Diego"},
			SrcAirports: []string{"LAX"},
			DstCities:   []string{"New York", "Philadelphia"},
			Options:     flights.OptionsDefault(),
		},
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(url)
}

func ExampleSession_GetOffers() {
	session, err := flights.New()
	if err != nil {
		log.Fatal(err)
	}

	offers, priceRange, err := session.GetOffers(
		context.Background(),
		flights.Args{
			Date:       time.Now().AddDate(0, 0, 30),
			ReturnDate: time.Now().AddDate(0, 0, 37),
			SrcCities:  []string{"Madrid"},
			DstCities:  []string{"Estocolmo"},
			Options: flights.Options{
				Travelers: flights.Travelers{Adults: 2},
				Currency:  currency.EUR,
				Stops:     flights.Stop1,
				Class:     flights.Economy,
				TripType:  flights.RoundTrip,
				Lang:      language.Spanish,
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
	session, err := flights.New()
	if err != nil {
		log.Fatal(err)
	}

	city, err := session.AbbrCity(context.Background(), "New York", language.English)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(city)
}
