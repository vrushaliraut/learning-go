package main

import "fmt"

func iteratorOverMapConcept() {

	// To iterate over map - use for-range syntax used for slices -
	// first variable receives the key, second variable receives the value

	m := map[string]int{
		"a": 1,
		"b": 3,
		"c": 2,
	}

	// Run the loop multiple times to see the order change
	for i := 0; i < 3; i++ {
		fmt.Println("Loop", i)

		for key, value := range m {
			fmt.Println(key, value)
		}
	}

	theFormattingMapException()

}

func theFormattingMapException() {
	m := map[string]int{"a": 1, "b": 2, "c": 3}

	// The output is STABLE - sorted by key
	fmt.Println(m)

	// This Output is UNSTABLE - random order
	for k := range m {
		fmt.Println(k)
	}
}
