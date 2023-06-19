package api

import (
	"testing"
	"time"
)

func TestCreateFullURL(t *testing.T) {
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
