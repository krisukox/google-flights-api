package main

import (
	"fmt"
	"log"
	"time"

	"github.com/krisukox/google-flights-api/flights"
	"golang.org/x/text/currency"
	"golang.org/x/text/language"
)

func getCheapOffers(rangeStartDate, rangeEndDate time.Time, tripLength int, srcCities, dstCities []string, lang language.Tag) {
	session := flights.New()

	for _, s := range srcCities {
		for _, d := range dstCities {
			offers, err := session.GetPriceGraph(
				rangeStartDate, rangeEndDate,
				[]string{s}, []string{}, []string{d}, []string{},
				1, currency.PLN, flights.AnyStops, flights.Economy, flights.RoundTrip,
				lang, tripLength)
			if err != nil {
				log.Fatal(err)
			}

			for _, o := range offers {
				_, priceRange, err := session.GetOffers(
					o.StartDate,
					o.ReturnDate,
					[]string{s},
					[]string{},
					[]string{d},
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

				if o.Price < priceRange.Low {
					fmt.Printf("%s %s\n", o.StartDate, o.ReturnDate)
					fmt.Printf("price %d\n", int(o.Price))
					url, err := session.SerializeUrl(
						o.StartDate,
						o.ReturnDate,
						[]string{s},
						[]string{},
						[]string{d},
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
			}
		}
	}
}

func main() {
	getCheapOffers(
		time.Now().AddDate(0, 0, 60),
		time.Now().AddDate(0, 0, 90),
		2,
		[]string{"Berlin", "Prague"},
		[]string{"Athens"},
		language.English)
}
