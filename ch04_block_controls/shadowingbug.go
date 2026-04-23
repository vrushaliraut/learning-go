package main

import "fmt"

func shadowingbug() {

	var total int

	for i := 0; i < 10; i++ {
		// BUG: := declares a NEW 'total' variable scoped to the loop block.
		// It initializes this new variable with (outer_total + i).

		total := total + i
		fmt.Println(total) // Prints the running sum (0, 1, 2...)

	}

	fmt.Println("Final Total :", total)
}

/*
The Result: The loop will print incrementing numbers, but the final print statement will output 0.

The Bug: The line total := total + i uses the short declaration operator (:=).
This creates a new, local variable named total inside the loop's block, shadowing the package-level (or function-level) total.
The outer total is never updated.
The fix is to use the assignment operator: total = total + i (or total += i).
*/
