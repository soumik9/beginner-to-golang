package main

import (
	"errors"
	"fmt"
	"math"
)

type NegativeNumberError struct {
	Value float64
}

func divide(value, valueDivideBy int) (int, error) {
	if valueDivideBy == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	return value / valueDivideBy, nil
}

// customer error handler
func (e NegativeNumberError) Error() string {
	return fmt.Sprintf("Custom error with value: %.2f", e.Value)
}

func squareRoot(x float64) (float64, error) {
	if x < 0 {
		return 0, NegativeNumberError{Value: x}
	}
	return math.Sqrt(x), nil
}

func updateValue(val *int, newValue int) {
	*val = *val + newValue
}

func main() {

	// pointer
	a := 10
	point := &a

	fmt.Println("Value of a:", a)
	fmt.Println("Address of point:", point)
	fmt.Println("Value at address of point:", *point)

	*point = 20
	fmt.Println("New value of a:", a)
	fmt.Println("New value at address of point:", point)
	fmt.Println("New value of point:", *point)

	// update point value with function
	x := 20
	fmt.Println("Value of x before update:", x)
	updateValue(&x, 20)
	fmt.Println("Value of x after update:", x)

	// error check
	result, err := divide(10, 0)
	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println("Result of division:", result)
	}

	// custom error check
	resultSqrt, errSqrt := squareRoot(-4)
	if errSqrt != nil {
		fmt.Println("Error: ", errSqrt)
	} else {
		fmt.Printf("Result of square root: %.2f\n", resultSqrt)
	}

	resultSqrt1, errSqrt1 := squareRoot(24)
	if errSqrt1 != nil {
		fmt.Println("Error: ", errSqrt1)
	} else {
		fmt.Printf("Result of square root: %.2f\n", resultSqrt1)
	}

}
