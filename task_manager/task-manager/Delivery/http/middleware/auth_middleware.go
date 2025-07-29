package middleware

import (
	"net/http"
	"strings"
	"task_manager/Domain/interfaces"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware creates authentication middleware using token service
func AuthMiddleware(tokenService interfaces.TokenService) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing or invalid"})
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		email, role, err := tokenService.ValidateToken(tokenString)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			return
		}

		// Set user data into context
		c.Set("userEmail", email)
		c.Set("userRole", role)

		c.Next()
	}
} 