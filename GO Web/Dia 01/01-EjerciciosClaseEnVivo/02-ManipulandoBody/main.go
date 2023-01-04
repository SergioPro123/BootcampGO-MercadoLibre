package main

import (
	"github.com/gin-gonic/gin"
)

type Client struct {
	Nombre   string `json:"nombre"`
	Apellido string `json:"apellido"`
}

func main() {
	router := gin.Default()

	router.POST("/saludo", func(ctx *gin.Context) {
		var client Client
		err := ctx.BindJSON(&client)
		if err != nil {
			ctx.String(500, "Ocurrio un error.")
			return
		}
		ctx.String(200, "Hola "+client.Nombre+" "+client.Apellido)
	})

	router.Run()
}
