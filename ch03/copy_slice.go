package main

import "fmt"

func copySlice() {
	x := []int{1, 2, 3, 4}

	y := make([]int, 4) // must create slice of size 4

	num := copy(y, x) // copy function take 2 param - destination source

	fmt.Println(y, num)

	// when destination is smaller
	fmt.Print("when destination is smaller :: ")
	x1 := []int{1, 2, 3, 4}
	y1 := make([]int, 2)
	num1 := copy(y1, x1)
	// y is [1 2], num is 2
	fmt.Println(y1, num1)

	// when source is smaller
	fmt.Print("when source is smaller :: ")
	x2 := []int{1, 2}
	y2 := make([]int, 4) // y is [0 0 0 0]
	num2 := copy(y2, x2)
	// y is [1 2 0 0], num is 2
	fmt.Println(y2, num2)

	proveIndependentCopies()

	copyByLimitedDestination()

	copyByLimitedSource()

	copySubSetOfSlice()

	overlappingSlice()

	convertingToAndFromArrays()

	convertArrayToSlice()

	slicesFromArraySharedMemory()

	convertSlicesToArray()

	proveIndependentCopy()

	convertingSubset()

	conversionRuleAndPanics()
}

func conversionRuleAndPanics() {
	// Array Size must be fixed
	xSlice := []int{5, 10, 15, 20}

	// invalid use of [...] array (outside a composite literal)
	// xArray := [...]int(xSlice)

	// This compiles, but will panic
	// We are asking for a 5-element array from a 4-element slice
	// panic: runtime error: cannot convert slice with length 4 to array or pointer to array with length 5
	panicArray := [5]int(xSlice)

	fmt.Println(panicArray)

	// Safe Subset
	// Request a [4]int from the slice
	a := [4]int(xSlice)

	fmt.Println(a)

}

// You have xSlice := []int{5, 10, 15, 20}.
// Write the code to create a new array named arr of size 3 that contains the first 3 elements of xSlice.
func convertingSubset() {
	xSlice := []int{5, 10, 15, 20}

	// This copies the first 3 elements [5,10,15]
	arr := [3]int(xSlice)

	fmt.Println("array contains the first 3 elements of xSlice :: ", arr)
}

// Declare s := []string{"a", "b"}. Convert this slice to a [2]string array named a.
// After the conversion, modify s[0] = "z". Print both s and a to prove they are independent.
func proveIndependentCopy() {
	s := []string{"a", "b"}

	// convert slice to array
	a := [2]string(s)

	// Modify original slice
	s[0] = "z"

	// print both - slice and a new array
	fmt.Print("Original slice :: ", s)
	fmt.Print("New Array :: ", a)
}

func convertSlicesToArray() {
	xSlice := []int{1, 2, 3, 4}

	// This *copies* the values from xSlice into a new array
	xArray := [4]int(xSlice)

	// Modify original slice
	xSlice[0] = 10

	// Ths slice is changed
	fmt.Println(xSlice)

	// The array which is a copy is NOT changed
	fmt.Println(xArray)
}

func slicesFromArraySharedMemory() {
	x := [4]int{5, 6, 7, 8} // x is an array
	y := x[:2]              // y is a slice viewing x
	z := x[2:]              // z is a slice viewing x

	// Modify the array
	x[0] = 10

	fmt.Println("x:", x)
	fmt.Println("y:", y)
	fmt.Println("z:", z)
}

func convertArrayToSlice() {
	// convert entire array into slice, use [:] , resulting slice will have len, cap equals to array size

	xArray := [4]int{5, 6, 7, 8}

	xSlice := xArray[:]
	fmt.Println(xSlice)

	// convert subset of an array

	x := [4]int{1, 2, 3, 4}
	y := x[:2] // 1,2
	z := x[2:] // 3,4

	fmt.Println(x, y, z)
}

/*
The copy function only accepts slice arguments.
However, you can use an array by taking a slice of it with the [:] operator.
This allows you to copy from an array to a slice, or to an array from a slice.
*/
func convertingToAndFromArrays() {
	x := []int{1, 2, 3, 4}
	y := [4]int{5, 6, 7, 8} // this is an array

	// create new slice of length 2
	z := make([]int, 2)

	// 1. Copy FROM array 'y' TO slice 'z'
	// We slice the array 'y' to make it a slice
	copy(z, y[:])
	fmt.Println("copying data from array to slice ::", z)

	// 2. Copy FROM slice 'x' TO array 'y'
	// We slice the array 'y' to make it a destination
	copy(y[:], x)
	fmt.Println("copying data from slice to array :: ", x)

	/*
		You have an array a := [5]int{1, 2, 3, 4, 5}. Create a slice s that is an independent copy of the array.
		Solution 2
		Why: The copy function's source argument must be a slice. You can create a slice "view" of an array by using the a[:] expression,
		which gives a slice of the entire array that can then be used as the source.
	*/

	a := [5]int{1, 2, 3, 4, 5}

	// Make a new slice of the same length
	s := make([]int, 5)

	// Copy from the array 'a' by slicing it
	copy(s, a[:])

	fmt.Println("copying data from array to slice :: ", s)
}

// The copy function is safe to use even when the source and destination slices overlap (point to the same underlying array).
func overlappingSlice() {
	x := []int{1, 2, 3, 4}
	num := copy(x[:3], x[1:])
	fmt.Println("overlappingSlice :: \n ", num, x)
}

// You can copy a subset of a slice by combining copy with a slice expression on the source.
func copySubSetOfSlice() {
	x := []int{1, 2, 3, 4}
	y := make([]int, 2)
	// We slice x[2:] to get a view [3, 4]
	// and copy that into y
	copy(y, x[2:])
	// y is now [3 4]
	fmt.Println("copySubSetofSlice :", y)
}

/*
- You have src := []int{10, 20}. You create dst := make([]int, 5).
- What will dst contain after copy(dst, src), and what integer will the copy function return?
*/
func copyByLimitedSource() {
	src := []int{10, 20}
	dst := make([]int, 5) // len=5, dst is [0 0 0 0 0]

	num := copy(dst, src)

	fmt.Println("dst:", dst)
	fmt.Println("num:", num)
}

/*
- You have src := []string{"a", "b", "c", "d"}. You create dst := make([]string, 3).
- What will dst contain after copy(dst, src), and what integer will the copy function return?
*/
func copyByLimitedDestination() {
	src := []string{"a", "b", "c", "d"}
	dst := make([]string, 3)
	num := copy(dst, src)
	fmt.Println("dst:", dst)
	fmt.Println("num:", num)

}

/*
Declare x := []int{1, 2, 3}. Create a new, independent slice y and copy the contents of x into it.
Then, change y[0] = 99 and print both x and y to prove they are separate.
*/
func proveIndependentCopies() {
	x := []int{1, 2, 3}
	y := make([]int, 3)

	copy(y, x)
	fmt.Print("prove Independent Copies ::")
	fmt.Println(y)

	y[0] = 99

	fmt.Print("prove Independent Copies after modifying::")
	fmt.Println(x, y)

}
