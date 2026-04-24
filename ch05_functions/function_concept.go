package main

import "fmt"

func division(num int, denom int) int {
	if denom == 0 {
		return 0
	}
	return num / denom
}

func add1(x int, y int) int {
	return x + y
}

// concise parameter style
func register(first string, last string, age int, city string, state string) int {
	return 0
}

// This tests grouping od adjacent types. Note that age breaks the chain of strings

func register2(first, last string, age int, city, state string) {
}

// The Void Function
// Write a function names log that accepts a string message, prints it to the console, and returns nothing

// This tests the syntax for functions with no return value
func log(message string) { // No return type specified
	fmt.Println(message)
	// No 'return' keyword needed at the end
}

// The Blank Identifier (_)
// In go you can implicitly ignore all values too --
