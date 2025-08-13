package main

import (
	"fmt"
	"math"
)

func calculate(num1, num2 float64, op string) (float64, bool) {
	switch op {
	case "+":
		return num1 + num2, true
	case "-":
		return num1 - num2, true
	case "*":
		return num1 * num2, true
	case "/":
		if num2 != 0 {
			return num1 / num2, true
		}
	}
	return 0, false
}

func main() {

	var num1, num2 float64
	var operator string

	fmt.Printf("Enter first number: ")
	fmt.Scanln(&num1)

	fmt.Printf("Enter second number: ")
	fmt.Scanln(&num2)

	fmt.Printf("Enter operator (+, -, *, /): ")
	fmt.Scanln(&operator)

	var result float64
	valid := true
	result, valid = calculate(num1, num2, operator)

	if valid {
		if result == math.Trunc(result) {
			fmt.Printf("Result: %.0f\n", result)
		} else {
			fmt.Printf("Result: %.2f\n", result)
		}
	}
}
