package main

import "fmt"

type Person struct {
	Age       int
	AgeWorked int
	IsWorking string
	Salary    float64
}

func main() {

	var datosClientes [3]Person

	//Obtener datos del cliente
	for i := 0; i < 3; i++ {
		var age int
		var isWorking string
		var ageWorked int
		var salary float64

		fmt.Printf(" ---------- Cliente #%d ---------- \n", i+1)

		fmt.Print("Ingrese la edad: ")
		fmt.Scanln(&age)
		datosClientes[i].Age = age

		fmt.Print("Se encuentra laborando actualmente? y/n : ")
		fmt.Scanln(&isWorking)
		datosClientes[i].IsWorking = isWorking

		if datosClientes[i].IsWorking == "y" {

			fmt.Print("Ingrese la antiguedad en el trabajo : ")
			fmt.Scanln(&ageWorked)
			datosClientes[i].AgeWorked = ageWorked

			fmt.Print("Ingrese el salario : ")
			fmt.Scanln(&salary)
			datosClientes[i].Salary = salary
		}
	}

	//Procesar datos
	for i := 0; i < 3; i++ {
		fmt.Printf("Cliente #%d : Años: %d ; Esta trabajando: %s ; Años laborando : %d ; Salario : %d -- Credito aprovado : %s \n ; Intereses: %s",
			i,
			datosClientes[i].Age,
			datosClientes[i].IsWorking,
			datosClientes[i].AgeWorked,
			datosClientes[i].Salary,
			IsApprovedCredit(datosClientes[i]),
			AplicaIntereses(datosClientes[i]))
	}
}

func IsApprovedCredit(person Person) bool {
	return person.Age > 22 &&
		person.AgeWorked >= 1 &&
		person.IsWorking == "y"
}

func AplicaIntereses(person Person) bool {
	return person.Salary < 100000
}
