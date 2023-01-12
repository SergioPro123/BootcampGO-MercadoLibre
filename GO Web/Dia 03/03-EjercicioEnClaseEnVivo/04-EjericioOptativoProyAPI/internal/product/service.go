package product

import (
	"fmt"
	"proyectoapisupermercado/internal/domain"
)

type ProductService struct {
	productRepository IProductRepository
}
type ErrorInformative struct {
	Message string
}

func (e *ErrorInformative) Error() string {
	return e.Message
}

// Service
type IProductService interface {
	GetProducts() ([]domain.Product, error)
	GetProductById(id int) (domain.Product, error)
	GetProductsByPrice(price float64) (products []domain.Product, err error)
	GetConsumerPrice(listIds []int) (productsReturn []domain.Product, priceTotal float64, err error)
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
	if err != nil {
		return
	}
	//Organizamos los objetos por id
	for i := 0; i < len(products); i++ {
		for j := i + 1; j < len(products)-1; j++ {
			if products[i].Id > products[j].Id {
				tmp := products[i]
				products[i] = products[j]
				products[j] = tmp
			}
		}
	}
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
func (p *ProductService) GetConsumerPrice(listIds []int) (productsReturn []domain.Product, priceTotal float64, err error) {

	products, err := p.GetProducts()
	if err != nil {
		err = fmt.Errorf("ocurrio un error: %w", err)
		return
	}
	productsReturn = []domain.Product{}
	//Recorremos los los productos existentes o hasta que se terminen los productos solicitados
	//NOTA: Este ciclo for se optimizo, para que se cumpla los objetos tipo []domain.Product deben estar ordenados.
	var indIdSolicitado int = 0
	var relProdQuantity map[int]int = make(map[int]int) // id-quantity
	for i := 0; i < len(products) && indIdSolicitado < len(listIds); i++ {
		//Validamos si el producto del ciclo actual es solicitado
		if products[i].Id != listIds[indIdSolicitado] {
			continue
		}
		//Validamos si aun queda cantidad en stock, de lo contrario retornamos error
		if relProdQuantity[products[i].Id] >= *products[i].Quantity {
			err = &ErrorInformative{
				Message: fmt.Sprintf("el producto '%s' con ID %d solo contiene %d en stock",
					products[i].Name, products[i].Id, *products[i].Quantity),
			}
			return
		}
		relProdQuantity[products[i].Id] = relProdQuantity[products[i].Id] + 1
		//Agregamos el producto que se quiere devolver
		productsReturn = append(productsReturn, products[i])
		priceTotal += products[i].Price

		if (indIdSolicitado+1) < len(listIds) && listIds[indIdSolicitado+1] == listIds[indIdSolicitado] {
			i--
		}
		indIdSolicitado++
	}
	//Calculamos los impuestos correspondientes al precio final segun la cantidad de productos
	var porImpuesto float64
	switch {
	case len(productsReturn) < 10:
		porImpuesto = 1.21
	case len(productsReturn) <= 20:
		porImpuesto = 1.17
	default:
		porImpuesto = 1.15
	}
	priceTotal = priceTotal * porImpuesto

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
