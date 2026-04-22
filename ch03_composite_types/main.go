package main

import (
	"fmt"
)

func main() {
	fmt.Println("Chapter 03: Composite Types")

	// Example: Slice
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println("Numbers:", numbers)

	// Example: Map
	person := map[string]string{
		"name": "Alice",
		"city": "New York",
	}
	fmt.Println("Person:", person)
}

