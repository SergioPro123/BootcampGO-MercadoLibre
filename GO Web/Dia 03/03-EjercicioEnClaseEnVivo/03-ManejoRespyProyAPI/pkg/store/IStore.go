package store

import "proyectoapisupermercado/internal/domain"

type IStore interface {
	GetProducts() ([]domain.Product, error)
	Addproduct(product domain.Product) (domain.Product, error)
	UpdateProduct(product domain.Product) (productUpdated domain.Product, err error)
	UpdateProductPartial(product domain.Product) (productUpdated domain.Product, err error)
	DeleteProduct(id int) (productDelete bool, err error)
}
