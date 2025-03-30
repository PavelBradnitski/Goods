package routes

import (
	"github.com/PavelBradnitski/Goods/book_service/handlers"
	"github.com/PavelBradnitski/Goods/book_service/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	api := router.Group("/api/v1/books")
	api.Use(middleware.AuthMiddleware())
	{
		api.POST("/", handlers.CreateBook)
		api.GET("/", handlers.GetBooks)
		api.GET("/:id", handlers.GetBookByID)
		api.PUT("/:id", handlers.UpdateBook)
		api.DELETE("/:id", handlers.DeleteBook)
	}

	return router
}
