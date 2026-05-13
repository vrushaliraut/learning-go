package main

import "fmt"

func main() {
	x := "hello"
	pointerToX := &x // Pointer to x now holds the memory address of x

	fmt.Println("pointerToX :", pointerToX)  // pointerToX : 0x5d63af08060 - prints the address
	fmt.Println("pointerToX :", *pointerToX) // prints the value at that address

	// You can also modify the value through the pointer
	*pointerToX = "Riyanshi"
	fmt.Println("pointerToX :", *pointerToX)

	// var value *int     // x is nil
	// fmt.Println(value) //prints <nil>
	//	fmt.Println("value :", *value) // PANIC: invalid memory address or nil pointer dereference

	// Pointer to Constant is not possible

	/* p := &5
	fmt.Println("p :", p) // prints the address of the value
	q := &"hello"         // compile error
	fmt.Println("q :", q) // prints the add
	*/
	
}
