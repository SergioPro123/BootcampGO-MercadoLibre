package main

import (
	"fmt"
	"time"
)

type Person struct {
	ID          int
	Name        string
	DateOfBirth time.Time
}
type Employee struct {
	ID       int
	Position string
	Person   Person
}

func (e *Employee) PrintEmployee() {
	fmt.Printf("%+v \n", e)
}

func main() {
	var employee Employee = Employee{
		ID:       1,
		Position: "Software Developer",
		Person: Person{
			ID:          2,
			Name:        "Sergio Mauricio",
			DateOfBirth: time.Date(2000, 2, 12, 0, 0, 0, 0, time.Local),
		},
	}

	employee.PrintEmployee()
}
