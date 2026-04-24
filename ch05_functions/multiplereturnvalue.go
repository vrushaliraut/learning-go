package main

import "errors"

func divAndRemainder(num, denom int) (int, int, error) {
	if denom == 0 {
		// Must return exactly three values
		return 0, 0, errors.New("cannot divide by zero")
	}

	// No parentheses needed around the return values
	return num / denom, num % denom, nil
}

// The return types must be enclosed in parenthesis - (int, int, error )
