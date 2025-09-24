package delivery

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"auth-service/internal/domain"
	"auth-service/internal/dto"
	"auth-service/internal/usecase"
)

func NewAuthenticationHandler(signInUc *usecase.SignInUsecase) *AuthenticationHandler {
	return &AuthenticationHandler{
		signInUc: signInUc,
	}
}

type AuthenticationHandler struct {
	signInUc *usecase.SignInUsecase
}

func (hand *AuthenticationHandler) RegisterHandlers(engine *gin.Engine) {
	engine.POST("/signin/username", hand.signInWithUsername)
}

// signInWithUsername godoc
//
//	@Summary		Sign in with username & password
//	@Description	Sign in with username & password, returns id and jwt.
//	@Tags			Authentication
//	@Accept			json
//	@Produce		json
//	@Param   	    user body       dto.UsernamePass true "user"
//	@Success		200	{object}	dto.UserIdAndTokens
//	@Failure		400	{object}	dto.ErrorResponse
//	@Failure		401	{object}	dto.ErrorResponse
//	@Failure		404	{object}	dto.ErrorResponse
//	@Failure		500	{object}	dto.ErrorResponse
//	@Router			/signin/username [post]
func (hand *AuthenticationHandler) signInWithUsername(c *gin.Context) {
	usernamePass := dto.UsernamePass{}
	parseErr := c.BindJSON(&usernamePass)

	if parseErr != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Message: parseErr.Error()})
		return
	}

	userIdAndTokens, err := hand.signInUc.SignInByUsername(&usernamePass)

	if errors.Is(err, domain.ErrNotFound) {
		msg := fmt.Sprintf("User with username=%s not found", usernamePass.Username)
		c.JSON(http.StatusNotFound, dto.ErrorResponse{Message: msg})
		return
	}

	if errors.Is(err, domain.ErrWrongPassword) {
		msg := fmt.Sprintf("Wrong password")
		c.JSON(http.StatusUnauthorized, dto.ErrorResponse{Message: msg})
		return
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, userIdAndTokens)
}
