package app

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"

	"user-service/api"
	"user-service/internal/handlers"
	"user-service/internal/repository"
	myjwt "user-service/internal/service/jwt"
	"user-service/internal/usecase"
)

func BuildGinEngine(db *gorm.DB) *gin.Engine {
	apiJwtSecret := os.Getenv("API_SECRET_KEY")
	userJwtSecret := os.Getenv("JWT_SECRET_KEY")
	signMethod := jwt.SigningMethodHS256
	jwtValidator := myjwt.NewValidator(userJwtSecret, apiJwtSecret, signMethod)

	userRepo := repository.NewUserRepo(db)

	userUseCase := usecase.NewUserUseCase(userRepo, jwtValidator)

	userHandler := handlers.NewUserHandler(userUseCase)

	engine := gin.Default()

	userHandler.RegisterHandlers(engine)

	engine.GET("/health", handlers.HealthCheck)
	addSwagger(engine)

	return engine
}

func addSwagger(engine *gin.Engine) {
	api.SwaggerInfo.BasePath = ""
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
