package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	fmt.Println("=== Network Request Demo ===")
	fmt.Println("Current time:", time.Now())

	fmt.Println("\nMaking HTTP request to google.com...")
	response, err := http.Get("https://google.com")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Response Status:", response.Status)

	fmt.Println("\nHello, world!")
	fmt.Println("Another line")
}
