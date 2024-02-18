package flights

import (
	"testing"
	"time"
)

const wrongAirportCode = "wrong"

func testValidateOffersArgs(t *testing.T, args Args, wantErr string) {
	gotErr := args.ValidateOffersArgs()

	if gotErr == nil {
		t.Fatalf("Validate call, should result in error args: %v", args)
	} else if gotErr.Error() != wantErr {
		t.Fatalf(`Wrong error want: "%s", got: "%s"`, wantErr, gotErr.Error())
	}
}

func testValidatePriceGraphArgs(t *testing.T, args PriceGraphArgs, wantErr string) {
	gotErr := args.Validate()

	if gotErr == nil {
		t.Fatalf("Validate call should result in error args: %v", args)
	} else if gotErr.Error() != wantErr {
		t.Fatalf(`Wrong error want: "%s", got: "%s"`, wantErr, gotErr.Error())
	}
}

func testValidateURLArg(t *testing.T, args Args, wantErr string) {
	gotErr := args.ValidateURLArgs()

	if gotErr == nil {
		t.Fatalf("Validate call should result in error args: %v", args)
	} else if gotErr.Error() != wantErr {
		t.Fatalf(`Wrong error want: "%s", got: "%s"`, wantErr, gotErr.Error())
	}
}

func TestValidateOffersArgs(t *testing.T) {
	args := Args{SrcCities: []string{"abc"}, SrcAirports: []string{}, DstCities: []string{}, DstAirports: []string{}}
	testValidateOffersArgs(t, args, "dst locations: number of locations should be at least 1, specified: 0")

	args = Args{SrcCities: []string{}, SrcAirports: []string{}, DstCities: []string{"abc"}, DstAirports: []string{}}
	testValidateOffersArgs(t, args, "src locations: number of locations should be at least 1, specified: 0")

	args = Args{SrcCities: []string{"abc"}, SrcAirports: []string{wrongAirportCode}, DstCities: []string{"abc"}, DstAirports: []string{}}
	testValidateOffersArgs(t, args, "src airport 'wrong' is not an airport code")

	args = Args{SrcCities: []string{"abc"}, SrcAirports: []string{}, DstCities: []string{"abc"}, DstAirports: []string{wrongAirportCode}}
	testValidateOffersArgs(t, args, "dst airport 'wrong' is not an airport code")

	args = Args{
		SrcCities: []string{"abc"}, SrcAirports: []string{}, DstCities: []string{"abc"}, DstAirports: []string{},
		Date:       time.Now().AddDate(0, 0, 3),
		ReturnDate: time.Now().AddDate(0, 0, 1),
	}
	testValidateOffersArgs(t, args, "returnDate is before date")

	args = Args{
		SrcCities: []string{"abc"}, SrcAirports: []string{}, DstCities: []string{"abc"}, DstAirports: []string{},
		Date:       time.Now().AddDate(0, 0, -1),
		ReturnDate: time.Now().AddDate(0, 0, 1),
	}
	testValidateOffersArgs(t, args, "date is before today's date")
}

func TestValidatePriceGraphArgs(t *testing.T) {
	args := PriceGraphArgs{SrcCities: []string{"abc"}, SrcAirports: []string{}, DstCities: []string{}, DstAirports: []string{}}
	testValidatePriceGraphArgs(t, args, "dst locations: number of locations should be at least 1, specified: 0")

	args = PriceGraphArgs{SrcCities: []string{}, SrcAirports: []string{}, DstCities: []string{"abc"}, DstAirports: []string{}}
	testValidatePriceGraphArgs(t, args, "src locations: number of locations should be at least 1, specified: 0")

	args = PriceGraphArgs{SrcCities: []string{"abc"}, SrcAirports: []string{wrongAirportCode}, DstCities: []string{"abc"}, DstAirports: []string{}}
	testValidatePriceGraphArgs(t, args, "src airport 'wrong' is not an airport code")

	args = PriceGraphArgs{SrcCities: []string{"abc"}, SrcAirports: []string{}, DstCities: []string{"abc"}, DstAirports: []string{wrongAirportCode}}
	testValidatePriceGraphArgs(t, args, "dst airport 'wrong' is not an airport code")

	args = PriceGraphArgs{
		SrcCities: []string{"abc"}, SrcAirports: []string{}, DstCities: []string{"abc"}, DstAirports: []string{},
		RangeStartDate: time.Now().AddDate(0, 0, 5),
		RangeEndDate:   time.Now().AddDate(0, 0, 170),
	}
	testValidatePriceGraphArgs(t, args, "number of days between dates is larger than 161, 165")

	args = PriceGraphArgs{
		SrcCities: []string{"abc"}, SrcAirports: []string{}, DstCities: []string{"abc"}, DstAirports: []string{},
		RangeStartDate: time.Now().AddDate(0, 0, 2),
		RangeEndDate:   time.Now().AddDate(0, 0, 2),
	}
	testValidatePriceGraphArgs(t, args, "rangeEndDate is the same as rangeStartDate")

	args = PriceGraphArgs{
		SrcCities: []string{"abc"}, SrcAirports: []string{}, DstCities: []string{"abc"}, DstAirports: []string{},
		RangeStartDate: time.Now().AddDate(0, 0, 5),
		RangeEndDate:   time.Now().AddDate(0, 0, 2),
	}
	testValidatePriceGraphArgs(t, args, "rangeEndDate is before rangeStartDate")

	args = PriceGraphArgs{
		SrcCities: []string{"abc"}, SrcAirports: []string{}, DstCities: []string{"abc"}, DstAirports: []string{},
		RangeStartDate: time.Now().AddDate(0, 0, -1),
		RangeEndDate:   time.Now().AddDate(0, 0, 2),
	}
	testValidatePriceGraphArgs(t, args, "rangeStartDate is before today's date")
}

func TestValidateURLArgs(t *testing.T) {
	args := Args{SrcCities: []string{"abc"}, SrcAirports: []string{}, DstCities: []string{}, DstAirports: []string{}}
	testValidateURLArg(t, args, "dst locations: number of locations should be at least 1, specified: 0")

	args = Args{SrcCities: []string{}, SrcAirports: []string{}, DstCities: []string{"abc"}, DstAirports: []string{}}
	testValidateURLArg(t, args, "src locations: number of locations should be at least 1, specified: 0")

	args = Args{SrcCities: []string{"abc"}, SrcAirports: []string{wrongAirportCode}, DstCities: []string{"abc"}, DstAirports: []string{}}
	testValidateURLArg(t, args, "src airport 'wrong' is not an airport code")

	args = Args{SrcCities: []string{"abc"}, SrcAirports: []string{}, DstCities: []string{"abc"}, DstAirports: []string{wrongAirportCode}}
	testValidateURLArg(t, args, "dst airport 'wrong' is not an airport code")
}
