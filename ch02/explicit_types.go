package main

import "fmt"

func main() {
	fmt.Println("=== Variable Declaration with Explicit Type ===\n")

	// Explicit type declaration
	var name string = "Go"
	var version int = 1
	var releaseYear int = 2009

	fmt.Printf("Language: %s\n", name)
	fmt.Printf("Version: %d\n", version)
	fmt.Printf("Released: %d\n", releaseYear)

	fmt.Println("\n=== Type Inference ===")
	var description = "A compiled, statically typed language"
	var isOpen = true
	fmt.Printf("Description: %s (type: %T)\n", description, description)
	fmt.Printf("Open Source: %v (type: %T)\n", isOpen, isOpen)

	fmt.Println("\n=== Short Declaration ===")
	author := "Robert Griesemer, Rob Pike, Ken Thompson"
	popularity := "High"
	fmt.Printf("Authors: %s\n", author)
	fmt.Printf("Popularity: %s\n", popularity)
}

