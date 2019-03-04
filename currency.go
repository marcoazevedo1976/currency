package currency

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// Currency represents financial values.
type Currency struct {
	value         int
	DecimalPlaces int
	Symbol        string
	DecimalSep    string
	ThousandSep   string
}

// NewCurrency creates a Currency using default values as follows:
// value: 0, DecimalPlaces: 2, Symbol: "", DecimalSep: ".", ThousandSep: ",".
func NewCurrency() Currency {
	return Currency{value: 0, DecimalPlaces: 2, Symbol: "", DecimalSep: ".", ThousandSep: ","}
}

// SetValue sets the content of u to c's value. An
// error occurs if the type of u is not string, int, or float64.
func (c *Currency) SetValue(u interface{}) error {
	v, err := c.wrap(u)
	if err != nil {
		return err
	}
	c.value = v
	return nil
}

// Add adds u to c's value. An error occurs if the type of u is
// not string, int, or float64.
func (c *Currency) Add(u interface{}) error {
	return c.doOperation('+', u)
}

// Subtract takes away u from c's value. An error occurs if the
// type of u is not string, int, or float64.
func (c *Currency) Subtract(u interface{}) error {
	return c.doOperation('-', u)
}

// Multiply multiplies c's value by u. An error occurs if the
// type of u is not string, int, or float64.
func (c *Currency) Multiply(u interface{}) error {
	return c.doOperation('*', u)
}

// Divide divides c's value by u. An error occurs if the
// type of u is not string, int, or float64.
func (c *Currency) Divide(u interface{}) error {
	return c.doOperation('/', u)
}

// AsFloat returns the current value as float.
func (c *Currency) AsFloat() float64 {
	return c.unwrap(c.value)
}

func (c Currency) String() string {
	fValue := c.unwrap(c.value)
	strValue := strconv.FormatFloat(fValue, 'f', c.DecimalPlaces, 64)
	strValue = letOnlyNumbers(strValue)

	bs := []byte(strValue)
	dec := bs[len(bs)-c.DecimalPlaces:]

	var thousandParts []string
	bs = bs[:len(bs)-c.DecimalPlaces]
	for {
		if p := len(bs); p <= 3 {
			if p > 0 {
				thousandParts = append(thousandParts, string(bs))
			}
			break
		}
		thousandParts = append(thousandParts, string(bs[len(bs)-3:]))
		bs = bs[:len(bs)-3]
	}

	var strBld strings.Builder
	strBld.WriteString(c.Symbol)
	for i := len(thousandParts) - 1; i >= 0; i-- {
		if i < len(thousandParts)-1 {
			strBld.WriteString(c.ThousandSep)
		}
		strBld.WriteString(thousandParts[i])
	}
	strBld.WriteString(c.DecimalSep)
	strBld.Write(dec)

	return strings.TrimSpace(strBld.String())
}

// wrap indentifies the type of u, extracts decimal and thousand separators,
// and return the result as an int.
// Examples u = 5.37 returns 537
//          u = 0.61 returns 61
//          u = 7 returns 700
//          u = 1,250.65 returns 125065
// The function is ready to receive string, int, or float64 as parameter. An error occurs otherwise.
func (c *Currency) wrap(u interface{}) (v int, err error) {
	switch u.(type) {
	case string:
		s := u.(string)
		s = strings.Replace(s, c.ThousandSep, "", -1)
		s = strings.Replace(s, c.Symbol, "", -1)
		s = strings.Replace(s, c.DecimalSep, ".", -1)
		f, e := strconv.ParseFloat(s, 64)
		if e != nil {
			return 0, e
		}
		return floatToInt(f, c.DecimalPlaces)
	case int:
		i := u.(int)
		f := float64(i)
		return floatToInt(f, c.DecimalPlaces)
	case float64:
		f := u.(float64)
		return floatToInt(f, c.DecimalPlaces)
	default:
		return 0, fmt.Errorf("Unknown type: %v", u)
	}
}

// unwrap returns the parameter named v as a float64.
// Examples v = 537 returns 5.37
//          v = 61 returns 0.61
//          v = 700 returns 7.00
//          v = 125065 returns 1250.65
func (c *Currency) unwrap(v int) float64 {
	return float64(v) / math.Pow10(c.DecimalPlaces)
}

// doOperantion does the operation defined by op.
func (c *Currency) doOperation(op byte, u interface{}) error {
	v, err := c.wrap(u)
	if err != nil {
		return err
	}

	total := c.unwrap(c.value)
	newValue := c.unwrap(v)

	switch op {
	case '+':
		total += newValue
	case '-':
		total -= newValue
	case '*':
		total *= newValue
	case '/':
		total /= newValue
	default:
		return fmt.Errorf("Operation not found: %v", u)
	}

	c.value, _ = c.wrap(total)
	return nil
}

func floatToInt(f float64, decimalPlaces int) (int, error) {
	base10 := math.Pow10(decimalPlaces)
	f = math.Round(f*base10) / base10
	s := strconv.FormatFloat(f, 'f', decimalPlaces, 64)
	s = letOnlyNumbers(s)
	return strconv.Atoi(s)
}

// letOnlyNumbers returns the string without the thousand separators and
// decimal separator. The sign is preserved.
func letOnlyNumbers(s string) string {
	s = strings.Replace(s, ",", "", -1)
	s = strings.Replace(s, ".", "", -1)
	return s
}
