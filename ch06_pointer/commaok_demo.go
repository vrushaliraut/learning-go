package main

import "fmt"

type UserStruct struct {
	Name string
}

// findUser now returns a concrete User and a 'found' boolean.
// This avoids heap allocation and makes the API much clearer.

func searchUserById(id int) (UserStruct, bool) {
	if id != 1 {
		// Return a "zero value" and false
		return UserStruct{}, false
	}

	// Return data and true
	return UserStruct{Name: "Bob"}, true
}

func main() {
	idToSearch := []int{1, 42}

	for _, id := range idToSearch {
		fmt.Printf("Searching for ID: %d...\n", id)

		// the comma ok idiom at the call site
		user, ok := searchUserById(id)

		if !ok {
			fmt.Println("User not found")
		} else {
			fmt.Printf("User found: %v\n", user.Name)
		}
		fmt.Println("---")
	}
}
