package main

import "fmt"

// Package-level variables (global)
var (
	host    = "localhost"
	port    int
	enabled = true
)

func main() {
	fmt.Println("=== Variable Declaration Patterns ===\n")

	fmt.Println("=== Explicit Type Declaration ===")
	var name string = "Go"
	var version int = 1
	var active bool = true
	fmt.Printf("Language: %s, Version: %d, Active: %v\n", name, version, active)

	fmt.Println("\n=== Type Inference ===")
	var city = "New York"     // Type inferred as string
	var count = 42            // Type inferred as int
	var rating = 4.5          // Type inferred as float64
	fmt.Printf("City: %s, Count: %d, Rating: %.1f\n", city, count, rating)

	fmt.Println("\n=== Short Declaration ===")
	greeting := "Hello"       // Type inference with :=
	age := 30                 // Type inference with :=
	temperature := 98.6       // Type inference with :=
	fmt.Printf("Greeting: %s, Age: %d, Temp: %.1f\n", greeting, age, temperature)

	fmt.Println("\n=== Multiple Declarations ===")
	var x, y, z int = 1, 2, 3
	fmt.Printf("x=%d, y=%d, z=%d\n", x, y, z)

	// Multiple types
	var firstName, lastName string = "John", "Doe"
	var userAge int = 25
	fmt.Printf("Name: %s %s, Age: %d\n", firstName, lastName, userAge)

	// Short declaration
	a, b := 10, 20
	fmt.Printf("a=%d, b=%d\n", a, b)

	fmt.Println("\n=== Variable Block ===")
	var (
		product  = "Laptop"
		price    = 999.99
		quantity = 5
	)
	fmt.Printf("Product: %s, Price: $%.2f, Quantity: %d\n", product, price, quantity)

	fmt.Println("\n=== Package-Level Variables ===")
	fmt.Printf("Package-level host: %s\n", host)
	fmt.Printf("Package-level port: %d (zero value)\n", port)
	fmt.Printf("Package-level enabled: %v\n", enabled)
}

