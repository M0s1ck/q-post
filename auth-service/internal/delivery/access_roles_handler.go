package delivery

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"auth-service/internal/domain"
	"auth-service/internal/dto"
	"auth-service/internal/usecase"
)

type AccessRolesHandler struct {
	uCase *usecase.AccessRolesUsecase
}

func NewAccessRolesHandler(uCase *usecase.AccessRolesUsecase) *AccessRolesHandler {
	return &AccessRolesHandler{uCase: uCase}
}

func (hand *AccessRolesHandler) RegisterHandlers(engine *gin.Engine) {
	engine.POST("/access-roles/:userId", hand.updateUserRole)
}

// updateUserRole godoc
//
//	@Summary		Update user role
//	@Description	Updates given user's role. Request should be sent by moder/admin (with jwt) to upgrade the role.
//	@Tags			Access Roles
//	@Accept			json
//	@Produce		json
//	@Param   	    user body       dto.AccessRole true "access_role"
//	@Param			userId	 path		string	true	"user id"
//	@Success		204
//	@Failure		400	{object}	dto.ErrorResponse
//	@Failure		401	{object}	dto.ErrorResponse
//	@Failure		403	{object}	dto.ErrorResponse
//	@Failure		404	{object}	dto.ErrorResponse
//	@Failure		500	{object}	dto.ErrorResponse
//	@Security		BearerAuth
//	@Router			/access-roles/{userId} [post]
func (hand *AccessRolesHandler) updateUserRole(c *gin.Context) {
	var dtoRole dto.AccessRole
	parseErr := c.BindJSON(&dtoRole)

	var idStr string = c.Param("userId") // TODO: test happy path(moder moder), test expired jwt
	userId, uuidFormErr := uuid.Parse(idStr)

	jwt := c.GetHeader("Authorization")
	if strings.HasPrefix(jwt, "Bearer ") {
		jwt = strings.TrimPrefix(jwt, "Bearer ")
	}

	if parseErr != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Message: parseErr.Error()})
		return
	} else if uuidFormErr != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Message: uuidFormErr.Error()})
		return
	}

	err := hand.uCase.UpdateUserRole(userId, dtoRole.Role, jwt)

	if errors.Is(err, domain.ErrInvalidToken) {
		c.JSON(http.StatusUnauthorized, dto.ErrorResponse{Message: err.Error()})
		return
	}

	if errors.Is(err, domain.ErrWeakRole) {
		c.JSON(http.StatusForbidden, dto.ErrorResponse{Message: err.Error()})
		return
	}

	if errors.Is(err, domain.ErrNotFound) {
		c.JSON(http.StatusNotFound, dto.ErrorResponse{Message: err.Error()})
		return
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Message: err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
