package flights

import (
	"testing"
	"time"
)

const wrongAirportCode = "wrong"

func _testValidateOffersArgs(t *testing.T, args Args, errorValue string) {
	err := args.ValidateOffersArgs()

	if err == nil {
		t.Fatalf("validate call, should result in error")
	} else if err.Error() != errorValue {
		t.Fatalf(`Wrong error "%s"`, errorValue)
	}
}

func _testValidatePriceGraphArgs(t *testing.T, args PriceGraphArgs, errorValue string) {
	err := args.Validate()

	if err == nil {
		t.Fatalf("validate call, should result in error")
	} else if err.Error() != errorValue {
		t.Fatalf(`Wrong error "%s"`, errorValue)
	}
}

func _testValidateURLArg(t *testing.T, args Args, errorValue string) {
	err := args.ValidateURLArgs()

	if err == nil {
		t.Fatalf("validate call, should result in error")
	} else if err.Error() != errorValue {
		t.Fatalf(`Wrong error "%s"`, errorValue)
	}
}

func TestValidateOffersArgs(t *testing.T) {
	args := Args{SrcCities: []string{"abc"}, SrcAirports: []string{}, DstCities: []string{}, DstAirports: []string{}}
	_testValidateOffersArgs(t, args, "dst locations: number of locations should be at least 1, specified: 0")

	args = Args{SrcCities: []string{}, SrcAirports: []string{}, DstCities: []string{"abc"}, DstAirports: []string{}}
	_testValidateOffersArgs(t, args, "src locations: number of locations should be at least 1, specified: 0")

	args = Args{SrcCities: []string{"abc"}, SrcAirports: []string{wrongAirportCode}, DstCities: []string{"abc"}, DstAirports: []string{}}
	_testValidateOffersArgs(t, args, "src airport 'wrong' is not an airport code")

	args = Args{SrcCities: []string{"abc"}, SrcAirports: []string{}, DstCities: []string{"abc"}, DstAirports: []string{wrongAirportCode}}
	_testValidateOffersArgs(t, args, "dst airport 'wrong' is not an airport code")

	args = Args{
		SrcCities: []string{"abc"}, SrcAirports: []string{}, DstCities: []string{"abc"}, DstAirports: []string{},
		Date:       time.Now().AddDate(0, 0, 3),
		ReturnDate: time.Now().AddDate(0, 0, 1),
		Options:    Options{TripType: RoundTrip},
	}
	_testValidateOffersArgs(t, args, "returnDate is before date")

	args = Args{
		SrcCities: []string{"abc"}, SrcAirports: []string{}, DstCities: []string{"abc"}, DstAirports: []string{},
		Date:       time.Now().AddDate(0, 0, -1),
		ReturnDate: time.Now().AddDate(0, 0, 1),
	}
	_testValidateOffersArgs(t, args, "date is before today's date")
}

func TestValidatePriceGraphArgs(t *testing.T) {
	args := PriceGraphArgs{SrcCities: []string{"abc"}, SrcAirports: []string{}, DstCities: []string{}, DstAirports: []string{}}
	_testValidatePriceGraphArgs(t, args, "dst locations: number of locations should be at least 1, specified: 0")

	args = PriceGraphArgs{SrcCities: []string{}, SrcAirports: []string{}, DstCities: []string{"abc"}, DstAirports: []string{}}
	_testValidatePriceGraphArgs(t, args, "src locations: number of locations should be at least 1, specified: 0")

	args = PriceGraphArgs{SrcCities: []string{"abc"}, SrcAirports: []string{wrongAirportCode}, DstCities: []string{"abc"}, DstAirports: []string{}}
	_testValidatePriceGraphArgs(t, args, "src airport 'wrong' is not an airport code")

	args = PriceGraphArgs{SrcCities: []string{"abc"}, SrcAirports: []string{}, DstCities: []string{"abc"}, DstAirports: []string{wrongAirportCode}}
	_testValidatePriceGraphArgs(t, args, "dst airport 'wrong' is not an airport code")

	args = PriceGraphArgs{
		SrcCities: []string{"abc"}, SrcAirports: []string{}, DstCities: []string{"abc"}, DstAirports: []string{},
		RangeStartDate: time.Now().AddDate(0, 0, 5),
		RangeEndDate:   time.Now().AddDate(0, 0, 170),
	}
	_testValidatePriceGraphArgs(t, args, "number of days between dates is larger than 161, 165")

	args = PriceGraphArgs{
		SrcCities: []string{"abc"}, SrcAirports: []string{}, DstCities: []string{"abc"}, DstAirports: []string{},
		RangeStartDate: time.Now().AddDate(0, 0, 2),
		RangeEndDate:   time.Now().AddDate(0, 0, 2),
	}
	_testValidatePriceGraphArgs(t, args, "rangeEndDate is the same as rangeStartDate")

	args = PriceGraphArgs{
		SrcCities: []string{"abc"}, SrcAirports: []string{}, DstCities: []string{"abc"}, DstAirports: []string{},
		RangeStartDate: time.Now().AddDate(0, 0, 5),
		RangeEndDate:   time.Now().AddDate(0, 0, 2),
	}
	_testValidatePriceGraphArgs(t, args, "rangeEndDate is before rangeStartDate")

	args = PriceGraphArgs{
		SrcCities: []string{"abc"}, SrcAirports: []string{}, DstCities: []string{"abc"}, DstAirports: []string{},
		RangeStartDate: time.Now().AddDate(0, 0, -1),
		RangeEndDate:   time.Now().AddDate(0, 0, 2),
	}
	_testValidatePriceGraphArgs(t, args, "rangeStartDate is before today's date")
}

func TestValidateURLArgs(t *testing.T) {
	args := Args{SrcCities: []string{"abc"}, SrcAirports: []string{}, DstCities: []string{}, DstAirports: []string{}}
	_testValidateURLArg(t, args, "dst locations: number of locations should be at least 1, specified: 0")

	args = Args{SrcCities: []string{}, SrcAirports: []string{}, DstCities: []string{"abc"}, DstAirports: []string{}}
	_testValidateURLArg(t, args, "src locations: number of locations should be at least 1, specified: 0")

	args = Args{SrcCities: []string{"abc"}, SrcAirports: []string{wrongAirportCode}, DstCities: []string{"abc"}, DstAirports: []string{}}
	_testValidateURLArg(t, args, "src airport 'wrong' is not an airport code")

	args = Args{SrcCities: []string{"abc"}, SrcAirports: []string{}, DstCities: []string{"abc"}, DstAirports: []string{wrongAirportCode}}
	_testValidateURLArg(t, args, "dst airport 'wrong' is not an airport code")
}
