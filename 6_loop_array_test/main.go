package main

import (
	"fmt"
)

func checkNumberType(number int) string {
	if number > 0 {
		return "positive"
	} else if number < 0 {
		return "negative"
	}
	return "zero"
}

func main() {

	// 1. Take an integer and print whether it’s positive, negative, or zero.
	var number int
	fmt.Printf("Enter a number: ")
	fmt.Scanln(&number)

	result := checkNumberType(number)
	fmt.Println("The number is:", result)

	fmt.Println()

	// 2. Create a slice of strings with 5 names and loop to print each.
	slice := []string{"A", "B", "C", "D", "E"}
	for _, sl := range slice {
		fmt.Println(sl)
	}

	fmt.Println()

	// 2. Make a map of 3 countries → capitals, and loop through them.
	countries := map[string]string{
		"Bangladesh": "Dhaka",
		"Korea":      "Seoul",
		"Germany":    "Berlin",
	}
	for country, capital := range countries {
		fmt.Printf("The capital of %s is %s\n", country, capital)
	}
}
