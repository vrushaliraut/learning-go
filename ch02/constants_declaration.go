package main

import "fmt"

// Package-level constants
const AppName = "MyApp"
const Version = "1.0.0"

func main() {
	fmt.Println("=== Constants Declaration ===\n")

	fmt.Println("=== Typed Constants ===")
	const maxUsers int = 100
	const minAge int = 18
	const pi float64 = 3.14159

	fmt.Printf("maxUsers: %d\n", maxUsers)
	fmt.Printf("minAge: %d\n", minAge)
	fmt.Printf("pi: %.5f\n", pi)

	fmt.Println("\n=== Untyped Constants ===")
	const defaultPort = 8080
	const dbName = "myapp"
	const isProduction = true

	fmt.Printf("defaultPort: %v (type: %T)\n", defaultPort, defaultPort)
	fmt.Printf("dbName: %s (type: %T)\n", dbName, dbName)
	fmt.Printf("isProduction: %v (type: %T)\n", isProduction, isProduction)

	fmt.Println("\n=== Constants Block ===")
	const (
		Sunrise = 6
		Noon    = 12
		Sunset  = 18
		Midnight = 24
	)
	fmt.Printf("Sunrise: %d, Noon: %d, Sunset: %d, Midnight: %d\n", Sunrise, Noon, Sunset, Midnight)

	fmt.Println("\n=== Package-Level Constants ===")
	fmt.Printf("App Name: %s\n", AppName)
	fmt.Printf("Version: %s\n", Version)

	fmt.Println("\n=== Computed Constants ===")
	const hoursPerDay = 24
	const minutesPerHour = 60
	const minutesPerDay = hoursPerDay * minutesPerHour
	fmt.Printf("Minutes per day: %d\n", minutesPerDay)

	fmt.Println("\n=== Constants are Immutable ===")
	const message = "Hello, Go!"
	fmt.Printf("Message: %s\n", message)
	// message = "Goodbye" // This would cause a compile error
	fmt.Println("Constants cannot be reassigned")
}

