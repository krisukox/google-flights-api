package api

import (
	"fmt"
	"log"
	"net/url"
	"time"
)

// flightDate
// returnFlightDate
// originCity
// targetCity

// f.req=%5Bnull%2C%22%5B%5Bnull%2Cnull%2Cnull%2C%5C%22HJRBIeHroQnsACCh4QBG--------lmbhp15AAAAAGScKcsGaWrAA%5C%22%5D%2C%5Bnull%2Cnull%2C1%2Cnull%2C%5B%5D%2C1%2C%5B1%2C0%2C0%2C0%5D%2Cnull%2Cnull%2Cnull%2Cnull%2Cnull%2Cnull%2C%5B%5B%5B%5B%5B%5C%22%2Fm%2F0845b%5C%22%2C4%5D%5D%5D%2C%5B%5B%5B%5C%22%2Fm%2F056_y%5C%22%2C5%5D%5D%5D%2Cnull%2C0%2C%5B%5D%2C%5B%5D%2C%5C%222023-12-01%5C%22%2Cnull%2C%5B%5D%2C%5B%5D%2C%5B%5D%2Cnull%2Cnull%2C%5B%5D%2C3%5D%2C%5B%5B%5B%5B%5C%22%2Fm%2F056_y%5C%22%2C5%5D%5D%5D%2C%5B%5B%5B%5C%22%2Fm%2F0845b%5C%22%2C4%5D%5D%5D%2Cnull%2C0%2C%5B%5D%2C%5B%5D%2C%5C%222023-12-08%5C%22%2Cnull%2C%5B%5D%2C%5B%5D%2C%5B%5D%2Cnull%2Cnull%2C%5B%5D%2C3%5D%5D%2Cnull%2Cnull%2Cnull%2C1%2Cnull%2Cnull%2Cnull%2Cnull%2Cnull%2C%5B%5D%5D%2C1%2C0%2C0%5D%22%5D&at=AAuQa1qjMakasqKYcQeoFJjN7RZ3%3A1687955915303&

func GetRawData(date time.Time, returnDate time.Time, originCity string, targetCity string) {
	encodedValue := "%5Bnull%2C%22%5B%5Bnull%2Cnull%2Cnull%2C%5C%22HJRBIeHroQnsACCh4QBG--------lmbhp15AAAAAGScKcsGaWrAA%5C%22%5D%2C%5Bnull%2Cnull%2C1%2Cnull%2C%5B%5D%2C1%2C%5B1%2C0%2C0%2C0%5D%2Cnull%2Cnull%2Cnull%2Cnull%2Cnull%2Cnull%2C%5B%5B%5B%5B%5B%5C%22%2Fm%2F0845b%5C%22%2C4%5D%5D%5D%2C%5B%5B%5B%5C%22%2Fm%2F056_y%5C%22%2C5%5D%5D%5D%2Cnull%2C0%2C%5B%5D%2C%5B%5D%2C%5C%222023-12-01%5C%22%2Cnull%2C%5B%5D%2C%5B%5D%2C%5B%5D%2Cnull%2Cnull%2C%5B%5D%2C3%5D%2C%5B%5B%5B%5B%5C%22%2Fm%2F056_y%5C%22%2C5%5D%5D%5D%2C%5B%5B%5B%5C%22%2Fm%2F0845b%5C%22%2C4%5D%5D%5D%2Cnull%2C0%2C%5B%5D%2C%5B%5D%2C%5C%222023-12-08%5C%22%2Cnull%2C%5B%5D%2C%5B%5D%2C%5B%5D%2Cnull%2Cnull%2C%5B%5D%2C3%5D%5D%2Cnull%2Cnull%2Cnull%2C1%2Cnull%2Cnull%2Cnull%2Cnull%2Cnull%2C%5B%5D%5D%2C1%2C0%2C0%5D%22%5D"
	decodedValue, err := url.QueryUnescape(encodedValue)
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Println(decodedValue)

	decodedMy := `[null,"[[null,null,null,\"HJRBIeHroQnsACCh4QBG--------lmbhp15AAAAAGScKcsGaWrAA\"],[null,null,1,null,[],1,[1,0,0,0],null,null,null,null,null,null,[[[[[\"/m/0845b\",4]]],[[[\"/m/056_y\",5]]],null,0,[],[],\"2023-12-01\",null,[],[],[],null,null,[],3],[[[[\"/m/056_y\",5]]],[[[\"/m/0845b\",4]]],null,0,[],[],\"2023-12-08\",null,[],[],[],null,null,[],3]],null,null,null,1,null,null,null,null,null,[]],1,0,0]"]`

	fmt.Println(url.QueryEscape(decodedMy))

	// encodedValue = "AAuQa1qjMakasqKYcQeoFJjN7RZ3%3A1687955915303"
	// decodedValue, err = url.QueryUnescape(encodedValue)
	// if err != nil {
	// 	log.Fatalf(err.Error())
	// }
	// fmt.Println(decodedValue)
}
