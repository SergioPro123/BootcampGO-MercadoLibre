package main

import (
	"fmt"
	"proyectoapisupermercado/controllers"
	"proyectoapisupermercado/data"
	"proyectoapisupermercado/handlers"

	"github.com/gin-gonic/gin"
)

var ProductController *controllers.ProductController

func main() {

	ProductController = controllers.GetInstance()
	err := ProductController.Init(&data.DatabaseProductJson{})
	if err != nil {
		panic(fmt.Sprint("No se pudo inicializar el controlador de ProductController: ", err))
	}

	router := gin.Default()

	productGroup := router.Group("Products")

	productGroup.GET("/", handlers.GetProducts)
	productGroup.GET("/:id", handlers.GetProductById)
	productGroup.GET("/search", handlers.GetProductByPrice)

	productGroup.POST("/", handlers.AddProduct)

	router.Run()
}
