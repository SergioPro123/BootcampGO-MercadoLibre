package main

import (
	"fmt"
)

func main() {
	var salario float64 = 190000
	impuesto, porcentaje := ImpuestoSalario(salario)

	fmt.Printf("Salario: %.2f | Impuesto : %.2f (%d)| Total: %.2f", salario, impuesto, porcentaje, (salario - impuesto))

}

func ImpuestoSalario(salario float64) (result float64, porcentajeImpuesto int) {

	if salario > 50000 {
		porcentajeImpuesto += 17
	}
	if salario > 150000 {
		porcentajeImpuesto += 10
	}

	result = salario * (float64(porcentajeImpuesto) / float64(100.))

	return
}
