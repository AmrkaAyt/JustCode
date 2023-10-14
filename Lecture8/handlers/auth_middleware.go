package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const token = "tokenXcxzcasdKLDSAdxc"

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader != token {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}
		c.Next()
	}
}
