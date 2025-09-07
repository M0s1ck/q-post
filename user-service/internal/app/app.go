package app

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"user-service/api"
	"user-service/internal/delivery"
)

func Build() *gin.Engine {
	engine := gin.Default()
	engine.GET("/health", delivery.HealthCheck)

	userHandler := delivery.NewUserHandler()
	userHandler.RegisterHandlers(engine)

	addSwagger(engine)

	return engine
}

func addSwagger(engine *gin.Engine) {
	api.SwaggerInfo.BasePath = ""
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
