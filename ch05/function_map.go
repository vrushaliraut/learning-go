package main

import "fmt"

// Define functions with matching signatures
func addition1(i, j int) int { return i + j }
func sub(i, j int) int       { return i - j }
func mul(i, j int) int       { return i * j }
func divide(i, j int) int    { return i / j }

// function_map
// Create a map of function
func main() {
	var opMap = map[string]func(int, int) int{
		"+": addition1,
		"-": sub,
		"*": mul,
		"/": divide,
	}

	// Execute via loop
	op := "+" // assume come from user input
	p1, p2 := 2, 3

	// lookup the function
	opFunc, ok := opMap[op]

	if ok {
		// call the function found in a map
		result := opFunc(p1, p2)
		fmt.Println(result) // prints 5
	} else {
		fmt.Println("Function not found in map")
	}
}
