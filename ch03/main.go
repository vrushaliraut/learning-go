package main

import (
	"fmt"
)

func main() {
	fmt.Println("Chapter 03: Composite Types")

	// Example: Slice
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println("Numbers:", numbers)

	// Example: Map
	person := map[string]string{
		"name": "Alice",
		"city": "New York",
	}
	fmt.Println("Person:", person)

	// arrayConcepts()
	// slicesConcept()

	// makeInSlice()

	// slicingSlices()

	// copySlice()

	// stringRuneByteConcepts()

	// mapConcept()

	// structConcept()

	// exercises()

	var name string
	fmt.Println("Enter your name:")
	fmt.Scan(&name)
	fmt.Println("Hello,", name)
}

func exercises() {
	sliceManipulation()

	stringAndRuneAccess()
}

// Problem: Write a program that defines a string variable called message with the value "Hi 👩 and 👨" and
// prints the fourth rune in it as a character, not a number.
func stringAndRuneAccess() {
	message := "Hi 👩 and 👨"
	// Convert the string to a slice of runes to access by character index
	runes := []rune(message)
	// Access the 4th rune (index 3)
	// Index 0: 'H'
	// Index 1: 'i'
	// Index 2: ' ' (space)
	// Index 3: '👩' (Woman emoji)
	r := runes[3]
	// Print as a character using %c
	fmt.Printf("The fourth rune is: %c\n", r)
}

/*
Write a program that defines a variable named greetings of type slice of strings with the following values:
"Hello", "Hola", "नमस्कार", "こんにちは", and "Привіт".
Create a subslice containing the first two values; a second subslice with the second, third, and fourth values;
and a third subslice with the fourth and fifth values. Print out all four slices.
*/
func sliceManipulation() {
	greetings := []string{"Hello", "Hola", "नमस्कार", "こんにちは", "Привіт"}

	// Create a subslice containing the first two values
	s1 := greetings[:2]

	// a second subslice with the second, third, and fourth values;
	s2 := greetings[1:4]

	// a third subslice with the fourth and fifth values. Print out all four slices.
	s3 := greetings[3:]

	fmt.Println(greetings, s1, s2, s3)
}
