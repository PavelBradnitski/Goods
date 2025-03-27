package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/PavelBradnitski/Goods/models"
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
	userID := c.GetString("userID") // Получаем ID пользователя из контекста

	user, err := models.FindByID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user"})
		return
	}

	// Не показываем пароль
	user.Password = ""

	c.JSON(http.StatusOK, user)
}
