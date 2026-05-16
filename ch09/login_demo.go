package main

import (
	"errors"
	"fmt"
)

func validateUser(email string, age int) error {
	if email == "" {
		return errors.New("Invalid email")
	}

	if age < 18 {
		return fmt.Errorf("provided age %d is below the legal limit of 18", age)
	}
	return nil
}
