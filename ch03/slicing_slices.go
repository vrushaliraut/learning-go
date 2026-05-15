package main

import "fmt"

func slicingSlices() {

	// slice[start:end]
	// A slice expression creates a new slice from an existing slice's backing array.
	// It’s written inside brackets and consists of a starting offset and an ending offset, separated by a colon (:).

	s := []string{"a", "b", "c", "d", "e", "f"}

	a := s[:2]  // [a,b] from index 0 up to, but not including 2
	b := s[1:]  // [b c d e f] from index 1 to end
	c := s[1:3] // [b c] from index 1 up to, but not including, 3
	d := s[:]   // [a b c d e f] - full copy of slice header

	fmt.Println("whole slice of string represent :: ", s)

	fmt.Println("a: ", a)
	fmt.Println("b: ", b)
	fmt.Println("c: ", c)
	fmt.Println("d: ", d)

	basicSlicing()

	slicingWithDefault()

	slicesSharedMemory()

	threePartSlice()
}

/*
- Problem 1: Calculate len and cap (Full Expression)
- Given x := []int{1, 2, 3, 4, 5}. You create a new slice y := x[1:3:4]. What are the len(y) and cap(y)?
*/
func threePartSlice() {
	/*
		Why: The rules for a three-part slice [start:end:max]
		are:
		len = end - start
		cap = max - start
	*/

	/*x := []int{1, 2, 3, 4, 5}
	y := x[1:3:4]

	fmt.Println("length of y :: ", y[])*/
	/*
		len = 3 - 1 = 2
		cap = 4 - 1 = 3
		The slice y will be [2, 3], with a len of 2 and a cap of 3.
	*/
}

func slicesSharedMemory() {
	x := []string{"a", "b", "c", "d"}
	y := x[:2] // y points to x's array, sees [a b]
	z := x[1:] // z points to x's array, sees [b c d]
	fmt.Print("slices Shared memory...!\n")

	fmt.Println("x:", x)
	// Modify the *shared* array
	x[1] = "y"
	y[0] = "x"
	z[1] = "z"

	fmt.Println("x:", x)
	fmt.Println("y:", y)
	fmt.Println("z:", z)
}

/*
Given letters := []string{"w", "x", "y", "z"}, use the default (empty) offsets to create two new slices:

head containing ["w", "x"]
tail containing ["y", "z"]
*/
func slicingWithDefault() {
	letters := []string{"w", "x", "y", "z"}
	head := letters[:2]
	tail := letters[2:]

	fmt.Println("slicingWithDefault :: \n", letters, head, tail)

	// The [:] Expression
	//	Declare a slice x using make([]int, 3, 5). Then, create a new slice y using y := x[:].
	//	Print the len and cap of both x and y.

	x := make([]int, 3, 5)
	y := x[:]

	fmt.Printf("x: len=%d, cap=%d\n", len(x), cap(x))
	fmt.Printf("y: len=%d, cap=%d\n", len(y), cap(y))
}

/*
Given nums := []int{10, 20, 30, 40, 50}, write the slice expressions to create two new slices:

firstTwo containing [10, 20]
lastThree containing [30, 40, 50]
*/
func basicSlicing() {
	nums := []int{10, 20, 30, 40, 50}

	fistTwo := nums[:2]
	lastThree := nums[2:5] // start at index 2 and ends at the end

	fmt.Println(nums, fistTwo, lastThree)
}
