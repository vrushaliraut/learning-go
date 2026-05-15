package main

import "fmt"

// Declare function variable
// The default zero value for a function variable is nil.
//Attempting to call a nil function variable results in a runtime panic.

// 1. take one string 2. return one int

var myFuncVariable func(string) int

// assign and  calling

func f1(a string) int { return len(a) }
func f2(a string) int { return 42 } // dummy logic
// function_are_values
func main() {
	var myFuncVariable func(string) int

	// Assign f1
	myFuncVariable = f1

	fmt.Println(myFuncVariable("Hello")) // 5

	// reassign
	myFuncVariable = f2
	fmt.Println(myFuncVariable("hello"))
}
