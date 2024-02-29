package healthcheck

import "github.com/gin-gonic/gin"

type controller struct{}

type Controller interface {
	HealthCheck(c *gin.Context)
}

func NewController() Controller {
	return &controller{}
}

func (c *controller) HealthCheck(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "success",
	})
}
