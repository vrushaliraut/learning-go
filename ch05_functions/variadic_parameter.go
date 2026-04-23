package main

import "fmt"

// 'vals' is a slice of integers ([]int) inside the function

func addTo(base int, vals ...int) []int {
	// We can use 'vals' just like any other slice (len, range, etc)

	out := make([]int, 0, len(vals))
	for _, v := range vals {
		out = append(out, base+v)
	}
	return out
}

func variadicFunctionConcept() {

	// If you have slice you cannot pass it directly because the function expects int values, not []int.
	// You must "unPack" or spread the slice using three dots (...) after the slice variable

	//a := []int{4, 3}

	// Unpack the slice 'a' into individual arguments
	//fmt.Println(3, a...)

	// Works with literals too
	fmt.Println(addTo(3, []int{1, 2, 3, 4, 5}...))

	// The sum function
	fmt.Println(sum())

	// passing mixed types
	printAll()

	// Function chaining example
	add(getNums())
}

func printAll(vals ...interface{}) {
	for _, v := range vals {
		fmt.Println(v)
	}
}

func sum(nums ...int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

// Function chaining

func add(a, b int) {
	fmt.Println(a + b)
}

func getNums() (int, int) {
	return 10, 20
}
