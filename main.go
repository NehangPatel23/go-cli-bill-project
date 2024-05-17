package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Function to get input from the user with a prompt
func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n')
	return strings.TrimSpace(input), err
}

// Function to create a new bill
func createBill() bill {
	reader := bufio.NewReader(os.Stdin)
	name, _ := getInput("Create a new bill name: ", reader)
	b := newBill(name)
	fmt.Println("Created the bill - ", b.name)
	return b
}

// Function to prompt options to the user and handle the input
func promptOptions(b bill) {
	reader := bufio.NewReader(os.Stdin)
	opt, _ := getInput("Choose an option (a - Add an item, s - Save the bill, t - Add a tip): ", reader)

	switch opt {
	case "a":
		name, _ := getInput("Item Name: ", reader)
		price, _ := getInput("Item Price: ", reader)
		p, err := strconv.ParseFloat(price, 64)
		if err != nil {
			fmt.Println("The price must be a number!")
			promptOptions(b)
		}
		b.addItem(name, p)
		fmt.Println("Item Added - ", name, p)
		promptOptions(b)
	case "t":
		tip, _ := getInput("Enter tip amount: $", reader)
		t, err := strconv.ParseFloat(tip, 64)
		if err != nil {
			fmt.Println("The tip must be a number!")
			promptOptions(b)
		}
		b.updateTip(t)
		fmt.Println("Tip Added - ", t)
		promptOptions(b)
	case "s":
		b.saveBill()
		fmt.Println("You chose to save the bill to the file - ", b.name)
	default:
		fmt.Println("Invalid option...")
		promptOptions(b)
	}
}

// Entry point of the application
func main() {
	myBill := createBill()
	promptOptions(myBill)
}
