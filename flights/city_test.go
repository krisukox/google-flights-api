package flights

import (
	"testing"

	"golang.org/x/text/language"
)

const (
	abbrA = "/m/0n2z"
	abbrB = "/m/081m_"
)

func TestAbbrCityEN(t *testing.T) {
	session := New()

	cityA := "Athens"
	cityB := "Warsaw"

	a, err := session.AbbrCity(cityA, language.English)
	if err != nil {
		t.Fatal(err)
	}
	if a != abbrA {
		t.Fatalf("wrong abbreviated city name, expected: %s received: %s", abbrA, a)
	}

	w, err := session.AbbrCity(cityB, language.English)
	if err != nil {
		t.Fatal(err)
	}
	if w != abbrB {
		t.Fatalf("wrong abbreviated city name, expected: %s received: %s", abbrB, w)
	}
}

func TestAbbrCityDE(t *testing.T) {
	session := New()

	cityA := "Athen"
	cityB := "Warschau"

	a, err := session.AbbrCity(cityA, language.German)
	if err != nil {
		t.Fatal(err)
	}
	if a != abbrA {
		t.Fatalf("wrong abbreviated city name, expected: %s received: %s", abbrA, a)
	}

	w, err := session.AbbrCity(cityB, language.German)
	if err != nil {
		t.Fatal(err)
	}
	if w != abbrB {
		t.Fatalf("wrong abbreviated city name, expected: %s received: %s", abbrB, w)
	}
}

func TestAbbrCityPL(t *testing.T) {
	session := New()

	cityA := "Ateny"
	cityB := "Warszawa"

	a, err := session.AbbrCity(cityA, language.Polish)
	if err != nil {
		t.Fatal(err)
	}
	if a != abbrA {
		t.Fatalf("wrong abbreviated city name, expected: %s received: %s", abbrA, a)
	}

	b, err := session.AbbrCity(cityB, language.Polish)
	if err != nil {
		t.Fatal(err)
	}
	if b != abbrB {
		t.Fatalf("wrong abbreviated city name, expected: %s received: %s", abbrB, b)
	}
}

func TestAbbrCity(t *testing.T) {
	httpClientMock, err := newHttpClientMock(
		t,
		"testdata/city_athens.resp",
		"testdata/city_warsaw.resp",
	)

	session := &Session{
		client: httpClientMock,
	}

	cityA := "Athens"
	cityB := "Warsaw"

	a, err := session.AbbrCity(cityA, language.English)
	if err != nil {
		t.Fatal(err)
	}
	if a != abbrA {
		t.Fatalf("wrong abbreviated city name, expected: %s received: %s", abbrA, a)
	}

	b, err := session.AbbrCity(cityB, language.English)
	if err != nil {
		t.Fatal(err)
	}
	if b != abbrB {
		t.Fatalf("wrong abbreviated city name, expected: %s received: %s", abbrB, b)
	}
}
