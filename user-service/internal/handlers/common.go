package handlers

import (
	"errors"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"user-service/internal/domain"
	"user-service/internal/dto"
	"user-service/internal/usecase/relationships"
)

func handleDefaultDomainErrors(c *gin.Context, err error) {
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

	if errors.Is(err, domain.ErrDuplicate) {
		respondErr(c, http.StatusConflict, err.Error())
		return
	}

	if err != nil {
		respondErr(c, http.StatusInternalServerError, err.Error())
		log.Println("Unexpected err: ", err)
		return
	}
}

// handleGettingToken gets token or responds 400 StatusBadRequest
func handleGettingToken(c *gin.Context) (token string, ok bool) {
	token, tokenErr := getAuthorizationToken(c)
	if tokenErr != nil {
		respondErr(c, http.StatusBadRequest, tokenErr.Error())
		return "", false
	}

	return token, true
}

// handleGettingId gets id from path, token from headers or responds 400 StatusBadRequest if smth went wrong
func handleGettingId(c *gin.Context, idParam string) (id uuid.UUID, ok bool) {
	var idStr = c.Param(idParam)
	id, uuidFormErr := uuid.Parse(idStr)

	if uuidFormErr != nil {
		respondErr(c, http.StatusBadRequest, uuidFormErr.Error())
		return uuid.Nil, false
	}

	return id, true
}

// handleGettingBodyAndToken parses body, gets token or responds 400 StatusBadRequest
func handleGettingBodyAndToken[T any](c *gin.Context) (*T, string, bool) {
	var body T
	bindErr := c.BindJSON(&body)

	if bindErr != nil {
		respondErr(c, http.StatusBadRequest, bindErr.Error())
		return nil, "", false
	}

	token, tokenErr := getAuthorizationToken(c)
	if tokenErr != nil {
		respondErr(c, http.StatusBadRequest, tokenErr.Error())
		return nil, "", false
	}

	return &body, token, true
}

// handleGettingIdAndToken looks for id in path, token in headers, responds with status bad request if smth went wrong
func handleGettingIdAndToken(c *gin.Context, idParam string) (id uuid.UUID, token string, ok bool) {
	var idStr = c.Param(idParam)
	id, uuidFormErr := uuid.Parse(idStr)

	if uuidFormErr != nil {
		respondErr(c, http.StatusBadRequest, uuidFormErr.Error())
		return uuid.Nil, "", false
	}

	token, tokenErr := getAuthorizationToken(c)
	if tokenErr != nil {
		respondErr(c, http.StatusBadRequest, tokenErr.Error())
		return uuid.Nil, "", false
	}

	return id, token, true
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

func getPaginationParams(c *gin.Context) (page int, pageSize int) {
	pageStr := c.DefaultQuery("page", "0")
	pageSizeStr := c.DefaultQuery("pageSize", "0")

	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 0 {
		page = 0
	}
	pSize, err := strconv.Atoi(pageSizeStr)
	if err != nil || pSize <= 0 {
		pSize = relationships.DefaultPageSize
	}

	return page, pSize
}
