package main

import "fmt"

func main() {
	// Declared 'enabled' without assigning value
	var enabled bool

	// Print its value - it will print 'false' because that is the zero value for bool
	fmt.Println("Boolean default (zero value):", enabled)

	var flag bool = true
	fmt.Println("Flag:", flag)
	fmt.Println("Not Flag:", !flag)
}
