package main

import "fmt"

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

func newBill(name string) bill {
	b := bill{
		name:  name,
		items: map[string]float64{"pie": 5.99, "cake": 3.99},
		tip:   0,
	}

	return b
}

// make receiver function
func (billObj bill) format() string {
	var phrase string
	var total float64

	for key, value := range billObj.items {
		phrase += fmt.Sprintf("%-10v ...$%v \n", key+":", value)
		total += value
	}

	return phrase + fmt.Sprintf("\nThe total is %v", total)
}
