package healthcheck

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type controller struct{}

type Controller interface {
	HealthCheck(c *gin.Context)
}

func NewController() Controller {
	return &controller{}
}

func (c *controller) HealthCheck(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}
