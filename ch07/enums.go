package main

import "fmt"

type MailCategory int

const name = "myTime"

const value = 123

const (
	Uncategorized MailCategory = iota
	Personal
	Spam
)

func main() {
	fmt.Println(name)
	fmt.Println(value)
}
