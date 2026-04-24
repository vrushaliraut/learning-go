package main

import "fmt"

func main() {
	fmt.Println("=== Type Conversion ===\n")

	fmt.Println("=== int to float64 conversion ===")
	var x int = 10
	var y float64 = 30.2

	// Convert x (int) to float64
	var sum1 float64 = float64(x) + y
	fmt.Printf("%.1f + %.1f = %.1f\n", float64(x), y, sum1)

	// Convert y (float64) to int
	var sum2 int = x + int(y)
	fmt.Printf("%d + %d = %d\n", x, int(y), sum2)

	fmt.Println("\n=== byte (uint8) to int conversion ===")
	var z int = 10
	var b byte = 100 // byte is an alias for uint8

	// We must convert 'b' to 'int' to add it to 'z'
	var sum3 int = z + int(b)
	fmt.Printf("%d + %d = %d\n", z, b, sum3)

	// Convert z to byte
	var sum4 byte = byte(z) + b
	fmt.Printf("%d + %d = %d\n", byte(z), b, sum4)

	fmt.Println("\n=== Literals and Variables (Type Coercion) ===")

	// Literals are flexible
	var f float64 = 20.5
	fmt.Printf("20.5 + 10 = %.1f\n", f+10) // 10 is promoted to 10.0

	// Variables are strict
	var intVar int = 10
	var floatVar float64 = 20.5

	// This would NOT work:
	// fmt.Println(floatVar + intVar) // Error: mismatched types

	// Must explicitly convert
	fmt.Printf("%.1f + %d (converted to float) = %.1f\n", floatVar, intVar, floatVar+float64(intVar))

	fmt.Println("\n=== Multiple Declarations with Type Coercion ===")
	port, service := 8080, "http"
	fmt.Printf("type of port: %T\n", port)
	fmt.Printf("type of service: %T\n", service)
}

