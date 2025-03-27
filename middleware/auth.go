package middleware

import (
	"net/http"

	"github.com/PavelBradnitski/Goods/utils"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware проверяет JWT токен и добавляет ID пользователя в контекст
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
			return
		}

		// Разделяем заголовок, чтобы получить только токен (предполагается "Bearer <token>")
		tokenString := authHeader[len("Bearer "):]
		if tokenString == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization header"})
			return
		}

		claims, err := utils.ValidateToken(tokenString)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		// Добавляем ID пользователя в контекст
		c.Set("userID", claims.UserID)
		c.Next()
	}
}
