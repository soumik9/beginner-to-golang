package main

import "fmt"

func calculation(a int, b int, conditionType string) int {
	if conditionType == "+" {
		return a + b
	} else if conditionType == "-" {
		return a - b
	} else if conditionType == "*" {
		return a * b
	} else if conditionType == "/" {
		return a / b
	} else if conditionType == "%" {
		return a % b
	} else {
		return 0
	}
}

func main() {
	sum := calculation(10, 20, "+")
	fmt.Println(sum)

	substract := calculation(10, 20, "-")
	fmt.Println(substract)

	multiply := calculation(10, 20, "*")
	fmt.Println(multiply)

	divide := calculation(10, 20, "/")
	fmt.Println(divide)

	modulus := calculation(10, 20, "%")
	fmt.Println(modulus)

	// array

	arr := []int{1, 2, 3, 4, 5} // without fixed size
	fmt.Println(arr)

	arr2 := [5]int{1, 2, 3, 4, 5} // fixed size
	fmt.Println(arr2)

	arr = append(arr, 6) // appending to the array
	fmt.Println(arr)

}
