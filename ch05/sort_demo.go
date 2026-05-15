package main

import (
	"fmt"
	"sort"
)

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {
	people := []Person{
		{"Vrushali", "Raut", 33},
		{"Priya", "Shelake", 31},
		{"Poonam", "Tilekar", 20},
	}

	fmt.Println("people: ", people)

	// sort by last name - pass an anonymous function as the second argument
	sort.Slice(people, func(i, j int) bool {
		return people[i].LastName < people[j].LastName
	})

	fmt.Println("people by lastname: ", people)

	// sort by age - pass an anonymous function as the second argument
	sort.Slice(people, func(i, j int) bool {
		return people[i].Age < people[j].Age
	})
	fmt.Println("people by age : ", people)
}
