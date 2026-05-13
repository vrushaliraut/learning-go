package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// User represent structure of our JSON data
type UserWithJson struct {
	ID       int    `json:"user_id"`
	Username string `json:"username"`
	IsAdmin  bool   `json:"is_admin"`
}

func main() {
	// Imagine this byte slice came from an API response or a file

	rawJSON := []byte(`{
		"user_id":101,
		"username":"Alice",
		"is_admin":true
	}`)

	// Initialize an empty struct instance
	var u UserWithJson

	// Use the pointer (&u)
	// If we passed 'u' instead of '&u', Unmarshal would receive a
	// copy and wouldn't be able to modify the original variable.

	err := json.Unmarshal(rawJSON, &u)
	if err != nil {
		log.Fatal("Oops! Unmarshal failed: ", err)
	}

	// success! the 'u' variable is now populated
	fmt.Printf("Successfully Unmarshaled:\n")
	fmt.Printf("ID:       %d\n", u.ID)
	fmt.Printf("Username: %s\n", u.Username)
	fmt.Printf("Admin:    %t\n", u.IsAdmin)
}
