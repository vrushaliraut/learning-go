package main

import (
	"fmt"
	"strings"
)

// Define rich error struct
type ValidationError struct {
	MissingFields []string
}

// Satisfy error interface
func (v ValidationError) Error() string {
	return fmt.Sprintf("validation failed, missing fields:%s",
		strings.Join(v.MissingFields, ", "))
}

func ValidateUser(payload map[string]string) error {
	var missing []string

	if payload["name"] == "" {
		missing = append(missing, "name")
	}
	if payload["email"] == "" {
		missing = append(missing, "email")
	}

	if len(missing) > 0 {
		// return custom error containing slice
		return ValidationError{MissingFields: missing}
	}

	return nil
}

func main() {
	user := map[string]string{
		"name":  "John Doe",
		"email": "",
	}

	err := ValidateUser(user)
	if err != nil {
		fmt.Println(err)

		// Type assertion to access MissingFields

		if validationErr, ok := err.(ValidationError); ok {
			fmt.Println("Missing Fields:")

			for _, field := range validationErr.MissingFields {
				fmt.Println("\t", field)
			}
		}
		return
	}
	fmt.Println("User OK")
}
