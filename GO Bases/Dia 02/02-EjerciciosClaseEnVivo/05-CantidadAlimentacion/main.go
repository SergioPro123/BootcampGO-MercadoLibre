package main

import (
	"errors"
	"fmt"
)

type OperacionAnimal func(cantidadAnimal int) (result float32, unidadMasa string)

const (
	perro     = "perro"
	gato      = "gato"
	hamster   = "hamster"
	tarantula = "tarantula"
)

func main2() {
	cantidadAnimal := 10

	animalPerro, errPerro := orquestadorAnimal(perro)
	animalGato, errGato := orquestadorAnimal(gato)
	animalHamster, errHamster := orquestadorAnimal(hamster)
	animalTarantula, errTarantula := orquestadorAnimal(tarantula)

	if errPerro != nil {
		panic(errPerro.Error())
	}
	if errGato != nil {
		panic(errGato.Error())
	}
	if errHamster != nil {
		panic(errHamster.Error())
	}
	if errTarantula != nil {
		panic(errTarantula.Error())
	}

	cantAlimentoPerro, uMasaPerro := animalPerro(cantidadAnimal)
	cantAlimentoGato, uMasaGato := animalGato(cantidadAnimal)
	cantAlimentoHamster, uMasaHamster := animalHamster(cantidadAnimal)
	cantAlimentoTarantula, uMasaTarantula := animalTarantula(cantidadAnimal)

	fmt.Printf("La cantidad de alimento para %d perros es de : %.2f %s\n", cantidadAnimal, cantAlimentoPerro, uMasaPerro)
	fmt.Printf("La cantidad de alimento para %d gatos es de : %.2f %s\n", cantidadAnimal, cantAlimentoGato, uMasaGato)
	fmt.Printf("La cantidad de alimento para %d hamster es de : %.2f %s\n", cantidadAnimal, cantAlimentoHamster, uMasaHamster)
	fmt.Printf("La cantidad de alimento para %d tarantula es de : %.2f %s\n", cantidadAnimal, cantAlimentoTarantula, uMasaTarantula)
}

func main() {
	var animales = []string{perro, gato, hamster, tarantula}
	cantidadAnimal := 10

	for i, animal := range animales {
		operacionAnimal, err := orquestadorAnimal(animal)

		if err != nil {
			panic(err.Error())
		}
		cantAlimento, uMasa := operacionAnimal(cantidadAnimal)

		fmt.Printf("%d) La cantidad de alimento para %d %s es de : %.2f %s\n", i+1, cantidadAnimal, animal, cantAlimento, uMasa)
	}
}

func orquestadorAnimal(nombreAnimal string) (operacionAnimal OperacionAnimal, err error) {

	switch nombreAnimal {
	case perro:
		operacionAnimal = operacionPerro
	case gato:
		operacionAnimal = operacionGato
	case hamster:
		operacionAnimal = operacionHamster
	case tarantula:
		operacionAnimal = operacionTarantula
	default:
		err = errors.New("Animal " + nombreAnimal + " no valido.")
	}

	return
}

func operacionPerro(cantidadAnimal int) (result float32, unidadMasa string) {
	unidadMasa = "Kg"
	result = float32(cantidadAnimal) * 10
	return
}
func operacionGato(cantidadAnimal int) (result float32, unidadMasa string) {
	unidadMasa = "Kg"
	result = float32(cantidadAnimal) * 5
	return
}
func operacionHamster(cantidadAnimal int) (result float32, unidadMasa string) {
	unidadMasa = "gramos"
	result = float32(cantidadAnimal) * 250
	return
}
func operacionTarantula(cantidadAnimal int) (result float32, unidadMasa string) {
	unidadMasa = "gramos"
	result = float32(cantidadAnimal) * 150
	return
}
