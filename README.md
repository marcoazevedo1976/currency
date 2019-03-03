# Currency

A simple library to deal with financial values written in Go. Currency aims to avoid well-known rounding problems when using float-point data types to represent decimal numbers. 


## Features 
- Rounding according to the defined decimal places
- Configurable thousand and decimal separators, currency symbol, and decimal places
- Four basic mathematical operations
- Implements Stringer interface


## Installation 

```
go get githum.com/marcoazevedo1976/currency
```

## Getting started

```
type expense struct {
	title  string
	amount float64
}

func main() {

	list := []expense{
		{"rent", 1600},
		{"water / sewer / garbage", 425.37},
		{"gas", 200}}

	total := currency.NewCurrency()
	total.Symbol = "$"

	for _, e := range list {
		total.Add(e.amount)
	}

	fmt.Println("Total Expenses:", total) // Total Expenses: $2,225.37

}
```

 