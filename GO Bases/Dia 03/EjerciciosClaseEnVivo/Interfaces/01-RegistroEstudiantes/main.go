package main

import (
	"fmt"
	"time"
)

type Estudiante struct {
	Nombre   string
	Apellido string
	DNI      int
	Fecha    time.Time
}

func (e *Estudiante) getDetalle() {
	fmt.Printf("%+v", e)
}

func main() {
	var estudiante Estudiante = Estudiante{
		Nombre:   "Sergio",
		Apellido: "Aparicio",
		DNI:      1007724255,
		Fecha:    time.Date(2000, 2, 12, 0, 0, 0, 0, time.Local),
	}
	estudiante.getDetalle()
}
