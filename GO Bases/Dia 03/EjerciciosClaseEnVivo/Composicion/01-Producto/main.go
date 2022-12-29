package main

import "fmt"

type Product struct {
	ID          int
	Name        string
	Price       float64
	Description string
	Category    string
}

var Products = []Product{
	{
		ID:          1,
		Name:        "Televisor",
		Price:       12.5,
		Description: "Televisor 45 pulgadas.",
		Category:    "A",
	},
}

func (p *Product) Save() {
	Products = append(Products, *p)
}
func (p *Product) GetAll() {
	for _, product := range Products {
		fmt.Printf("%+v \n", product)
	}
}
func getById(id int) *Product {
	for _, product := range Products {
		if product.ID == id {
			return &product
		}
	}
	panic("Producto no existe.")
}

func main() {
	nevera := Product{
		ID:          2,
		Name:        "Nevera",
		Price:       39.99,
		Description: "Nevera Samsung",
		Category:    "B",
	}
	//Utilizar metodo "Save"
	nevera.Save()
	//Utilizar metodo "GetAll"
	nevera.GetAll()
	//Utilizar función "getById"
	fmt.Printf("%+v \n", getById(1))
}
