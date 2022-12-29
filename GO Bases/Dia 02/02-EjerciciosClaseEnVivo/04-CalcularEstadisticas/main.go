package main

import (
	"errors"
	"fmt"
)

const (
	minimo   = "minimo"
	promedio = "promedio"
	maximo   = "maximo"
)

func main() {
	var notas = []float64{3.5, 7.8, 1.2, 9.7}

	operacionMin, errMin := OrquestadorOperaciones(minimo)
	operacionPro, errMax := OrquestadorOperaciones(promedio)
	operacionMax, errPro := OrquestadorOperaciones(maximo)

	if errMin != nil {
		panic(errMin.Error())
	}
	if errMax != nil {
		panic(errMax.Error())
	}
	if errPro != nil {
		panic(errPro.Error())
	}

	fmt.Println("El resultado de la operacion minimo es: ", operacionMin(notas...))
	fmt.Println("El resultado de la operacion del promedio es: ", operacionPro(notas...))
	fmt.Println("El resultado de la operacion maxima  es: ", operacionMax(notas...))
}

func OrquestadorOperaciones(nombreOperacion string) (operacion func(notas ...float64) (result float64), err error) {
	switch nombreOperacion {
	case minimo:
		operacion = calculateMin
	case promedio:
		operacion = calculateProm
	case maximo:
		operacion = calculateMax
	default:
		err = errors.New("Operacion invalida")
	}

	return
}

func calculateMin(notas ...float64) (result float64) {
	result = notas[0]

	for _, nota := range notas {
		if nota < result {
			result = nota
		}
	}
	return
}
func calculateProm(notas ...float64) (result float64) {
	var sumaNotas float64
	for _, nota := range notas {
		sumaNotas += nota
	}

	result = sumaNotas / float64(len(notas))

	return
}
func calculateMax(notas ...float64) (result float64) {
	result = notas[0]

	for _, nota := range notas {
		if nota > result {
			result = nota
		}
	}

	return
}
