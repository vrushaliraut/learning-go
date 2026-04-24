package main

import (
	"fmt"
)

// Define struct with all potential parameters
type MyFuncOpts struct {
	FirstName string
	LastName  string
	Age       int
}

func MyFuncError(opts MyFuncOpts) {
	if opts.LastName == " " {
		// return fmt.Errorf("LastName is required")
	}
	// ...
	// return nil
}

func MyFuncLastnameRequired(opts MyFuncOpts) error {
	if opts.LastName == "" {
		return fmt.Errorf("LastName is required") // Errorf
	}
	// ...
	return nil
}

// Accept the struct as a single parameter
func MyFunc(opts MyFuncOpts) {
	// Set Default Age to be 18
	// Why: This addresses the main weakness of the struct pattern: distinguishing "user set 0" from "user didn't set it".
	if opts.Age == 0 {
		opts.Age = 18
	}
	fmt.Printf("User: %s %s, Age:%d\n", opts.FirstName, opts.LastName, opts.Age)
}

// Distinguishing Zero from Unset
/*
type MyFuncOpts struct {
	Age *int
}
*/

/*
// In the function
if opts.Age == nil{
	// User ommited it; use default
}else{
	// User provided a value even 0
	use(*opts.Age)
}
*/
