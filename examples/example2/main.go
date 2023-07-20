package main

import (
	"fmt"
	"log"
	"time"

	"github.com/krisukox/google-flights-api/flights"
	"golang.org/x/text/currency"
	"golang.org/x/text/language"
)

func getBestOffer(rangeStartDate, rangeEndDate time.Time, tripLength int, srcCities, dstCities string, lang language.Tag) {
	session := flights.New()
	var bestOffer flights.Offer

	offers, err := session.GetPriceGraph(rangeStartDate, rangeEndDate, tripLength, srcCities, dstCities, currency.PLN, lang)
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
	url, err := session.SerializeUrl(
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
		lang,
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(url)
}

func main() {
	getBestOffer(
		time.Now().AddDate(0, 0, 60),
		time.Now().AddDate(0, 0, 90),
		2,
		"Warsaw",
		"Athens",
		language.English)
}
