package middlewares

import (
	"github.com/Masih-Ghasri/GolangBackend/api/helper"
	"github.com/didip/tollbooth"
	"github.com/gin-gonic/gin"
	"net/http"
)

func LimitByRequest() gin.HandlerFunc {
	lmt := tollbooth.NewLimiter(1, nil)
	return func(c *gin.Context) {
		err := tollbooth.LimitByRequest(lmt, c.Writer, c.Request)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusTooManyRequests,
				helper.GenerateBaseResponseWithError(nil, false, -100, err))
			return
		} else {
			c.Next()
		}
	}
}
