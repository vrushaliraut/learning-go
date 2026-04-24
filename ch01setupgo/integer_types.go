package main

import "fmt"

func main() {
	fmt.Println("=== Integer Types ===\n")

	var packetCount int32
	// No value is assigned so it defaults to its zero value
	fmt.Println("int32 zero value:", packetCount)

	var i int = 10 // 'int' type (size depends on platform)
	fmt.Println("int value:", i)

	var i64 int64 = 20 // 'int64' type (always 64 bits)
	fmt.Println("int64 value:", i64)

	// Type conversion required
	i = int(i64)
	fmt.Println("After conversion int(i64):", i)

	fmt.Println("\n=== Integer Operations ===")
	x := 10
	y := 3
	fmt.Println("10 + 3 =", x+y)   // 13
	fmt.Println("10 * 3 =", x*y)   // 30
	fmt.Println("10 / 3 =", x/y)   // 3
	fmt.Println("10 % 3 =", x%y)   // 1

	fmt.Println("\n=== Division Behavior ===")
	a := 10
	b := 3
	fmt.Println("10 / 3 =", a/b) // Prints 3

	c := 5
	d := 2
	fmt.Println("5 / 2 =", c/d) // Prints 2

	e := -5
	f := 2
	fmt.Println("-5 / 2 =", e/f) // Prints -2 (truncates toward zero)

	fmt.Println("\n=== Unsigned Integer Types ===")
	var u uint8 = 100
	fmt.Printf("uint8 value: %d, Type: %T\n", u, u)

	var b1 byte = 255
	fmt.Printf("byte (uint8 alias) value: %d, Type: %T\n", b1, b1)
}

