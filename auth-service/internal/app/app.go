package app

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"

	"auth-service/api"
	"auth-service/internal/delivery"
	infradb "auth-service/internal/infra/db"
	"auth-service/internal/repository"
	"auth-service/internal/security"
	"auth-service/internal/usecase"
)

func BuildGinEngine() *gin.Engine {
	var db *gorm.DB = infradb.ConnectToPostgres()

	authenRepo := repository.NewAuthenticationRepo(db)

	passHasher := security.NewArgonHasher()
	tokenIssuer := security.NewTokenIssuer()

	authenUCase := usecase.NewAuthenticationUsecase(authenRepo, passHasher, tokenIssuer)

	authenHandler := delivery.NewAuthenticationHandler(authenUCase)

	engine := gin.Default()

	authenHandler.RegisterHandlers(engine)

	engine.GET("/health", delivery.HealthCheck)

	addSwagger(engine)

	return engine
}

func addSwagger(engine *gin.Engine) {
	api.SwaggerInfo.BasePath = ""
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
