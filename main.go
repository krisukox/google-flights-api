package main

import (
	"fmt"
	"log"
	"time"

	"github.com/krisukox/google-flights-api/api"
	"golang.org/x/text/currency"
)

func main() {
	date, _ := time.Parse("2006-01-02", "2023-11-01")
	returnDate, _ := time.Parse("2006-01-02", "2023-11-08")

	trips, err := api.GetFlightsV2(date, returnDate, "Wrocław", "Rzym", currency.PLN)
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Println(trips)
}
