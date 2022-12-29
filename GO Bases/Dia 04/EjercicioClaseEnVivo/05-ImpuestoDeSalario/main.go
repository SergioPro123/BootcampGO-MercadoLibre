package main

import (
	"errors"
	"fmt"
)

type ErrorEmployee struct{}

func (e ErrorEmployee) Error() string {
	return "el trabajador no puede haber trabajado menos de 80 hs mensuales"
}

func main() {
	salary, err := calculateSalaray(100, 2000)
	if err != nil {
		if errors.Is(err, ErrorEmployee{}) {
			fmt.Println("Error por empleado.")
		}
		panic(err)
	}
	fmt.Printf("El salario del trabajador es de: $%.2f", salary)
}

func calculateSalaray(hoursWorked int, valueHour float64) (salary float64, err error) {
	if hoursWorked <= 80 {
		err = ErrorEmployee{}
		return
	}
	salary = float64(hoursWorked) * valueHour
	if salary >= 150000 {
		salary -= salary * 0.1
	}

	return
}
