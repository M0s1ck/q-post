package handlers

import (
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"user-service/internal/domain"
	"user-service/internal/usecase/friends"
)

type FriendHandler struct {
	getFriendsUc *friends.GetFriendsUseCase
}

func NewFriendHandler(getFriendsUc *friends.GetFriendsUseCase) *FriendHandler {
	return &FriendHandler{
		getFriendsUc: getFriendsUc,
	}
}

func (h *FriendHandler) RegisterHandlers(engine *gin.Engine) {
	engine.GET("/users/:id/friends", h.GetFriends)
}

// GetFriends godoc
// @Summary      Get user's friends
// @Description  Returns paginated list of user's friends
// @Tags         friends
// @Accept       json
// @Produce      json
// @Param        id         path      string  true  "User ID"
// @Param        page       query     int     false "Page number (default 0)"
// @Param        pageSize   query     int     false "Page size (default 20)"
// @Success      200 {object} []dto.UserSummary
// @Failure      400 {object} dto.ErrorResponse
// @Failure      403 {object} dto.ErrorResponse
// @Failure      500 {object} dto.ErrorResponse
// @Router       /users/{id}/friends [get]
// @Security	 BearerAuth
func (h *FriendHandler) GetFriends(c *gin.Context) {
	var idStr = c.Param("id")
	userId, uuidFormErr := uuid.Parse(idStr)

	if uuidFormErr != nil {
		respondErr(c, http.StatusBadRequest, uuidFormErr.Error())
		return
	}

	pageStr := c.DefaultQuery("page", "0")
	pageSizeStr := c.DefaultQuery("pageSize", "0")

	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 0 {
		page = 0
	}
	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil || pageSize <= 0 {
		pageSize = friends.DefaultPageSize
	}

	token, tokenErr := getAuthorizationToken(c)
	if tokenErr != nil {
		respondErr(c, http.StatusBadRequest, tokenErr.Error())
		return
	}

	homies, err := h.getFriendsUc.GetFriends(userId, page, pageSize, token)

	if errors.Is(err, domain.ErrInvalidToken) {
		respondErr(c, http.StatusForbidden, err.Error())
		return
	}

	if err != nil {
		respondErr(c, http.StatusInternalServerError, err.Error())
		log.Println("Unexpected err: ", err)
		return
	}

	c.IndentedJSON(200, homies)
}
