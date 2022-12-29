package main

import "fmt"

const (
	prodPequenio = "pequenio"
	prodMediano  = "mediano"
	prodGrande   = "grande"
)

// ------------- Interfaces -------------
type IProducto interface {
	Precio() float64
}

// ------------- Estructuras -------------
type ProductoPequenio struct {
	IProducto
	Costo float64
}
type ProductoMediano struct {
	IProducto
	Costo        float64
	CostoEnStock float64
}
type ProductoGrande struct {
	IProducto
	Costo        float64
	CostoEnStock float64
	CostoEnvio   float64
}

// ------------- Metodos -------------
func (p *ProductoPequenio) Precio() float64 {
	return p.Costo
}
func (p *ProductoMediano) Precio() float64 {
	return p.Costo + (p.Costo * p.CostoEnStock)
}
func (p *ProductoGrande) Precio() float64 {
	return p.Costo + (p.Costo * p.CostoEnStock) + p.CostoEnvio
}

func FactoryProducto(tipoProducto string, precio float64) (instance IProducto, err bool) {
	switch tipoProducto {
	case prodPequenio:
		instance = &ProductoPequenio{
			Costo: precio,
		}
	case prodMediano:
		instance = &ProductoMediano{
			Costo:        precio,
			CostoEnStock: 0.03,
		}
	case prodGrande:
		instance = &ProductoGrande{
			Costo:        precio,
			CostoEnStock: 0.06,
			CostoEnvio:   2500,
		}
	default:
		instance = nil
		err = true
	}
	return
}

func main() {
	var tipoProductos []string = []string{prodPequenio, prodMediano, prodGrande, "ErrorProducto"}

	//Recorremos todos los tipos de producto en el array
	for _, tipoProducto := range tipoProductos {

		producto, err := FactoryProducto(tipoProducto, 10000)

		if err {
			fmt.Printf("Hubo un error con el tipo producto: '%s'", tipoProducto)
			continue
		}

		fmt.Printf("Precio del producto %s es: $%.2f \n", tipoProducto, producto.Precio())
	}
}
