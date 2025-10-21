package app

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"

	"user-service/api"
	"user-service/internal/delivery"
	"user-service/internal/repository"
	"user-service/internal/usecase"
)

func BuildGinEngine(db *gorm.DB) *gin.Engine {
	engine := gin.Default()
	engine.GET("/health", delivery.HealthCheck)

	userRepo := repository.NewUserRepo(db)
	userUseCase := usecase.NewUserUseCase(userRepo)

	userHandler := delivery.NewUserHandler(userUseCase)
	userHandler.RegisterHandlers(engine)

	addSwagger(engine)

	return engine
}

func addSwagger(engine *gin.Engine) {
	api.SwaggerInfo.BasePath = ""
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
