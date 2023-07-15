package main

import (
	"fmt"
	"log"
	"time"

	"golang.org/x/text/currency"
)

func main() {
	date, _ := time.Parse("2006-01-02", "2023-07-04")
	returnDate, _ := time.Parse("2006-01-02", "2023-07-08")

	offer, err := GetOffers(date, returnDate, "Wrocław", "Rzym", currency.PLN)
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Println(offer)
}
