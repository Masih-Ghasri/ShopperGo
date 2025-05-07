package api

import (
	"fmt"
	"github.com/Masih-Ghasri/GolangBackend/api/middlewares"
	"github.com/Masih-Ghasri/GolangBackend/api/routers"
	"github.com/Masih-Ghasri/GolangBackend/api/validation"
	"github.com/Masih-Ghasri/GolangBackend/config"
	"github.com/Masih-Ghasri/GolangBackend/docs"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
)

func Initserver(cfg *config.Config) {
	engine := gin.New()
	RegisterValidator()

	engine.Use(middlewares.Cors(cfg))
	engine.Use(middlewares.DefualtStructuredLogger(cfg))
	engine.Use(gin.Logger(), gin.Recovery() /*middlewares.TestMiddleware()*/, middlewares.LimitByRequest())

	RegisterRouter(engine)
	RegisterSwagger(engine, cfg)

	err := engine.Run(fmt.Sprintf(":%s", cfg.Server.Port))
	if err != nil {
		return
	}
}

func RegisterValidator() {
	val, ok := binding.Validator.Engine().(*validator.Validate)
	if ok {
		val.RegisterValidation("mobile", validation.IranianMobileNumberValidator, true)
		val.RegisterValidation("password", validation.PasswordValidator, true)
	}
}

func RegisterRouter(engine *gin.Engine) {
	api := engine.Group("/api")

	v1 := api.Group("/v1")
	{
		health := v1.Group("/health")
		testRouter := v1.Group("/test")

		routers.Health(health)
		routers.TestRouter(testRouter)
	}
}

func RegisterSwagger(engine *gin.Engine, cfg *config.Config) {
	docs.SwaggerInfo.Title = "Golang Backend"
	docs.SwaggerInfo.Description = "Golang Backend"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.BasePath = "/api"
	docs.SwaggerInfo.Schemes = []string{"http"}
	docs.SwaggerInfo.Host = fmt.Sprintf("localhost:%s", cfg.Server.Port)
	//docs.SwaggerInfo.Schemes = []string{"https"}

	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
