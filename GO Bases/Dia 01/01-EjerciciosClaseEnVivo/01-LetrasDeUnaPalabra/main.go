package main

import (
	"fmt"
)

func main() {
	//Declariación de variables
	var palabra string

	fmt.Print("Ingrese una palabra: ")
	fmt.Scanln(&palabra)

	fmt.Println("longitud: ", len(palabra))
	//Recorremos cada palabra
	for _, letra := range palabra {
		fmt.Println(string(letra))
	}
}
