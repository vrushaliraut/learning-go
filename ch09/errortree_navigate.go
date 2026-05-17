package main

import (
	"errors"
	"fmt"
	"os"
)

// error.is - searching  - find specific values
// error.is - extract specific types

func checkFile() error {
	_, err := os.Open("missing.txt")
	// We wrap sentinel error
	return fmt.Errorf("initialization failed: %w", err)
}

type StatusErr struct {
	Code int
}

func (s StatusErr) Error() string {
	return "status error"
}

func main() {
	err := checkFile()

	if errors.Is(err, os.ErrNotExist) {
		fmt.Println("file does not exist")
	}

	// Error.AS
	// A custom error, wrapped in a standard format error
	rootErr := StatusErr{
		Code: 404,
	}
	err = fmt.Errorf("root error: %w", rootErr)
	fmt.Println(err)
	// Declare a zero value variable of the TYPE we want to find
	var statusErr StatusErr

	// pass the address of that variables to errors.AS
	if errors.As(err, &statusErr) {
		fmt.Println("status error: %d \n", statusErr.Code)
	}
}
