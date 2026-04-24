package main

import "fmt"

func main() {
	fmt.Println("=== Type Inference ===\n")

	// The compiler infers the type from the right-hand side
	var city = "San Francisco"      // Type: string
	var population = 875000         // Type: int
	var avgTemp = 61.5              // Type: float64
	var isCapital = false           // Type: bool

	fmt.Printf("City: %s (type: %T)\n", city, city)
	fmt.Printf("Population: %d (type: %T)\n", population, population)
	fmt.Printf("Avg Temp: %.1f (type: %T)\n", avgTemp, avgTemp)
	fmt.Printf("Is Capital: %v (type: %T)\n", isCapital, isCapital)

	fmt.Println("\n=== Short Variable Declaration ===")
	country := "USA"
	year := 2024
	density := 92.1
	developed := true

	fmt.Printf("Country: %s (type: %T)\n", country, country)
	fmt.Printf("Year: %d (type: %T)\n", year, year)
	fmt.Printf("Density: %.1f (type: %T)\n", density, density)
	fmt.Printf("Developed: %v (type: %T)\n", developed, developed)

	fmt.Println("\n=== Reassignment ===")
	value := 100
	fmt.Printf("Initial value: %d (type: %T)\n", value, value)

	value = 200
	fmt.Printf("After reassignment: %d (type: %T)\n", value, value)

	fmt.Println("\n=== Multiple Declarations ===")
	first, second, third := 1, "two", 3.0
	fmt.Printf("first: %v (type: %T)\n", first, first)
	fmt.Printf("second: %s (type: %T)\n", second, second)
	fmt.Printf("third: %.1f (type: %T)\n", third, third)
}

