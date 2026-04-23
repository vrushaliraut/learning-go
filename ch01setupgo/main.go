package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Chapter 01: Setup Go")

	networklimitation()

	typesDeclaration()

	booleanDeclaration()

	integerTypes()

	var u uint8 = 100
	processByte(u)

	r := 'a'
	fmt.Printf("Type: %T, Value: %d \n", r, r)

	tasteOfStringAndRunes()

	declareImmutableVariable()
}

func declareImmutableVariable() {
	// package level typed constant
	const x int64 = 100

	// Package-level constant block
	// These are "untyped" string constants
	const (
		idKey   = "id"
		nameKey = "name"
	)

	// The value is figured out at compile time
	const z = 20 * 10
}

func tasteOfStringAndRunes() {
	var s string = "Hello, world!"
	s = "goodbye"
	fmt.Println(s)
}

// this function accept the alias 'byte'
func processByte(b byte) {
	fmt.Println("this function accept the alias byte ", b)
}

// This block is now grouped and more readable
var (
	host    = "localhost"
	port    int
	enabled = true
)

func integerTypes() {
	var packetCount int32
	// No value is assigned so it defaults to its zero value

	fmt.Println(packetCount)

	// Integer overflow  - 256 is one more than the max value of uint8
	/*var myByte uint8 = 256
	fmt.Println(myByte)*/

	// cannot use -1 (untyped int constant) as uint32 value in variable declaration (overflows)
	/*var counter uint32 = -1
	fmt.Println(counter)*/

	var i int = 10 // 'int' type (size depends on platform)

	fmt.Println("The value of i is", i)
	var i64 int64 = 20 // 'int64' type (always 64 bits)

	// This command would fail to compile
	// i = i64

	i = int(i64)
	fmt.Println("The value of newly assigned i is", i)

	x := 10
	y := 3

	fmt.Println("The value of x is", x+y) // 13
	fmt.Println("The value of x is", x*y) //30

	a := 10
	b := 3
	fmt.Println(a / b) // Prints 3

	c := 5
	d := 2
	fmt.Println(c / d) // Prints 2

	e := -5
	f := 2
	fmt.Println(e / f) // Prints -2 (truncates toward zero)

	checkEpsilon()

	runeVsIntThirtyTwo()

	stringVsRune()

	theZeroValue()

	concatenation()

	convertTypes()

	x_verb := 10   // defaults to int
	y_verb := 10.5 // defaults to float64

	fmt.Printf("x_verb is %T\n", x_verb)
	fmt.Printf("y_verb is %T\n", y_verb)

	literalsAndVariables()

	fmt.Println(host, port, enabled)

}

// Literals (Flexible) vs. Variables (Strict)
func literalsAndVariables() {
	var f float64 = 20.5
	//var i int = 10

	//  Error: invalid operation: f + i (mismatched types)
	//fmt.Println(f + i)

	fmt.Println(f + 10) // Prints: 30.5 10 gets promoted to 10.5

	var b1 byte = 255
	fmt.Println(b1)

	//var byte_overflow byte = 256 // (untyped int constant) as byte value in variable declaration (overflows)
	//fmt.Println(byte_overflow)

	/*var l m int = 10 , 20

	var x, y int

	var j, k = 10, "hello"*/

	var port, service = 8080, "http"

	// Use %T verb to print the type
	fmt.Printf("The type of port is %T\n", port)
	fmt.Printf("The type of service is %T\n", service)
}

// type cast
func convertTypes() {
	var x int = 10
	var y float64 = 30.2

	// Convert x (int) to float64
	var sum1 float64 = float64(x) + y

	// convert y (float64) to int
	var sum2 int = x + int(y)

	fmt.Println("sum1 ::", sum1, "sum2 ::", sum2)

	var z int = 10
	var b byte = 100 // byte is an alias for uint8

	// We must convert 'x' to 'byte' to add it to 'b'
	var sum3 int = z + int(b)
	var sum4 byte = byte(x) + b

	fmt.Println("sum3 with int conversion :: ", sum3, "\n sum4 with byte conversion :: ", sum4)

}

func concatenation() {
	// write a program that declares two string s1 = "hello" and s2 = "world" and a rune r = '!'
	// create a new string s3 that contains "hello world!".

	s1 := "hello"
	s2 := "world"
	r := '!'

	// You must cast the rune to a string
	s3 := s1 + " " + s2 + string(r)
	fmt.Println("concatenation example :: ", s3)
}

func theZeroValue() {
	var user string
	var initial rune

	fmt.Printf("String zero value: '%s'\n", user)
	fmt.Printf("Rune zero value: '%d'\n", initial)
}

func stringVsRune() {
	x := 'a' // quotes create a string literal - the type of y is rune
	y := "a" // quotes create a rune literal

	fmt.Println("string x value ::-- ", string(x), "string y value ::-- ", string(y))
}

func runeVsIntThirtyTwo() {
	var r rune = 'a'  // This is the number 97
	var i int32 = 'b' // This is the number 98

	result := r + i
	fmt.Println("integer math on rune and int32 :: ", result)
}

func checkEpsilon() {
	a := 0.1
	b := 0.2

	sum := a + b // 0.3000000004

	// This will print "NOT equal"

	if sum == 0.3 {
		fmt.Println("The sum is 0.3")
	} else {
		fmt.Println("The sum is not 0.3")
	}

	//The Right way
	const epsilon = 0.00001

	// Check if the absolute difference is less than our tolerance
	if math.Abs(sum-0.3) < epsilon {
		fmt.Println("Epsilon comparison: Equal")
	} else {
		fmt.Println("Epsilon comparison: NotEqual")
	}

	//The Precision trap
	var coord float32 = 123.456789

	// printing will show a rounded value
	fmt.Println(coord)

	// Integer Vs Float Division by zero
	// what is the key difference in behaviour between running fmt.Println(10/0) integers and fmt.Println(10.0/0.0)
	// fmt.Println(10/0) and fmt.Println(10.0/0.0)

	// nan := 0.0 / 0.0

	// fmt.Println(nan == nan) // false

	// Because of this you must check for NAN
	// fmt.Println(math.IsNaN(nan)) // Prints: true
}

func booleanDeclaration() {
	// Declared 'enabled' without assigning value
	var enabled bool

	// Print its value - it will print 'false' because that is the zero value for bool
	fmt.Println(enabled)

	/*var myBool bool
	myBool = 1 // this is the line that will fail */
}

func typesDeclaration() {
	s := "Greetings \n\"Salutions\""
	fmt.Println(s)

	str := 'A'  // this is assigning ASCII value
	str1 := "A" // this is assigning actual character
	fmt.Println(str)
	fmt.Println(str1)
}

// go fmt ./... ==> silently reformatted the file
