package main

import "fmt"

func makeInSlice() {
	// make function has 2 forms for slices
	// make(type, length)
	// make(type, length, capacity)

	// the elements x[0] through x[4] are valid and are all initialized to the zero value for their type
	x := make([]int, 5) // type length
	fmt.Println(x)

	// create an int slice with a length of 5 and a capacity of 10
	y := make([]int, 5, 10)
	fmt.Println(len(y))
	fmt.Println(cap(y))

	// Create a slice of strings with length 3
	sliceOfString := make([]string, 3)
	fmt.Println(sliceOfString)

	// Create a slice with 5 elements (all 0)
	k := make([]int, 5)
	// x is [0 0 0 0 0], len=5, cap=5

	// Append 10
	k = append(k, 10)
	fmt.Println(len(k)) // length = 6 can capacity was 5 so append will increase the capacity by 25%  which is double
	fmt.Println(cap(k))

	makeAndAppendPattern()

	clearBuiltInFunction()
}

// The clear Built-in Function
func clearBuiltInFunction() {

	fmt.Print("Clearing built-in function...")
	s := []string{"first", "second", "third"}
	fmt.Println("printing slice s and its length :: ", s, len(s))

	clear(s)

	fmt.Println(s, len(s))

	// What happens when you clear a slice of integers? Write a program that creates an []int slice with the values {10, 20, 30},
	// clears it, and then prints the slice and its length.

	num := []int{10, 20, 30}
	fmt.Println("printing numbers :: ", num, len(num))

	clear(num)
	fmt.Println(num, len(num)) // prints [0,0,0] 3

	// Write a program that creates a []bool slice with the values {true, true, true}, clears it, and then prints the slice and its length.
	flags := []bool{true, true, true}
	fmt.Println(flags, len(flags)) // Prints [true true true] 3

	clear(flags)

	fmt.Println(flags, len(flags)) // Prints [false false false] 3

	// What is the difference in len and cap between a slice that has been cleared and a newly declared nil slice?
	// The primary goal is to minimize the number of times the slice needs to grow, as growing is an expensive operation.

	// The Nil Slice
	// var data []int

	// nil vs empty slice literal
	// var x = []int{}
	// This creates a slice that is not nil but still has a length and capacity of 0.
	
}

func makeAndAppendPattern() {
	// Problem 1: The make and append Pattern
	// Create a []byte slice with a length of 0 and a capacity of 5. Append the bytes 0x01 and 0x02.
	// Print the final slice, its length, and its capacity.

	// This is the most common and efficient pattern for building a slice when you have an estimate of the final size.
	// You pre-allocate the capacity, then use append to fill it.

	byteSlice := make([]byte, 0, 5)     // 0,0,0,0,0
	byteSlice = append(byteSlice, 0x01) // 0,0,0,0,0, 1
	byteSlice = append(byteSlice, 0x02) //  0,0,0,0,0, 1, 2

	fmt.Print("The make and append Pattern :: ")
	fmt.Println(byteSlice)
	fmt.Println("byteSlice length ::", len(byteSlice))   // 2
	fmt.Println("byteSlice capacity ::", cap(byteSlice)) // 5

	// panic: runtime error: index out of range [0] with length 0
	/* k := make([]int, 0, 5)  // make(type, len, cap)
	   k[0] = 5
	   fmt.Println(len(k))*/
	// Why: The code will panic (crash) at runtime. The len of the slice is 0. Valid indices are 0 up to len-1.
	// Since len is 0, there are no valid indices to write to.

	// You need to store 10,000 log lines. What is the efficient way to declare the slice before you start appending in a loop?
	// (e.g., for ... { lines = append(lines, aNewLine) })

	// inefficient way
	/*var lines []string
	for i := 0; i < 1000; i++ {
		lines = append(lines, string(i))
	}*/

	// most efficient way is
	// lines := make([]string, 0, 1000)

	fmt.Print("The Runtime Panic (Cap < Len):: ")
	l := 10
	c := 15 // if i use capacity as 5
	x := make([]int, l, c)

	fmt.Println(x)

}
