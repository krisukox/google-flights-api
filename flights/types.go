package flights

import (
	"fmt"
	"time"
)

type flight struct {
	Departure string
	Arrival   string
	Price     string
}

type flightV2 struct {
	DepartureAirportCode string
	DepartureAirportName string
	ArrivalAirportName   string
	ArrivalAirportCode   string
	DepartureTime        time.Time
	ArrivalTime          time.Time
	Duration             time.Duration
	Airplane             string
	FlightNumber         string
	Unknown              []interface{}
	AirlineName          string
	Legroom              string
}

type Offer struct {
	StartDate  time.Time
	ReturnDate time.Time
	Price      float64
}

type FullOffer struct {
	Offer
	Flight                 []flightV2
	ReturnFlight           []flightV2 // Not implemented yet
	OriginAirportCode      string
	DestinationAirportCode string
	OriginCity             string
	DestinationCity        string
	Duration               time.Duration
}

type PriceRange struct {
	Low  float64
	High float64
}

func (f flightV2) String() string {
	out := ""
	out += fmt.Sprintf("departureAirportCode: %s ", f.DepartureAirportCode)
	out += fmt.Sprintf("departureAirportName: %s ", f.DepartureAirportName)
	out += fmt.Sprintf("arrivalAirportName: %s ", f.ArrivalAirportName)
	out += fmt.Sprintf("arrivalAirportCode: %s ", f.ArrivalAirportCode)
	out += fmt.Sprintf("departureTime: %s ", f.DepartureTime)
	out += fmt.Sprintf("arrivalTime: %s ", f.ArrivalTime)
	out += fmt.Sprintf("duration: %s ", f.Duration)
	out += fmt.Sprintf("airplane: %s ", f.Airplane)
	out += fmt.Sprintf("flightNumber: %s ", f.FlightNumber)
	// out += fmt.Sprintf("unknown: %v ", f.unknown)
	out += fmt.Sprintf("airlineName: %s ", f.AirlineName)
	out += fmt.Sprintf("legroom: %s ", f.Legroom)
	return out
}

func (t FullOffer) String() string {
	out := ""
	out += fmt.Sprintf("flight: %s \n", t.Flight)
	out += fmt.Sprintf("returnFlight: %s \n", t.ReturnFlight)
	out += fmt.Sprintf("originAirportCode: %s \n", t.OriginAirportCode)
	out += fmt.Sprintf("destinationAirportCode: %s \n", t.DestinationAirportCode)
	out += fmt.Sprintf("startDate: %s \n", t.StartDate)
	out += fmt.Sprintf("returnDate: %s \n", t.ReturnDate)
	out += fmt.Sprintf("duration: %s \n", t.Duration)
	out += fmt.Sprintf("price: %f \n", t.Price)
	return out
}
