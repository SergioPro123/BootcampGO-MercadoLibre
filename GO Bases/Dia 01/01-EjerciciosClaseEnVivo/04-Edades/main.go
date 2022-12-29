package main

import (
	"fmt"
)

func main() {
	var employees = map[string]int{
		"Benjamin": 20,
		"Nahuel":   26,
		"Brenda":   19,
		"Darío":    44,
		"Pedro":    30}
	var nombreBuscar string = "Benjamin"

	//Imprimimos la edad de Benjamin
	fmt.Printf("La edad de %s es de %d \n", nombreBuscar, employees[nombreBuscar])
	//Imprimir usuario que tengan mas de 21 años
	fmt.Println("------------------ Usuario mayores de 21 años ----------------")
	for nombre, edad := range employees {
		if edad > 21 {
			fmt.Printf("El usuario %s tiene %d años \n", nombre, edad)
		}
	}
	//Agregamos un nuevo usuario
	employees["Federico"] = 25
	//Eliminamos a Pedro
	delete(employees, "Pedro")
	//Imprimimos todos los usuario existentes hasta este punto.
	fmt.Println("------------------ Usuario Actualizados ----------------")
	for nombre, edad := range employees {
		fmt.Printf("El usuario %s tiene %d años \n", nombre, edad)
	}
}
