package flights

import (
	"fmt"
	"time"
)

type flight struct {
	DepAirportCode string
	DepAirportName string
	ArrAirportName string
	ArrAirportCode string
	DepTime        time.Time
	ArrTime        time.Time
	Duration       time.Duration
	Airplane       string
	FlightNumber   string
	Unknown        []interface{}
	AirlineName    string
	Legroom        string
}

type Offer struct {
	StartDate  time.Time
	ReturnDate time.Time
	Price      float64
}

type FullOffer struct {
	Offer
	Flight          []flight
	ReturnFlight    []flight // Not implemented yet
	SrcAirportCode  string
	DstAirportCode  string
	OriginCity      string
	DestinationCity string
	Duration        time.Duration
}

type PriceRange struct {
	Low  float64
	High float64
}

func (f flight) String() string {
	out := ""
	out += fmt.Sprintf("DepAirportCode: %s ", f.DepAirportCode)
	out += fmt.Sprintf("DepAirportName: %s ", f.DepAirportName)
	out += fmt.Sprintf("ArrAirportName: %s ", f.ArrAirportName)
	out += fmt.Sprintf("ArrAirportCode: %s ", f.ArrAirportCode)
	out += fmt.Sprintf("DepTime: %s ", f.DepTime)
	out += fmt.Sprintf("ArrTime: %s ", f.ArrTime)
	out += fmt.Sprintf("Duration: %s ", f.Duration)
	out += fmt.Sprintf("Airplane: %s ", f.Airplane)
	out += fmt.Sprintf("FlightNumber: %s ", f.FlightNumber)
	// out += fmt.Sprintf("unknown: %v ", f.unknown)
	out += fmt.Sprintf("AirlineName: %s ", f.AirlineName)
	out += fmt.Sprintf("Legroom: %s ", f.Legroom)
	return out
}

func (t FullOffer) String() string {
	out := ""
	out += fmt.Sprintf("StartDate: %s \n", t.StartDate)
	out += fmt.Sprintf("ReturnDate: %s \n", t.ReturnDate)
	out += fmt.Sprintf("Price: %f \n", t.Price)
	out += fmt.Sprintf("Flight: %s \n", t.Flight)
	out += fmt.Sprintf("ReturnFlight: %s \n", t.ReturnFlight)
	out += fmt.Sprintf("SrcAirportCode: %s \n", t.SrcAirportCode)
	out += fmt.Sprintf("DstAirportCode: %s \n", t.DstAirportCode)
	out += fmt.Sprintf("OriginCity: %s \n", t.OriginCity)
	out += fmt.Sprintf("DestinationCity: %s \n", t.DestinationCity)
	out += fmt.Sprintf("Duration: %s \n", t.Duration)
	return out
}
