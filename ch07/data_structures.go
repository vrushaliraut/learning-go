package main

import "fmt"

type LLNode struct {
	value int
	next  *LLNode
}

func (n *LLNode) append(val int) *LLNode {
	if n == nil {
		return &LLNode{val, nil}
	}
	n.next = n.next.append(val)
	return n
}

func (n *LLNode) print() {
	if n == nil {
		return
	}
	fmt.Print(n.value, "->")
	n.next.print()
}
