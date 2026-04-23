package main

import (
	"fmt"
	"strconv"
)

func stringRuneByteConcepts() {

	// Go uses sequence of bytes to represent a string

	var s string = "Hello there"
	var b byte = s[6] // the 7th byte 't' (ASCII 116)
	fmt.Println("Go uses sequence of bytes to represent a string :: ", b)

	var s2 string = s[4:7]
	var s3 string = s[:5]
	var s4 string = s[6:]
	fmt.Println("Go uses sequence of bytes to represent a string :: ", s2, s3, s4)
	fmt.Println(len(s))

	fmt.Println("UTF-8 Danger zone when you use emojis which contain multiple byte as a single emoji...!")
	// "Hello " + Sun Emoji (4 bytes)
	var s1 string = "Hello 🌞"
	// This works: "Hello" is 5 bytes
	var sk string = s1[:5] // "Hello"
	fmt.Println(sk)
	// This fails! The sun emoji starts at index 6.
	// Index 4 is 'o', index 5 is space.
	// Indices 6, 7, 8, 9 are the sun emoji.
	var stringWithEmojiFirstByte string = s1[4:7] // "o " + [first byte of sun]
	var stringWithEmoji string = s1[6:]

	fmt.Println(stringWithEmojiFirstByte)
	fmt.Println(stringWithEmoji)

	typeConversionsStringRunesBytes()

	stringToByteSlice()

	countingCharactersCorrectly()

}

func countingCharactersCorrectly() {
	str := "Hello 🌞"
	count := len([]rune(str))

	fmt.Println("countingCharactersCorrectly using rune slice :: ", count)
}

// Write code to convert the string "ABC" into a slice of bytes. Print the slice.
func stringToByteSlice() {
	str := "ABC"

	byteSlice := []byte(str)

	fmt.Println("convert the string \"ABC\" into a slice of bytes ::", byteSlice)
}

func typeConversionsStringRunesBytes() {
	// Rune/Byte to string
	var a rune = 'x'
	var s string = string(a)
	fmt.Println(s)

	// String to Integer
	// To convert the number 65 to the string "65", you must use strconv.Itoa.
	var x int = 65
	var y = string(x)
	fmt.Println("Integer to string conversion but with ASCII :: ", y) // "A"

	var convertAtoString string = strconv.Itoa(x)
	fmt.Println("Integer to string conversion with built in strconv.Itoa function :: ", convertAtoString)

	var str string = "hello, 🌞"
	var byteString []byte = []byte(str)
	var runeString []rune = []rune(str)

	fmt.Println("byteString :: ", byteString)
	fmt.Println("runeString :: ", runeString)
}
