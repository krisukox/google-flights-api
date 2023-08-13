package flights

import (
	"context"
	"testing"
	"time"

	"golang.org/x/text/currency"
	"golang.org/x/text/language"
)

func TestSerializeURL1(t *testing.T) {
	session, err := New()
	if err != nil {
		t.Fatal(err)
	}

	expectedURL := "https://www.google.com/travel/flights/search?tfs=Gj4SCjIwMjMtMTEtMDYoAWoOCAISCi9tLzAzMHFiM3RqBwgBEgNTRk9yDAgCEggvbS8wNGpwbHIHCAESA0NERxo-EgoyMDIzLTExLTEzKAFqDAgCEggvbS8wNGpwbGoHCAESA0NER3IOCAISCi9tLzAzMHFiM3RyBwgBEgNTRk9CAQFIAZgBAQ&curr=USD&hl=en"

	date, _ := time.Parse("2006-01-02", "2023-11-06")
	returnDate, _ := time.Parse("2006-01-02", "2023-11-13")

	url, err := session.SerializeURL(
		context.Background(),
		Args{
			date,
			returnDate,
			[]string{"Los Angeles"},
			[]string{"SFO"},
			[]string{"London"},
			[]string{"CDG"},
			Options{Travelers{Adults: 1}, currency.USD, Stop1, Economy, RoundTrip, language.English},
		},
	)

	if err != nil {
		t.Fatalf("error during serialization: %s", err.Error())
	}

	if expectedURL != url {
		t.Fatalf("wrong serialized url, expected: %v serialized: %v", expectedURL, url)
	}
}

func TestSerializeURL2(t *testing.T) {
	session, err := New()
	if err != nil {
		t.Fatal(err)
	}

	expectedURL := "https://www.google.com/travel/flights/search?tfs=GjMSCjIwMjMtMTItMTAoAmoMCAISCC9tLzA0anBsagcIARIDU0ZPcgwIAhIIL20vMGYydjAaMxIKMjAyMy0xMi0yMCgCagwIAhIIL20vMGYydjByDAgCEggvbS8wNGpwbHIHCAESA1NGT0IEAQECA0gBmAEB&curr=USD&hl=en"

	date, _ := time.Parse("2006-01-02", "2023-12-10")
	returnDate, _ := time.Parse("2006-01-02", "2023-12-20")

	url, err := session.SerializeURL(
		context.Background(),
		Args{
			date,
			returnDate,
			[]string{"London"},
			[]string{"SFO"},
			[]string{"Miami"},
			[]string{},
			Options{Travelers{Adults: 2, Children: 1, InfantOnLap: 1}, currency.USD, Stop2, Economy, RoundTrip, language.English},
		},
	)

	if err != nil {
		t.Fatalf("error during serialization: %s", err.Error())
	}

	if expectedURL != url {
		t.Fatalf("wrong serialized url, expected: %v serialized: %v", expectedURL, url)
	}
}
