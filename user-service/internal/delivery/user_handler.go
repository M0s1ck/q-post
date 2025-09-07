package delivery

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func NewUserHandler() *UserHandler {
	uHandler := UserHandler{}
	return &uHandler
}

type UserHandler struct {
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
//	@Success		200	{object}	string
//	@Failure		400	{object}	map[string]string
//	@Failure		500	{object}	map[string]string
//	@Router			/users/{id} [get]
func (uHand *UserHandler) Get(c *gin.Context) {
	id := c.Param("id")
	c.String(http.StatusOK, "Hello %s", id)
}
