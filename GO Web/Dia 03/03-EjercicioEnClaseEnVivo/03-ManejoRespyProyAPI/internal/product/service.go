package product

import (
	"fmt"
	"proyectoapisupermercado/internal/domain"
)

type ProductService struct {
	productRepository IProductRepository
}

// Service
type IProductService interface {
	GetProducts() ([]domain.Product, error)
	GetProductById(id int) (domain.Product, error)
	GetProductsByPrice(price float64) (products []domain.Product, err error)
	UpdateProduct(product domain.Product) (productUpdated domain.Product, err error)
	UpdateProductPartial(product domain.Product) (productUpdated domain.Product, err error)
	DeleteProduct(id int) (productDelete bool, err error)
	Addproduct(product domain.Product) (domain.Product, error)
}

func NewProductService(productRepository IProductRepository) IProductService {
	return &ProductService{productRepository: productRepository}
}

func (p *ProductService) GetProducts() (products []domain.Product, err error) {
	products, err = p.productRepository.GetProducts()
	return
}
func (p *ProductService) GetProductById(id int) (product domain.Product, err error) {
	product, err = p.productRepository.GetProductById(id)
	return
}
func (p *ProductService) GetProductsByPrice(price float64) (products []domain.Product, err error) {
	products, err = p.productRepository.GetProductsByPrice(price)
	return
}
func (p *ProductService) Addproduct(product domain.Product) (productAdded domain.Product, err error) {
	existProduct := p.productRepository.ExistProduct(product.CodeValue)
	if existProduct {
		err = fmt.Errorf("producto ya existe")
		return
	}
	productAdded, err = p.productRepository.Addproduct(product)

	return
}
func (p *ProductService) UpdateProduct(product domain.Product) (productUpdated domain.Product, err error) {
	productUpdated, err = p.productRepository.UpdateProduct(product)

	return
}
func (p *ProductService) UpdateProductPartial(product domain.Product) (productUpdated domain.Product, err error) {
	productUpdated, err = p.productRepository.UpdateProductPartial(product)

	return
}
func (p *ProductService) DeleteProduct(id int) (productDelete bool, err error) {
	productDelete, err = p.productRepository.DeleteProduct(id)

	return
}
