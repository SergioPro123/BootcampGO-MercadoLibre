package main

import (
	"errors"
	"fmt"
)

type SalaryErrorStruct struct {
}

func (s SalaryErrorStruct) Error() string {
	return "Error: el salario es menor a 10.000"
}

var ErrorSalary SalaryErrorStruct = SalaryErrorStruct{}

func main() {
	var salary int = 2000
	err := LowSalary(salary)
	if err != nil {
		if errors.Is(err, ErrorSalary) {
			fmt.Println("Error debido al salario.")
		}
		panic(err)
	}

	fmt.Println("Salario ok.")
}

func LowSalary(salary int) (err error) {
	if salary < 10000 {
		err = SalaryErrorStruct{}
	}
	return
}
