package main

import (
	"fmt"
	"math/rand"
)

func goto_demo() {

	illegalgoto()

	validgoto()
}

func validgoto() {
	a := rand.Intn(10)

	for a < 100 {
		if a%5 == 0 {
			// Logic dictates we are done with the loop AND
			// the normal post-loop logic. Jump to cleanup.
			goto done
		}
		a = a*2 + 1
	}

	fmt.Println("do something when the loop completes normally")

done:
	fmt.Println("do complicated stuff no matter why we left the loop")
	fmt.Println(a)

}

func illegalgoto() {
	/*a := 10

		// Illegal - this jumps over the declaration of 'b'
		goto skip
		b := 20

	skip :
		c := 30
		fmt.Println(a, b, c) // 'b' would be undefined here

		if c > a{
			// ILLEGAL : This jumps INTO the block below
			goto inner
		}

		if a < b {
			inner:
				fmt.Println("a is less than b")
		}
	*/
}
