package handlers

import (
	"errors"
	"net/http"
	"proyectoapisupermercado/internal/domain"
	"proyectoapisupermercado/internal/product"
	"proyectoapisupermercado/pkg/web"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var (
	ErrorGenerated = "ocurrio un error"
	ErrorInvalidId = "id invalido"
)

type ProductHandler struct {
	productService product.IProductService
}

func NewProductHandler(productService product.IProductService) *ProductHandler {
	return &ProductHandler{productService: productService}
}

func (p *ProductHandler) AddProduct() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		var product domain.ProductRequest
		err := ctx.BindJSON(&product)

		if err != nil {
			web.Failure(ctx, http.StatusBadRequest, errors.New(ErrorGenerated))
			return
		}

		validate := validator.New()
		err = validate.Struct(product)
		if err != nil {
			web.Failure(ctx, http.StatusBadRequest, err)
			return
		}

		productResponse, err := p.productService.Addproduct(domain.Product{
			Name:        product.Name,
			Quantity:    product.Quantity,
			CodeValue:   product.CodeValue,
			IsPublished: product.IsPublished,
			Expiration:  product.Expiration,
			Price:       product.Price,
		})
		if err != nil {
			web.Failure(ctx, http.StatusBadRequest, err)
			return
		}
		web.Success(ctx, http.StatusCreated, productResponse)
	}
}

func (p *ProductHandler) GetProducts() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		products, err := p.productService.GetProducts()
		if err != nil {
			web.Failure(ctx, http.StatusInternalServerError, err)
			return
		}
		web.Success(ctx, http.StatusAccepted, products)
	}
}

func (p *ProductHandler) GetProductById() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			web.Failure(ctx, http.StatusBadRequest, errors.New(ErrorInvalidId))
			return
		}
		product, err := p.productService.GetProductById(id)
		if err != nil {
			web.Failure(ctx, http.StatusBadRequest, err)
			return
		}
		web.Success(ctx, http.StatusAccepted, product)
	}
}
func (p *ProductHandler) GetProductByPrice() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		priceGt, err := strconv.ParseFloat(ctx.Query("priceGt"), 64)
		if err != nil {
			web.Failure(ctx, http.StatusBadRequest, errors.New("invalido 'priceGt'"))
			return
		}

		products, err := p.productService.GetProductsByPrice(priceGt)

		if err != nil {
			web.Failure(ctx, http.StatusInternalServerError, errors.New(ErrorGenerated))
			return
		}
		web.Success(ctx, http.StatusAccepted, products)
	}
}
func (p *ProductHandler) UpdateAllProduct() gin.HandlerFunc {

	return func(ctx *gin.Context) {

		var product domain.ProductRequest
		err := ctx.BindJSON(&product)

		if err != nil {
			web.Failure(ctx, http.StatusBadRequest, errors.New(ErrorGenerated))
			return
		}

		validate := validator.New()
		err = validate.Struct(product)
		if err != nil {
			web.Failure(ctx, http.StatusBadRequest, err)
			return
		}

		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			web.Failure(ctx, http.StatusBadRequest, errors.New(ErrorInvalidId))
			return
		}

		products, err := p.productService.UpdateProduct(domain.Product{
			Id:          id,
			Name:        product.Name,
			Quantity:    product.Quantity,
			CodeValue:   product.CodeValue,
			IsPublished: product.IsPublished,
			Expiration:  product.Expiration,
			Price:       product.Price,
		})

		if err != nil {
			web.Failure(ctx, http.StatusInternalServerError, errors.New(ErrorGenerated))
			return
		}
		web.Success(ctx, http.StatusAccepted, products)
	}
}
func (p *ProductHandler) UpdateProductPartial() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		var product domain.ProductRequestUpdatePartial
		err := ctx.BindJSON(&product)

		if err != nil {
			web.Failure(ctx, http.StatusBadRequest, errors.New(ErrorGenerated))
			return
		}

		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			web.Failure(ctx, http.StatusBadRequest, errors.New(ErrorInvalidId))
			return
		}

		products, err := p.productService.UpdateProductPartial(domain.Product{
			Id:          id,
			Name:        product.Name,
			Quantity:    product.Quantity,
			CodeValue:   product.CodeValue,
			IsPublished: product.IsPublished,
			Expiration:  product.Expiration,
			Price:       product.Price,
		})

		if err != nil {
			web.Failure(ctx, http.StatusInternalServerError, errors.New(ErrorGenerated))
			return
		}
		web.Success(ctx, http.StatusAccepted, products)
	}
}

func (p *ProductHandler) DeleteProduct() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		var product domain.ProductRequestUpdatePartial
		err := ctx.BindJSON(&product)

		if err != nil {
			web.Failure(ctx, http.StatusBadRequest, errors.New(ErrorGenerated))
			return
		}

		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			web.Failure(ctx, http.StatusBadRequest, errors.New(ErrorInvalidId))
			return
		}

		products, err := p.productService.DeleteProduct(id)

		if err != nil {
			web.Failure(ctx, http.StatusInternalServerError, errors.New(ErrorGenerated))
			return
		}
		web.Success(ctx, http.StatusAccepted, products)
	}
}
