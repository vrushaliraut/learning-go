package main

import "fmt"

func addition_example(a, b float64) float64 { return a + b }
func substraction(a, b float64) float64     { return a - b }
func multiplication(a, b float64) float64   { return a * b }
func division_example(a, b float64) float64 {
	if b == 0 {
		fmt.Println("Error: cannot divide by zero")
		return 0
	}
	return a / b
}

// simulate_do_while
func main() {
	// variable to hold our inputs
	var num1, num2 float64
	var operator string
	var choice string

	for {
		// Step 1 - Get inputs
		fmt.Println("\n Enter first number: ")

		fmt.Scan(&num1) // Dont forget the &

		fmt.Println("Enter operator (+, -, *, /): ")
		fmt.Scan(&operator)

		fmt.Println("\n Enter second number: ")
		fmt.Scan(&num2)

		// Step 2: Process Logic
		result := 0.0

		// We use a switch statement to choose the right function

		switch operator {
		case "+":
			result = addition_example(num1, num2)
		case "-":
			result = substraction(num1, num2)
		case "*":
			result = multiplication(num1, num2)
		case "/":
			result = division_example(num1, num2)
		default:
			fmt.Println("Error: unknown operator")
			// 'continue' skips the rest of the loop and start over at the top
			continue

		}

		// Step 3 = Output
		fmt.Printf("Result: %.2f\n", result)

		// Step 4 - The "while" condition
		fmt.Print("Do you want to calculate again? (y/n): ")
		fmt.Scan(&choice)

		if choice == "n" || choice == "N" {
			fmt.Println("Goodbye!")
		}
	}
}
