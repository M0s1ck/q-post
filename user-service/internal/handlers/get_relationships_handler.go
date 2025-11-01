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

type GetRelationshipsHandler struct {
	getRelsUc *relationships.GetRelationshipsUseCase
}

func NewGetRelationshipsHandler(getFriendsUc *relationships.GetRelationshipsUseCase) *GetRelationshipsHandler {
	return &GetRelationshipsHandler{
		getRelsUc: getFriendsUc,
	}
}

func (h *GetRelationshipsHandler) RegisterHandlers(engine *gin.Engine) {
	engine.GET("/users/:id/friends", h.GetFriends)
	engine.GET("/users/:id/followers", h.GetFollowers)
	engine.GET("/users/:id/followees", h.GetFollowees)
	engine.GET("/users/:id/relationship", h.GetRelationship)
}

// GetFriends godoc
// @Summary      Get user's friends
// @Description  Returns paginated list of user's friends
// @Tags         relationships
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
func (h *GetRelationshipsHandler) GetFriends(c *gin.Context) {
	var idStr = c.Param("id")
	userId, uuidFormErr := uuid.Parse(idStr)

	if uuidFormErr != nil {
		respondErr(c, http.StatusBadRequest, uuidFormErr.Error())
		return
	}

	page, pageSize := getPaginationParams(c)

	token, tokenErr := getAuthorizationToken(c)
	if tokenErr != nil {
		respondErr(c, http.StatusBadRequest, tokenErr.Error())
		return
	}

	homies, err := h.getRelsUc.GetFriends(userId, page, pageSize, token)

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

// GetFollowers godoc
// @Summary      Get user's followers
// @Description  Returns paginated list of user's followers
// @Tags         relationships
// @Accept       json
// @Produce      json
// @Param        id         path      string  true  "User ID"
// @Param        page       query     int     false "Page number (default 0)"
// @Param        pageSize   query     int     false "Page size (default 20)"
// @Success      200 {object} []dto.UserSummary
// @Failure      400 {object} dto.ErrorResponse
// @Failure      403 {object} dto.ErrorResponse
// @Failure      500 {object} dto.ErrorResponse
// @Router       /users/{id}/followers [get]
// @Security	 BearerAuth
func (h *GetRelationshipsHandler) GetFollowers(c *gin.Context) {
	var idStr = c.Param("id")
	userId, uuidFormErr := uuid.Parse(idStr)

	if uuidFormErr != nil {
		respondErr(c, http.StatusBadRequest, uuidFormErr.Error())
		return
	}

	page, pageSize := getPaginationParams(c)

	token, tokenErr := getAuthorizationToken(c)
	if tokenErr != nil {
		respondErr(c, http.StatusBadRequest, tokenErr.Error())
		return
	}

	followers, err := h.getRelsUc.GetFollowers(userId, page, pageSize, token)

	if errors.Is(err, domain.ErrInvalidToken) {
		respondErr(c, http.StatusForbidden, err.Error())
		return
	}

	if err != nil {
		respondErr(c, http.StatusInternalServerError, err.Error())
		log.Println("Unexpected err: ", err)
		return
	}

	c.IndentedJSON(200, followers)
}

// GetFollowees godoc
// @Summary      Get user's followees
// @Description  Returns paginated list of user's followees (those who user follows)
// @Tags         relationships
// @Accept       json
// @Produce      json
// @Param        id         path      string  true  "User ID"
// @Param        page       query     int     false "Page number (default 0)"
// @Param        pageSize   query     int     false "Page size (default 20)"
// @Success      200 {object} []dto.UserSummary
// @Failure      400 {object} dto.ErrorResponse
// @Failure      403 {object} dto.ErrorResponse
// @Failure      500 {object} dto.ErrorResponse
// @Router       /users/{id}/followees [get]
// @Security	 BearerAuth
func (h *GetRelationshipsHandler) GetFollowees(c *gin.Context) {
	var idStr = c.Param("id")
	userId, uuidFormErr := uuid.Parse(idStr)

	if uuidFormErr != nil {
		respondErr(c, http.StatusBadRequest, uuidFormErr.Error())
		return
	}

	page, pageSize := getPaginationParams(c)

	token, tokenErr := getAuthorizationToken(c)
	if tokenErr != nil {
		respondErr(c, http.StatusBadRequest, tokenErr.Error())
		return
	}

	followers, err := h.getRelsUc.GetFollowees(userId, page, pageSize, token)

	if errors.Is(err, domain.ErrInvalidToken) {
		respondErr(c, http.StatusForbidden, err.Error())
		return
	}

	if err != nil {
		respondErr(c, http.StatusInternalServerError, err.Error())
		log.Println("Unexpected err: ", err)
		return
	}

	c.IndentedJSON(200, followers)
}

// GetRelationship godoc
// @Summary      Get relationship status
// @Description  Returns relationship status of user of given id to user-sender
// @Tags         relationships
// @Accept       json
// @Produce      json
// @Param        id         path      string  true  "User ID"
// @Success      200 {object} []dto.RelationshipStatus
// @Failure      400 {object} dto.ErrorResponse
// @Failure      403 {object} dto.ErrorResponse
// @Failure      500 {object} dto.ErrorResponse
// @Router       /users/{id}/relationship [get]
// @Security	 BearerAuth
func (h *GetRelationshipsHandler) GetRelationship(c *gin.Context) {
	var idStr = c.Param("id")
	userId, uuidFormErr := uuid.Parse(idStr)

	if uuidFormErr != nil {
		respondErr(c, http.StatusBadRequest, uuidFormErr.Error())
		return
	}

	token, tokenErr := getAuthorizationToken(c)
	if tokenErr != nil {
		respondErr(c, http.StatusBadRequest, tokenErr.Error())
		return
	}

	relStatus, err := h.getRelsUc.GetRelationshipStatus(userId, token)

	if errors.Is(err, domain.ErrInvalidToken) {
		respondErr(c, http.StatusForbidden, err.Error())
		return
	}

	if err != nil {
		respondErr(c, http.StatusInternalServerError, err.Error())
		log.Println("Unexpected err: ", err)
		return
	}

	c.IndentedJSON(200, relStatus)
}
