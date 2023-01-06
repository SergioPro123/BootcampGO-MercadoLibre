package data

import (
	"encoding/json"
	"fmt"
	"os"
	"proyectoapisupermercado/interfaces"
	"proyectoapisupermercado/models"
)

type DatabaseProductJson struct {
	interfaces.IDatabaseProduct
	fileBytes []byte
}

func (d *DatabaseProductJson) GetProducts() (products []models.Product, err error) {
	defer func() {
		existPanic := recover()
		if existPanic != nil {
			fmt.Println("Metodo GetProducts [DatabaseProductJson]: Finalizo mal.")
			return
		}
	}()

	d.fileBytes, err = os.ReadFile("data/products.json")
	if err != nil {
		err = fmt.Errorf("el archivo indicado no fue encontrado o está dañado")
		return
	}

	json.Unmarshal(d.fileBytes, &products)
	return
}
