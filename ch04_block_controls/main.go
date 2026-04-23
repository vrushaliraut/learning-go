package main

import "fmt"

func main() {
	fmt.Println("Chapter 04: Block Shadow Control structure")

	// slice_generate()

	//precedence()

	shadowingbug()

	// blocks and shadow
	x := 10

	if x > 5 {
		fmt.Println(x) // Prints 10 (Outer x)

		// This declaration SHADOWS the outer 'x'
		x := 5

		fmt.Println(x) // Prints 5 (Inner x)
	}
	// The inner block has ended. The inner 'x' is gone.
	// The outer 'x' becomes visible again.
	fmt.Println(x) // Prints 10 (Outer x)
	/*
		fmt := "log"
		println(fmt)
	*/

	/*
		nil := "empty"
		if err != nil {

	}*/

	n := 10
	if n == 0 {
		fmt.Println("That's too low")
	} else if n > 5 {
		fmt.Println("That's too big: ", n)
	} else {
		fmt.Println("That's  a good number: ", n)
	}

	count := 0
	if count == 0 {
		fmt.Println("Empty")
	} else {
		fmt.Println("not empty")
	}

	// if declaration; condition { ... }

	if n := 10; n == 0 {
		fmt.Println("That's too low")
	} else if n > 5 {
		fmt.Println("That's too big")
	} else {
		fmt.Println("That's  a good number: ", n)
	}

	ifConditionConcept()

	forConcept()

	conditionOnlyForStatement()

	doWhileSimulation()

	searchAndBreak()

	// do while loop simulation using continue
	// Filter Odd Numbers
	filterOddNumbers()

	// Iterators like resembles is in go is achieve using for each range
	forRangeStatement()

	summingSlice()

	mapIteration()

	indexOnlyLoop()

	// map concepts ::
	//iteratorOverMapConcept()

	iteratingOverString()

}

func iteratingOverString() {
	// When you use for-range on a string, the loop attempts to decode the string into runes (Unicode code points), not bytes.

	fmt.Println("The For-Range Loop On Strings\n")

	// hello contains only 1-byte characters
	for i, r := range "Hello" {
		fmt.Printf("%d %q \n", i, r)
	}

	str := "Golang"
	for _, r := range str {
		// Use %c to print the rune as a character

		fmt.Printf("%c \n", r)
	}
}

func indexOnlyLoop() {
	fmt.Println("index Only Loop :: ")
	s := []string{"a", "b", "c"}
	for i := range s {
		fmt.Println(i)
	}
}

// Create a map prices := map[string]int{"apple": 2, "banana": 3}.
// Use a for-range loop to print strings in the format "Item: [key], Cost: [value]".
func mapIteration() {
	prices := map[string]int{"apple": 1, "banana": 2, "cherry": 3}
	fmt.Print("Map Iteration :: \n ")
	// k gets the key, v gets the value
	for k, v := range prices {
		fmt.Printf("Item: %s, Cost: %d\n", k, v)
	}
}

/*
- Write a function that accepts a slice of int and uses a for-range loop to calculate and return the sum of all elements.
- You do not need the index.
*/

func summingSlice() {
	sliceOfInt := []int{1, 2, 3, 4, 5, 6}

	sum := 0
	for _, val := range sliceOfInt {
		sum += val
	}
	fmt.Println(sum)
}

func forRangeStatement() {
	/*
		for key, value := range collection {
			// block of code
		}
	*/

	// Slice - iterating over slices

	fmt.Println("For range statement -iterating over slices \n ")
	evenVals := []int{2, 4, 6, 8, 10}

	for i, val := range evenVals {
		fmt.Println(i, val)
	}

	fmt.Println("You can ignore the indexes by using _ ")
	for _, val := range evenVals {
		fmt.Println(val)
	}

	// Ignoring values
	uniqueName := map[string]bool{"vrushali": true, "Riyanshi": true, "Abhi": false}

	// Only capture the key 'k'. The value is implicitly ignored.
	for k := range uniqueName {
		fmt.Println(k)
	}
}

/*
Write a loop from 1 to 10. Use continue to skip the fmt.Println(i) statement for any odd number.
(The output should be 2, 4, 6, 8, 10).
*/
func filterOddNumbers() {
	fmt.Print("Print even numbers ::  ")
	for i := 1; i < 10; i++ {
		// if odd, skip the rest
		if i%2 != 0 {
			continue
		}

		fmt.Println(i)
	}
}

/*
You have a slice nums := []int{10, 20, 30, 40}. Loop through it.
If you find the value 30, print "Found" and exit the loop immediately.
Ensure you do not print anything for 40.
*/
func searchAndBreak() {
	nums := []int{10, 20, 30, 40}

	for i := 0; i < len(nums); i++ {
		if nums[i] == 30 {
			fmt.Println("Found :: ", nums[i])
			break
		}
		// Just to prove we didn't process 40
		fmt.Println("Checked", nums[i])
	}
}

/*
Go does not have a do-while loop (a loop that always runs at least once).
How can you use an infinite for loop to simulate this behavior?
Write code that prints "Work" once, checks a condition (e.g., false), and then exits.
*/
func doWhileSimulation() {
	condition := false
	for {
		// DO the work
		fmt.Println("Work")

		// Check condition at the end
		if !condition {
			break
		}
	}
}

// go doesn't have while
func conditionOnlyForStatement() {
	// for ; 1 < 100; ;{} --> go fmt will remove it for you - for condition {...}

	i := 1
	// This runs while i is less than 100
	for i < 100 {
		fmt.Println(i)
		i = i * 2
	}

	/* the code will fail to compile - while is not a keyword in Go.
	   k := 0
	   while k < 5{
	   	fmt.Println(k)
	   	k++
	   }*/

	// Writing infinite for loop
	/*
		for {
			fmt.Println("hello")
		}*/
}

/*
A complete, C-style for
A condition-only for (similar to while in other languages)
An infinite for
for-range (for iterating over collections)
*/
func forConcept() {
	// complete for statement
	for i := 0; i < 10; i++ {
		fmt.Println("complete for statement :: ", i)
	}

	// initialization; comparison; increment {...}

	// Ommiting parts
	i := 0
	for ; i < 11; i++ {
		fmt.Println("complete for statement :: ", i)
	}

	for j := 0; j < 10; {
		fmt.Println("complete for statement :: ", j)
		if j%2 == 0 {
			j++
		} else {
			j = j + 2
		}
	}
}

func ifConditionConcept() {
	// scope isolation
	x := 100
	if x := 1; x > 0 {
		fmt.Println("Inside function x scope :: ", x)
	}
	fmt.Println("Outside function x scope :: ", x)

	// The comma Ok Idiom in if
	/*
		You have a map config := map[string]string{"env": "dev"}.
		Write an if statement that checks if the key "env" exists. If it does, print the value.
		Ensure the variables val and ok do not exist after the if statement.
	*/
	/*
		if val, ok := config["env"]; ok {
			fmt.Println("Environment variable found: ", val)
		}

		// outside function access
		if i := 10; i > 5 {
			fmt.Println("Big")
		}
		fmt.Println(i)
	*/
}
