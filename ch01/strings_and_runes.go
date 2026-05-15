package main

import "fmt"

func main() {
	fmt.Println("=== Strings and Runes ===\n")

	fmt.Println("=== Rune vs Int32 ===")
	var r rune = 'a'  // This is the number 97
	var i int32 = 'b' // This is the number 98

	fmt.Printf("Rune 'a': %T, Value: %d\n", r, r)
	fmt.Printf("Int32 'b': %T, Value: %d\n", i, i)

	result := r + i
	fmt.Println("Integer math on rune and int32: 'a' + 'b' =", result)

	fmt.Println("\n=== String vs Rune ===")
	x := 'a'  // Single quotes create a rune (int32)
	y := "a"  // Double quotes create a string
	z := 'A'  // Another rune

	fmt.Printf("x := 'a' is type %T, value %d\n", x, x)
	fmt.Printf("y := \"a\" is type %T, value %q\n", y, y)
	fmt.Printf("z := 'A' is type %T, value %d\n", z, z)

	fmt.Println("\nConverting rune to string:")
	fmt.Printf("string(x): %s\n", string(x))
	fmt.Printf("string(z): %s\n", string(z))

	fmt.Println("\n=== String Concatenation ===")
	s1 := "hello"
	s2 := "world"
	r_exclaim := '!'

	s3 := s1 + " " + s2 + string(r_exclaim)
	fmt.Println("Concatenated:", s3)

	fmt.Println("\n=== String Operations ===")
	var s string = "Hello, world!"
	fmt.Println("Original:", s)

	s = "goodbye"
	fmt.Println("Modified:", s)

	// String escape sequences
	escaped := "Greetings\n\"Salutations\""
	fmt.Println("Escaped string:")
	fmt.Println(escaped)
}

