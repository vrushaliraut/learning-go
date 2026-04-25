package main

import "fmt"

func main() {
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
	shadowingCheck()

	// the counter closure
	counter := createCounter()
	fmt.Println(counter()) // 1
	fmt.Println(counter()) // 2

}

var x = 10

func shadowingCheck() {
	func() {
		x = 20
		x := 5
		x = 50

		fmt.Println("inner x: ", x)
	}()
	fmt.Println("outer x: ", x)
}

func createCounter() func() int {
	count := 0 // 'count' is captured by the closure

	return func() int {
		count++
		return count
	}
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
