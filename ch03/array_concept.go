package main

import "fmt"

func arrayConcepts() {
	// Array Declaration
	// create an array of 3 integers all set to 0
	var x [3]int
	fmt.Println(x)

	// Array Literal
	var arrayLiteral = [3]int{10, 20, 30}
	fmt.Println("arrayLiteral ::", arrayLiteral)

	// Ellipsis

	var ellipsisExample = [...]int{10, 20, 30}
	fmt.Println("ellipsisExample ::", ellipsisExample)

	// Sparse Array Literal
	var sparseArrayLiteral = [7]int{0: 1, 5: 10, 10}
	fmt.Println("sparseArrayLiteral ::", sparseArrayLiteral)
	fmt.Println("sparseArrayLiteral length ::", len(sparseArrayLiteral))

	fmt.Println("sparseArrayLiteral length ::", sparseArrayLiteral[2])

	// Multidimensional arrays
	var multiArray [2][3]int // An array of 2 elements, where each element is an array of 3 ints

	fmt.Println("multiArray ::", multiArray)

	// This brings us to the main reason arrays are rarely used. In Go, the size of the array is part of its type.

	zeroValueDeclaration()

	// Array Literal and ... syntax
	// Declare an array rgb using an array literal with the string values "red", "green", and "blue".
	// Use the ... syntax so you don't have to specify the element count. Print the array.

	var rgb = [...]string{"red", "green", "blue"}
	fmt.Println("rgb ::", rgb)

	// Sparse array
	//	Declare an array grades to hold 8 int values.
	//	Use a sparse array literal to set the 1st element (at index 0) to 100 and the 5th element (at index 4) to 90.
	//	Print the resulting array.

	// Initializes a [8]int array
	// All unspecified indices (1,2,3,5,6,7) are set to 0

	var grades = [8]int{0: 100, 4: 90}
	fmt.Println("grades ::", grades)

	// the type mismatch

	/*var a = [3]int{1, 2, 3}
	var b = [4]int{1, 2, 3, 4}
	*/
	// Try to check for equality
	// fmt.Println("equality ::", a == b) fail to compile

	// Compile time bound check
	var boundCheckArray = [3]int{10, 20, 30}

	// try to access index that doesn't exists
	fmt.Println(boundCheckArray[2])
}

func zeroValueDeclaration() {
	// var name [size]type
	// Declare an array of 5 ints
	var scores [5]int

	// All elements will be 0
	fmt.Println(scores)

	// the length function
	fmt.Println(len(scores))
}
