package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func TestMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		apikey := ctx.GetHeader("api-key")
		if apikey == "1" {
			ctx.Next()
		}
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"result": "api key is required",
		})
		return
	}
}
