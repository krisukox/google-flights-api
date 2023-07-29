package flights

import (
	"encoding/base64"
	"time"

	"golang.org/x/text/currency"
	"golang.org/x/text/language"
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

func serializeStops(stops Stops) []byte {
	switch stops {
	case Nonstop:
		return []byte{40, 0}
	case Stop1:
		return []byte{40, 1}
	case Stop2:
		return []byte{40, 2}
	}
	return []byte{}
}

func serializeClass(class Class) []byte {
	switch class {
	case Economy:
		return []byte{72, 1}
	case PremiumEconomy:
		return []byte{72, 2}
	case Buisness:
		return []byte{72, 3}
	}
	return []byte{72, 4}
}

func serializeFlight(
	date time.Time,
	srcCities, srcAirports, dstCities, dstAirports []string,
	stops Stops,
) []byte {
	bytes := serializeDate(date)
	bytes = append(bytes, serializeStops(stops)...)
	bytes = append(bytes, serializeLocations(srcCities, serializeSrcCity)...)
	bytes = append(bytes, serializeLocations(srcAirports, serializeSrcAirport)...)
	bytes = append(bytes, serializeLocations(dstCities, serializeDstCity)...)
	bytes = append(bytes, serializeLocations(dstAirports, serializeDstAirport)...)
	return append([]byte{26, byte(len(bytes))}, bytes...)
}

func serializeAdults(adults int) []byte {
	bytes := []byte{}
	for i := 0; i < adults; i++ {
		bytes = append(bytes, 64, 1)
	}
	return bytes
}

func serializeTripType(tripType TripType) byte {
	if tripType == RoundTrip {
		return 1
	}
	return 2
}

func (s *Session) SerializeUrl(
	date, returnDate time.Time,
	srcCities, srcAirports, dstCities, dstAirports []string,
	adults int,
	curr currency.Unit,
	stops Stops,
	class Class,
	tripType TripType,
	lang language.Tag,
) (string, error) {
	if err := checkLocations(srcCities, srcAirports, dstCities, dstAirports); err != nil {
		return "", err
	}

	srcCities, err := s.AbbrCities(srcCities, lang)
	if err != nil {
		return "", err
	}

	dstCities, err = s.AbbrCities(dstCities, lang)
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

	bytes = append(bytes, serializeFlight(date, srcCities, srcAirports, dstCities, dstAirports, stops)...)

	if tripType == RoundTrip {
		bytes = append(bytes, serializeFlight(returnDate, dstCities, dstAirports, srcCities, srcAirports, stops)...)
	}

	bytes = append(bytes, serializeAdults(adults)...)

	bytes = append(bytes, serializeClass(class)...)

	bytes = append(bytes, additionalBytes...)

	bytes = append(bytes, serializeTripType(tripType))

	RawURLEncoding := base64.URLEncoding.WithPadding(base64.NoPadding)

	url := "https://www.google.com/travel/flights/search?tfs=" + RawURLEncoding.EncodeToString(bytes) + "&curr=" + curr.String()

	return url, nil
}
