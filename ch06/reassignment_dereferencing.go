package main

import "fmt"

func main() {
	x := 10
	failedUpdateWithAssignment(&x)
	fmt.Println("assign address value of x :: ", x) // Prints 10 no change

	update(&x)
	fmt.Println("Pointer to value of x :: ", x) // Prints 20 success!
}

func update(i *int) {
	*i = 20
}

func failedUpdateWithAssignment(i *int) {
	a := 20
	i = &a
}
