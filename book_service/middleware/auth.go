package middleware

import (
	"io"
	"net/http"

	"log"

	"github.com/PavelBradnitski/Goods/book_service/config"
	"github.com/gin-gonic/gin"
)

// AuthMiddleware проверяет токен через auth-service
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing token"})
			c.Abort()
			return
		}

		log.Println(config.AuthServiceURL)
		req, err := http.NewRequest("GET", config.AuthServiceURL, nil)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Server error"})
			c.Abort()
			return
		}

		req.Header.Set("Authorization", token)
		log.Println(req)
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Auth service error"})
			c.Abort()
			return
		}
		defer resp.Body.Close()

		body, _ := io.ReadAll(resp.Body)

		if resp.StatusCode != http.StatusOK {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token", "details": string(body)})
			c.Abort()
			return
		}

		c.Next()
	}
}
