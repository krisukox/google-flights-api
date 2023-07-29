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

	args := flights.Args{
		Adults:   1,
		Curr:     currency.PLN,
		Stops:    flights.AnyStops,
		Class:    flights.Economy,
		TripType: flights.RoundTrip,
		Lang:     lang,
	}

	priceGraphOffers, err := session.GetPriceGraph(
		flights.PriceGraphArgs{
			RangeStartDate: rangeStartDate,
			RangeEndDate:   rangeEndDate,
			TripLength:     tripLength,
			SrcCities:      srcCities,
			DstCities:      dstCities,
			Args:           args,
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	for _, priceGraphOffer := range priceGraphOffers {
		offers, _, err := session.GetOffers(
			flights.OffersArgs{
				Date:       priceGraphOffer.StartDate,
				ReturnDate: priceGraphOffer.ReturnDate,
				SrcCities:  srcCities,
				DstCities:  dstCities,
				Args:       args,
			},
		)
		if err != nil {
			log.Fatal(err)
		}

		var bestOffer flights.FullOffer

		for _, o := range offers {
			if bestOffer.Price == 0 || o.Price < bestOffer.Price {
				bestOffer = o
			}
		}

		_, priceRange, err := session.GetOffers(
			flights.OffersArgs{
				Date:        bestOffer.StartDate,
				ReturnDate:  bestOffer.ReturnDate,
				SrcAirports: []string{bestOffer.SrcAirportCode},
				DstAirports: []string{bestOffer.DstAirportCode},
				Args:        args,
			},
		)
		if err != nil {
			log.Fatal(err)
		}
		if priceRange == nil {
			log.Fatal("missing priceRange")
		}

		if bestOffer.Price < priceRange.Low {
			url, err := session.SerializeUrl(
				flights.UrlArgs{
					Date:        bestOffer.StartDate,
					ReturnDate:  bestOffer.ReturnDate,
					SrcAirports: []string{bestOffer.SrcAirportCode},
					DstAirports: []string{bestOffer.DstAirportCode},
					Args:        args,
				},
			)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("%s %s\n", bestOffer.StartDate, bestOffer.ReturnDate)
			fmt.Printf("price %d\n", int(bestOffer.Price))
			fmt.Println(url)
		}
	}
}

func main() {
	t := time.Now()

	getCheapOffers(
		time.Now().AddDate(0, 0, 60),
		time.Now().AddDate(0, 0, 90),
		2,
		[]string{"Berlin", "Prague"},
		[]string{"Athens", "Rome"},
		language.English,
	)

	fmt.Println(time.Since(t))
}
