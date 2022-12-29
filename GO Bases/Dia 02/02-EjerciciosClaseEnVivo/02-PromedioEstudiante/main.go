package main

import "fmt"

func main() {
	var notas = []float64{3.44, 4.56, 1.5}
	var promedio float64 = calcularNotas(notas...)

	fmt.Println("--------- Notas ---------")
	for i, nota := range notas {
		fmt.Printf("Nota #%d = %.2f \n", i, nota)
	}
	fmt.Printf("--------- Promedio = %.2f ---------", promedio)

}

func calcularNotas(notas ...float64) (promedio float64) {

	var sumaNotas float64
	for _, nota := range notas {
		sumaNotas += nota
	}

	promedio = sumaNotas / float64(len(notas))

	return
}
