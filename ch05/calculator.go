package main

import (
	"errors"
	"fmt"
	"strconv"
)

/*
The Question: The simple calculator program (from the "Functions Are Values" section)
doesn’t handle one error case:division by zero.

Change the function signature for the math operations to return both an int and an error.
In the div function, if the divisor is 0, return errors.New("division by zero") for the error.
In all other cases (add, sub, mul), return nil for the error.
Adjust the main function to check for this error and print it if it occurs.

*/

// Define a  function type that returns an int and an error
type opFunctionType func(int, int) (int, error)

// Math functions now match the opFuncType signature
func add(i, j int) (int, error)      { return i + j, nil }
func subtract(i, j int) (int, error) { return i - j, nil }
func multiply(i, j int) (int, error) { return i * j, nil }

func div(i, j int) (int, error) {
	if j == 0 {
		return 0, errors.New("division by zero")
	}
	return i / j, nil
}

// Map now uses the new signature
var opMap = map[string]opFunctionType{
	"+": add,
	"-": subtract,
	"*": multiply,
	"/": div,
}

func main() {
	expressions := [][]string{
		{"2", "+", "3"},
		{"2", "/", "3"}, // This will trigger the error
		{"5"},
	}

	for _, expression := range expressions {
		if len(expression) != 3 {
			fmt.Println("Invalid expression: ", expression)
			continue
		}

		p1, err := strconv.Atoi(expression[0])
		if err != nil {
			fmt.Println(err)
			continue
		}

		op := expression[1]
		opFunc, ok := opMap[op]
		if !ok {
			fmt.Println("unsupported operator: ", op)
			continue
		}

		p2, err := strconv.Atoi(expression[2])
		if err != nil {
			fmt.Println(err)
			continue
		}

		// Call the function and handler the new error return
		result, err := opFunc(p1, p2)
		if err != nil {
			fmt.Println(err)
			continue
		}

		fmt.Println(result)
	}

}
