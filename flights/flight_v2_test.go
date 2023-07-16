package flights

import (
	"fmt"
	"testing"
	"time"

	"golang.org/x/text/currency"
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

	offersPLN, priceRange, err := session.GetOffers(departureDate, returnDate, "Wrocław", "Madryt", currency.PLN)
	if err != nil {
		t.Fatalf(err.Error())
	}

	if len(offersPLN) < 5 {
		t.Fatalf("received flights array contains less than 5 flights: %d", len(offersPLN))
	}
	fmt.Println(priceRange)

	// offersUSD, err := GetOffers(departureDate, returnDate, "Wrocław", "Madryt", currency.USD)
	// if err != nil {
	//  t.Fatalf(err.Error())
	// }

	// if len(offersPLN) != len(offersUSD) {
	//  t.Fatalf("received offers array has different length: %d %d", len(offersPLN), len(offersUSD))
	// }

	// for i := range offersPLN {
	//  if offersPLN[i].price < offersUSD[i].price || offersPLN[i].price < offersUSD[i].price*3 {
	//      t.Fatalf("wrong flight price: PLN %f USD %f", offersPLN[i].price, offersUSD[i].price)
	//  }
	// }
}
