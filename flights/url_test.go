package flights

import (
	"bytes"
	"testing"
	"time"

	"golang.org/x/text/currency"
	"golang.org/x/text/language"
)

func TestSerializeSrcCity(t *testing.T) {
	expectedSerializedCity1 := []byte{
		106, 12,
		8, 3,
		18, 8,
		47, 109, 47, 48, 56, 52, 53, 98}
	serializedCity1 := serializeSrcCity("/m/0845b")

	if !bytes.Equal(expectedSerializedCity1, serializedCity1) {
		t.Fatalf("wrong serialized source city, expected: %v serialized: %v", expectedSerializedCity1, serializedCity1)
	}

	expectedSerializedCity2 := []byte{
		106, 11,
		8, 3,
		18, 7,
		47, 109, 47, 48, 110, 50, 122}
	serializedCity2 := serializeSrcCity("/m/0n2z")

	if !bytes.Equal(expectedSerializedCity2, serializedCity2) {
		t.Fatalf("wrong serialized source city, expected: %v serialized: %v", expectedSerializedCity2, serializedCity2)
	}
}

func TestSerializeDstCity(t *testing.T) {
	expectedSerializedCity1 := []byte{
		114, 12,
		8, 3,
		18, 8,
		47, 109, 47, 48, 56, 52, 53, 98}
	serializedCity1 := serializeDstCity("/m/0845b")

	if !bytes.Equal(expectedSerializedCity1, serializedCity1) {
		t.Fatalf("wrong serialized city, expected: %v serialized: %v", expectedSerializedCity1, serializedCity1)
	}

	expectedSerializedSrcCity2 := []byte{
		114, 11,
		8, 3,
		18, 7,
		47, 109, 47, 48, 110, 50, 122}
	serializedCity2 := serializeDstCity("/m/0n2z")

	if !bytes.Equal(expectedSerializedSrcCity2, serializedCity2) {
		t.Fatalf("wrong serialized city, expected: %v serialized: %v", expectedSerializedSrcCity2, serializedCity2)
	}
}

func TestSerializeSrcAirport(t *testing.T) {
	expectedSerializedCity := []byte{
		106, 7,
		8, 1,
		18, 3,
		76, 67, 89}
	serializedCity := serializeSrcAirport("LCY")

	if !bytes.Equal(expectedSerializedCity, serializedCity) {
		t.Fatalf("wrong serialized city, expected: %v serialized: %v", serializedCity, serializedCity)
	}
}

func TestSerializeDstAirport(t *testing.T) {
	expectedSerializedCity := []byte{
		114, 7,
		8, 1,
		18, 3,
		76, 67, 89}
	serializedCity := serializeDstAirport("LCY")

	if !bytes.Equal(expectedSerializedCity, serializedCity) {
		t.Fatalf("wrong serialized city, expected: %v serialized: %v", serializedCity, serializedCity)
	}
}

func TestSerializeUrl(t *testing.T) {
	session := New()

	expectedUrl := "https://www.google.com/travel/flights/search?tfs=CBwQAho-EgoyMDIzLTExLTA2KAFqDggDEgovbS8wMzBxYjN0agcIARIDU0ZPcgwIAxIIL20vMDRqcGxyBwgBEgNDREcaPhIKMjAyMy0xMS0xMygBagwIAxIIL20vMDRqcGxqBwgBEgNDREdyDggDEgovbS8wMzBxYjN0cgcIARIDU0ZPQAFAAUgBcAGCAQsI____________AZgBAQ&curr=USD"

	date, _ := time.Parse("2006-01-02", "2023-11-06")
	returnDate, _ := time.Parse("2006-01-02", "2023-11-13")

	url, err := session.SerializeUrl(
		date,
		returnDate,
		[]string{"Los Angeles"},
		[]string{"SFO"},
		[]string{"London"},
		[]string{"CDG"},
		2,
		currency.USD,
		Stop1,
		Economy,
		RoundTrip,
		language.English,
	)

	if err != nil {
		t.Fatalf("error during serialization: %s", err.Error())
	}

	if expectedUrl != url {
		t.Fatalf("wrong serialized url, expected: %v serialized: %v", expectedUrl, url)
	}
}
