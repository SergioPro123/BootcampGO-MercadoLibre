package main

import (
	"fmt"
	"time"
)

func main() {
	var diaMes int
	var mesesEspañol = [12]string{
		"Enero",
		"Febrero",
		"Marzo",
		"Abril",
		"Mayo",
		"Junio",
		"Julio",
		"Agosto",
		"Septiembre",
		"Octubre",
		"Noviembre",
		"Diciembre",
	}

	fmt.Print("Digite el numero del mes: ")
	fmt.Scanln(&diaMes)

	//Validamos que el numero del mes sea el correcto
	if !(diaMes >= 1 && diaMes <= 12) {
		fmt.Println("¡Numero invalido!")
		return
	}

	fmt.Println("Mes en ingles: ", time.Month(diaMes))
	fmt.Println("Mes en español: ", mesesEspañol[diaMes-1])

}
