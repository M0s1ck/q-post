package delivery

import (
	"github.com/gin-gonic/gin"
)

func NewAuthHandler() *AuthHandler {
	handler := AuthHandler{}
	return &handler
}

type AuthHandler struct {
}

func RegisterHandlers(engine *gin.Engine) {
}
