package main

import (
	"fmt"
)

// const & iota
const (
	Low = iota
	Medium
	High
)

// swap string fn
func swapString(a, b string) (string, string) {
	return b, a
}

func main() {

	first, second := swapString("Hello", "World")
	fmt.Println(first, second)

	// blank identifier
	_, last := swapString("Hello", "World")
	fmt.Println(last)

	arr := [5]string{"Hello", "World"}
	fmt.Println(len(arr), cap(arr))

	arr2 := []string{"Hello", "World", "Go", "Lang"}
	arr2 = append(arr2, "New")
	fmt.Println(len(arr2), cap(arr2))

	arr3 := [...]string{"Hello", "World", "Go", "Lang"}
	fmt.Println(len(arr3), cap(arr3))

	// copy the array 3 then append work, and arr3 has fixed size so that can not append
	slice3 := arr3[:]
	slice3 = append(slice3, "New")
	fmt.Println(slice3)
}
