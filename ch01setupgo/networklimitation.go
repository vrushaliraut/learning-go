package main

import (
	"fmt"
	"net/http"
	"time"
)

func networklimitation() {
	fmt.Println(time.Now())

	response, err := http.Get("https://google.com")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(response.Status)

	fmt.Println("Hello, world!")
	fmt.Println("Another line")
}
