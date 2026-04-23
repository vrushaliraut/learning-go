package main

import "fmt"

func structConcept() {

	// Declare map with empty struct values
	intSet := map[int]struct{}{}

	// Assignment is clumsy
	intSet[5] = struct{}{}
	// Checking requires comma ok

	_, ok := intSet[5]
	if ok {
		fmt.Println("5 is in the set")
	}

	// Write the line of code to add the string "admin" to a set defined as roles := map[string]struct{}{}.
	//roles["admin"] = struct{}{}

	defineStruct()

	copyingStruct()

	declareAnonymousStruct()

	structConversions()
}

func structConversions() {
	// Names Order Types

	type firstPerson struct {
		firstName string
		age       int
	}

	type secondPerson struct {
		firstName string
		age       int
	}

	var firstPerson1 firstPerson
	var secondPerson1 = secondPerson(firstPerson1)

	fmt.Println("First Person Name:", firstPerson1.firstName)
	fmt.Println("First Person Age:", firstPerson1.age)
	fmt.Println("Second Person Name:", secondPerson1.firstName)
	fmt.Println("Second Person Age:", secondPerson1.age)

}

func declareAnonymousStruct() {
	var person struct {
		name string
		age  int
		pet  string
	}

	// You access fields just like a named struct
	person.name = "bob"
	person.age = 50
	person.pet = "dog"

	fmt.Println("anonymous struct :: ", person)

	// Use a literal
	/*pet := struct {
		name string
		kind string
	}{
		name: "bob",
		kind: "dog",
	}*/

	/*
		You are writing a script and need a single variable config to hold two settings: APIKey (string) and Retries (int).
		You don't need this structure anywhere else. Declare and initialize this as an anonymous struct.
	*/
	config := struct {
		APIKey  string
		Retries int
	}{
		APIKey:  "SECRET_123",
		Retries: 3,
	}
	fmt.Println(config.APIKey, config.Retries)

	/* Create a variable shapes. It should be a slice. The elements of the slice are anonymous structs,
	each having a name (string) and sides (int). Initialize it with a "Triangle" (3 sides) and a "Square" (4 sides).
	*/

	shapes := []struct {
		name  string
		sides int
	}{
		{name: "Triangle", sides: 3},
		{name: "Square", sides: 4},
	}

	fmt.Println(shapes)
}

/*
* If you have c1 := Car{Make: "Ford"} and you do c2 := c1, then change c2.Make = "Chevy",
*  what is the value of c1.Make?
 */
func copyingStruct() {

	var c1 car
	c1 = car{Make: "Ford"}

	c2 := c1
	c2 = car{Make: "Chevy"}

	fmt.Println("Go makes a full copy of the struct fields. c2 is a distinct piece of memory :: \n ", c1, c2)
}

// Define a struct named Car with three fields: Make (string), Model (string), and Year (int).
// Using the Car struct from Problem 1, declare a variable myCar using var. Print its Year field. What is the value?
// Why: This tests the basic syntax of type definitions. Note the lack of commas between fields.
type car struct {
	Make  string
	Model string
	Year  int
}

// Write a snippet that checks if myCar.Year is less than 2000. If it is, print "Classic Car".
func defineStruct() {

	fmt.Println("car struct :: ", car{})

	var myCar car
	fmt.Println("car struct year:: ", myCar.Year)

	var c car
	c.Make = "Honda"
	c.Year = 1990
	fmt.Println("car make:: ", c.Make)

	if c.Year < 2000 {
		fmt.Println("Classic car")
	}
}
