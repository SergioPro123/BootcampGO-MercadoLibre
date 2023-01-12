package middlewares

import (
	"errors"
	"net/http"
	"os"
	"proyectoapisupermercado/pkg/web"

	"github.com/gin-gonic/gin"
)

var (
	ErrorAccessToken = "token invalido"
)

func ValidateToken() gin.HandlerFunc {

	return func(ctx *gin.Context) {

		token := ctx.GetHeader("token")
		if token != os.Getenv("TOKEN") {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized,
				web.GetStructErrorResponse(http.StatusUnauthorized, errors.New(ErrorAccessToken)))
			return
		}
		ctx.Next()
	}
}
