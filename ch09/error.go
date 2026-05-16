package main

import (
	"errors"
	"fmt"
)

type error interface {
	Error() string
}

func divide(numerator, denominator int) (int, int, error) {
	if denominator == 0 {
		// Lower case no trailing punctuation
		return 0, 0, errors.New("cannot divide by zero")
	}

	return numerator / denominator, numerator % denominator, nil
}

func main() {
	remainder, mod, err := divide(10, 0)

	if err != nil {
		fmt.Println("Math failed:", err)
		return
	}

	fmt.Println(remainder, mod)
}
