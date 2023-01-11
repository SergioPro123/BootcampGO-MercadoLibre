package handlers

import (
	"net/http"
	"os"
	"proyectoapisupermercado/internal/domain"
	"proyectoapisupermercado/internal/product"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type ProductHandler struct {
	productService product.IProductService
}

func NewProductHandler(productService product.IProductService) *ProductHandler {
	return &ProductHandler{productService: productService}
}

func (p *ProductHandler) AddProduct() gin.HandlerFunc {

	return func(ctx *gin.Context) {

		token := ctx.GetHeader("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"Error": "Token invalido",
			})
			return
		}

		var product domain.ProductRequest
		err := ctx.BindJSON(&product)

		if err != nil {
			ctx.String(http.StatusBadRequest, "Ocurrio un error.")
			return
		}

		validate := validator.New()
		err = validate.Struct(product)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"Error": err.Error(),
			})
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
			ctx.JSON(400, gin.H{
				"err":  "Ha ocurrido un error.",
				"info": err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusCreated, productResponse)
	}
}

func (p *ProductHandler) GetProducts() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		products, err := p.productService.GetProducts()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"err":  "Ha ocurrido un error.",
				"info": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusAccepted, products)
	}
}

func (p *ProductHandler) GetProductById() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"Error": "Invalido ID",
			})
			return
		}
		product, err := p.productService.GetProductById(id)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"Error": err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusAccepted, product)
	}
}
func (p *ProductHandler) GetProductByPrice() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		priceGt, err := strconv.ParseFloat(ctx.Query("priceGt"), 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"Error": "Invalido priceGt",
			})
			return
		}

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"Error": err.Error(),
			})
			return
		}

		products, err := p.productService.GetProductsByPrice(priceGt)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"err":  "Ha ocurrido un error.",
				"info": err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusAccepted, products)
	}
}
func (p *ProductHandler) UpdateAllProduct() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"Error": "Token invalido",
			})
			return
		}
		var product domain.ProductRequest
		err := ctx.BindJSON(&product)

		if err != nil {
			ctx.String(http.StatusBadRequest, "Ocurrio un error.")
			return
		}

		validate := validator.New()
		err = validate.Struct(product)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"Error": err.Error(),
			})
			return
		}

		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"Error": "Invalido ID",
			})
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
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"err":  "Ha ocurrido un error.",
				"info": err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusAccepted, products)
	}
}
func (p *ProductHandler) UpdateProductPartial() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"Error": "Token invalido",
			})
			return
		}
		var product domain.ProductRequestUpdatePartial
		err := ctx.BindJSON(&product)

		if err != nil {
			ctx.String(http.StatusBadRequest, "Ocurrio un error.")
			return
		}

		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"Error": "Invalido ID",
			})
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
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"err":  "Ha ocurrido un error.",
				"info": err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusAccepted, products)
	}
}

func (p *ProductHandler) DeleteProduct() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"Error": "Token invalido",
			})
			return
		}
		var product domain.ProductRequestUpdatePartial
		err := ctx.BindJSON(&product)

		if err != nil {
			ctx.String(http.StatusBadRequest, "Ocurrio un error.")
			return
		}

		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"Error": "Invalido ID",
			})
			return
		}

		products, err := p.productService.DeleteProduct(id)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"err":  "Ha ocurrido un error.",
				"info": err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusAccepted, products)
	}
}
