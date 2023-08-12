# Google Flights API

[![Go Reference](https://pkg.go.dev/badge/github.com/krisukox/google-flights-api/flights.svg)](https://pkg.go.dev/github.com/krisukox/google-flights-api/flights)

This project is a Go client library for the Google Flights API. The client produces direct requests to the Google Flights API, which is much faster than using WebDriver. 

The Google Flights API doesn't have official documentation, so the project relies on analyzing how the [Google Flights website](https://www.google.com/travel/flights/) communicates with the backend.

The project uses [go-retryablehttp](https://github.com/hashicorp/go-retryablehttp) under the hood. Every request to the Google Flights API is retried five times in case of an error.

## Installation

```
go get -u github.com/krisukox/google-flights-api/flights
```

## Usage

### Session
Session is the main object that contains all the API-related functions.
```
session := flights.New()
```
### 1. Price graph
Find the cheapest flights from San Francisco to New York. The call uses the "Price graph" subsection of the Google Flights website.
#### Google Flights API:
```
offers, err := session.GetPriceGraph(
    context.Background(),
    flights.PriceGraphArgs{
        RangeStartDate: time.Now().AddDate(0, 0, 30),
        RangeEndDate:   time.Now().AddDate(0, 0, 60),
        TripLength:     7,
        SrcCities:      []string{"San Francisco"},
        DstCities:      []string{"New York"},
        Args:           flights.ArgsDefault(),
    },
)
if err != nil {
    log.Fatal(err)
}
fmt.Println(offers)
```
Example output:
```
[{2023-08-31 2023-09-07 245} {2023-09-01 2023-09-08 216} {2023-09-02 2023-09-09 180} {2023-09-03 2023-09-10 201} {2023-09-04 2023-09-11 225} {2023-09-05 2023-09-12 173}...
```

#### Google Flights website:

<img src="https://github.com/krisukox/google-flights-api/assets/38990293/c2cc5c50-2401-4c93-8e76-d7e8073c9f11" alt="drawing" width="800"/>

### 2. Serialize URL
```
url, err := session.SerializeURL(
    context.Background(),
    flights.URLArgs{
        Date:        time.Now().AddDate(0, 0, 30),
        ReturnDate:  time.Now().AddDate(0, 0, 37),
        SrcCities:   []string{"San Diego"},
        SrcAirports: []string{"LAX"},
        DstCities:   []string{"New York", "Philadelphia"},
        Args:        flights.ArgsDefault(),
    },
)
if err != nil {
    log.Fatal(err)
}
fmt.Println(url)
```
Example output:
```
https://www.google.com/travel/flights/search?tfs=CBwQAhpAEgoyMDIzLTA4LTMxagwIAxIIL20vMDcxdnJqBwgBEgNMQVhyDQgDEgkvbS8wMl8yODZyDAgDEggvbS8wZGNsZxpAEgoyMDIzLTA5LTA3ag0IAxIJL20vMDJfMjg2agwIAxIIL20vMGRjbGdyDAgDEggvbS8wNzF2cnIHCAESA0xBWEABSAFwAYIBCwj___________8BmAEB&curr=USD&hl=en
```
### 3. Get offers

The call below uses Spanish city names:
```
offers, priceRange, err := session.GetOffers(
    context.Background(),
    flights.OffersArgs{
        Date:       time.Now().AddDate(0, 0, 30),
        ReturnDate: time.Now().AddDate(0, 0, 37),
        SrcCities:  []string{"Madrid"},
        DstCities:  []string{"Estocolmo"},
        Args: flights.Args{
            Travelers: flights.Travelers{Adults: 2},
            Currency:  currency.EUR,
            Stops:     flights.Stop1,
            Class:     flights.Economy,
            TripType:  flights.RoundTrip,
            Lang:      language.Spanish,
        },
    },
)
if err != nil {
    log.Fatal(err)
}

if priceRange != nil {
    fmt.Printf("High price %d\n", int(priceRange.High))
    fmt.Printf("Low price %d\n", int(priceRange.Low))
}
fmt.Println(offers)
```

Example output:
```
High price 710
Low price 415
[{StartDate: 2023-08-31 22:30:00 +0000 UTC
ReturnDate: 2023-09-07 00:00:00 +0000 UTC
Price: 296
Flight: [DepAirportCode: MAD DepAirportName: Adolfo Suárez Madrid–Barajas Airport ArrAirportName: Josep Tarradellas Barcelona-El Prat Airport ArrAirportCode: BCN DepTime: 2023-08-31 22:30:00 +0000 UTC ArrTime: 2023-08-31 23:50:00 +0000 UTC Duration: 1h20m0s Airplane: Airbus A321 FlightNumber: VY 1009 AirlineName: Vueling Legroom: 29 inches DepAirportCode: BCN DepAirportName: Josep Tarradellas Barcelona-El Prat Airport ArrAirportName: Stockholm Arlanda Airport ArrAirportCode: ARN DepTime: 2023-09-01 06:35:00 +0000 UTC ArrTime: 2023-09-01 10:20:00 +0000 UTC Duration: 3h45m0s Airplane: Airbus A320 FlightNumber: VY 1265 AirlineName: Vueling Legroom: 29 inches]
SrcAirportCode: MAD
DstAirportCode: ARN
FlightDuration: 11h50m0s}
 {StartDate: 2023-08-31 06:20:00 +0000 UTC
ReturnDate: 2023-09-07 00:00:00 +0000 UTC
Price: 355
Flight: [DepAirportCode: MAD DepAirportName: Adolfo Suárez Madrid–Barajas Airport ArrAirportName: Brussels Airport ArrAirportCode: BRU DepTime: 2023-08-31 06:20:00 +0000 UTC ArrTime: 2023-08-31 08:35:00 +0000 UTC Duration: 2h15m0s Airplane: Airbus A320 FlightNumber: SN 3732 AirlineName: Brussels Airlines Legroom: 30 inches DepAirportCode: BRU DepAirportName: Brussels Airport ArrAirportName: Bromma Stockholm Airport ArrAirportCode: BMA DepTime: 2023-08-31 09:50:00 +0000 UTC ArrTime: 2023-08-31 12:00:00 +0000 UTC Duration: 2h10m0s Airplane: Airbus A319 FlightNumber: SN 2303 AirlineName: Brussels Airlines Legroom: 30 inches]
SrcAirportCode: MAD
DstAirportCode: BMA
FlightDuration: 5h40m0s}
 {StartDate: 2023-08-31 10:15:00 +0000 UTC
ReturnDate: 2023-09-07 00:00:00 +0000 UTC
Price: 370
Flight: [DepAirportCode: MAD DepAirportName: Adolfo Suárez Madrid–Barajas Airport ArrAirportName: Stockholm Arlanda Airport ArrAirportCode: ARN DepTime: 2023-08-31 10:15:00 +0000 UTC ArrTime: 2023-08-31 14:10:00 +0000 UTC Duration: 3h55m0s Airplane: Airbus A320 FlightNumber: IB 3314 AirlineName: Iberia Legroom: 28 inches]
SrcAirportCode: MAD
DstAirportCode: ARN
FlightDuration: 3h55m0s}
...
```

#### Google Flights website

<img src="https://github.com/krisukox/google-flights-api/assets/38990293/8a47ed4a-e54a-4917-9459-8994230b5d94" alt="drawing" width="800"/>


### More advanced examples:
```
go run ./examples/example1/main.go
go run ./examples/example2/main.go
go run ./examples/example3/main.go
```

## Bug / Feature / Suggestion

If you've found a bug, have a suggestion, or a feature you're looking for is not yet implemented, please feel free to [open an issue](https://github.com/krisukox/google-flights-api/issues). I'll try to handle it ASAP.
