package middleware

import (
	"net/http"

	"github.com/Ma1y0/gist-clone/helpers"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware(c *gin.Context) {
	// Extract JWT
	token, err := c.Cookie("jwt")
	if err != nil || token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
		c.Abort()
		return
	}

	// Validate JWt
	claims, err := helpers.VerifyJWT(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid token",
			"error":   err.Error(),
		})
		c.Abort()
		return
	}

	id, err := helpers.ExtractIdFromJWT(claims)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid token",
			"error":   err.Error(),
		})
		c.Abort()
		return
	}

	// Set a global state
	c.Set("ID", id)

	c.Next()
}
