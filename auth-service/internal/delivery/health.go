package delivery

import (
	"github.com/gin-gonic/gin"
	"net/http"
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
