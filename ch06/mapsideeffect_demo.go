package main

import "fmt"

func main() {
	// Initialise map
	data := map[string]int{"initial": 0}
	fmt.Printf("Initial state : %v\n", data)

	// call modifyMap
	// We are passing a copy of the 'header', which still points to the same data
	modifyMap(data)

	// Output: map[active:1 initial:0 status_code:200]

	// call reassign Map
	// We are reassigning the LOCAL copy of the header. The original 'data'
	// header in main() still points to the old hash table.

	reassignMap(data)
	fmt.Printf("After reassignMap: %v\n", data)
	// Output: map[active:1 initial:0 status_code:200] (No change!)
}

// reassignMap tries to point the map to a brand new memory location.
// This mutation IS NOT visible to the caller
func reassignMap(data map[string]int) {
	data = map[string]int{"new_data": 999}
	fmt.Printf("Inside reassignMap: %v\n", data)
}

// modifyMap adds/updates keys. This mutation IS visible to the caller.
func modifyMap(m map[string]int) {
	m["active"] = 1
	m["status_code"] = 200
}
