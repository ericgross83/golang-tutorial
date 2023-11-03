package main

import "fmt"

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

// make new bills
func newBill(name string) bill {
	newBill := bill{
		name:  name,
		items: map[string]float64{"pie": 5.99, "cake": 3.99},
		tip:   0,
	}
	return newBill
}

// format the bill
func (thisBill bill) format() string {
	formatedString := "Bill breakdown: \n"
	var total float64 = 0

	// list items
	for key, item := range thisBill.items {
		formatedString += fmt.Sprintf("%-25v ...%v€ \n", key+":", item)
		total += item
	}

	// total
	formatedString += fmt.Sprintf("%-25v ...%0.2f€", "total:", total)
	return formatedString
}
