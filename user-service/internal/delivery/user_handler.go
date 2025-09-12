package delivery

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"user-service/internal/domain"
	"user-service/internal/dto"
	"user-service/internal/usecase"
)

func NewUserHandler(userUseCase *usecase.UserUseCase) *UserHandler {
	uHandler := UserHandler{userUseCase: userUseCase}
	return &uHandler
}

type UserHandler struct {
	userUseCase *usecase.UserUseCase
}

func (uHand *UserHandler) RegisterHandlers(engine *gin.Engine) {
	engine.GET("/users/:id", uHand.Get)
}

// Get godoc
//
//	@Summary		Get user by id
//	@Description	Get user by id
//	@Tags			Users
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string	true	"user id"
//	@Success		200	{object}	dto.UserDto
//	@Failure		400	{object}	dto.ErrorResponse
//	@Failure		404	{object}	dto.ErrorResponse
//	@Failure		500	{object}	dto.ErrorResponse
//	@Router			/users/{id} [get]
func (uHand *UserHandler) Get(c *gin.Context) {
	var idStr string = c.Param("id")
	id, uuidFormErr := uuid.Parse(idStr)

	if uuidFormErr != nil {
		respondErr(c, http.StatusBadRequest, uuidFormErr.Error())
		return
	}

	userDto, err := uHand.userUseCase.GetById(id)

	if errors.Is(err, domain.ErrorNotFound) {
		respondErr(c, http.StatusNotFound, fmt.Sprintf("User with id=%s was not found", idStr))
		return
	}

	if err != nil {
		respondErr(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, userDto)
}

func respondErr(c *gin.Context, code int, message string) {
	errResponse := dto.ErrorResponse{
		Message: message,
	}

	c.JSON(code, errResponse)
}
