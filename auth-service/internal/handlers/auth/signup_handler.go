package handlers

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"auth-service/internal/domain"
	"auth-service/internal/dto"
	"auth-service/internal/usecase/auth"
)

type SignUpHandler struct {
	signUpUc *auth.SignUpUsecase
}

func NewSignUpHandler(signUpUc *auth.SignUpUsecase) *SignUpHandler {
	return &SignUpHandler{signUpUc: signUpUc}
}

func (hand *SignUpHandler) RegisterHandlers(engine *gin.Engine) {
	engine.POST("/signup/username", hand.signUpWithUsername)
}

// signUpWithUsername godoc
//
//	@Summary		Sign up with username & password
//	@Description	Signs up with username & password, saves to db, returns created id.
//	@Tags			Sign up
//	@Accept			json
//	@Produce		json
//	@Param   	    user body       dto.UsernamePass true "user"
//	@Success		201	{object}	dto.UserIdAndTokens
//	@Failure		400	{object}	dto.ErrorResponse
//	@Failure		409	{object}	dto.ErrorResponse
//	@Failure		500	{object}	dto.ErrorResponse
//	@Router			/signup/username [post]
func (hand *SignUpHandler) signUpWithUsername(c *gin.Context) {
	usernamePass := dto.UsernamePass{}
	parseErr := c.BindJSON(&usernamePass)

	if parseErr != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Message: parseErr.Error()})
		return
	}

	userIdAndTokens, err := hand.signUpUc.SignUpWithUsername(&usernamePass)

	if errors.Is(err, domain.ErrDuplicate) {
		msg := fmt.Sprintf("User with username=%s already exists", usernamePass.Username)
		c.JSON(http.StatusConflict, dto.ErrorResponse{Message: msg})
		return
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, userIdAndTokens)
}
