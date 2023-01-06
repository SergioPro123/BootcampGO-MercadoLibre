package handlers

import (
	"net/http"
	"proyectoapisupermercado/controllers"
	"proyectoapisupermercado/models"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var productController *controllers.ProductController = controllers.GetInstance()

func AddProduct(ctx *gin.Context) {
	var product models.ProductRequest
	err := ctx.BindJSON(&product)

	if err != nil {
		ctx.String(500, "Ocurrio un error.")
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

	exist := productController.ExistProduct(product.CodeValue)
	if exist {
		ctx.JSON(http.StatusConflict, gin.H{
			"Error": "El producto ya existe.",
		})
		return
	}
	productResponse := productController.Addproduct(models.Product{
		Name:        product.Name,
		Quantity:    product.Quantity,
		CodeValue:   product.CodeValue,
		IsPublished: product.IsPublished,
		Expiration:  product.Expiration,
		Price:       product.Price,
	})
	ctx.JSON(http.StatusCreated, productResponse)
}

func GetProducts(ctx *gin.Context) {
	ctx.JSON(http.StatusAccepted, productController.GetProducts())
}

func GetProductById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Error": "Invalido ID",
		})
		return
	}
	product, err := productController.GetProductById(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusAccepted, product)
}
func GetProductByPrice(ctx *gin.Context) {
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
	ctx.JSON(http.StatusAccepted, productController.GetProductsByPrice(priceGt))
}
