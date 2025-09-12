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
	engine.POST("/users/create", uHand.Create)
}

// Get godoc
//
//	@Summary		Get user by id
//	@Description	Get user by id
//	@Tags			Users
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string	true	"user id"
//	@Success		200	{object}	dto.UserResponse
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

	if errors.Is(err, domain.ErrNotFound) {
		respondErr(c, http.StatusNotFound, fmt.Sprintf("User with id=%s was not found", idStr))
		return
	}

	if err != nil {
		respondErr(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, *userDto)
}

// Create godoc
//
//	@Summary		Create new user
//	@Description	Creates a new user, saves him to db, returns created id.
//	@Tags			Users
//	@Accept			json
//	@Produce		json
//	@Param   	    user body       dto.UserToCreate true "user"
//	@Success		201	{object}	dto.UuidOnlyResponse
//	@Failure		400	{object}	dto.ErrorResponse
//	@Failure		409	{object}	dto.ErrorResponse
//	@Failure		500	{object}	dto.ErrorResponse
//	@Router			/users/create [post]
func (uHand *UserHandler) Create(c *gin.Context) {
	userToCreate := dto.UserToCreate{}
	parseErr := c.BindJSON(&userToCreate)

	if parseErr != nil {
		respondErr(c, http.StatusBadRequest, parseErr.Error())
		return
	}

	uuidResponse, err := uHand.userUseCase.Create(&userToCreate)

	if errors.Is(err, domain.ErrDuplicate) {
		respondErr(c, http.StatusConflict, fmt.Sprintf("User with username=%s already exists", userToCreate.Username))
		return
	}

	if err != nil {
		respondErr(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.IndentedJSON(http.StatusCreated, uuidResponse)
}

func respondErr(c *gin.Context, code int, message string) {
	errResponse := dto.ErrorResponse{
		Message: message,
	}

	c.JSON(code, errResponse)
}
