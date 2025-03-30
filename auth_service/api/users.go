package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/PavelBradnitski/Goods/auth_service/models"
)

// @Summary Get current user information
// @Description Get information about the currently authenticated user
// @Tags users
// @Produce json
// @Success 200 {object} models.User
// @Failure 401 {object} map[string]interface{}
// @Security BearerAuth
// @Router /users/me [get]
func GetCurrentUser(c *gin.Context) {
	userID := c.GetString("userID")

	user, err := models.FindByID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user"})
		return
	}

	user.Password = ""

	c.JSON(http.StatusOK, user)
}

// @Summary Validate User with BearerToken
// @Description Validate User with BearerToken
// @Tags users
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Security BearerAuth
// @Router /users/validate [get]
func ValidateToken(c *gin.Context) {
	userID := c.GetString("userID")

	_, err := models.FindByID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Token is valid"})
}
