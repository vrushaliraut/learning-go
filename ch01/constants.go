package main

import "fmt"

func main() {
	fmt.Println("=== Constants ===\n")

	// Package level typed constant
	const x int64 = 100
	fmt.Printf("Typed constant x: %v, Type: %T\n", x, x)

	// Untyped constant
	const y = 200
	fmt.Printf("Untyped constant y: %v, Type: %T\n", y, y)

	// Constant block with typed constants
	const (
		idKey   = "id"
		nameKey = "name"
	)
	fmt.Printf("idKey: %v, nameKey: %v\n", idKey, nameKey)

	// The value is figured out at compile time
	const z = 20 * 10
	fmt.Printf("Computed constant z (20 * 10): %v\n", z)

	fmt.Println("\n=== Constant vs Variable ===")

	// Variables must have explicit type conversion
	var intVar int = 10
	var floatVar float64 = 10 // This works: untyped literal promoted
	fmt.Printf("intVar: %v (type %T)\n", intVar, intVar)
	fmt.Printf("floatVar: %v (type %T)\n", floatVar, floatVar)

	// Constants are more flexible with type coercion
	const intConst int = 10
	const floatConst = 10.5
	fmt.Printf("intConst: %v (type %T)\n", intConst, intConst)
	fmt.Printf("floatConst: %v (type %T)\n", floatConst, floatConst)

	fmt.Println("\n=== Constant Rules ===")
	fmt.Println("- Constants cannot be assigned with variables")
	fmt.Println("- Constants cannot be modified after declaration")
	fmt.Println("- Only basic types: numbers, strings, booleans")
	// const bad = intVar  // This would fail - cannot use variable as constant value
}
