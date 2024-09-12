// This example gets PriceGraph offers and prints only those whose
// price is cheaper than the low price of the offer
// (the price is considered as low by the Google Flights)
package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/krisukox/google-flights-api/flights"
	"github.com/scylladb/termtables"
	"golang.org/x/text/currency"
	"golang.org/x/text/language"
)

func main() {
	session, err := flights.New()
	if err != nil {
		log.Fatal(err)
	}

	date := time.Now().AddDate(0, 0, 60)
	returnDate := time.Now().AddDate(0, 0, 90)
	currency := currency.USD

	args := flights.Args{
		Date:        date,
		ReturnDate:  returnDate,
		SrcCities:   []string{},
		SrcAirports: []string{"XMN"},
		DstCities:   []string{},
		DstAirports: []string{"NRT", "HND"},
		Options: flights.Options{
			Travelers: flights.Travelers{
				Adults: 1,
			},
			Currency: currency,
			Stops:    flights.AnyStops,
			Class:    flights.Economy,
			TripType: flights.RoundTrip,
			Lang:     language.Chinese,
		},
	}

	offers, _, err := session.GetOffers(context.Background(), args)
	for _, offer := range offers {
		fmt.Printf("Price: %0.2f %s\n", offer.Price, currency)
		t := termtables.CreateTable()
		t.AddHeaders("DepTime", "DepCity", "DepAirport", "Flight", "Duration", "ArrAirport", "ArrCity", "ArrTime")
		t.AddRow(
			offer.Flight[0].DepTime, offer.SrcCity, offer.SrcAirportCode,
			"", offer.FlightDuration.String(),
			offer.DstAirportCode, offer.DstCity, offer.Flight[len(offer.Flight)-1].ArrTime,
		)
		for _, flight := range offer.Flight {
			t.AddRow(
				flight.DepTime, flight.DepCity, fmt.Sprintf("%s(%s)", flight.DepAirportName, flight.DepAirportCode),
				fmt.Sprintf("%s - %s(%s)", flight.AirlineName, flight.FlightNumber, flight.Airplane), flight.Duration.String(),
				fmt.Sprintf("%s(%s)", flight.ArrAirportName, flight.ArrAirportCode), flight.ArrCity, flight.ArrTime)
		}
		t.AddRow("", "", "", "", "", "", "", "")
		t.AddRow(
			offer.ReturnFlight[0].DepTime, offer.DstCity, offer.DstAirportCode,
			"", offer.ReturnFlightDuration.String(),
			offer.SrcAirportCode, offer.SrcCity, offer.ReturnFlight[len(offer.ReturnFlight)-1].ArrTime,
		)
		for _, returnFlight := range offer.ReturnFlight {
			t.AddRow(
				returnFlight.DepTime, returnFlight.DepCity, fmt.Sprintf("%s(%s)", returnFlight.DepAirportName, returnFlight.DepAirportCode),
				fmt.Sprintf("%s - %s(%s)", returnFlight.AirlineName, returnFlight.FlightNumber, returnFlight.Airplane), returnFlight.Duration.String(),
				fmt.Sprintf("%s(%s)", returnFlight.ArrAirportName, returnFlight.ArrAirportCode), returnFlight.ArrCity, returnFlight.ArrTime,
			)
		}
		fmt.Println(t.Render())
	}
}
