package product

import (
	"encoding/json"
	"errors"
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

	ErrorNoFoundProduct = errors.New("no found product")
)

func NewProductRepository(store store.IStore) IProductRepository {
	return &ProductRepository{store: store}
}
func (p *ProductRepository) GetProducts() (products []domain.Product, err error) {
	products, err = p.store.GetProducts()
	if err != nil {
		err = fmt.Errorf("ocurrio un error: %w", err)
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
		err = ErrorNoFoundProduct
		return
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
func (p *ProductRepository) GetConsumerPrice(listIds []int) (productsReturn []domain.Product, priceTotal float64, err error) {

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
	for i := 0; i < len(products) || indIdSolicitado < len(listIds); i++ {
		//Validamos si el producto del ciclo actual es solicitado
		if products[i].Id != listIds[indIdSolicitado] {
			continue
		}
		//Validamos si aun queda cantidad en stock
		if relProdQuantity[products[i].Id] >= *products[i].Quantity {
			continue
		}
		relProdQuantity[products[i].Id] = relProdQuantity[products[i].Id] + 1
		//Agregamos el producto que se quiere devolver
		productsReturn = append(productsReturn, products[i])
		priceTotal += products[i].Price

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
		err = ErrorNoFoundProduct
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
		err = ErrorNoFoundProduct
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
