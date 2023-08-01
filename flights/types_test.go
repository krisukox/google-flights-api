package flights

import (
	"testing"
	"time"
)

const wrongAirportCode = "wrong"

type validatable interface {
	Validate() error
}

func _testValidate(t *testing.T, args validatable, errorValue string) {
	err := args.Validate()

	if err == nil {
		t.Fatalf("validate call, should result in error")
	} else if err.Error() != errorValue {
		t.Fatalf(`Wrong error "%s"`, errorValue)
	}
}

func TestValidateOffersArgs(t *testing.T) {
	args := OffersArgs{SrcCities: []string{"abc"}, SrcAirports: []string{}, DstCities: []string{}, DstAirports: []string{}}
	_testValidate(t, &args, "dst locations: number of locations should be at least 1, specified: 0")

	args = OffersArgs{SrcCities: []string{}, SrcAirports: []string{}, DstCities: []string{"abc"}, DstAirports: []string{}}
	_testValidate(t, &args, "src locations: number of locations should be at least 1, specified: 0")

	args = OffersArgs{SrcCities: []string{"abc"}, SrcAirports: []string{wrongAirportCode}, DstCities: []string{"abc"}, DstAirports: []string{}}
	_testValidate(t, &args, "src airport 'wrong' is not an airport code")

	args = OffersArgs{SrcCities: []string{"abc"}, SrcAirports: []string{}, DstCities: []string{"abc"}, DstAirports: []string{wrongAirportCode}}
	_testValidate(t, &args, "dst airport 'wrong' is not an airport code")

	args = OffersArgs{
		SrcCities: []string{"abc"}, SrcAirports: []string{}, DstCities: []string{"abc"}, DstAirports: []string{},
		Date:       time.Now().AddDate(0, 0, 3),
		ReturnDate: time.Now().AddDate(0, 0, 1),
	}
	_testValidate(t, &args, "returnDate is before date")

	args = OffersArgs{
		SrcCities: []string{"abc"}, SrcAirports: []string{}, DstCities: []string{"abc"}, DstAirports: []string{},
		Date:       time.Now().AddDate(0, 0, -1),
		ReturnDate: time.Now().AddDate(0, 0, 1),
	}
	_testValidate(t, &args, "date is before today's date")
}

func TestValidatePriceGraphArgs(t *testing.T) {
	args := PriceGraphArgs{SrcCities: []string{"abc"}, SrcAirports: []string{}, DstCities: []string{}, DstAirports: []string{}}
	_testValidate(t, &args, "dst locations: number of locations should be at least 1, specified: 0")

	args = PriceGraphArgs{SrcCities: []string{}, SrcAirports: []string{}, DstCities: []string{"abc"}, DstAirports: []string{}}
	_testValidate(t, &args, "src locations: number of locations should be at least 1, specified: 0")

	args = PriceGraphArgs{SrcCities: []string{"abc"}, SrcAirports: []string{wrongAirportCode}, DstCities: []string{"abc"}, DstAirports: []string{}}
	_testValidate(t, &args, "src airport 'wrong' is not an airport code")

	args = PriceGraphArgs{SrcCities: []string{"abc"}, SrcAirports: []string{}, DstCities: []string{"abc"}, DstAirports: []string{wrongAirportCode}}
	_testValidate(t, &args, "dst airport 'wrong' is not an airport code")

	args = PriceGraphArgs{
		SrcCities: []string{"abc"}, SrcAirports: []string{}, DstCities: []string{"abc"}, DstAirports: []string{},
		RangeStartDate: time.Now().AddDate(0, 0, 5),
		RangeEndDate:   time.Now().AddDate(0, 0, 170),
	}
	_testValidate(t, &args, "number of days between dates is larger than 161, 165")

	args = PriceGraphArgs{
		SrcCities: []string{"abc"}, SrcAirports: []string{}, DstCities: []string{"abc"}, DstAirports: []string{},
		RangeStartDate: time.Now().AddDate(0, 0, 2),
		RangeEndDate:   time.Now().AddDate(0, 0, 2),
	}
	_testValidate(t, &args, "rangeEndDate is the same as rangeStartDate")

	args = PriceGraphArgs{
		SrcCities: []string{"abc"}, SrcAirports: []string{}, DstCities: []string{"abc"}, DstAirports: []string{},
		RangeStartDate: time.Now().AddDate(0, 0, 5),
		RangeEndDate:   time.Now().AddDate(0, 0, 2),
	}
	_testValidate(t, &args, "rangeEndDate is before rangeStartDate")

	args = PriceGraphArgs{
		SrcCities: []string{"abc"}, SrcAirports: []string{}, DstCities: []string{"abc"}, DstAirports: []string{},
		RangeStartDate: time.Now().AddDate(0, 0, -1),
		RangeEndDate:   time.Now().AddDate(0, 0, 2),
	}
	_testValidate(t, &args, "rangeStartDate is before today's date")
}

func TestValidateURLArgs(t *testing.T) {
	args := URLArgs{SrcCities: []string{"abc"}, SrcAirports: []string{}, DstCities: []string{}, DstAirports: []string{}}
	_testValidate(t, &args, "dst locations: number of locations should be at least 1, specified: 0")

	args = URLArgs{SrcCities: []string{}, SrcAirports: []string{}, DstCities: []string{"abc"}, DstAirports: []string{}}
	_testValidate(t, &args, "src locations: number of locations should be at least 1, specified: 0")

	args = URLArgs{SrcCities: []string{"abc"}, SrcAirports: []string{wrongAirportCode}, DstCities: []string{"abc"}, DstAirports: []string{}}
	_testValidate(t, &args, "src airport 'wrong' is not an airport code")

	args = URLArgs{SrcCities: []string{"abc"}, SrcAirports: []string{}, DstCities: []string{"abc"}, DstAirports: []string{wrongAirportCode}}
	_testValidate(t, &args, "dst airport 'wrong' is not an airport code")
}
