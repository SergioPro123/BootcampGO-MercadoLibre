package routes

import (
	"proyectoapisupermercado/cmd/handlers"
	"proyectoapisupermercado/internal/product"

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
	rp := product.NewProductRepository()
	sv := product.NewProductService(rp)
	h := handlers.NewProductHandler(sv)

	ws := r.en.Group("/products")

	ws.GET("/", h.GetProducts())
	ws.GET("/:id", h.GetProductById())
	ws.GET("/search", h.GetProductByPrice())
	ws.POST("/", h.AddProduct())
}
