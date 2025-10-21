package delivery

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HealthCheck godoc
//
//	@Summary		Service healthcheck
//	@Tags			Health
//	@Success		200    "OK"
//	@Router			/health [get]
func HealthCheck(c *gin.Context) {
	c.String(http.StatusOK, "OK")
}
