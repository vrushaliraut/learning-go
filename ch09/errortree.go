package main

import (
	"errors"
	"fmt"
)

func main() {
	err := errors.New("database offline")

	fmt.Println(err)

	dbErr := errors.New("database offline")

	// %v copies only text
	err1 := fmt.Errorf("login failed: %v", dbErr)
	fmt.Println(err1)

	// %w preserves the origin error object

	err2 := fmt.Errorf("login failed: %w", dbErr)
	fmt.Println(err2)
	fmt.Println(errors.Is(err, dbErr))

	// wrapping and unwrapping
	rootErr := errors.New("database offline")
	wrappedErr := fmt.Errorf("login failed: %w", rootErr)
	fmt.Println(wrappedErr)
	fmt.Println()

	// errors.Unwrap() removes one layer.
	unWrapped := errors.Unwrap(wrappedErr)
	fmt.Println(unWrapped)

	HandleLogin()
}

// Handler - Service - Repository

var ErrDatabaseOffline = errors.New("database offline")

func getUserFromDB() error {
	return ErrDatabaseOffline
}

// Service layer

func Authenticate() error {
	err := getUserFromDB()
	if err != nil {
		return fmt.Errorf(
			"authentication process aborted: %w",
			err,
		)
	}
	return nil
}

func HandleLogin() {
	err := Authenticate()

	if err != nil {
		fmt.Printf("HTTP 500: %v\n", err)
	}
}
