package main

import "fmt"

func closures() {
	a := 20
	// This anonymous function is a closure
	f := func() {
		fmt.Println(a) // Read 'a' from outer scope
		a = 30         // modify 'a' from outer scope
	}

	f()
	fmt.Println(a) // 'a' has been changed

	// shadowing in closure
	shadowing()
}

func shadowing() {
	x := 20
	f := func() {
		fmt.Println(x) // Reads outer 'a' (20)

		// := create a NEW variable shadowing the outer one
		x := 30
		fmt.Println("Prints local 'a' (30): ", x) // Prints local 'a' (30)
	}
	f()
	fmt.Println("prints outer x still (20) : ", x) // prints outer x still 20
}
