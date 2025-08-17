package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

type Rectangle struct {
	Width  float64
	Height float64
}

type Car struct {
	Brand string
	Model string
	Year  int
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r *Rectangle) Scale(factor float64) {
	r.Width *= factor
	r.Height *= factor
}

// practice fun
func (c Car) Details() string {
	return fmt.Sprintf("Car: %s %s (%d)", c.Brand, c.Model, c.Year)
}

func (c *Car) ChangeYear(newYear int) {
	c.Year = newYear
}

func main() {

	p1 := Person{Name: "Soumik", Age: 25}
	p2 := Person{Name: "Mun", Age: 22}

	fmt.Printf("Person 1: %s is %d years old\n", p1.Name, p1.Age)
	fmt.Println("Person 2:", p2.Name, "is", p2.Age, "years old.")

	react := Rectangle{Width: 10, Height: 5}
	fmt.Printf("Area of rectangle: %.2f\n", react.Area())

	react.Scale(2)
	fmt.Printf("Scaled area of rectangle: %.2f\n", react.Area())

	// practices test

	/**
	1. Create a Car struct with Brand, Model, and Year fields.
	2. Add a method Details() that prints the car details.
	3. Add a method ChangeYear(newYear int) that updates the car year.
	4. Create multiple Car instances and call their methods.
	*/

	car1 := Car{Brand: "Toyota", Model: "Corolla", Year: 2020}
	car2 := Car{Brand: "Honda", Model: "Civic", Year: 2019}

	fmt.Println(car1.Details())
	fmt.Println(car2.Details())

	car1.ChangeYear(2021)
	fmt.Println(car1.Details())
}
