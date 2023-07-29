package flights

import (
	"fmt"
	"time"
	"unicode"

	"golang.org/x/text/currency"
	"golang.org/x/text/language"
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
	out += fmt.Sprintf("FlightDuration: %s \n", t.FlightDuration)
	out += fmt.Sprintf("ReturnFlightDuration: %s \n", t.ReturnFlightDuration)
	return out
}

func truncateToDay(date time.Time) time.Time {
	return date.Truncate(24 * time.Hour)
}

func validateNumberOfLocations(cities, airports []string) error {
	length := len(cities) + len(airports)
	if length < 1 {
		return fmt.Errorf("number of locations should be at least 1, specified: %d", length)
	}

	return nil
}

func validateDate(date, returnDate time.Time) error {
	now := time.Now().Truncate(time.Hour * 24)

	if returnDate.Before(date) {
		return fmt.Errorf("returnDate is before date")
	}
	if date.Before(now) {
		return fmt.Errorf("date is before today's date")
	}
	return nil
}

func validateRangeDate(rangeStartDate time.Time, rangeEndDate time.Time) error {
	now := time.Now().Truncate(time.Hour * 24)

	days := int(rangeEndDate.Sub(rangeStartDate).Hours() / 24)
	if days > 161 {
		return fmt.Errorf("number of days between dates is larger than 161, %d", days)
	}
	if rangeEndDate.Equal(rangeStartDate) {
		return fmt.Errorf("rangeEndDate is the same as rangeStartDate")
	}
	if rangeEndDate.Before(rangeStartDate) {
		return fmt.Errorf("rangeEndDate is before rangeStartDate")
	}
	if rangeStartDate.Before(now) {
		return fmt.Errorf("rangeStartDate is before today's date")
	}
	return nil
}

func isAirportCode(code string) bool {
	if len(code) != 3 {
		return false
	}
	for _, r := range code {
		if !unicode.IsUpper(r) {
			return false
		}
	}
	return true
}

func validateLocations(srcCities, srcAirports, dstCities, dstAirports []string) error {
	if err := validateNumberOfLocations(srcCities, srcAirports); err != nil {
		return fmt.Errorf("src locations: %s", err)
	}
	if err := validateNumberOfLocations(dstCities, dstAirports); err != nil {
		return fmt.Errorf("dst locations: %s", err)
	}

	for _, s := range srcAirports {
		if !isAirportCode(s) {
			return fmt.Errorf("src airport '%s' is not an airport code", s)
		}
	}
	for _, d := range dstAirports {
		if !isAirportCode(d) {
			return fmt.Errorf("dst airport '%s' is not an airport code", d)
		}
	}

	return nil
}

type Args struct {
	Adults   int
	Curr     currency.Unit
	Stops    Stops
	Class    Class
	TripType TripType
	Lang     language.Tag
}

type PriceGraphArgs struct {
	RangeStartDate, RangeEndDate                   time.Time
	TripLength                                     int
	SrcCities, SrcAirports, DstCities, DstAirports []string
	Args
}

func (a *PriceGraphArgs) validate() error {
	if err := validateLocations(a.SrcCities, a.SrcAirports, a.DstCities, a.DstAirports); err != nil {
		return err
	}

	a.RangeStartDate = truncateToDay(a.RangeStartDate)
	a.RangeEndDate = truncateToDay(a.RangeEndDate)

	if err := validateRangeDate(a.RangeStartDate, a.RangeEndDate); err != nil {
		return err
	}

	return nil
}

type _args struct {
	Date, ReturnDate                               time.Time
	SrcCities, SrcAirports, DstCities, DstAirports []string
	Args
}

type OffersArgs _args
type UrlArgs _args

func (a *OffersArgs) validate() error {
	if err := validateLocations(a.SrcCities, a.SrcAirports, a.DstCities, a.DstAirports); err != nil {
		return err
	}

	a.Date = truncateToDay(a.Date)
	a.ReturnDate = truncateToDay(a.ReturnDate)

	if err := validateDate(a.Date, a.ReturnDate); err != nil {
		return err
	}

	return nil
}

func (a *UrlArgs) validate() error {
	return validateLocations(a.SrcCities, a.SrcAirports, a.DstCities, a.DstAirports)
}
