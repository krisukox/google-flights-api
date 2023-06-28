package main

import (
	"time"

	"github.com/krisukox/google-flights-api/api"
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

	api.GetRawData(date, returnDate, "Wrocław", "Rzym")
}
