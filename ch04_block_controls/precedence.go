package main

import "fmt"

/*
Loop over the slice created in Problem 1. For each value, apply the following rules:

 - If the value is divisible by 2 and 3, print "Six!". Don't print anything else.
 - If the value is divisible by 2, print "Two!".
 - If the value is divisible by 3, print "Three!".
 - Otherwise, print "Never mind".

Solution 2
Why: This tests control flow logic. Specifically, it tests handling overlapping conditions
(divisibility by 6 implies divisibility by 2 and 3). You must check the most specific condition first.

How: A for-range loop is the best choice for iterating over the slice.
An if-else if chain or a switch statement can handle the logic.

*/

func precedence() {

	nums := []int{6, 2, 3, 7}

	for _, v := range nums {
		// check the most specific condition first
		switch {
		case v%2 == 0 && v%3 == 0:
			fmt.Println("Six")
		case v%2 == 0:
			fmt.Println("Two")
		case v%3 == 0:
			fmt.Println("Three")
		default:
			fmt.Println("Never mind")

		}
	}
}
