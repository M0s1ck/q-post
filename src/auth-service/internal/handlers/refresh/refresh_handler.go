package handlers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	"auth-service/internal/domain"
	"auth-service/internal/dto"
	"auth-service/internal/usecase/refresh"
)

type RefreshHandler struct {
	ucase *refresh.RefreshUsecase
}

func NewRefreshHandler(ucase *refresh.RefreshUsecase) *RefreshHandler {
	return &RefreshHandler{
		ucase: ucase,
	}
}

func (hand *RefreshHandler) RegisterHandlers(engine *gin.Engine) {
	engine.POST("/refresh", hand.refresh)
}

// refresh godoc
//
//	@Summary		Refresh pipeline
//	@Description	Takes refresh token, returns jwt if it's valid
//	@Tags			Refresh
//	@Accept			json
//	@Produce		json
//	@Param   	    refresh body    dto.RefreshDto true "refresh"
//	@Failure		400	{object}	dto.ErrorResponse
//	@Failure		401	{object}	dto.ErrorResponse
//	@Failure		500	{object}	dto.ErrorResponse
//	@Router			/refresh [post]
func (hand *RefreshHandler) refresh(c *gin.Context) {
	refreshDto := dto.RefreshDto{}
	parseErr := c.BindJSON(&refreshDto)

	if parseErr != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Message: parseErr.Error()})
		return
	}

	response, err := hand.ucase.Refresh(refreshDto.Token)

	if errors.Is(err, domain.ErrInvalidToken) || errors.Is(err, domain.ErrNotFound) {
		c.JSON(http.StatusUnauthorized, dto.ErrorResponse{Message: err.Error()})
		return
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}
