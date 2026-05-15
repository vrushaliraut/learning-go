package main

import "fmt"

func main() {
	fmt.Println("=== Multiple Variable Declarations ===\n")

	fmt.Println("=== var with multiple assignments ===")
	var x, y, z int = 10, 20, 30
	fmt.Printf("x: %d, y: %d, z: %d\n", x, y, z)

	var firstName, lastName string = "John", "Doe"
	fmt.Printf("Name: %s %s\n", firstName, lastName)

	fmt.Println("\n=== var with type inference (multiple) ===")
	var a, b = 100, 200
	var c, d = "apple", "banana"
	fmt.Printf("a: %v, b: %v\n", a, b)
	fmt.Printf("c: %s, d: %s\n", c, d)

	fmt.Println("\n=== Short declaration (multiple) ===")
	name1, age1 := "Alice", 25
	name2, age2 := "Bob", 30
	fmt.Printf("%s is %d years old\n", name1, age1)
	fmt.Printf("%s is %d years old\n", name2, age2)

	fmt.Println("\n=== Mixed types in short declaration ===")
	id, title, enabled := 1, "Admin", true
	fmt.Printf("ID: %d, Title: %s, Enabled: %v\n", id, title, enabled)

	fmt.Println("\n=== Variable block declaration ===")
	var (
		server   = "localhost"
		portNum  = 8080
		isActive = true
	)
	fmt.Printf("Server: %s:%d (Active: %v)\n", server, portNum, isActive)

	fmt.Println("\n=== Short declaration with blank identifier ===")
	result, _ := divideWithRemainder(10, 3)
	fmt.Printf("10 / 3 = %d\n", result)
}

func divideWithRemainder(a, b int) (int, int) {
	return a / b, a % b
}

