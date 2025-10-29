package handlers

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

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
	engine.GET("/users/:id", uHand.GetById)
	engine.POST("/users/create", uHand.Create)
	engine.PUT("/users/me", uHand.UpdateDetails)
	engine.DELETE("/users/:id", uHand.Delete)
	engine.GET("/users/me", uHand.GetMe)
}

// GetById godoc
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
func (uHand *UserHandler) GetById(c *gin.Context) {
	var idStr string = c.Param("id")
	id, uuidFormErr := uuid.Parse(idStr)

	if uuidFormErr != nil {
		respondErr(c, http.StatusBadRequest, uuidFormErr.Error())
		return
	}

	userDto, err := uHand.userUseCase.GetById(id)

	if errors.Is(err, domain.ErrNotFound) {
		respondErr(c, http.StatusNotFound, err.Error())
		return
	}

	if err != nil {
		respondErr(c, http.StatusInternalServerError, err.Error())
		log.Println("Unexpected err: ", err)
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
//	@Failure		403	{object}	dto.ErrorResponse
//	@Failure		409	{object}	dto.ErrorResponse
//	@Failure		500	{object}	dto.ErrorResponse
//	@Security		BearerAuth
//	@Router			/users/create [post]
func (uHand *UserHandler) Create(c *gin.Context) {
	userToCreate := dto.UserToCreate{}
	parseErr := c.BindJSON(&userToCreate)

	if parseErr != nil {
		respondErr(c, http.StatusBadRequest, parseErr.Error())
		return
	}

	token, err := getAuthorizationToken(c)
	if err != nil {
		respondErr(c, http.StatusBadRequest, err.Error())
		return
	}

	uuidResponse, err := uHand.userUseCase.Create(&userToCreate, token)

	if errors.Is(err, domain.ErrInvalidToken) {
		respondErr(c, http.StatusForbidden, err.Error())
		return
	}

	if errors.Is(err, domain.ErrDuplicate) {
		respondErr(c, http.StatusConflict, fmt.Sprintf("User with username=%s already exists", userToCreate.Username))
		return
	}

	if err != nil {
		respondErr(c, http.StatusInternalServerError, err.Error())
		log.Println("Unexpected err: ", err)
		return
	}

	c.IndentedJSON(http.StatusCreated, uuidResponse)
}

// UpdateDetails godoc
//
//	@Summary		Update user details
//	@Description	Updates user whose id is in jwt, date is in the YYYY-MM-DD format
//	@Tags			Users
//	@Accept			json
//	@Produce		json
//	@Param   	    user body       dto.UserDetailStr true "details"
//	@Success		204
//	@Failure		400	{object}	dto.ErrorResponse
//	@Failure		403	{object}	dto.ErrorResponse
//	@Failure		404	{object}	dto.ErrorResponse
//	@Failure		500	{object}	dto.ErrorResponse
//	@Security		BearerAuth
//	@Router			/users/me [put]
func (uHand *UserHandler) UpdateDetails(c *gin.Context) {
	var userDetails dto.UserDetailStr
	bindErr := c.BindJSON(&userDetails)

	if bindErr != nil {
		respondErr(c, http.StatusBadRequest, bindErr.Error())
		return
	}

	token, tokenErr := getAuthorizationToken(c)
	if tokenErr != nil {
		respondErr(c, http.StatusBadRequest, tokenErr.Error())
		return
	}

	err := uHand.userUseCase.UpdateDetails(&userDetails, token)

	if errors.Is(err, domain.ErrInvalidDto) {
		respondErr(c, http.StatusBadRequest, err.Error())
		return
	}

	if errors.Is(err, domain.ErrInvalidToken) {
		respondErr(c, http.StatusForbidden, err.Error())
		return
	}

	if errors.Is(err, domain.ErrNotFound) {
		respondErr(c, http.StatusNotFound, err.Error())
		return
	}

	if err != nil {
		respondErr(c, http.StatusInternalServerError, err.Error())
		log.Println("Unexpected err: ", err)
		return
	}

	c.Status(http.StatusNoContent)
}

// Delete godoc
//
//	@Summary		Removes user
//	@Description	Removes user by their id
//	@Tags			Users
//	@Accept			json
//	@Produce		json
//	@Param			id	 path		string	true	"user id"
//	@Success		204
//	@Failure		400	{object}	dto.ErrorResponse
//	@Failure		403	{object}	dto.ErrorResponse
//	@Failure		404	{object}	dto.ErrorResponse
//	@Failure		500	{object}	dto.ErrorResponse
//	@Security		BearerAuth
//	@Router			/users/{id} [delete]
func (uHand *UserHandler) Delete(c *gin.Context) {
	var idStr string = c.Param("id")
	id, uuidFormErr := uuid.Parse(idStr)

	if uuidFormErr != nil {
		respondErr(c, http.StatusBadRequest, uuidFormErr.Error())
		return
	}

	token, tokenErr := getAuthorizationToken(c)
	if tokenErr != nil {
		respondErr(c, http.StatusBadRequest, tokenErr.Error())
		return
	}

	err := uHand.userUseCase.Delete(id, token)

	if errors.Is(err, domain.ErrInvalidToken) {
		respondErr(c, http.StatusForbidden, err.Error())
		return
	}

	if errors.Is(err, domain.ErrNotFound) {
		respondErr(c, http.StatusNotFound, fmt.Sprintf("User with id=%s was not found", idStr))
		return
	}

	if err != nil {
		respondErr(c, http.StatusInternalServerError, err.Error())
		log.Println("Unexpected err: ", err)
		return
	}

	c.Status(http.StatusNoContent)
}

// GetMe godoc
//
//	@Summary		Get me
//	@Description	Gets user whose id is in given token
//	@Tags			Users
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	dto.UserResponse
//	@Failure		400	{object}	dto.ErrorResponse
//	@Failure		403	{object}	dto.ErrorResponse
//	@Failure		404	{object}	dto.ErrorResponse
//	@Failure		500	{object}	dto.ErrorResponse
//	@Router			/users/me [get]
//	@Security		BearerAuth
func (uHand *UserHandler) GetMe(c *gin.Context) {
	token, tokenErr := getAuthorizationToken(c)
	if tokenErr != nil {
		respondErr(c, http.StatusBadRequest, tokenErr.Error())
		return
	}

	userDto, err := uHand.userUseCase.GetMe(token)

	if errors.Is(err, domain.ErrInvalidToken) {
		respondErr(c, http.StatusForbidden, err.Error())
		return
	}

	if errors.Is(err, domain.ErrNotFound) {
		respondErr(c, http.StatusNotFound, err.Error())
		return
	}

	if err != nil {
		respondErr(c, http.StatusInternalServerError, err.Error())
		log.Println("Unexpected err: ", err)
		return
	}

	c.IndentedJSON(http.StatusOK, *userDto)
}

func respondErr(c *gin.Context, code int, message string) {
	errResponse := dto.ErrorResponse{
		Message: message,
	}

	c.JSON(code, errResponse)
}

func getAuthorizationToken(c *gin.Context) (string, error) {
	jwt := c.GetHeader("Authorization")

	if strings.HasPrefix(jwt, "Bearer ") {
		jwt = strings.TrimPrefix(jwt, "Bearer ")
	}

	if jwt == "" {
		return "", errors.New("no authorization token found")
	}

	return jwt, nil
}
