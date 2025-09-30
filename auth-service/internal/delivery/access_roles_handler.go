package delivery

import (
	"auth-service/internal/dto"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type AccessRolesHandler struct {
}

func NewAccessRolesHandler() *AccessRolesHandler {
	return &AccessRolesHandler{}
}

func (hand *AccessRolesHandler) RegisterHandlers(engine *gin.Engine) {
	engine.POST("/access-roles/:userId", hand.updateUserRole)
}

// updateUserRole godoc
//
//	@Summary		Update user role
//	@Description	Updates given user's role. Request should be sent by moder/admin to upgrade the role.
//	@Tags			Access Roles
//	@Accept			json
//	@Produce		json
//	@Param   	    user body       dto.AccessRole true "access_role"
//	@Success		200	{object}	dto.UserIdAndTokens
//	@Failure		400	{object}	dto.ErrorResponse
//	@Failure		401	{object}	dto.ErrorResponse
//	@Failure		404	{object}	dto.ErrorResponse
//	@Failure		500	{object}	dto.ErrorResponse
//	@Router			/access-roles/{userId} [post]
func (hand *AccessRolesHandler) updateUserRole(c *gin.Context) {
	var dtoRole dto.AccessRole
	parseErr := c.BindJSON(&dtoRole)
	jwt := c.GetHeader("Authorization")
	log.Println(jwt)

	if parseErr != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Message: parseErr.Error()})
		return
	}

}
