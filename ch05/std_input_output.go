package main

import "fmt"

func main() {

	var name string

	fmt.Println(name)
	fmt.Println("Enter your name: ")
	_, err := fmt.Scan(&name)
	if err != nil {
		return
	}

	fmt.Println("Hello, ", name)
	fmt.Println(&name)
}
