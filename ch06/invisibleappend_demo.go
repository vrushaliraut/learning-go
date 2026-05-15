package main

import "fmt"

func main() {
	// slice with length 0 and cap 5
	items := make([]int, 0, 5)

	fmt.Printf("Before:  val=%v, len=%d, cap=%d\n", items, len(items), cap(items))

	tryAppend(items)

	// Even though the underlying array was modified,
	// 'items' in main still thinks its length is 0.
	fmt.Printf("After:   val=%v, len=%d, cap=%d\n", items, len(items), cap(items))

	// Proof: The data IS there, we just can't see it without reslicing!
	fullView := items[:1]
	fmt.Printf("Proof:   Hidden element is %d\n", fullView[0])
}

func tryAppend(items []int) {
	// append modifies the LOCAL 's' headers length and writes to the underlying array

	items = append(items, 100)
	fmt.Printf("Inside:  val=%v, len=%d, cap=%d\n", items, len(items), cap(items))
}
