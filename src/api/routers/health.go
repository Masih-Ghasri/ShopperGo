package routers

import (
	"github.com/Masih-Ghasri/GolangBackend/api/handlers"
	"github.com/gin-gonic/gin"
)

func Health(r *gin.RouterGroup) {
	handler := handlers.NewHealthHandler()
	r.GET("", handler.Health)
}
