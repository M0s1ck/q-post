package handlers

import (
	"errors"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"

	"user-service/internal/dto"
	"user-service/internal/usecase/relationships"
)

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
