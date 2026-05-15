package main

import (
	"fmt"
	"io"
	log2 "log"
	"os"
)

func main() {
	fmt.Println(os.Args)

	if len(os.Args) < 2 {
		log2.Fatal("no file specified")
	}

	// Acquire the resource
	file, err := os.Open(os.Args[1])
	if err != nil {
		log2.Fatal(err)
	}

	// 2. Schedule cleanup immediately after successful open
	// This will run when main() exits.
	defer file.Close()

	// 3. Use the resource
	data := make([]byte, 2048)
	for {
		count, err := file.Read(data)
		os.Stdout.Write(data[:count])
		if err != nil {
			if err != io.EOF {
				log2.Fatal(err)
			}
			break
		}
	}

	defer fmt.Println("First Defer (Runs Last)")
	defer fmt.Println("Second Defer (Runs Second)")
	defer fmt.Println("Third Defer (Runs First)")

	fmt.Println("Main body")

}
