package main

import "fmt"

func main() {
	fmt.Println("=== Zero Values ===\n")

	// Zero value for string
	var user string
	fmt.Printf("String zero value: '%s' (empty)\n", user)

	// Zero value for int
	var count int
	fmt.Printf("Int zero value: %v\n", count)

	// Zero value for float
	var price float64
	fmt.Printf("Float64 zero value: %v\n", price)

	// Zero value for bool
	var active bool
	fmt.Printf("Bool zero value: %v\n", active)

	// Zero value for rune
	var initial rune
	fmt.Printf("Rune zero value: %d\n", initial)

	// Zero value for pointer
	var ptr *int
	fmt.Printf("Pointer zero value: %v\n", ptr)

	fmt.Println("\n=== Using Zero Values ===")
	var total int // 0
	total += 5
	total += 10
	fmt.Println("Total (initialized to 0):", total)

	var message string // ""
	message += "Hello"
	message += " "
	message += "World"
	fmt.Println("Message (initialized to ''):", message)

	var enabled bool // false
	if !enabled {
		fmt.Println("Feature is disabled by default")
	}
}

