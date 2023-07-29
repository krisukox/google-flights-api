package flights

import "fmt"

func checkNumberOfLocations(cities, airports []string) error {
	length := len(cities) + len(airports)
	if length > 7 {
		return fmt.Errorf("number of locations should not exceed 7 specified: %d", length)
	}

	if length < 1 {
		return fmt.Errorf("number of locations should be at least 1, specified: %d", length)
	}

	return nil
}

func checkDupplicate(location1, location2 []string) error {
	for _, l1 := range location1 {
		for _, l2 := range location2 {
			if l1 == l2 {
				return fmt.Errorf("duplicated locations: %s %s", l1, l2)
			}
		}
	}

	return nil
}

func checkLocations(srcCities, srcAirports, dstCities, dstAirports []string) error {
	if err := checkNumberOfLocations(srcCities, srcAirports); err != nil {
		return fmt.Errorf("src locations: %s", err)
	}
	if err := checkNumberOfLocations(dstCities, dstAirports); err != nil {
		return fmt.Errorf("dst locations: %s", err)
	}
	if err := checkDupplicate(srcCities, dstCities); err != nil {
		return fmt.Errorf("cities: %s", err)
	}
	if err := checkDupplicate(dstCities, dstAirports); err != nil {
		return fmt.Errorf("airports: %s", err)
	}

	return nil
}
