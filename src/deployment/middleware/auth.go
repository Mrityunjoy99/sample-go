package middleware

import (
	"net/http"
	"strings"

	"github.com/Mrityunjoy99/sample-go/src/common/constant"
	"github.com/Mrityunjoy99/sample-go/src/domain/service/jwt"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware(jwtService jwt.JwtService, accessThrshold constant.UserType) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Auth")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Missing Auth header"})
			return
		}

		tokenStr := strings.TrimSpace(authHeader)

		token, gerr := jwtService.ValidateToken(tokenStr)
		if gerr != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": gerr.Error()})
			return
		}

		if token.UserType > accessThrshold {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "User is forbidden"})
			return
		}

		c.Next()
	}
}
