package product

import (
	"encoding/json"
	"fmt"
	"os"
	"proyectoapisupermercado/internal/domain"
)

type ProductRepositoryJson struct {
}

var (
	stringPathFileProducts = "../internal/product/products.json"

	ErrorNoFoundProduct = "No found product"
)

type IProductRepository interface {
	GetProducts() ([]domain.Product, error)
	GetProductById(id int) (domain.Product, error)
	GetProductsByPrice(price float64) (products []domain.Product, err error)
	Addproduct(product domain.Product) (domain.Product, error)
	UpdateAllProduct(id int, product domain.Product) (productUpdated domain.Product, err error)
	UpdateProductPartial(id int, product domain.Product) (productUpdated domain.Product, err error)
	DeleteProduct(id int) (productDelete bool, err error)
	ExistProduct(codeValue string) (exist bool)
}

func NewProductRepository() IProductRepository {
	return &ProductRepositoryJson{}
}
func (d *ProductRepositoryJson) GetProducts() (products []domain.Product, err error) {
	defer func() {
		existPanic := recover()
		if existPanic != nil {
			fmt.Println("Metodo GetProducts [ProductRepositoryJson]: Finalizo mal.")
			return
		}
	}()

	fileBytes, err := os.ReadFile(stringPathFileProducts)
	if err != nil {
		err = fmt.Errorf("el archivo indicado no fue encontrado o está dañado: %w", err)
		return
	}

	json.Unmarshal(fileBytes, &products)
	return
}
func (p *ProductRepositoryJson) GetProductById(id int) (product domain.Product, err error) {

	products, err := p.GetProducts()
	if err != nil {
		err = fmt.Errorf("ocurrio un error: %w", err)
		return
	}

	var existProduct bool = false
	for _, productFor := range products {
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
func (p *ProductRepositoryJson) GetProductsByPrice(price float64) (productsReturn []domain.Product, err error) {

	products, err := p.GetProducts()
	if err != nil {
		err = fmt.Errorf("ocurrio un error: %w", err)
		return
	}

	for _, productFor := range products {
		if productFor.Price >= price {
			productsReturn = append(productsReturn, productFor)
		}
	}
	return
}

func (p *ProductRepositoryJson) ExistProduct(codeValue string) (exist bool) {

	products, err := p.GetProducts()
	if err != nil {
		exist = false
		return
	}

	for _, productFor := range products {
		if productFor.CodeValue == codeValue {
			exist = true
			break
		}
	}
	return
}

func (p *ProductRepositoryJson) Addproduct(product domain.Product) (productAdded domain.Product, err error) {
	products, err := p.GetProducts()
	if err != nil {
		err = fmt.Errorf("ocurrio un error: %w", err)
		return
	}
	product.Id, err = p.generateId()
	if err != nil {
		err = fmt.Errorf("ocurrio un error: %w", err)
		return
	}

	products = append(products, product)

	result, err := json.Marshal(products)
	if err != nil {
		err = fmt.Errorf("ocurrio un error: %w", err)
		return
	}
	err = os.WriteFile((stringPathFileProducts), result, os.ModeAppend)
	if err != nil {
		err = fmt.Errorf("ocurrio un error: %w", err)
		return
	}

	productAdded = product
	return
}

func (p *ProductRepositoryJson) UpdateAllProduct(id int, product domain.Product) (productUpdated domain.Product, err error) {

	products, err := p.GetProducts()
	if err != nil {
		err = fmt.Errorf("ocurrio un error: %w", err)
		return
	}
	exist := false

	for i, productFor := range products {
		if productFor.Id != id {
			continue
		}

		products[i] = product
		products[i].Id = id
		productUpdated = products[i]
		exist = true
		break
	}
	if !exist {
		err = fmt.Errorf(ErrorNoFoundProduct)
		return
	}

	result, err := json.Marshal(products)
	if err != nil {
		err = fmt.Errorf("ocurrio un error: %w", err)
		return
	}
	err = os.WriteFile((stringPathFileProducts), result, os.ModeAppend)
	if err != nil {
		err = fmt.Errorf("ocurrio un error: %w", err)
		return
	}

	return
}
func (p *ProductRepositoryJson) UpdateProductPartial(id int, product domain.Product) (productUpdated domain.Product, err error) {

	products, err := p.GetProducts()
	if err != nil {
		err = fmt.Errorf("ocurrio un error: %w", err)
		return
	}
	exist := false
	var productBefore *domain.Product

	for i, productFor := range products {
		if productFor.Id != id {
			continue
		}
		productBefore = &products[i]
		exist = true
		break
	}
	if !exist {
		err = fmt.Errorf(ErrorNoFoundProduct)
		return
	}

	if product.CodeValue != "" {
		productBefore.CodeValue = product.CodeValue
	}
	if product.Expiration != "" {
		productBefore.Expiration = product.Expiration
	}
	if product.Name != "" {
		productBefore.Name = product.Name
	}
	if product.Price != 0 {
		productBefore.Price = product.Price
	}
	if product.Quantity != nil {
		productBefore.Quantity = product.Quantity
	}
	if product.IsPublished != nil {
		productBefore.IsPublished = product.IsPublished
	}
	productUpdated = *productBefore
	result, err := json.Marshal(products)
	if err != nil {
		err = fmt.Errorf("ocurrio un error: %w", err)
		return
	}
	err = os.WriteFile((stringPathFileProducts), result, os.ModeAppend)
	if err != nil {
		err = fmt.Errorf("ocurrio un error: %w", err)
		return
	}

	return
}
func (p *ProductRepositoryJson) DeleteProduct(id int) (productDelete bool, err error) {

	products, err := p.GetProducts()
	if err != nil {
		err = fmt.Errorf("ocurrio un error: %w", err)
		return
	}
	exist := false

	for i, productFor := range products {
		if productFor.Id != id {
			continue
		}

		products = append(products[:i], products[i+1:]...)
		exist = true
		break
	}
	if !exist {
		err = fmt.Errorf(ErrorNoFoundProduct)
		return
	}

	result, err := json.Marshal(products)
	if err != nil {
		err = fmt.Errorf("ocurrio un error: %w", err)
		return
	}
	err = os.WriteFile((stringPathFileProducts), result, os.ModeAppend)
	if err != nil {
		err = fmt.Errorf("ocurrio un error: %w", err)
		return
	}
	productDelete = true

	return
}

func (p *ProductRepositoryJson) generateId() (id int, err error) {
	products, err := p.GetProducts()
	if err != nil {
		err = fmt.Errorf("ocurrio un error: %w", err)
		return
	}
	lastId := 0

	for _, productFor := range products {
		if productFor.Id > lastId {
			lastId = productFor.Id
		}
	}

	id = lastId + 1
	return
}
