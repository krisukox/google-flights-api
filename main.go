package main

import (
	"fmt"
	"log"
	"time"

	"github.com/krisukox/google-flights-api/api"
	"golang.org/x/text/currency"
)

func main() {
	departureDate, _ := time.Parse("2006-01-02", "2023-08-12")
	returnDate, _ := time.Parse("2006-01-02", "2023-08-19")
	// url, err := api.CreateFullURL(departureDate, returnDate, "Wrocław", "Rzym")
	// if err != nil {
	// 	log.Fatalf(err.Error())
	// }
	// fmt.Println(url)

	flights, err := api.GetFlights(departureDate, returnDate, "Wrocław", "Rzym", currency.PLN)
	if err != nil {
		log.Fatalf(err.Error())
	}
	for _, flight := range flights {
		fmt.Printf("%+v\n", flight)
	}

}
