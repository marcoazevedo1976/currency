package currency_test

import (
	"testing"

	"github.com/marcoazevedo1976/currency"
)

type dataTest struct {
	input    interface{}
	expected string
}

const errorMsg = "Expected %v, found %v"

func TestNew(t *testing.T) {
	twoDecPlacesList := [...]dataTest{
		{12, "12.00"},
		{-12, "-12.00"},
		{7, "7.00"},
		{2.5, "2.50"},
		{"5.37", "5.37"},
		{"5350", "5,350.00"},
		{"-5350", "-5,350.00"},
		{"1250.647", "1,250.65"},
		{"1,250.647", "1,250.65"},
		{0.61, "0.61"}}

	for _, test := range twoDecPlacesList {
		val := currency.NewCurrency()
		err := val.SetValue(test.input)
		if val.String() != test.expected {
			t.Logf("Error: %v", err)
			t.Errorf(errorMsg, test.expected, val)
		}
	}
}

func TestNewUsingSymbol(t *testing.T) {
	type st struct {
		thouSep string
		decSep  string
		symbol  string
		dataTest
	}
	twoDecPlacesList := [...]st{
		{",", ".", "$", dataTest{12, "$12.00"}},
		{",", ".", "$", dataTest{7, "$7.00"}},
		{",", ".", "$", dataTest{2.5, "$2.50"}},
		{",", ".", "$", dataTest{"5.37", "$5.37"}},
		{",", ".", "$", dataTest{"5350", "$5,350.00"}},
		{",", ".", "$", dataTest{"-5350", "$-5,350.00"}},
		{".", ",", "R$", dataTest{"1250,647", "R$1.250,65"}},
		{",", ".", "$", dataTest{"1,250.647", "$1,250.65"}}}

	for _, test := range twoDecPlacesList {
		val := currency.NewCurrency()

		val.ThousandSep = test.thouSep
		val.DecimalSep = test.decSep
		val.Symbol = test.symbol
		err := val.Add(test.input)

		if val.String() != test.expected {
			t.Logf("Error: %v", err)
			t.Errorf(errorMsg, test.expected, val)
		}
	}
}
func TestAddUsing2DecimalPlaces(t *testing.T) {
	twoDecPlacesAdding := [...]dataTest{
		{12, "12.00"},
		{7, "19.00"},
		{2.5, "21.50"},
		{"asdasdas", "21.50"},
		{"5.37", "26.87"},
		{-5.01, "21.86"}}

	val := currency.NewCurrency()
	for _, test := range twoDecPlacesAdding {
		val.Add(test.input)
		if val.String() != test.expected {
			t.Errorf(errorMsg, test.expected, val)
		}
	}
}

func TestAddUsing3DecimalPlaces(t *testing.T) {
	twoDecPlacesAdding := [...]dataTest{
		{12, "12.000"},
		{7, "19.000"},
		{2.5, "21.500"},
		{"asdasdas", "21.500"},
		{"5.3733", "26.873"},
		{-5.010, "21.863"}}

	val := currency.NewCurrency()
	val.DecimalPlaces = 3
	for _, test := range twoDecPlacesAdding {
		val.Add(test.input)
		if val.String() != test.expected {
			t.Errorf(errorMsg, test.expected, val)
		}
	}
}
func TestSubtractUsing2DecimalPlaces(t *testing.T) {
	twoDecPlacesSubtracting := [...]dataTest{
		{120, "120.00"},
		{7, "113.00"},
		{2.5, "110.50"},
		{"asdasdas", "110.50"},
		{"5.37", "105.13"},
		{"-0.13", "105.26"}}

	val := currency.NewCurrency()

	for i, test := range twoDecPlacesSubtracting {
		if i == 0 {
			val.SetValue(test.input)
		} else {
			val.Subtract(test.input)
		}
		if val.String() != test.expected {
			t.Errorf(errorMsg, test.expected, val)
		}
	}
}

func TestSubtractUsing3DecimalPlaces(t *testing.T) {
	twoDecPlacesSubtracting := [...]dataTest{
		{120, "120.000"},
		{7, "113.000"},
		{2.5, "110.500"},
		{"asdasdas", "110.500"},
		{"5.370", "105.130"},
		{-1.135, "106.265"},
		{112.300, "-6.035"}}

	val := currency.NewCurrency()
	val.DecimalPlaces = 3
	for i, test := range twoDecPlacesSubtracting {
		if i == 0 {
			val.SetValue(test.input)
		} else {
			val.Subtract(test.input)
		}
		if val.String() != test.expected {
			t.Errorf(errorMsg, test.expected, val)
		}
	}
}
func TestMultiplyUsing2DecimalPlaces(t *testing.T) {
	twoDecPlacesMult := [...]dataTest{
		{120, "120.00"},
		{2, "240.00"},
		{10, "2,400.00"},
		{"asdasdas", "2,400.00"},
		{"5.37", "12,888.00"}}

	val := currency.NewCurrency()

	for i, test := range twoDecPlacesMult {
		if i == 0 {
			val.SetValue(test.input)
		} else {
			val.Multiply(test.input)
		}
		if val.String() != test.expected {
			t.Errorf(errorMsg, test.expected, val)
		}
	}
}

