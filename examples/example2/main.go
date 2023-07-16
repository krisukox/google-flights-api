package main

import (
	"fmt"
	"log"
	"time"

	"github.com/krisukox/google-flights-api/flights"
	"golang.org/x/text/currency"
)

func getBestOffer(rangeStartDate, rangeEndDate time.Time, tripLength int, srcCities, dstCities string) {
	session := flights.New()
	var bestOffer flights.Offer

	offers, err := session.GetPriceGraph(rangeStartDate, rangeEndDate, tripLength, srcCities, dstCities, currency.PLN)
	if err != nil {
		log.Fatal(err)
	}

	for _, o := range offers {
		if bestOffer.Price == 0 || o.Price < bestOffer.Price {
			bestOffer = o
		}
	}

	fmt.Printf("%s %s\n", bestOffer.StartDate, bestOffer.ReturnDate)
	fmt.Printf("price %d\n", int(bestOffer.Price))
	url, err := flights.SerializeUrl(
		bestOffer.StartDate,
		bestOffer.ReturnDate,
		[]string{srcCities},
		[]string{},
		[]string{dstCities},
		[]string{},
		1,
		currency.PLN,
		flights.AnyStops,
		flights.Economy,
		flights.RoundTrip,
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(url)
}

func main() {
	rangeStartDate, _ := time.Parse("2006-01-02", "2023-10-01")
	rangeEndDate, _ := time.Parse("2006-01-02", "2023-10-30")

	getBestOffer(rangeStartDate, rangeEndDate, 2, "Wrocław", "Ateny")
}
