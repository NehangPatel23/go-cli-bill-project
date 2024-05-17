package main

import (
	"fmt"
	"os"
)

// Bill struct to hold bill data
type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

// Function to create a new bill
func newBill(name string) bill {
	b := bill{
		name:  name,
		items: map[string]float64{},
		tip:   0,
	}
	return b
}

// Method to format the bill as a string
func (b *bill) format() string {
	fs := "Bill Breakdown: \n\n"
	var total float64 = 0

	// List the items
	for k, v := range b.items {
		fs += fmt.Sprintf("%-25v ...$%v \n", k+":", v)
		total += v
	}

	// Tip
	fs += fmt.Sprintf("\n%-25v ...$%0.2f \n", "Tip:", b.tip)

	// Total
	fs += fmt.Sprintf("\n%-25v ...$%0.2f", "Total:", total+b.tip)

	return fs
}

// Method to update the tip
func (b *bill) updateTip(tip float64) {
	b.tip = tip
}

// Method to add an item to the bill
func (b *bill) addItem(name string, price float64) {
	b.items[name] = price
}

// Method to save the bill to a file
func (b *bill) saveBill() {
	data := []byte(b.format())

	err := os.WriteFile("bills/"+b.name+".txt", data, 0644)
	if err != nil {
		panic(err)
	}
	fmt.Println("The bill was saved successfully!")
}
