package main

import "fmt"

// functions are values, you can declare package-level variables and assign anonymous functions to them.

var (
	addition2     = func(i, j int) int { return i + j }
	substraction2 = func(i, j int) int { return i - j }
)

// Unlike standard function declarations (which are constant),
// a package-level variable holding a function can be modified at runtime.
// package_level_anonymous_function
func main() {
	x := addition2(3, 6)
	fmt.Println(x) // print 9

	changeAdditionFunction()

	y := addition2(3, 6) // prints 15
	fmt.Println(y)
}

func changeAdditionFunction() {
	// We are replacing the logic of 'add' dynamically!
	addition2 = func(i, j int) int {
		return i + j + j
	}
}
