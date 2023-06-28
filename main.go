package main

import (
	"log"
	"time"

	"github.com/krisukox/google-flights-api/api"
	"golang.org/x/text/currency"
)

func main() {
	date, _ := time.Parse("2006-01-02", "2023-12-01")
	returnDate, _ := time.Parse("2006-01-02", "2023-12-08")
	// // url, err := api.CreateFullURL(departureDate, returnDate, "Wrocław", "Rzym")
	// // if err != nil {
	// // 	log.Fatalf(err.Error())
	// // }
	// // fmt.Println(url)

	// flights, err := api.GetFlights(departureDate, returnDate, "Wrocław", "Madryt", currency.PLN)
	// if err != nil {
	// 	log.Fatalf(err.Error())
	// }
	// for _, flight := range flights {
	// 	fmt.Printf("%+v\n", flight)
	// }

	// city, _ := api.GetSerializedCityName("Wrocław")
	// fmt.Println(city)
	// city2, _ := api.GetSerializedCityName("Madryt")
	// fmt.Println(city2)

	_, err := api.GetFlightsV2(date, returnDate, "Wrocław", "Madryt", currency.PLN)
	if err != nil {
		log.Fatalf(err.Error())
	}
}
