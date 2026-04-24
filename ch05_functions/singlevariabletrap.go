package main

import "fmt"

type UserData struct {
	Name string
	Age  int
}

// change function to return the struct OR assign to immediate variables

func main() {

	name, age := getUser()
	userData := UserData{Name: name, Age: age}

	fmt.Println(userData)

	// another clean option is
	userData = getUserWithStruct()
}

func getUser() (string, int) {
	return "vrushali", 30
}

func getUserWithStruct() UserData {
	return UserData{Name: "vrushali", Age: 30}
}
