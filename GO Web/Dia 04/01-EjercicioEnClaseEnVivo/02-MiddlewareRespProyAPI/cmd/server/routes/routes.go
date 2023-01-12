package routes

import (
	"proyectoapisupermercado/cmd/server/handlers"
	"proyectoapisupermercado/cmd/server/middlewares"
	"proyectoapisupermercado/internal/product"
	"proyectoapisupermercado/pkg/store"

	"github.com/gin-gonic/gin"
)

type Router struct {
	en *gin.Engine
}

func NewRouter(en *gin.Engine) *Router {
	return &Router{en: en}
}

func (r *Router) SetRoutes() {
	r.SetProduct()
}

// product
func (r *Router) SetProduct() {
	// instances
	js := store.NewStore("../../internal/product/products.json")
	rp := product.NewProductRepository(js)
	sv := product.NewProductService(rp)
	h := handlers.NewProductHandler(sv)

	r.en.Use(middlewares.Logger())
	ws := r.en.Group("/products")

	ws.GET("/", h.GetProducts())
	ws.GET("/:id", h.GetProductById())
	ws.GET("/search", h.GetProductByPrice())

	ws.POST("/", middlewares.ValidateToken(), h.AddProduct())

	ws.PUT("/:id", middlewares.ValidateToken(), h.UpdateAllProduct())
	ws.PATCH("/:id", middlewares.ValidateToken(), h.UpdateProductPartial())

	ws.DELETE("/:id", middlewares.ValidateToken(), h.DeleteProduct())

}
