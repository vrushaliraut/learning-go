package main

import "fmt"

func main() {
	categorizingNumbers()

	fizzBuzz()

	refactorSwitch("open")
}

func refactorSwitch(status string) {
	/*	switch {
		case status == "open":
			// handle open
		case status == "closed":
			// handle closed
		case status == "error":
			// handle error
		}*/

	switch status {
	case "open":
		// ...
	case "closed":
		// ...
	case "error":
		// ...
	}

}

// Re-implement the classic FizzBuzz logic
// (print "Fizz" if divisible by 3, "Buzz" if by 5, "FizzBuzz" if by both)
// using a blank switch inside a loop from 1 to 15.
func fizzBuzz() {
	for i := 1; i <= 15; i++ {
		switch {
		case i%3 == 0 && i%5 == 0:
			fmt.Println("FizzBuzz")
		case i%3 == 0:
			fmt.Println("Fizz")
		case i%5 == 0:
			fmt.Println("Buzz")
		default:
			fmt.Println(i)
		}
	}
}

// Write a function that takes an int and prints "Negative", "Zero", or "Positive" using a blank switch.
func categorizingNumbers() {
	checkSign(-5)
	checkSign(0)
	checkSign(10)
}

func checkSign(n int) {
	switch {
	case n < 0:
		fmt.Println("Negative")
	case n == 0:
		fmt.Println("Zero")
	default:
		fmt.Println("Positive")
	}
}
