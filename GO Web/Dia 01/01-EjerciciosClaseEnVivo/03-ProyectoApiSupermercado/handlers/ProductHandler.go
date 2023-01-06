package handlers

import (
	"net/http"
	"proyectoapisupermercado/controllers"
	"strconv"

	"github.com/gin-gonic/gin"
)

var productController *controllers.ProductController = controllers.GetInstance()

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
