package app

import (
	"auth-service/internal/domain/refresh"
	"auth-service/internal/domain/user"
	"auth-service/internal/service/hash"
	jwt2 "auth-service/internal/service/jwt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"

	"auth-service/api"
	"auth-service/internal/delivery"
	infradb "auth-service/internal/infra/db"
	"auth-service/internal/repository"
	"auth-service/internal/usecase"
)

func BuildGinEngine() *gin.Engine {
	var db *gorm.DB = infradb.ConnectToPostgres()

	authenRepo := repository.NewAuthenticationRepo(db)
	refreshRepo := repository.NewRefreshTokenRepo(db)

	strHasher := hash.NewArgonHasher()

	jwtSecret := os.Getenv("JWT_SECRET_KEY")
	signMethod := jwt.SigningMethodHS256
	tokenIssuer := jwt2.NewJwtIssuer(jwtSecret, signMethod)
	tokenValidator := jwt2.NewJwtValidator(jwtSecret, signMethod)

	userServ := user.NewAuthUserService(authenRepo, strHasher)
	refreshServ := refresh.NewRefreshTokenService(refreshRepo, strHasher)

	signUpUc := usecase.NewSignUpUsecase(userServ, refreshServ, tokenIssuer)
	signInUc := usecase.NewSignInUsecase(authenRepo, tokenIssuer, strHasher)
	accessRolesUc := usecase.NewAccessRolesUsecase(authenRepo, tokenValidator)

	signUpHandler := delivery.NewSignUpHandler(signUpUc)
	authenHandler := delivery.NewAuthenticationHandler(signInUc)
	accessRoleHandler := delivery.NewAccessRolesHandler(accessRolesUc)

	engine := gin.Default()

	authenHandler.RegisterHandlers(engine)
	signUpHandler.RegisterHandlers(engine)
	accessRoleHandler.RegisterHandlers(engine)

	engine.GET("/health", delivery.HealthCheck)
	addSwagger(engine)

	return engine
}

func addSwagger(engine *gin.Engine) {
	api.SwaggerInfo.BasePath = ""
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