func TestMultiplyUsing3DecimalPlaces(t *testing.T) {
	twoDecPlacesMult := [...]dataTest{
		{120, "120.000"},
		{2, "240.000"},
		{10, "2,400.000"},
		{"asdasdas", "2,400.000"},
		{"5.37", "12,888.000"}}

	val := currency.NewCurrency()
	val.DecimalPlaces = 3

	for i, test := range twoDecPlacesMult {
		if i == 0 {
			val.SetValue(test.input)
		} else {
			val.Multiply(test.input)
		}
		if val.String() != test.expected {
			t.Errorf(errorMsg, test.expected, val)
		}
	}
}

func TestDivideUsing2DecimalPlaces(t *testing.T) {
	twoDecPlacesMult := [...]dataTest{
		{120.57, "120.57"},
		{2, "60.29"},
		{"asdasdas", "60.29"}}

	val := currency.NewCurrency()

	for i, test := range twoDecPlacesMult {
		if i == 0 {
			val.SetValue(test.input)
		} else {
			val.Divide(test.input)
		}
		if val.String() != test.expected {
			t.Errorf(errorMsg, test.expected, val)
		}
	}
}

func TestDivideUsing3DecimalPlaces(t *testing.T) {
	twoDecPlacesMult := [...]dataTest{
		{120.573, "120.573"},
		{2, "60.287"},
		{"asdasdas", "60.287"}}

	val := currency.NewCurrency()
	val.DecimalPlaces = 3

	for i, test := range twoDecPlacesMult {
		if i == 0 {
			val.SetValue(test.input)
		} else {
			val.Divide(test.input)
		}
		if val.String() != test.expected {
			t.Errorf(errorMsg, test.expected, val)
		}
	}
}

func TestInvalidValuesA(t *testing.T) {
	invalidData := [...]dataTest{
		{"qwe", "0.00"},
		{"5,352,525", "5,352,525.00"},
		{"5350..,", "0.00"}}

	for _, test := range invalidData {
		val := currency.NewCurrency()
		err := val.Add(test.input)
		if val.String() != test.expected {
			t.Logf("Error: %v", err)
			t.Errorf(errorMsg, test.expected, val)
		}
	}
}

func TestInvalidValuesB(t *testing.T) {
	val := currency.NewCurrency()
	err := val.Add("dsdfsdfsd")

	if err == nil {
		t.Error("err is not null")
	}

	expected := "0.00"
	if val.String() != expected {
		t.Logf("Error: %v", err)
		t.Errorf(errorMsg, expected, val)
	}
}

func TestAsFloat(t *testing.T) {
	type st struct {
		c   currency.Currency
		in  interface{}
		out float64
	}

	data := [...]st{
		{currency.Currency{DecimalPlaces: 2, ThousandSep: ",", DecimalSep: "."}, 12.69, 12.69},
		{currency.Currency{DecimalPlaces: 2, ThousandSep: ",", DecimalSep: "."}, 12.696, 12.70},
		{currency.Currency{DecimalPlaces: 3, ThousandSep: ",", DecimalSep: "."}, -12.696, -12.696},
		{currency.Currency{DecimalPlaces: 2, ThousandSep: ".", DecimalSep: ","}, 12.696, 12.70},
		{currency.Currency{DecimalPlaces: 2, ThousandSep: ",", DecimalSep: "."}, 12.696, 12.70},
		{currency.Currency{DecimalPlaces: 2, ThousandSep: ".", DecimalSep: ","}, "-12,696", -12.70},
		{currency.Currency{DecimalPlaces: 3, ThousandSep: ",", DecimalSep: "."}, "4,456.995", 4456.995},
		{currency.Currency{DecimalPlaces: 2, ThousandSep: ".", DecimalSep: ","}, "4.456,99", 4456.99},
		{currency.Currency{DecimalPlaces: 2, ThousandSep: ".", DecimalSep: ","}, "1.124.456,99", 1124456.99},
	}

	for _, d := range data {
		d.c.SetValue(d.in)
		expected := d.out
		found := d.c.AsFloat()
		if found != expected {
			t.Errorf(errorMsg, expected, found)
		}
	}

}
