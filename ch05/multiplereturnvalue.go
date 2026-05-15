package main

import (
	"errors"
	"fmt"
	"os"
)

func divAndRemainder(num, denom int) (int, int, error) {
	if denom == 0 {
		// Must return exactly three values
		return 0, 0, errors.New("cannot divide by zero")
	}

	// No parentheses needed around the return values
	return num / denom, num % denom, nil
}

// The return types must be enclosed in parenthesis - (int, int, error )
func main() {
	// when calling a function with multiple return values, you see multiple assignment on
	// the left-hand side of the := or = operator

	result, remainder, err := divAndRemainder(5, 2)

	// standard error checking
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(result, remainder) // Prints 2 1
}
