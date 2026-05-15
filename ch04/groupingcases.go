package main

import "fmt"

func main() {
	/*day := 2 // Wednesday
	switch day {
	case 0, 6:
		fmt.Println("Weekend")
	case 1, 2, 3, 4, 5:
		fmt.Println("Workday")
	default:
		fmt.Println("Invalid")
	}*/

	// empty case
	n := 2
	switch n {
	case 1:
		fmt.Println("one")
	case 2:
	case 3:
		fmt.Println("three")
	}

	// Variable scope
	/*fmt.Println("Variable scope: ")
	switch x:= 10; x {
	case 10:
		y := 5
		fmt.Println(y)
	case 20:
		fmt.Println(y) // error here
	}
	*/

}
