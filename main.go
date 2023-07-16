package main

import (
	"fmt"
	"time"

	"github.com/krisukox/google-flights-api/flights"
	"golang.org/x/text/currency"
)

func GetSpecialPrices(rangeStartDate, rangeEndDate time.Time, tripLength int, originCities, targetCities []string) {
	session := flights.New()

	offers, err := session.GetPriceGraph(rangeStartDate, rangeEndDate, tripLength, originCities[0], targetCities[0], currency.PLN)
	if err != nil {
		fmt.Println(err)
	}
	// fmt.Println(offers)

	for _, o := range offers {
		// priceRange := api.PriceRange{Start: 600, End: 1100}
		_, priceRange, err := session.GetOffers(o.StartDate, o.ReturnDate, originCities[0], targetCities[0], currency.PLN)
		// fmt.Println(priceRange)
		if err != nil {
			fmt.Println(err)
		}

		// fmt.Printf("price: %f, low price: %f\n", o.Price, priceRange.Start)

		if o.Price < priceRange.Low {
			fmt.Printf("%s %s\n", o.StartDate, o.ReturnDate)
			// fmt.Printf("priceRange %v\n", priceRange)
			fmt.Printf("price %f\n", o.Price)
			fmt.Println(flights.SerializeUrl(
				o.StartDate,
				o.ReturnDate,
				[]string{originCities[0]},
				[]string{},
				[]string{targetCities[0]},
				[]string{},
				1,
				currency.PLN,
				flights.AnyStops,
				flights.Economy,
				flights.RoundTrip,
			))
			// fmt.Printf("price %f\n", o.price)
		}
	}
}

func main() {
	rangeStartDate, _ := time.Parse("2006-01-02", "2023-10-01")
	rangeEndDate, _ := time.Parse("2006-01-02", "2023-10-30")

	// GetSpecialPrices(rangeStartDate, rangeEndDate, 2, []string{"Wrocław"}, []string{"Ateny"})

	session := flights.New()

	session.GetPriceGraph(rangeStartDate, rangeEndDate, 2, "Wrocław", "Ateny", currency.PLN)
}
