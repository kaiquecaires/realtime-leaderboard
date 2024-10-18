package auth

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthRequired(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, map[string]string{"error": "Authorization header missing"})
		c.Abort()
		return
	}

	fields := strings.Fields(authHeader)
	if len(fields) != 2 || strings.ToLower(fields[0]) != "bearer" {
		c.JSON(http.StatusUnauthorized, map[string]string{"error": "invalid authorization format"})
		c.Abort()
		return
	}

	token := fields[1]

	_, err := VerifyToken(token)

	if err != nil {
		c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid token"})
		c.Abort()
		return
	}

	c.Next()
}
