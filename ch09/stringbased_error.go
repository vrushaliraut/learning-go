package main

import (
	"errors"
	"fmt"
)

// static failure - errors.New

// dynamic failure - fmt.Errorf

func processEven(i int) (int, error) {

	if i%2 != 0 {
		// The failure reason is identical every time this is triggered
		return 0, errors.New("processEven")
	}
	return i * 2, nil
}

func processEven1(i int) (int, error) {
	if i%2 != 0 {
		return 0, fmt.Errorf("number %d is not an even number", i)
	}
	return i * 2, nil
}

func main() {
	y, err := processEven(3)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(y)

}
