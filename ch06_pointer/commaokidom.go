package main

import "fmt"

func main() {
	/*val := getValue()
	fmt.Println(val)
	fmt.Println(&val)
	*/
	v, ok := getValue2()
	if ok {
		fmt.Println(v)
	}

	/*u1 := findUser(2)
	fmt.Println("nil user: ", u1)

	u11 := findUser(2)
	fmt.Println("nil user: ", u11.Name)
	*/
	u2, ok := findUser2(2)
	if ok {
		fmt.Println("nil user: ", u2.Name)
	}
	fmt.Println(u2)
}

func getValue() *int {
	return nil
}

func getValue2() (int, bool) {
	return 0, false
}

/*
	{
		"name": "jack",
		"age":0  ---> explicitly send age 0
		// if we dont send age
	}
*/
type Request struct {
	// Age *int `json:"age"`
	Age int `json:"age"`
}
type User struct {
	Name string
}

func findUser(id int) *User {
	if id != 1 {
		return nil
	}
	return &User{Name: "bob"}
}

func findUser2(id int) (User, bool) {
	if id != 1 {
		return User{}, false
	}
	return User{Name: "bob"}, true
}
