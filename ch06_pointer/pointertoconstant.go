package main

import "fmt"

type UserWithNamePointer struct {
	Name *string // Pointer to string
}

func main() {
	/*
		p := &5
		q := &"hello"

		// ERROR: Cannot use "Alice" (type string) as type *string
		u := User{Name: "Alice"}

		// ERROR: Cannot take the address of "Alice"
		u2 := User{Name: &"Alice"}
	*/

	// Correct way: Create a variable and take its address
	name := "Alice"
	u3 := UserWithNamePointer{Name: &name}

	// Print the user
	println(*u3.Name) // Output: Alice

	// Another way to get around this - use a helper function to create a pointer to a string
	u4 := UserWithNamePointer{Name: makePointer("Vrushali with pointer passed as a string ")}
	fmt.Println(*u4.Name) // Output: Alice
}

// Take any value 't'  and returns a pointer to it
func makePointer[T any](t T) *T {
	return &t
}
