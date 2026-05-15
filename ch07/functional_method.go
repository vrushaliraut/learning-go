package main

import "fmt"

type Adder struct {
	start int
}

func (a Adder) addTo(val int) int {
	return a.start + val
}

func addToWithoutStruct(x int) int {
	return 10 * x
}

func main() {
	myAdder := Adder{start: 10}
	f1 := myAdder.addTo

	fmt.Println(f1(5))

	fmt.Println(f1(25))

	fmt.Println(addToWithoutStruct(5))
	fmt.Println(addToWithoutStruct(10))
}
