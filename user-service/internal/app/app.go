package app

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"

	"user-service/api"
	"user-service/internal/handlers"
	infradb "user-service/internal/infra/db"
	"user-service/internal/infra/env"
	"user-service/internal/repository"
	myjwt "user-service/internal/service/jwt"
	"user-service/internal/usecase/relationships"
	"user-service/internal/usecase/users"
)

func BuildGinEngine(envConf *env.Config) *gin.Engine {
	var db *gorm.DB = infradb.ConnectToPostgres(envConf.PsgConf)

	signMethod := jwt.SigningMethodHS256
	jwtValidator := myjwt.NewValidator(envConf.JWTSecret, envConf.ApiSecret, signMethod)

	userRepo := repository.NewUserRepo(db)
	friendsRepo := repository.NewFriendRepo(db)

	userUseCase := users.NewUserUseCase(userRepo, jwtValidator)
	getRelationshipsUseCase := relationships.NewGetRelationshipsUseCase(friendsRepo, userRepo, jwtValidator)

	userHandler := handlers.NewUserHandler(userUseCase)
	friendHandler := handlers.NewGetRelationshipsHandler(getRelationshipsUseCase)

	engine := gin.Default()

	userHandler.RegisterHandlers(engine)
	friendHandler.RegisterHandlers(engine)

	engine.GET("/health", handlers.HealthCheck)
	addSwagger(engine)

	return engine
}

func addSwagger(engine *gin.Engine) {
	api.SwaggerInfo.BasePath = ""
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
