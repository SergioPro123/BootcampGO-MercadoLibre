package main

import (
	"log"
	"os"
	"proyectoapisupermercado/cmd/server/docs"
	"proyectoapisupermercado/cmd/server/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error al intentar cargar archivo .env")
	}

	server := gin.New()
	server.Use(gin.Recovery())

	docs.SwaggerInfo.Host = os.Getenv("HOST")
	server.GET("docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	rt := routes.NewRouter(server)

	rt.SetRoutes()

	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
	server.Run()
}
