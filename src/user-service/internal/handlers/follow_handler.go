package handlers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

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
	engine.POST("/users/:id/unfollow", h.Unfollow)
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
	followeeId, token, ok := handleGettingIdAndToken(c, "id")
	if !ok {
		return
	}

	err := h.useCase.Follow(followeeId, token)

	if errors.Is(err, domain.ErrSelfFollow) {
		respondErr(c, http.StatusBadRequest, err.Error())
		return
	}

	if err != nil {
		handleDefaultDomainErrors(c, err)
		return
	}

	c.Status(http.StatusNoContent)
}

// Unfollow godoc
// @Summary      Unfollow a user
// @Description  Sender unfollows the user with followeeId, they're not friends anymore if they were
// @Tags         relationships
// @Accept       json
// @Produce      json
// @Param        id         path      string  true  "Followee ID"
// @Success      204
// @Failure      400 {object} dto.ErrorResponse
// @Failure      403 {object} dto.ErrorResponse
// @Failure      404 {object} dto.ErrorResponse
// @Failure      500 {object} dto.ErrorResponse
// @Router       /users/{id}/unfollow [post]
// @Security	 BearerAuth
func (h *FollowHandler) Unfollow(c *gin.Context) {
	followeeId, token, ok := handleGettingIdAndToken(c, "id")
	if !ok {
		return
	}

	err := h.useCase.Unfollow(followeeId, token)

	if errors.Is(err, domain.ErrSelfFollow) {
		respondErr(c, http.StatusBadRequest, err.Error())
		return
	}

	if err != nil {
		handleDefaultDomainErrors(c, err)
		return
	}

	c.Status(http.StatusNoContent)
}
