package routers

import (
	"github.com/Masih-Ghasri/GolangBackend/api/handlers"
	"github.com/gin-gonic/gin"
)

func TestRouter(r *gin.RouterGroup) {
	h := handlers.NewTestHandler()

	r.GET("", h.Test)
	r.GET("/users", h.Users)
	r.GET("/user/:id", h.UserById)
	r.GET("/get-user-by-username/:username", h.UserByUsername)
	r.GET("/user/:id/accounts", h.Accounts)

	r.POST("/add-user", h.AddUser)

	r.POST("/binder/v1", h.HeaderBinder1)
	r.POST("/binder/v2", h.HeaderBinder2)

	r.POST("/query/v1", h.QueryBinder1)
	r.POST("/query/v2", h.QueryBinder2)

	r.POST("/uri/:id/:name", h.UriBinder)

	r.POST("/bodyBinder", h.BodyBinder)

	r.POST("/form", h.FormBinder)

	r.POST("/file", h.FileBinder)
}
