package url

import (
	"net/http"

	"github.com/Mrityunjoy99/sample-go/src/common/constant"
	"github.com/gin-gonic/gin"
)

type Controller interface {
	RegisterUrl(c *gin.Context)
	Redirect(c *gin.Context)
}

type controller struct {
	service Service
}

func NewController(service Service) Controller {
	return &controller{
		service: service,
	}
}

func (ctrl *controller) RegisterUrl(c *gin.Context) {
	var dto RegisterUrlReqDto
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request",
		})

		return
	}

	url, gerr := ctrl.service.RegisterUrl(dto.OriginalURL, dto.TTLinSec)
	if gerr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": gerr.Error(),
		})

		return
	}

	c.JSON(http.StatusCreated, url)
}

func (ctrl *controller) Redirect(c *gin.Context) {
	url := c.Param("url")

	originalUrl, gerr := ctrl.service.Redirect(url)
	if gerr != nil {
		if gerr.GetCode() == constant.ErrorCodeResourceNotFound {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "url does not exists",
			})

			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": gerr.Error(),
		})

		return
	}

	c.JSON(http.StatusPermanentRedirect, originalUrl)
}
