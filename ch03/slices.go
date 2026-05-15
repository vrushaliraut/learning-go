package main

import "fmt"

import "slices" // import new package

func slicesConcept() {

	// Create a Slice of 3 ints

	var x = []int{10, 20, 30}
	fmt.Println("Slices Example ::", x)

	// Using [...] makes an array. Using [] makes a slice

	// [1, 0, 0, 0, 0, 4, 6, 0, 0, 0, 100, 15]
	var createSliceWithIndex = []int{1, 5: 4, 6, 10: 100, 15}
	fmt.Println("Slices createSliceWithIndex ::", createSliceWithIndex)

	// simulate multidimensional slices by making a slice of slices :
	var multidimensionalSlice [][]int
	fmt.Println("Slices Multidimensional ::", multidimensionalSlice)
	/*
		arrayVsSlicesEx()

		sparseSliceEx()

		sliceOfSlices()

		sliceComparison()

		sliceEqual()

		lenFunctionExample()

		appendSlices()*/

	capacitySlice()

}

func capacitySlice() {
	// Length - The number of elements currently in a slice. This is returned by built-in len() function.
	// Capacity - The number of consecutive memory locations reserved in the slice's underlying array.
	// This is total space available, which can be larger than the length

	// A slice has a len of 10 and a cap of 32. What does this mean in practical terms?
	// It holds 10 element - 0 - 9 and have reserved space for 32 elements in total
	// This is an efficient state - You can append 22 more elements to this slice very quickly, without causing any
	// new memory allocations.

	// What are the three "expensive" operations the Go runtime performs when len == cap and you append a new element?
	// memory allocation - find and  allocate new large block of memory
	// memory copy - Runtime must copy every single element from old array to new
	// and Garbage collection -  Old array unreferenced, the gb will have to work to free its memory

	//
	fmt.Print("You can also pass array to cap function, cap return same value as len for array.\n")
	var x []int
	fmt.Println(x, len(x), cap(x))

	x = append(x, 10)
	fmt.Println(x, len(x), cap(x))

	x = append(x, 20)
	fmt.Println(x, len(x), cap(x))

	x = append(x, 30)
	fmt.Println(x, len(x), cap(x))

	x = append(x, 40)
	fmt.Println(x, len(x), cap(x))

	x = append(x, 50)
	fmt.Println(x, len(x), cap(x))

	/*
		Analysis of the output:
		[] 0 0: The nil slice (zero value) has a length and capacity of 0.
		[10] 1 1: append allocates a new array of size 1. len is 1, cap is 1.
		[10 20] 2 2: len == cap, so append allocates a new, larger array (size 2).
		[10 20 30] 3 4: len == cap, so append allocates a new, larger array.
		It doubles the old capacity (2) to 4. len is 3, cap is 4.
		[10 20 30 40] 4 4: len < cap (3 < 4), so no allocation needed. append just adds the element.
		len is 4, cap is 4.
		[10 20 30 40 50] 5 8: len == cap (4 == 4), so append allocates a new, larger array.
		It doubles the old capacity (4) to 8.
	*/

	// len and cap of nil
	// What does len(x) and cap(x) return if you declare var x []int?
	// ans - Why: The zero value for a slice is nil. A nil slice has no elements and no backing array.
	//How: Both len(x) and cap(x) will return 0. This is safe to check and is a common way to see if a slice is uninitialized.

}

func appendSlices() {
	var x []int
	x = append(x, 1)
	fmt.Println("The value of appended empty slice is now :: ", x)

	// append to the slice that already has an elements
	var y = []int{1, 2, 3}
	y = append(y, 4)
	fmt.Println("The value of appended new element to slice is now :: ", y)

	s1 := []int{1, 2}
	s2 := []int{3, 4}

	// compile time error
	//- The append function expects a series of int values, but you are trying to pass it a single []int (slice of int) value
	//s1 = append(s1, s2)
	s1 = append(s1, s2...)

	var n []int
	//append(n, 10) // ERROR: append(x, 10) evaluated but not used
	// go is call by value, it creates a copy and returns potentially new copy
	n = append(n, 20)
	fmt.Println("reassign is mandatory :: ", n)

	// Write a main function that declares a slice x := []int{1} and then tries to append 2 to it without re-assigning the result.
	//What is the exact compiler error?

	/*x := []int{1}
	append(x, 2)*/
}

