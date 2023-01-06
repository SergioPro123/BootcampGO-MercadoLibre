package controllers

import (
	"fmt"
	"proyectoapisupermercado/interfaces"
	"proyectoapisupermercado/models"
	"sync"
)

type ProductController struct {
	products []models.Product
	lastId   int
	err      error
}

var lock = &sync.Mutex{}

var singleInstance *ProductController

func GetInstance() *ProductController {
	if singleInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if singleInstance == nil {
			fmt.Println("[ProductController]Creating single instance now.")
			singleInstance = &ProductController{}
		} else {
			fmt.Println("[ProductController]:Single instance already created.")
		}
	} else {
		fmt.Println("[ProductController]:Single instance already created.")
	}

	return singleInstance
}

func (p *ProductController) Init(databaseProduct interfaces.IDatabaseProduct) (err error) {

	defer func() {
		existPanic := recover()
		if existPanic != nil {
			fmt.Println("Metodo Init: Finalizo mal.")
			return
		}
	}()

	p.products, err = databaseProduct.GetProducts()
	if err != nil {
		p.err = err
		return
	}
	p.assignLastId()

	return
}
func (p *ProductController) GetProducts() (products []models.Product) {
	products = p.products
	return
}
func (p *ProductController) GetProductById(id int) (product models.Product, err error) {

	var existProduct bool = false
	for _, productFor := range p.products {
		if productFor.Id == id {
			product = productFor
			existProduct = true
			break
		}
	}
	if !existProduct {
		err = fmt.Errorf("no se encontro producto por el id '%d'", id)
	}

	return
}
func (p *ProductController) GetProductsByPrice(price float64) (products []models.Product) {
	products = []models.Product{}

	for _, productFor := range p.products {
		if productFor.Price >= price {
			products = append(products, productFor)
		}
	}
	return
}

func (p *ProductController) ExistProduct(codeValue string) (exist bool) {

	for _, productFor := range p.products {
		if productFor.CodeValue == codeValue {
			exist = true
			break
		}
	}
	return
}

func (p *ProductController) Addproduct(product models.Product) (productAdded models.Product) {
	product.Id = p.GenerateId()
	p.products = append(p.products, product)
	return product
}

func (p *ProductController) GenerateId() (id int) {
	id = p.lastId + 1
	p.lastId = id
	return
}

func (p *ProductController) assignLastId() {

	if len(p.products) <= 0 {
		return
	}
	p.lastId = p.products[0].Id
	for _, productFor := range p.products {
		if productFor.Id > p.lastId {
			p.lastId = productFor.Id
		}
	}
}
