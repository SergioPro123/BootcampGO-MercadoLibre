package main

import (
	"errors"
	"fmt"
)

var (
	ErrorMinSalary = "Error: el salario es menor a 10.000"
)
var ErrorSalary error

func main() {
	var salary int = 20000
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
	if salary < 150000 {
		ErrorSalary = fmt.Errorf("error: el mínimo imponible es de 150.000 y el salario ingresado es de: %d", salary)
		err = ErrorSalary
	}
	return
}
