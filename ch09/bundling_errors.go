package main

import (
	"errors"
	"fmt"
)

// Go dictate returning singular error interface - not an []error slice
// Need a way to bundle a slice of errors into one cohesive object.
// -> Two Built in ways to achieve this - Dynamic Bundling (errors.join)
// and Static bundling - multiple %w verbs

type Person struct {
	FirstName, LastName string
	Age                 int
}

func ValidatePerson(p Person) error {
	var errs []error // Collect failures here

	if p.FirstName == "" {
		errs = append(errs, errors.New("FirstName cannot be empty"))
	}
	if p.LastName == "" {
		errs = append(errs, errors.New("LastName cannot be empty"))
	}
	if p.Age < 0 {
		errs = append(errs, errors.New("Age cannot be negative"))
	}

	// Join returns nil if len(errs) == 0, otherwise it returns the bundled tree
	return errors.Join()
}

func main() {
	// Static Bundling - multiple %w

	err1 := errors.New("database timeout")
	err2 := errors.New("cache unreachable")

	// Bundles both errors under a new parent error
	err := fmt.Errorf("system boot failed: %w, and %w", err1, err2)
	fmt.Println(err)
}
