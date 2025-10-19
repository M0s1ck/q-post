package app

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"

	"auth-service/api"
	"auth-service/internal/domain/refresh"
	"auth-service/internal/domain/user"
	roleshand "auth-service/internal/handlers/access_roles"
	authhand "auth-service/internal/handlers/auth"
	healthhand "auth-service/internal/handlers/health"
	refreshHand "auth-service/internal/handlers/refresh"
	infradb "auth-service/internal/infra/db"
	"auth-service/internal/repository"
	"auth-service/internal/service/hash"
	myjwt "auth-service/internal/service/jwt"
	authcase "auth-service/internal/usecase/auth"
	refreshcase "auth-service/internal/usecase/refresh"
	rolescase "auth-service/internal/usecase/roles"
)

func BuildGinEngine() *gin.Engine {
	var db *gorm.DB = infradb.ConnectToPostgres()

	authenRepo := repository.NewAuthenticationRepo(db)
	refreshRepo := repository.NewRefreshTokenRepo(db)

	argonHasher := hash.NewArgonHasher()
	sha256Hasher := hash.NewSha256Hasher()

	jwtSecret := os.Getenv("JWT_SECRET_KEY")
	signMethod := jwt.SigningMethodHS256
	tokenIssuer := myjwt.NewJwtIssuer(jwtSecret, signMethod)
	tokenValidator := myjwt.NewJwtValidator(jwtSecret, signMethod)

	userServ := user.NewAuthUserService(authenRepo, argonHasher)
	refreshServ := refresh.NewRefreshTokenService(refreshRepo, sha256Hasher)

	signUpUc := authcase.NewSignUpUsecase(userServ, refreshServ, tokenIssuer)
	signInUc := authcase.NewSignInUsecase(userServ, refreshServ, tokenIssuer)
	refreshUc := refreshcase.NewRefreshUsecase(refreshServ, authenRepo, tokenIssuer)
	accessRolesUc := rolescase.NewAccessRolesUsecase(authenRepo, tokenValidator)

	signUpHandler := authhand.NewSignUpHandler(signUpUc)
	authenHandler := authhand.NewSignInHandler(signInUc)
	accessRoleHandler := roleshand.NewAccessRolesHandler(accessRolesUc)
	refreshHandler := refreshHand.NewRefreshHandler(refreshUc)

	engine := gin.Default()

	authenHandler.RegisterHandlers(engine)
	signUpHandler.RegisterHandlers(engine)
	accessRoleHandler.RegisterHandlers(engine)
	refreshHandler.RegisterHandlers(engine)

	engine.GET("/health", healthhand.HealthCheck)
	addSwagger(engine)

	return engine
}

func addSwagger(engine *gin.Engine) {
	api.SwaggerInfo.BasePath = ""
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
