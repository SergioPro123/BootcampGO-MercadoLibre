package main

import (
	"log"
	"proyectoapisupermercado/cmd/server/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	server := gin.Default()

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error al intentar cargar archivo .env")
	}

	rt := routes.NewRouter(server)
	rt.SetRoutes()

	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
	server.Run()
}
