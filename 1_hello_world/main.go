package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
	fmt.Println(1)
	fmt.Println(1 + 3)
	fmt.Println(true)
	fmt.Println(false)

	var name string = "Soumik Ahammed"
	fmt.Println(name)

	var name2 = "Soumik Ahammed" // automatically infers the type
	fmt.Println(name2)

	var name3 string         // no need to initialize the variable
	name3 = "Soumik Ahammed" // but you have to initialize it before using it
	fmt.Println(name3)

	name4 := "Soumik Ahammed" // short variable declaration, only works inside functions
	fmt.Println(name4)

	var price int = 1000
	fmt.Println(price)

	var price2 float32 = 1000.50
	fmt.Println(price2)

	var price3 float64 = 1000.50
	fmt.Println(price3)

	var isAdult bool = true
	fmt.Println(isAdult)
}
