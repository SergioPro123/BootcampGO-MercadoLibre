package main

import (
	"fmt"
	"os"
)

func main() {
	leerArchivo()
}

func leerArchivo() {
	var file *os.File
	var err error

	defer func() {
		file.Close()
		existPanic := recover()
		if existPanic != nil {
			fmt.Println(existPanic)
		}

		fmt.Println("ejecución finalizada")
	}()

	file, err = os.Open("archivo.txt")
	if err != nil {
		panic("el archivo indicado no fue encontrado o está dañado")
	}
}
