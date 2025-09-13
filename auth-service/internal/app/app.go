package app

import (
	"auth-service/api"
	"auth-service/internal/delivery"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func BuildGinEngine() *gin.Engine {
	engine := gin.Default()

	engine.GET("/health", delivery.HealthCheck)

	addSwagger(engine)

	return engine
}

func addSwagger(engine *gin.Engine) {
	api.SwaggerInfo.BasePath = ""
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
