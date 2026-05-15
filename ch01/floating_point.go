package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("=== Floating Point Types ===\n")

	var f float32 = 123.456789
	fmt.Printf("float32 value: %v, Type: %T\n", f, f)

	var f64 float64 = 123.456789123456
	fmt.Printf("float64 value: %v, Type: %T\n", f64, f64)

	fmt.Println("\n=== Floating Point Operations ===")
	x := 10.5
	y := 3.2
	fmt.Println("10.5 + 3.2 =", x+y)
	fmt.Println("10.5 * 3.2 =", x*y)
	fmt.Println("10.5 / 3.2 =", x/y)

	fmt.Println("\n=== Epsilon Comparison (Floating Point Precision) ===")
	a := 0.1
	b := 0.2
	sum := a + b // 0.3000000004

	// Direct comparison fails due to precision
	if sum == 0.3 {
		fmt.Println("Direct comparison: The sum is 0.3")
	} else {
		fmt.Println("Direct comparison: The sum is NOT 0.3 (precision issue)")
	}

	// The Right way: Use epsilon
	const epsilon = 0.00001
	if math.Abs(sum-0.3) < epsilon {
		fmt.Println("Epsilon comparison: Equal (within tolerance)")
	} else {
		fmt.Println("Epsilon comparison: Not Equal")
	}

	fmt.Println("\n=== Special Float Values ===")
	posInf := math.Inf(1)
	negInf := math.Inf(-1)
	nan := math.NaN()

	fmt.Printf("Positive Infinity: %v\n", posInf)
	fmt.Printf("Negative Infinity: %v\n", negInf)
	fmt.Printf("NaN: %v\n", nan)
	fmt.Printf("NaN == NaN: %v (always false)\n", nan == nan)
	fmt.Printf("math.IsNaN(nan): %v\n", math.IsNaN(nan))
}

