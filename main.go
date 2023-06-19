package main

import (
	"fmt"
	"log"
	"time"

	"github.com/krisukox/google-flights-api/api"
)

func main() {
	// a := []byte{8, 28, 16, 2, 26, 40, 18, 10,
	// 	50, 48, 50, 51, 45, 48, 55, 45, 48, 51, // date
	// 	106, 12, 8, 2, 18, 8, // ??
	// 	47, 109, 47, 48, 56, 52, 53, 98, 114, 12, 8, 3, 18, 8, // departure
	// 	47, 109, 47, 48, 53, 121, 119, 103, 26, 40, 18, 10, // arrival
	// 	50, 48, 50, 51, 45, 48, 55, 45, 48, 55, // date
	// 	106, 12, 8, 3, 18, 8, // ??
	// 	47, 109, 47, 48, 53, 121, 119, 103, 114, 12, 8, 2, 18, 8, // departure
	// 	47, 109, 47, 48, 56, 52, 53, 98, 64, 1, 72, 1, 112, 1, 130, // arrival
	// 	1, 11, 8, 255, 255, 255, 255, 255, 255, 255, 255, 255, 1, 152, 1, 1, // ??
	// }
	departureDate, _ := time.Parse("2006-01-02", "2023-07-11")
	returnDate, _ := time.Parse("2006-01-02", "2023-07-17")
	// serializedDepartureCity, err := api.GetSerializedCityName("Wrocław")
	// if err != nil {
	// 	log.Fatalf(err.Error())
	// }
	// serializedArrivalCity, err := api.GetSerializedCityName("Madryt")
	// if err != nil {
	// 	log.Fatalf(err.Error())
	// }

	// fmt.Println(serializedDepartureCity)
	// fmt.Println(serializedArrivalCity)

	// fmt.Println("https://www.google.com/travel/flights/search?tfs=" + createURL(subSerialize(departureDate, returnDate, serializedDepartureCity, serializedArrivalCity)))
	url, err := api.CreateFullURL(departureDate, returnDate, "Wrocław", "Madryt")
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Println(url)
	// log.SetFlags(log.LstdFlags | log.Lshortfile)
	// api.GetSerializedCityName("Praga")
	// api.GetSerializedCityName("Wrocław")
	// body, err := io.ReadAll(resp.Body)
	// if err != nil {
	// 	log.Fatalf("Couldn't read body")
	// }
	// fmt.Print(body)
	// var j interface{}
	// err = json.NewDecoder(resp.Body).Decode(&j)
	// if err != nil {
	// 	log.Fatalf("Couldn't decode response")
	// }
	// fmt.Print(j)
}