func lenFunctionExample() {
	// len function works for both arrays as well as slices
	k := []int{10, 20, 30}
	fmt.Println(len(k))

	// nil slice
	a := []int{}
	fmt.Println("passing  nil to slice return 0", len(a))

	s := []string{"a", "b", "c", "d"}
	fmt.Println("length of string slice:", len(s))

	// nil vs empty slice
	var x []int
	var y = []int{}

	fmt.Printf("Length of nil slice x: %d\n", len(x))
	fmt.Printf("Length of empty slice y: %d\n", len(y))

	// You can also see the nil difference
	fmt.Println("x is nil:", x == nil) // true
	fmt.Println("y is nil:", y == nil) // false

	// The len function is built-in but is not universal.
	// It only works on specific types (arrays, slices, strings, maps, channels). An int is not one of those types.
	/*i := 1024
	fmt.Println(len(i))*/

}

// Declare two slices, s1 = []int{1, 2, 3} and s2 = []int{1, 2, 3}. Then, declare a third slice, s3 = []int{1, 2, 4}.
// Use the slices.Equal function to compare s1 to s2 and s1 to s3.
func sliceEqual() {
	s1 := []int{1, 2, 3}
	s2 := []int{1, 2, 3}
	s3 := []int{1, 2, 4}

	fmt.Println("slices equality s1 = []int{1, 2, 3} and s2 = []int{1, 2, 3} :: ", slices.Equal(s1, s2))
	fmt.Println("slices equality s1 = []int{1, 2, 3} and s3 = []int{1, 2, 4} :: ", slices.Equal(s1, s3))
}

// What is the "zero value" of a slice? Write a program that declares var s []string and then checks if it is nil.
func sliceComparison() {
	var x []int           // x is a nil slice
	fmt.Println(x == nil) // prints true

	var y = []int{} // y is an empty but not nil, slice
	fmt.Println(y == nil)

	//fmt.Println(x == y)

	a := []int{1, 2, 3}
	b := []int{1, 2, 3}
	c := []int{1, 2, 3, 4, 5}

	fmt.Println(slices.Equal(a, b)) // prints true
	fmt.Println(slices.Equal(b, c)) // prints false

	//z := []string{}

	// compare a slice of ints and a slice of strings
	//fmt.Println(slices.Equal(z, a))

	/*x1 := []int{1}
	y1 := []int{1}
	// Slices are not a comparable type in Go
	fmt.Println(x1 == y1) */

}

// Declare a 2D slice matrix to hold int values, initialized as a 2x2 matrix with values [[1, 2], [3, 4]].
func sliceOfSlices() {
	var sliceMatrix = [][]int{
		{1, 2},
		{3, 4},
	}

	fmt.Println(sliceMatrix)
}

// Declare a slice of string values named config with a total length of 5.
// Use a sparse literal to set the first element (index 0) to "host" and the last element (index 4) to "port".
func sparseSliceEx() {
	var config = []string{0: "host", 4: "port"}
	fmt.Println("SparseSlice ::", config)
	fmt.Println("SparseSlice length ::", len(config))
}

// Write code that declares arr as an array of 3 int values and slc as a slice of 3 int values, both initialized with 1, 2, 3.
func arrayVsSlicesEx() {
	// only difference is in declaration [...] vs []
	// ... will let know system count element and declare as an array vs [] will create a slice

	var arr = [...]int{1, 2, 3}
	fmt.Println("Slices Array ::", arr)

	var slc = []int{1, 2, 3}
	fmt.Println("Slices Slice ::", slc)
}
