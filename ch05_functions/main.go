package main

import (
	"fmt"
	"os"
)

/*
os.Open: We try to open the file. If this fails (e.g., file doesn't exist), we return the error immediately.
defer f.Close(): The moment we have a valid file handle f, we defer its closure.
This ensures that even if f.Stat() fails later, the file handle won't leak.
f.Stat(): This is the standard way to get metadata about a file, including its size in bytes (info.Size()).
*/

func main() {
	fmt.Println("Chapter 05: functions")

	//calculator()

	// create dummy file for testing
	err := os.WriteFile("test.txt", []byte("Hello World"), 0644)

	if err != nil {
		fmt.Println(err)
		return
	}

	size, err := filelength_with_defer("test.txt")

	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("File size:", size, "bytes")
	}

	// Cleanup the test file
	os.Remove("test.txt")

	declaringAndCallingFunction()

	namedAndOptionalParam()

	// Passing Individual Argument
	fmt.Println(addTo(3))             // vals is []int{}, output:[]
	fmt.Println(addTo(3, 2))          // vals is []int{2}, output:[5]
	fmt.Println(addTo(3, 2, 4, 6, 8)) // vals is []int{2,4,6,8}, output:[5,7,9,11]

	// when calling a function with multiple return values, you see multiple assignment on
	// the left-hand side of the := or = operator

	result, remainder, err := divAndRemainder(5, 2)

	// standard error checking
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(result, remainder) // Prints 2 1

	// MisMatched Assignment
	// If a function returns 3 values, can you assign them to 2 variables
	// Go requires strict assignment counts

	// partial ignore
	f, _ := os.Open("file.txt")
	defer f.Close()

	// Shadowing Risk
	example()
}

func example() (count int) {
	count = 10
	if true {
		count := 20 // Look closely at the operator
		return count
	}
	return count
}

func namedAndOptionalParam() {
	// Call using named Fields. Unset fields get zero values

	// Omit FirstName -> becomes " "
	MyFunc(MyFuncOpts{
		LastName: "Raut",
		Age:      20,
	})

	// Omit Age -> becomes 0
	MyFunc(MyFuncOpts{
		FirstName: "Joe",
		LastName:  "Smith",
	})
}

func declaringAndCallingFunction() {
	//result := div(10, 5)
	//fmt.Println(result)
}
