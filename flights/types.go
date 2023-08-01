package flights

import (
	"fmt"
	"time"
	"unicode"

	"golang.org/x/text/currency"
	"golang.org/x/text/language"
)

// Flight describes single one-way flight.
//
// NOTE: currently, departure and arrival times are always provided with the UTC time zone instead of the
// correct time zone for the airport.
type Flight struct {
	DepAirportCode string        // departure airport code
	DepAirportName string        // departure airport name
	ArrAirportName string        // arrival airport name
	ArrAirportCode string        // arrival airport code
	DepTime        time.Time     // departure time
	ArrTime        time.Time     // arrival time
	Duration       time.Duration // duration of the flight
	Airplane       string        // airplane
	FlightNumber   string        // flight number
	Unknown        []interface{} // it contains all unknown data which are parsed from the Google Flight API
	AirlineName    string        // airline name
	Legroom        string        // legroom in the airplane seats
}

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
	// out += fmt.Sprintf("Unknown: %v ", f.Unknown)
	out += fmt.Sprintf("AirlineName: %s ", f.AirlineName)
	out += fmt.Sprintf("Legroom: %s", f.Legroom)
	return out
}

// It describes offer of a trip. [Session.GetPriceGraph] returns slice of Offers.
// [Session.GetOffers] returns slice of [FullOffer] which contain Offer.
type Offer struct {
	StartDate  time.Time // start date of the offer
	ReturnDate time.Time // return date of the offer
	Price      float64   // price of the offer
}

func (o Offer) String() string {
	out := ""
	out += fmt.Sprintf("{%s ", o.StartDate.Format(time.DateOnly))
	out += fmt.Sprintf("%s ", o.ReturnDate.Format(time.DateOnly))
	out += fmt.Sprintf("%d}", int(o.Price))
	return out
}

// It describes full offer of a trip. [Session.GetOffers] returns slice of FullOffers.
type FullOffer struct {
	Offer
	Flight               []Flight      // contains all flights in the trip
	ReturnFlight         []Flight      // not implemented yet
	SrcAirportCode       string        // code of the airport where the trip starts
	DstAirportCode       string        // destination airport
	SrcCity              string        // not implemented yet
	DstCity              string        // not implemented yet
	FlightDuration       time.Duration // duration of whole Flight
	ReturnFlightDuration time.Duration // not implemented yet
}

func (o FullOffer) String() string {
	out := ""
	out += fmt.Sprintf("{StartDate: %s\n", o.StartDate)
	out += fmt.Sprintf("ReturnDate: %s\n", o.ReturnDate)
	out += fmt.Sprintf("Price: %d\n", int(o.Price))
	out += fmt.Sprintf("Flight: %s\n", o.Flight)
	// out += fmt.Sprintf("ReturnFlight: %s\n", o.ReturnFlight)
	out += fmt.Sprintf("SrcAirportCode: %s\n", o.SrcAirportCode)
	out += fmt.Sprintf("DstAirportCode: %s\n", o.DstAirportCode)
	// out += fmt.Sprintf("SrcCity: %s\n", o.SrcCity)
	// out += fmt.Sprintf("DstCity: %s\n", o.DstCity)
	out += fmt.Sprintf("FlightDuration: %s}\n", o.FlightDuration)
	// out += fmt.Sprintf("ReturnFlightDuration: %s}\n", o.ReturnFlightDuration)
	return out
}

type PriceRange struct {
	Low  float64
	High float64
}

// Stops specifies how many stops the trip should contains.
type Stops int64

const (
	AnyStops Stops = iota // any number of stops
	Nonstop               // nonstop only
	Stop1                 // 1 stop or fewer
	Stop2                 // 2 stops or fewer
)

// Class describes the class of a flight.
type Class int64

const (
	Economy Class = iota
	PremiumEconomy
	Buisness
	First
)

// TripType specifies whether trip is round trip or one way.
type TripType int64

const (
	RoundTrip TripType = iota
	OneWay
)

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

// Args contains common arguments used in [OffersArgs], [PriceGraphArgs] and [URLArgs].
type Args struct {
	Adults   int // number of adults
	Currency currency.Unit
	Stops    Stops
	Class    Class
	TripType TripType
	Lang     language.Tag
}

func ArgsDefault() Args {
	return Args{
		Adults:   1,
		Currency: currency.USD,
		Stops:    AnyStops,
		Class:    Economy,
		TripType: RoundTrip,
		Lang:     language.English,
	}
}

// Arguments used in the [Session.GetPriceGraph].
type PriceGraphArgs struct {
	RangeStartDate, RangeEndDate                   time.Time // days range of the price graph
	TripLength                                     int       // number of days between start trip date and return date
	SrcCities, SrcAirports, DstCities, DstAirports []string  // source and destination; cities and airports of the trip
	Args
}

// Validates PriceGraphArgs requirements:
//   - at least one source location (srcCities / srcAirports)
//   - at least one destination location (dstCities / dstAirports)
//   - srcAirports and dstAirports have to be in the right IATA format https://en.wikipedia.org/wiki/IATA_airport_code
//   - dates have to be in the chronological order - today's date -> RangeStartDate -> RangeEndDate
//   - the difference between RangeStartDate and RangeEndDate cannot be higher than 161 days
func (a *PriceGraphArgs) Validate() error {
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

// Arguments used in the [Session.GetOffers]
type OffersArgs struct {
	Date, ReturnDate                               time.Time // start trip date and return date
	SrcCities, SrcAirports, DstCities, DstAirports []string  // source and destination; cities and airports of the trip
	Args
}

// Validates OffersArgs requirements:
//   - at least one source location (srcCities / srcAirports)
//   - at least one destination location (dstCities / dstAirports)
//   - srcAirports and dstAirports have to be in the right IATA format https://en.wikipedia.org/wiki/IATA_airport_code
//   - dates have to be in the chronological order - today's date -> Date -> ReturnDate
func (a *OffersArgs) Validate() error {
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

// Arguments used in the [Session.SerializeURL]
type URLArgs struct {
	Date, ReturnDate                               time.Time // start trip date and retrun date
	SrcCities, SrcAirports, DstCities, DstAirports []string  // source and destination; cities and airports of the trip
	Args
}

// Validates URLArgs requirements:
//   - at least one source location (srcCities / srcAirports)
//   - at least one destination location (dstCities / dstAirports)
//   - srcAirports and dstAirports have to be in the right IATA format https://en.wikipedia.org/wiki/IATA_airport_code
func (a *URLArgs) Validate() error {
	return validateLocations(a.SrcCities, a.SrcAirports, a.DstCities, a.DstAirports)
}
