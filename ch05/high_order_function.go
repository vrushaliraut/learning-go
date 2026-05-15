package main

import "fmt"

// Return Functions from functions

// In addition to passing functions as parameters, Go allows you to return a closure from a function.
// This pattern effectively allows you to create "function factories",
// functions that configure and generate new, specialized functions.

func main() {
	twoBase := makeMult(2)
	threeBase := makeMult(3)

	for i := 0; i < 3; i++ {
		fmt.Println(twoBase(i), threeBase(i))
	}
}

// makeMult takes int and return a function : func(int) int
func makeMult(base int) func(int) int {
	// return anonymous function closure

	return func(factor int) int {
		return base * factor
	}
}
