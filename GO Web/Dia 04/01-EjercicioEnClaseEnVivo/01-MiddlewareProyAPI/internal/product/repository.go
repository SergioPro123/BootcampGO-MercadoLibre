package product

import (
	"encoding/json"
	"fmt"
	"os"
	"proyectoapisupermercado/internal/domain"
	"proyectoapisupermercado/pkg/store"
)

type ProductRepository struct {
	store store.IStore
}

var (
	stringPathFileProducts = "../../internal/product/products.json"

	ErrorNoFoundProduct = "No found product"
)

func NewProductRepository(store store.IStore) IProductRepository {
	return &ProductRepository{store: store}
}
func (p *ProductRepository) GetProducts() (products []domain.Product, err error) {
	return p.store.GetProducts()
}
func (p *ProductRepository) GetProductById(id int) (product domain.Product, err error) {

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
func (p *ProductRepository) GetProductsByPrice(price float64) (productsReturn []domain.Product, err error) {

	products, err := p.GetProducts()
	if err != nil {
		err = fmt.Errorf("ocurrio un error: %w", err)
		return
	}
	productsReturn = []domain.Product{}
	for _, productFor := range products {
		if productFor.Price >= price {
			productsReturn = append(productsReturn, productFor)
		}
	}
	return
}

func (p *ProductRepository) ExistProduct(codeValue string) (exist bool) {

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

func (p *ProductRepository) Addproduct(product domain.Product) (domain.Product, error) {
	var err error
	product.Id, err = p.generateId()
	if err != nil {
		return product, fmt.Errorf("ocurrio un error: %w", err)
	}

	return p.store.Addproduct(product)
}

func (p *ProductRepository) UpdateProduct(product domain.Product) (productUpdated domain.Product, err error) {

	return p.store.UpdateProduct(product)
}
func (p *ProductRepository) UpdateProductPartial(product domain.Product) (productUpdated domain.Product, err error) {

	products, err := p.GetProducts()
	if err != nil {
		err = fmt.Errorf("ocurrio un error: %w", err)
		return
	}
	exist := false
	var productBefore *domain.Product

	for i, productFor := range products {
		if productFor.Id != product.Id {
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
func (p *ProductRepository) DeleteProduct(id int) (productDelete bool, err error) {

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

func (p *ProductRepository) generateId() (id int, err error) {
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
