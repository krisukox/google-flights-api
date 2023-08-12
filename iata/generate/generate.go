package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"sort"
	"sync"
	"time"

	"github.com/krisukox/google-flights-api/flights"
)

type airport struct {
	Iata string
	Tz   string
	City string
}

type result struct {
	line string
	err  error
}

func getAirports(commitHash string) (map[string]airport, error) {
	resp, err := http.Get(fmt.Sprintf("https://raw.githubusercontent.com/mwgg/Airports/%s/airports.json", commitHash))
	if err != nil {
		return nil, err
	}

	airports := map[string]airport{}
	err = json.NewDecoder(resp.Body).Decode(&airports)
	if err != nil {
		return nil, err
	}
	return airports, err
}

func main() {
	commitHash := "f259c38566a5acbcb04b64eb5ad01d14bf7fd07c"

	airports, err := getAirports(commitHash)
	if err != nil {
		log.Fatal(err)
	}

	session, err := flights.New()
	if err != nil {
		log.Fatal(err)
	}

	// We have to iterate over the map every time in the same order because the airport
	// database has a bug where it has the same iata code twice with different timezones.
	keys := make([]string, 0)
	for k := range airports {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	out := make(chan result)
	checked := map[string]struct{}{}
	var wg sync.WaitGroup

	caseTmpl := `	case "%s":
		return Location{"%s", "%s"}
`

	var a airport
	for _, k := range keys {
		a = airports[k]
		if a.Iata == "" || a.Iata == "0" {
			continue
		}
		if _, ok := checked[a.Iata]; ok {
			continue
		}
		checked[a.Iata] = struct{}{}

		wg.Add(1)
		go func(iata, tz, city string) {
			defer wg.Done()

			ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
			defer cancel()

			ok, err := session.IsIATASupported(ctx, iata)
			if err != nil {
				out <- result{err: err}
			}

			if ok {
				res := result{line: fmt.Sprintf(caseTmpl, iata, city, tz), err: nil}
				out <- res
			}
		}(a.Iata, a.Tz, a.City)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	iataFileContent := fmt.Sprintf(
		`// Package iata contains IATA airport codes, which are supported by the Google Flights API, along with time zones.
// This package was generated using an airport list (which can be found at this address: [airports.json])
// and the Google Flights API.
//
// Command: go run ./iata/generate/generate.go
//
// Generation date: %s
//
// [airports.json]: https://github.com/mwgg/Airports/blob/%s/airports.json
package iata

type Location struct{ City, Tz string }

// IATATimeZone turns IATA airport codes into the time zone where the airport is located.
// If IATATimeZone can't find an IATA airport code, then it returns "Not supported IATA Code".
func IATATimeZone(iata string) Location {
	switch iata {
`, time.Now().Format(time.DateOnly), commitHash)

	lines := []string{}

	for res := range out {
		if res.err != nil {
			log.Fatal(res.err)
		}
		lines = append(lines, res.line)
	}

	sort.Strings(lines)

	for _, line := range lines {
		iataFileContent += line
	}

	iataFileContent += `	}
	return Location{"Not supported IATA Code", "Not supported IATA Code"}
}
`

	iataFile, err := os.Create("./iata/iata.go")
	if err != nil {
		log.Fatal(err)
	}
	defer iataFile.Close()

	_, err = iataFile.WriteString(iataFileContent)
	if err != nil {
		log.Fatal(err)
	}
}
