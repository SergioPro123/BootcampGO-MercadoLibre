package main

import "fmt"

type SalaryErrorStruct struct {
	error
	salary int
}

func (s *SalaryErrorStruct) Error() string {
	return fmt.Sprintf("Error: el salario $%d ingresado no alcanza el mínimo imponible", s.salary)
}

func main() {
	var salary int = 200000
	err := ApplyToTax(salary)
	if err != nil {
		panic(err)
	}

	fmt.Println("Debe pagar impuestos.")
}

func ApplyToTax(salary int) (err error) {
	if salary < 150000 {
		err = &SalaryErrorStruct{salary: salary}
	}
	return
}
