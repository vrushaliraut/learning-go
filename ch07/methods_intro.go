package main

import "fmt"

type Employee struct {
	name       string
	salary     float64
	department string
}

func main() {
	employee := Employee{"", 1000, ""}
	updateEmp := updateEmployee(employee, 13000)
	fmt.Println(updateEmp)

	// method calling - behaviour attaching
	emp2 := Employee{"", 3000, ""}
	emp2.updateEmployee(35000)
	fmt.Println(emp2)
}

// method
func (e Employee) updateEmployee(newSalary float64) Employee {
	e.salary = newSalary
	return e
}

// function
func updateEmployee(emp Employee, newSalary float64) Employee {
	emp.salary = newSalary
	return emp
}
