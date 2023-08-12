package flights

import (
	"context"
	"testing"

	"golang.org/x/text/language"
)

const (
	abbrA = "/m/0n2z"
	abbrB = "/m/081m_"
)

func TestAbbrCityEN(t *testing.T) {
	session, err := New()
	if err != nil {
		t.Fatal(err)
	}

	cityA := "Athens"
	cityB := "Warsaw"

	a, err := session.AbbrCity(context.Background(), cityA, language.English)
	if err != nil {
		t.Fatal(err)
	}
	if a != abbrA {
		t.Fatalf("wrong abbreviated city name, expected: %s received: %s", abbrA, a)
	}

	w, err := session.AbbrCity(context.Background(), cityB, language.English)
	if err != nil {
		t.Fatal(err)
	}
	if w != abbrB {
		t.Fatalf("wrong abbreviated city name, expected: %s received: %s", abbrB, w)
	}
}

func TestAbbrCityDE(t *testing.T) {
	session, err := New()
	if err != nil {
		t.Fatal(err)
	}

	cityA := "Athen"
	cityB := "Warschau"

	a, err := session.AbbrCity(context.Background(), cityA, language.German)
	if err != nil {
		t.Fatal(err)
	}
	if a != abbrA {
		t.Fatalf("wrong abbreviated city name, expected: %s received: %s", abbrA, a)
	}

	w, err := session.AbbrCity(context.Background(), cityB, language.German)
	if err != nil {
		t.Fatal(err)
	}
	if w != abbrB {
		t.Fatalf("wrong abbreviated city name, expected: %s received: %s", abbrB, w)
	}
}

func TestAbbrCityPL(t *testing.T) {
	session, err := New()
	if err != nil {
		t.Fatal(err)
	}

	cityA := "Ateny"
	cityB := "Warszawa"

	a, err := session.AbbrCity(context.Background(), cityA, language.Polish)
	if err != nil {
		t.Fatal(err)
	}
	if a != abbrA {
		t.Fatalf("wrong abbreviated city name, expected: %s received: %s", abbrA, a)
	}

	b, err := session.AbbrCity(context.Background(), cityB, language.Polish)
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
	if err != nil {
		t.Fatal(err)
	}

	session := &Session{
		client: httpClientMock,
	}

	cityA := "Athens"
	cityB := "Warsaw"

	// Since httpClientMock has only two responses, run it twice to check whether
	// the cache for abbreviated city names works properly
	for i := 0; i < 2; i++ {
		a, err := session.AbbrCity(context.Background(), cityA, language.English)
		if err != nil {
			t.Fatal(err)
		}
		if a != abbrA {
			t.Fatalf("wrong abbreviated city name, expected: %s received: %s", abbrA, a)
		}

		b, err := session.AbbrCity(context.Background(), cityB, language.English)
		if err != nil {
			t.Fatal(err)
		}
		if b != abbrB {
			t.Fatalf("wrong abbreviated city name, expected: %s received: %s", abbrB, b)
		}
	}

	if abbrCity, ok := session.Cities.Load("Athens"); !ok || abbrCity != abbrA {
		t.Fatalf("Athens abbreviated city name not stored in cache")
	}

	if abbrCity, ok := session.Cities.Load("Warsaw"); !ok || abbrCity != abbrB {
		t.Fatalf("Warsaw abbreviated city name not stored in cache")
	}
}

func TestAbbrCityLatin(t *testing.T) {
	session, err := New()
	if err != nil {
		t.Fatal(err)
	}

	expectedAbbrCity := "/m/0c6fz"

	abbrCity, err := session.AbbrCity(context.Background(), "Łódź", language.English)
	if err != nil {
		t.Fatal(err)
	}
	if abbrCity != expectedAbbrCity {
		t.Fatalf("wrong abbreviated city name, expected: %s received: %s", expectedAbbrCity, abbrCity)
	}

	abbrCity, err = session.AbbrCity(context.Background(), "łódź", language.English)
	if err != nil {
		t.Fatal(err)
	}
	if abbrCity != expectedAbbrCity {
		t.Fatalf("wrong abbreviated city name, expected: %s received: %s", expectedAbbrCity, abbrCity)
	}

	abbrCity, err = session.AbbrCity(context.Background(), "lodz", language.English)
	if err != nil {
		t.Fatal(err)
	}
	if abbrCity != expectedAbbrCity {
		t.Fatalf("wrong abbreviated city name, expected: %s received: %s", expectedAbbrCity, abbrCity)
	}
}
