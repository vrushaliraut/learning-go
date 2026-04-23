package main

import (
	"fmt"
	"os"
)

/*
os.Open: We try to open the file. If this fails (e.g., file doesn't exist), we return the error immediately.
defer f.Close(): The moment we have a valid file handle f, we defer its closure.
This ensures that even if f.Stat() fails later, the file handle won't leak.
f.Stat(): This is the standard way to get metadata about a file, including its size in bytes (info.Size()).
*/

func main() {
	fmt.Println("Chapter 05: functions")

	//calculator()

	// create dummy file for testing
	err := os.WriteFile("test.txt", []byte("Hello World"), 0644)

	if err != nil {
		fmt.Println(err)
		return
	}

	size, err := filelength_with_defer("test.txt")

	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("File size:", size, "bytes")
	}

	// Cleanup the test file
	os.Remove("test.txt")
}
