package main

type IntTree struct {
	value int
	left  *IntTree
	right *IntTree
}

func (it *IntTree) insert(value int) *IntTree {
	// Base case when tree is empty
	if it == nil {
		// Create a new tree
		return &IntTree{value: value}
	}

	if value < it.value {
		it.left = it.left.insert(value)
		return it
	} else if value > it.value {
		it.right = it.right.insert(value)
		return it
	}
	return it
}

func main() {
	var it *IntTree = nil
	it = it.insert(5)
	it = it.insert(3)
	it = it.insert(10)
	it = it.insert(2)

}
