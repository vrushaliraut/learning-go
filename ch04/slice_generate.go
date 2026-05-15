package main

import (
	"fmt"
	"math/rand"
)

// Write a for loop that puts 100 random numbers between 0 and 100 into an int slice.

// Why: This tests slice initialization (make), looping, and using external packages (math/rand).
func slice_generate() {
	// create slice with length of 0 and capacity 100

	nums := make([]int, 0, 100)

	// Loop 100 times
	for i := 0; i < 100; i++ {
		// Generate random no - rand.Intn
		// usually you want 0 - 100 inclusive to Intn(101)
		val := rand.Intn(101)

		// Append to the slice
		nums = append(nums, val)

		fmt.Println(nums)
	}
}
