package main

import (
	"fmt"
	"log"
	"time"

	"github.com/krisukox/google-flights-api/api"
)

func main() {
	departureDate, _ := time.Parse("2006-01-02", "2023-07-11")
	returnDate, _ := time.Parse("2006-01-02", "2023-07-17")
	url, err := api.CreateFullURL(departureDate, returnDate, "Wrocław", "Madryt")
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Println(url)
}
