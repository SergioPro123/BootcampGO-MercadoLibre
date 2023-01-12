package middlewares

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		fecha := time.Now().Format("02/01/06 - 15:04:05")
		fmt.Printf("[%s] | %v | %s\n", ctx.Request.Method, fecha, ctx.Request.RequestURI)
		ctx.Next()
	}
}
