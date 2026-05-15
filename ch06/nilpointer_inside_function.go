package main

import "fmt"

// You Cannot Initialize a nil Pointer from Inside a Function
func main() {

	// If you pass a nil pointer variable to a function, the function receives a copy containing nil
	// If you point that copy to a new object, you are only changing the local copy.
	// The original variable in main remains nil.

	var f *int // f is nil
	failedUpdate(f)
	fmt.Println("Value of f after failedUpdate:", f) // prints: Value of f after failedUpdate: <nil>
}

func failedUpdate(f *int) {
	x := 10
	// We are pointing the local copy of f to a new object (x),
	// but this does not affect the original variable in main
	f = &x
	fmt.Println("Value of f inside failedUpdate:", *f)
	// prints: Value of f inside failedUpdate: 10
}
