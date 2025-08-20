package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	data := "Hola, Let's learn Go with Soumik!"

	/**
	 * write file if not exists
	 * overwrite if exists
	 */

	err1 := os.WriteFile("example.txt", []byte(data), 0644)
	if err1 != nil {
		fmt.Println("Error writing file:", err1)
		return
	}

	fmt.Println("File written successfully")

	// read a file
	content, err2 := os.ReadFile("example.txt")
	if err2 != nil {
		fmt.Println("Error reading file:", err2)
		return
	}

	fmt.Println("File content:", string(content))

	// append file
	f, err3 := os.OpenFile("example.txt", os.O_APPEND|os.O_WRONLY, 0644)
	if err3 != nil {
		fmt.Println("Error opening file for appending:", err3)
		return
	}
	defer f.Close()

	if _, err4 := f.WriteString(" New line appended successfully\n"); err4 != nil {
		fmt.Println("Error appending to file:", err4)
		return
	}

	fmt.Println("File appended successfully")

	// json marshal
	p1 := Person{Name: "Soumik", Age: 28}

	// struct to json
	jsonData, err5 := json.Marshal(p1)
	if err5 != nil {
		fmt.Println("Error marshaling struct to JSON:", err5)
		return
	}

	fmt.Println("JSON data:", string(jsonData))

	jsonStr := string(jsonData)
	var p2 Person
	err6 := json.Unmarshal([]byte(jsonStr), &p2)
	if err6 != nil {
		fmt.Println("Error unmarshaling JSON to struct:", err6)
		return
	}

	fmt.Printf("Person: %+v\n", p2)
}
