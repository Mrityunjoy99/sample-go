package user

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type controller struct {
	service Service
}

type Controller interface {
	GetUserById(c *gin.Context)
	CreateUser(c *gin.Context)
}

func NewController(service Service) Controller {
	return &controller{
		service: service,
	}
}

func (ctrl *controller) GetUserById(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "invalid id",
		})

		return
	}

	user, err := ctrl.service.GetUserById(id)
	if err != nil {
		if err.Error() == "record not found" {
			c.JSON(404, gin.H{
				"error": "user not found",
			})

			return
		}

		c.JSON(500, gin.H{
			"error": "internal server error",
		})

		return
	}

	c.JSON(200, user)
}

func (ctrl *controller) CreateUser(c *gin.Context) {
	var dto CreateUserDto
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(400, gin.H{
			"error": "invalid request",
		})

		return
	}

	user, err := ctrl.service.CreateUser(dto)
	if err != nil {
		c.JSON(500, gin.H{
			"error": "internal server error",
		})

		return
	}

	c.JSON(201, user)
}
