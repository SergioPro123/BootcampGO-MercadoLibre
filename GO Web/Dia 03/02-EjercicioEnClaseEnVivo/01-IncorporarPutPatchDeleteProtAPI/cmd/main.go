package main

import (
	"log"
	"proyectoapisupermercado/cmd/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	rt := routes.NewRouter(server)
	rt.SetRoutes()

	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
	server.Run()
}
