package main

import "fmt"

type Person struct {
	name string
	age  int
}

func main() {
	vrushali := Person{"vrushali", 20}
	fmt.Println(vrushali)
	processPerson(vrushali)
	fmt.Println(vrushali)
	processPersonWithPointer(&vrushali)

	fmt.Println(vrushali)

	var vrushaliP *Person = nil
	processPersonNil(vrushaliP)
}

func processPersonNil(p *Person) {
	newP := &Person{"sally", 20}
	p = newP
}

// Why this method is a costly operation
func processPerson(p Person) {
	p.name = "Riyanshi"
	p.age = 33
}

func processPersonWithPointer(p *Person) {
	p.name = "Riyanshi"
	p.age = 4

	p = &Person{"sally", 20}
}
