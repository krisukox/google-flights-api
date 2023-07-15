package main

import (
	"testing"
	"time"

	"golang.org/x/text/currency"
)

func TestGetFlightsPLN(t *testing.T) {
	departureDate, err := time.Parse("2006-01-02", "2023-10-01")
	if err != nil {
		t.Fatalf("Error while creating departure date")
	}
	returnDate, err := time.Parse("2006-01-02", "2023-10-08")
	if err != nil {
		t.Fatalf("Error while creating return date")
	}

	flights, err := GetFlights(departureDate, returnDate, "Wrocław", "Madryt", currency.PLN)
	if err != nil {
		t.Fatalf(err.Error())
	}

	if len(flights) < 5 {
		t.Fatalf("received flights array contains less than 5 flights: %d", len(flights))
	}
}

func TestGetFlightsUSD(t *testing.T) {
	departureDate, err := time.Parse("2006-01-02", "2023-10-01")
	if err != nil {
		t.Fatalf("Error while creating departure date")
	}
	returnDate, err := time.Parse("2006-01-02", "2023-10-08")
	if err != nil {
		t.Fatalf("Error while creating return date")
	}

	flights, err := GetFlights(departureDate, returnDate, "Wrocław", "Madryt", currency.USD)
	if err != nil {
		t.Fatalf(err.Error())
	}

	if len(flights) < 5 {
		t.Fatalf("received flights array contains less than 5 flights: %d", len(flights))
	}
}
