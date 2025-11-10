package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"user-service/internal/dto"
	"user-service/internal/usecase/users"
)

func NewUserHandler(userUseCase *users.UserUseCase) *UserHandler {
	uHandler := UserHandler{userUseCase: userUseCase}
	return &uHandler
}

type UserHandler struct {
	userUseCase *users.UserUseCase
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
	id, ok := handleGettingId(c, "id")
	if !ok {
		return
	}

	userDto, err := uHand.userUseCase.GetById(id)

	if err != nil {
		handleDefaultDomainErrors(c, err)
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
	userToCreate, token, ok := handleGettingBodyAndToken[dto.UserToCreate](c)
	if !ok {
		return
	}

	uuidResponse, err := uHand.userUseCase.Create(userToCreate, token)

	if err != nil {
		handleDefaultDomainErrors(c, err)
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
	userDetails, token, ok := handleGettingBodyAndToken[dto.UserDetailStr](c)
	if !ok {
		return
	}

	err := uHand.userUseCase.UpdateDetails(userDetails, token)

	if err != nil {
		handleDefaultDomainErrors(c, err)
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
	userId, token, ok := handleGettingIdAndToken(c, "id")
	if !ok {
		return
	}

	err := uHand.userUseCase.Delete(userId, token)

	if err != nil {
		handleDefaultDomainErrors(c, err)
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
	token, ok := handleGettingToken(c)
	if !ok {
		return
	}

	userDto, err := uHand.userUseCase.GetMe(token)

	if err != nil {
		handleDefaultDomainErrors(c, err)
		return
	}

	c.IndentedJSON(http.StatusOK, *userDto)
}
