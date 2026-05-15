package main

import (
	"fmt"
	"time"
)

type MyTime time.Time

func (t MyTime) String() string {
	return time.Time(t).String()
}

func main() {
	fmt.Println("hello world")

}
