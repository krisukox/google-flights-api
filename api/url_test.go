package api

import (
	"testing"
	"time"
)

func TestCreateFullURL1(t *testing.T) {
	expectedUrl := "https://www.google.com/travel/flights/search?tfs=CBwQAhooEgoyMDIzLTA3LTExagwIAhIIL20vMDg0NWJyDAgDEggvbS8wNTZfeRooEgoyMDIzLTA3LTE3agwIAxIIL20vMDU2X3lyDAgCEggvbS8wODQ1YkABSAFwAYIBCwj___________8BmAEB"

	departureDate, err := time.Parse("2006-01-02", "2023-07-11")
	if err != nil {
		t.Fatalf("Error while creating departure date")
	}
	returnDate, err := time.Parse("2006-01-02", "2023-07-17")
	if err != nil {
		t.Fatalf("Error while creating return date")
	}
	url, err := CreateFullURL(departureDate, returnDate, "Wrocław", "Madryt")
	if err != nil {
		t.Fatalf(err.Error())
	}
	if url != expectedUrl {
		t.Fatalf("Created url is different than expected. Created: %s Expected: %s", url, expectedUrl)
	}
}

func TestCreateFullURL2(t *testing.T) {
	expectedUrl := "https://www.google.com/travel/flights/search?tfs=CBwQAhooEgoyMDIzLTA4LTIwagwIAhIIL20vMDV5d2dyDAgDEggvbS8wMTU2cRooEgoyMDIzLTA5LTAxagwIAxIIL20vMDE1NnFyDAgCEggvbS8wNXl3Z0ABSAFwAYIBCwj___________8BmAEB"

	departureDate, err := time.Parse("2006-01-02", "2023-08-20")
	if err != nil {
		t.Fatalf("Error while creating departure date")
	}
	returnDate, err := time.Parse("2006-01-02", "2023-09-01")
	if err != nil {
		t.Fatalf("Error while creating return date")
	}
	url, err := CreateFullURL(departureDate, returnDate, "Praga", "Berlin")
	if err != nil {
		t.Fatalf(err.Error())
	}
	if url != expectedUrl {
		t.Fatalf("Created url is different than expected. Created: %s Expected: %s", url, expectedUrl)
	}
}
