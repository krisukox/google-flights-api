package flights

import (
	"encoding/base64"
	"time"
)

const (
	airportConst byte = 1
	cityConst    byte = 3
	dstConst     byte = 114
	srcConst     byte = 106
)

func serializeLocation(city string, locationNo byte) []byte {
	cityBytes := []byte(city)
	bytes := append([]byte{8, locationNo, 18}, append([]byte{byte(len(cityBytes))}, cityBytes...)...)
	return append([]byte{byte(len(bytes))}, bytes...)
}

func serializeSrcCity(city string) []byte {
	return append([]byte{srcConst}, serializeLocation(city, cityConst)...)
}

func serializeDstCity(city string) []byte {
	return append([]byte{dstConst}, serializeLocation(city, cityConst)...)
}

func serializeSrcAirport(airport string) []byte {
	return append([]byte{srcConst}, serializeLocation(airport, airportConst)...)
}

func serializeDstAirport(airport string) []byte {
	return append([]byte{dstConst}, serializeLocation(airport, airportConst)...)
}

func serializeLocations(locations []string, f func(string) []byte) []byte {
	ret := []byte{}
	for _, l := range locations {
		ret = append(ret, f(l)...)
	}
	return ret
}

func serializeDate(date time.Time) []byte {
	return append([]byte{18, 10}, []byte(date.Format("2006-01-02"))...)
}

func serializeStops(Stops Stops) []byte {
	switch Stops {
	case Nonstop:
		return []byte{40, 0}
	case Stop1:
		return []byte{40, 1}
	case Stop2:
		return []byte{40, 2}
	}
	return []byte{}
}

func serializeClass(Class Class) []byte {
	switch Class {
	case Economy:
		return []byte{72, 1}
	case PremiumEconomy:
		return []byte{72, 2}
	case Business:
		return []byte{72, 3}
	}
	return []byte{72, 4}
}

func serializeFlight(
	date time.Time,
	srcCities, srcAirports, DstCities, DstAirports []string,
	Stops Stops,
) []byte {
	bytes := serializeDate(date)
	bytes = append(bytes, serializeStops(Stops)...)
	bytes = append(bytes, serializeLocations(srcCities, serializeSrcCity)...)
	bytes = append(bytes, serializeLocations(srcAirports, serializeSrcAirport)...)
	bytes = append(bytes, serializeLocations(DstCities, serializeDstCity)...)
	bytes = append(bytes, serializeLocations(DstAirports, serializeDstAirport)...)
	return append([]byte{26, byte(len(bytes))}, bytes...)
}

func serializeAdults(Adults int) []byte {
	bytes := []byte{}
	for i := 0; i < Adults; i++ {
		bytes = append(bytes, 64, 1)
	}
	return bytes
}

func serializeTripType(TripType TripType) byte {
	if TripType == RoundTrip {
		return 1
	}
	return 2
}

// The function serializes arguments to the Google Flight URL. The city names should be provided in the
// language described by args.Lang. The args.Lang language is also used to set the language of the
// website to which the serialized URL leads.
//
// GetPriceGraph returns an error if any of the requests fail or if any of the city names are misspelled.
//
// Requirements are described by the [URLArgs.Validate] function.
func (s *Session) SerializeURL(args URLArgs) (string, error) {
	var err error

	if err = args.Validate(); err != nil {
		return "", err
	}

	args.SrcCities, err = s.abbrCities(args.SrcCities, args.Lang)
	if err != nil {
		return "", err
	}

	args.DstCities, err = s.abbrCities(args.DstCities, args.Lang)
	if err != nil {
		return "", err
	}

	bytes := []byte{8, 28, 16, 2}

	additionalBytes := []byte{
		112, 1,
		130, 1,
		11,
		8, 255, 255, 255, 255, 255, 255, 255, 255, 255, 1,
		152, 1,
	}

	bytes = append(
		bytes,
		serializeFlight(args.Date, args.SrcCities, args.SrcAirports, args.DstCities, args.DstAirports, args.Stops)...)

	if args.TripType == RoundTrip {
		bytes = append(
			bytes,
			serializeFlight(args.ReturnDate, args.DstCities, args.DstAirports, args.SrcCities, args.SrcAirports, args.Stops)...)
	}

	bytes = append(bytes, serializeAdults(args.Adults)...)

	bytes = append(bytes, serializeClass(args.Class)...)

	bytes = append(bytes, additionalBytes...)

	bytes = append(bytes, serializeTripType(args.TripType))

	RawURLEncoding := base64.URLEncoding.WithPadding(base64.NoPadding)

	url := "https://www.google.com/travel/flights/search?tfs=" + RawURLEncoding.EncodeToString(bytes) + "&curr=" + args.Currency.String() + "&hl=" + args.Lang.String()

	return url, nil
}
