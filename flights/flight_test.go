package flights

import (
	"testing"
	"time"

	"golang.org/x/text/currency"
	"golang.org/x/text/language"
)

func TestGetOffers(t *testing.T) {
	session := New()

	departureDate, err := time.Parse("2006-01-02", "2024-02-10")
	if err != nil {
		t.Fatalf("Error while creating departure date")
	}
	returnDate, err := time.Parse("2006-01-02", "2024-02-20")
	if err != nil {
		t.Fatalf("Error while creating return date")
	}

	offersPLN, _, err := session.GetOffers(departureDate, returnDate, "Berlin", "Madrid", currency.PLN, language.English)
	if err != nil {
		t.Fatalf(err.Error())
	}

	if len(offersPLN) < 5 {
		t.Fatalf("received flights array contains less than 5 flights: %d", len(offersPLN))
	}

	offersUSD, _, err := session.GetOffers(departureDate, returnDate, "Berlin", "Madrid", currency.USD, language.English)
	if err != nil {
		t.Fatalf(err.Error())
	}

	if len(offersPLN) != len(offersUSD) {
		t.Fatalf("received offers array has different length: %d %d", len(offersPLN), len(offersUSD))
	}

	for i := range offersPLN {
		if offersPLN[i].Price < offersUSD[i].Price || offersPLN[i].Price < offersUSD[i].Price*3 {
			t.Fatalf("wrong flight price: PLN %f USD %f", offersPLN[i].Price, offersUSD[i].Price)
		}
	}
}
