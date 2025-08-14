package main

import (
	"fmt"
)

func calculateGrade(passMark int) string {
	if passMark < 0 || passMark > 100 {
		return "Invalid"
	} else if passMark > 80 {
		return "A+"
	} else if passMark > 33 {
		return "Pass"
	}
	return "Fail"
}

func main() {

	var passMark int

	fmt.Printf("Enter the Pass Mark: ")
	fmt.Scanln(&passMark)

	result := calculateGrade(passMark)
	fmt.Printf("The grade is: %s\n", result)

	fmt.Println("")

	// for loop
	for i := 0; i < 5; i++ {
		fmt.Printf("Loop iteration: %d\n", i)
	}

	fmt.Println("")

	// while loop
	count := 0
	for count < 5 {
		fmt.Printf("Loop iteration: %d\n", count)
		count++
	}

	fmt.Println("")

	// loop infinite
	for {
		fmt.Println("Infinite loop")
		break
	}
}
