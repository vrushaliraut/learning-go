package main

import "fmt"

func main() {
	fmt.Println("Chapter 06: pointer")
	// pointer_behaviour()
	type array2 struct {
		ints []int
		len  int
		cap  int
	}

	var x int32 = 10
	var y bool = true
	pointerX := &x
	pointerY := &y
	var pointerZ *string

	fmt.Println(pointerX, pointerY, pointerZ)

	x1:= "hello"
	pointerToX1 := &x1 // pointerToX now holds the memory address of x

	fmt.Println("pointerToX1 :", pointerToX1)

}
