package main

import (
	"errors"
	"fmt"
)

type ValueErr struct {
	msg string
}

func (v ValueErr) Error() string {
	return v.msg
}

func main() {
	// Pointer semantics

	err1 := errors.New("database failure")
	err2 := errors.New("database failure")

	fmt.Println("errors.New matches?", err1 == err2) // FALSE

	// VALUE Semantics
	val1 := ValueErr{msg: "database failure"}
	val2 := ValueErr{msg: "database failure"}

	fmt.Println("ValueErr Matches?", val1 == val2) // TRUE

}
