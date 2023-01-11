package store

import (
	"encoding/json"
	"fmt"
	"os"
	"proyectoapisupermercado/internal/domain"
)

type Store struct {
	fileName string
}

var (
	ErrorNoFoundProduct = "No found product"
)

func NewStore(fileName string) IStore {
	return &Store{fileName: fileName}
}
func (s *Store) saveProducts(products []domain.Product) (err error) {

	result, err := json.Marshal(products)
	if err != nil {
		err = fmt.Errorf("ocurrio un error: %w", err)
		return
	}
	err = os.WriteFile((s.fileName), result, os.ModeAppend)
	if err != nil {
		err = fmt.Errorf("ocurrio un error: %w", err)
		return
	}
	return
}
func (s *Store) GetProducts() (products []domain.Product, err error) {
	defer func() {
		existPanic := recover()
		if existPanic != nil {
			fmt.Println("Metodo GetProducts [Store]: Finalizo mal.")
			return
		}
	}()

	fileBytes, err := os.ReadFile(s.fileName)
	if err != nil {
		err = fmt.Errorf("el archivo indicado no fue encontrado o está dañado: %w", err)
		return
	}

	json.Unmarshal(fileBytes, &products)
	return
}
func (s *Store) Addproduct(product domain.Product) (productAdded domain.Product, err error) {
	products, err := s.GetProducts()
	if err != nil {
		err = fmt.Errorf("ocurrio un error: %w", err)
		return
	}
	products = append(products, product)
	err = s.saveProducts(products)
	if err != nil {
		err = fmt.Errorf("ocurrio un error: %w", err)
		return
	}

	productAdded = product
	return
}

func (s *Store) UpdateProduct(product domain.Product) (productUpdated domain.Product, err error) {

	products, err := s.GetProducts()
	if err != nil {
		err = fmt.Errorf("ocurrio un error: %w", err)
		return
	}
	exist := false

	for i, productFor := range products {
		if productFor.Id != product.Id {
			continue
		}

		products[i] = product
		productUpdated = products[i]
		exist = true
		break
	}
	if !exist {
		err = fmt.Errorf(ErrorNoFoundProduct)
		return
	}

	err = s.saveProducts(products)
	if err != nil {
		err = fmt.Errorf("ocurrio un error: %w", err)
		return
	}

	return
}
func (s *Store) UpdateProductPartial(product domain.Product) (productUpdated domain.Product, err error) {

	products, err := s.GetProducts()
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

	err = s.saveProducts(products)
	if err != nil {
		err = fmt.Errorf("ocurrio un error: %w", err)
		return
	}

	return
}
func (s *Store) DeleteProduct(id int) (productDelete bool, err error) {

	products, err := s.GetProducts()
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

	err = s.saveProducts(products)
	if err != nil {
		err = fmt.Errorf("ocurrio un error: %w", err)
		return
	}
	productDelete = true

	return
}
