package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Controller interface {
	GenerateToken(ctx *gin.Context)
	ValidateToken(ctx *gin.Context)
}

type controller struct {
	service Service
}

func NewController(service Service) Controller {
	return &controller{service: service}
}

func (c *controller) GenerateToken(ctx *gin.Context) {
	var reqDto GenerateTokenReqDto

	if err := ctx.ShouldBindJSON(&reqDto); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := c.service.GenerateToken(reqDto.UserId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	respDto := GenerateTokenRespDto{
		Token: token,
	}

	ctx.JSON(http.StatusOK, respDto)
}

func (c *controller) ValidateToken(ctx *gin.Context) {
	var reqDto ValidateTokenReqDto
	if err := ctx.ShouldBindJSON(&reqDto); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	respDto, err := c.service.ValidateToken(reqDto.Token)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, respDto)
}
