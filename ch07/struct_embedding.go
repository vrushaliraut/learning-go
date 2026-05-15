package main

import "fmt"

type employee struct {
	name string
	id   int
}

// earlier name was description-  go has interfaces
func (e *employee) String() string {
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
	fmt.Println(m.String())
	fmt.Println(m.name)
	fmt.Println(m.employee.id)
	fmt.Println(m.id)
	fmt.Println(m.reports[0].String())

	// we cannot provide like this - as both m and e are peer and not inheritance
	// processId(m)
}

func processId(e employee) {
}

func processSalary(e employee, salary float64) (employee, float64) {
	return e, salary
}

func processSalary2(m manager, salary float64) (manager, float64) {
	return m, salary
}
