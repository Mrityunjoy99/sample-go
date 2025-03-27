package user

import (
	"net/http"

	"github.com/Mrityunjoy99/sample-go/src/common/constant"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type controller struct {
	service Service
}

type Controller interface {
	GetUserById(c *gin.Context)
	CreateUser(c *gin.Context)
	UpdateUser(c *gin.Context)
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
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid id",
		})

		return
	}

	user, gerr := ctrl.service.GetUserById(id)
	if gerr != nil {
		if gerr.GetCode() == constant.ErrorCodeResourceNotFound {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "user not found",
			})

			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": gerr.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, user)
}

func (ctrl *controller) CreateUser(c *gin.Context) {
	var dto CreateUserDto
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request",
		})

		return
	}

	user, gerr := ctrl.service.CreateUser(dto)
	if gerr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": gerr.Error(),
		})

		return
	}

	c.JSON(http.StatusCreated, user)
}

func (ctrl *controller) UpdateUser(c *gin.Context) {
	idStr := c.Param("id")

	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid id",
		})

		return
	}

	var dto UpdateUserDto
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request",
		})

		return
	}

	user, gerr := ctrl.service.UpdateUser(id, dto)
	if gerr != nil {
		if gerr.GetCode() == constant.ErrorCodeResourceNotFound {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "user not found",
			})

			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": gerr.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, user)
}
