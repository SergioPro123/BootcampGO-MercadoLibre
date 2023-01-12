package product

import "proyectoapisupermercado/internal/domain"

type IProductRepository interface {
	GetProducts() ([]domain.Product, error)
	GetProductById(id int) (domain.Product, error)
	GetProductsByPrice(price float64) (products []domain.Product, err error)
	Addproduct(product domain.Product) (domain.Product, error)
	UpdateProduct(product domain.Product) (productUpdated domain.Product, err error)
	UpdateProductPartial(product domain.Product) (productUpdated domain.Product, err error)
	DeleteProduct(id int) (productDelete bool, err error)
	ExistProduct(codeValue string) (exist bool)
}
