package main

import "fmt"

// anonymous function
// Define the anonymous function and assign it to f
func anonymous() {
	f := func(j int) {
		fmt.Println("printing", j, "from inside of an anonymous function")
	}

	// call 'f' like any other function
	for i := 1; i < 10; i++ {
		f(i)
	}

	// Immediate Invocation
	for j := 1; j < 5; j++ {
		// define and call function immediately
		func(j int) {
			fmt.Println("printing", j, "from inside of an anonymous function")
		}(j)
	}
}
