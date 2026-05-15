package main

import (
	"fmt"
	"time"
)

type RateLimiter struct {
	numOfRequest int
	lastUpdate   time.Time
}

// Value receiver
func (r RateLimiter) update() {
	r.numOfRequest++
	r.lastUpdate = time.Now()
}

// Pointer receiver
func (r *RateLimiter) updatePointer() {
	// simply you can drop * here and use r
	(*r).numOfRequest++
	r.lastUpdate = time.Now()
}
func main() {
	s := RateLimiter{
		numOfRequest: 0,
	}

	s.update()
	fmt.Println(s.numOfRequest)

	// when you call pointer you always pass a pointer  -
	(&s).updatePointer()
	fmt.Println(s.numOfRequest)

	var a int = 10
	var b *int = &a
	fmt.Println(a, b)
}
