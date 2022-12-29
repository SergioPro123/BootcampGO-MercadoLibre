package main

import "fmt"

func main() {

	var minutosTrabajados int = 1200
	var categoria string = "C"
	var salario float64 = CalcularSalario(minutosTrabajados, categoria)

	fmt.Printf("El trabajor X, con %d minutos trabajados y categoria %s, su salario correspondiente es: %f",
		minutosTrabajados,
		categoria,
		salario)
}

func CalcularSalario(minutosTrabajados int, categoria string) (result float64) {
	var horasTrabajadas float64 = float64(minutosTrabajados) / float64(60)

	switch categoria {
	case "A":
		salario := float64(3000) * horasTrabajadas
		result = salario + (salario * 0.5)
	case "B":
		salario := float64(1500) * horasTrabajadas
		result = salario + (salario * 0.2)

	case "C":
		result = float64(1000) * horasTrabajadas
	}
	return
}
