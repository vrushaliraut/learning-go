package main

import "fmt"

type employee struct {
	name string
	id   int
}

func (e *employee) description() string {
	return fmt.Sprintf("%s -> %d", e.name, e.id)
}

type manager struct {
	employee
	id      int
	reports []employee
}

func main() {
	e := employee{"John", 42}
	fmt.Println(e)

	m := manager{employee{"Jane", 6}, 30, []employee{{"cha", 0}}}
	fmt.Println(m)
	fmt.Println(m.description())
	fmt.Println(m.name)
	fmt.Println(m.employee.id)
	fmt.Println(m.id)
	fmt.Println(m.reports[0].description())

	// we cannot provide like this - as both m and e are peer and not inheritance
	// processId(m)
}

func processId(e employee) {
}
