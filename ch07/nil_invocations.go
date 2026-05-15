package main

import (
	"fmt"
	"time"
)

type Counter struct {
	total       int
	lastUpdated time.Time
}

// Value Receiver - will panic when we pass nil
func (c Counter) Print() {
	fmt.Println(c.total)
	c.lastUpdated = time.Now()
	fmt.Println(c.lastUpdated)
}
func (c *Counter) PrintPointer() {
	fmt.Println(c)
	//fmt.Println(c.lastUpdated)
}
func main() {
	var s *Counter = nil
	s.PrintPointer()

	/*s.Print()
	s.PrintPointer()*/
}
