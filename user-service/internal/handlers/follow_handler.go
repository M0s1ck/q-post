package handlers

import (
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"user-service/internal/domain"
	"user-service/internal/usecase/relationships"
)

type FollowHandler struct {
	useCase *relationships.FollowUseCase
}

func NewFollowHandler(uc *relationships.FollowUseCase) *FollowHandler {
	return &FollowHandler{useCase: uc}
}

func (h *FollowHandler) RegisterHandlers(engine *gin.Engine) {
	engine.POST("/users/:id/follow", h.Follow)
}

// Follow godoc
// @Summary      Follow a user
// @Description  Sender follows user with followeeId, they might become friends if followee already follows
// @Tags         relationships
// @Accept       json
// @Produce      json
// @Param        id         path      string  true  "Followee ID"
// @Success      204
// @Failure      400 {object} dto.ErrorResponse
// @Failure      403 {object} dto.ErrorResponse
// @Failure      404 {object} dto.ErrorResponse
// @Failure      500 {object} dto.ErrorResponse
// @Router       /users/{id}/follow [post]
// @Security	 BearerAuth
func (h *FollowHandler) Follow(c *gin.Context) {
	var idStr = c.Param("id")
	followeeId, uuidFormErr := uuid.Parse(idStr)

	if uuidFormErr != nil {
		respondErr(c, http.StatusBadRequest, uuidFormErr.Error())
		return
	}

	token, tokenErr := getAuthorizationToken(c)
	if tokenErr != nil {
		respondErr(c, http.StatusBadRequest, tokenErr.Error())
		return
	}

	err := h.useCase.Follow(followeeId, token)

	if errors.Is(err, domain.ErrSelfFollow) {
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
