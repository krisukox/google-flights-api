package main

import (
	"fmt"
	"log"
	"time"

	"github.com/krisukox/google-flights-api/flights"
	"golang.org/x/text/currency"
)

func getCheapOffers(rangeStartDate, rangeEndDate time.Time, tripLength int, srcCities, dstCities []string) {
	session := flights.New()

	for _, s := range srcCities {
		for _, d := range dstCities {
			offers, err := session.GetPriceGraph(rangeStartDate, rangeEndDate, tripLength, s, d, currency.PLN)
			if err != nil {
				log.Fatal(err)
			}

			for _, o := range offers {
				_, priceRange, err := session.GetOffers(o.StartDate, o.ReturnDate, s, d, currency.PLN)
				if err != nil {
					log.Fatal(err)
				}

				if o.Price < priceRange.Low {
					fmt.Printf("%s %s\n", o.StartDate, o.ReturnDate)
					fmt.Printf("price %d\n", int(o.Price))
					url, err := flights.SerializeUrl(
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
	rangeStartDate, _ := time.Parse("2006-01-02", "2023-10-01")
	rangeEndDate, _ := time.Parse("2006-01-02", "2023-10-30")

	getCheapOffers(rangeStartDate, rangeEndDate, 2, []string{"Wrocław", "Katowice"}, []string{"Ateny"})
}
