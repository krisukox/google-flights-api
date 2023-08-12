// This example iterates over PriceGraph offers concurrently and prints only
// those whose price is cheaper than the low price of the offer.
// This example is the same as Example 2, but it sends requests concurrently.
package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/krisukox/google-flights-api/flights"
	"golang.org/x/text/currency"
	"golang.org/x/text/language"
)

func getCheapOffersConcurrent(
	session *flights.Session,
	rangeStartDate, rangeEndDate time.Time,
	tripLength int,
	srcCities, dstCities []string,
	lang language.Tag,
) {
	logger := log.New(os.Stdout, "", 0)

	args := flights.Args{
		Travelers: flights.Travelers{Adults: 1},
		Currency:  currency.USD,
		Stops:     flights.AnyStops,
		Class:     flights.Economy,
		TripType:  flights.RoundTrip,
		Lang:      lang,
	}

	priceGraphOffers, err := session.GetPriceGraph(
		context.Background(),
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
		logger.Fatal(err)
	}

	var wg sync.WaitGroup
	wg.Add(len(priceGraphOffers))

	for _, priceGraphOffer := range priceGraphOffers {
		go func(offer flights.Offer) {
			defer wg.Done()
			offers, _, err := session.GetOffers(
				context.Background(),
				flights.OffersArgs{
					Date:       offer.StartDate,
					ReturnDate: offer.ReturnDate,
					SrcCities:  srcCities,
					DstCities:  dstCities,
					Args:       args,
				},
			)
			if err != nil {
				logger.Fatal(err)
			}

			var bestOffer flights.FullOffer
			for _, o := range offers {
				if o.Price != 0 && (bestOffer.Price == 0 || o.Price < bestOffer.Price) {
					bestOffer = o
				}
			}

			_, priceRange, err := session.GetOffers(
				context.Background(),
				flights.OffersArgs{
					Date:        bestOffer.StartDate,
					ReturnDate:  bestOffer.ReturnDate,
					SrcAirports: []string{bestOffer.SrcAirportCode},
					DstAirports: []string{bestOffer.DstAirportCode},
					Args:        args,
				},
			)
			if err != nil {
				logger.Fatal(err)
			}
			if priceRange == nil {
				logger.Fatal("missing priceRange")
			}

			if bestOffer.Price < priceRange.Low {
				url, err := session.SerializeURL(
					context.Background(),
					flights.URLArgs{
						Date:        bestOffer.StartDate,
						ReturnDate:  bestOffer.ReturnDate,
						SrcAirports: []string{bestOffer.SrcAirportCode},
						DstAirports: []string{bestOffer.DstAirportCode},
						Args:        args,
					},
				)
				if err != nil {
					logger.Fatal(err)
				}
				logger.Printf("%s %s\n", bestOffer.StartDate, bestOffer.ReturnDate)
				logger.Printf("price %d\n", int(bestOffer.Price))
				logger.Println(url)
			}
		}(priceGraphOffer)
	}
	wg.Wait()
}

func main() {
	t := time.Now()

	session, err := flights.New()
	if err != nil {
		log.Fatal(err)
	}

	getCheapOffersConcurrent(
		session,
		time.Now().AddDate(0, 0, 60),
		time.Now().AddDate(0, 0, 90),
		7,
		[]string{"San Francisco", "San Jose"},
		[]string{"New York", "Philadelphia", "Washington"},
		language.English,
	)

	fmt.Println(time.Since(t))
}
