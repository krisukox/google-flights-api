package flights

import (
	"fmt"
	"time"
)

type Flight struct {
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
	Flight               []Flight
	ReturnFlight         []Flight // Not implemented yet
	SrcAirportCode       string
	DstAirportCode       string
	SrcCity              string
	DstCity              string
	FlightDuration       time.Duration
	ReturnFlightDuration time.Duration // Not implemented yet
}

type PriceRange struct {
	Low  float64
	High float64
}

type Stops int64

const (
	AnyStops Stops = iota
	Nonstop
	Stop1
	Stop2
)

type Class int64

const (
	Economy Class = iota
	PremiumEconomy
	Buisness
	First
)

type TripType int64

const (
	RoundTrip TripType = iota
	OneWay
)

func (f Flight) String() string {
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
	out += fmt.Sprintf("SrcCity: %s \n", t.SrcCity)
	out += fmt.Sprintf("DstCity: %s \n", t.DstCity)
	out += fmt.Sprintf("Duration: %s \n", t.FlightDuration)
	return out
}
