package main

import (
	"log"
	"os"

	"github.com/PavelBradnitski/Goods/book_service/config"
	"github.com/PavelBradnitski/Goods/book_service/routes"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/PavelBradnitski/Goods/book_service/docs"
)

// @title Book Service API
// @version 1.0
// @description API для управления книгами
// @host localhost:8081
// @BasePath /api/v1
func main() {
	config.InitConfig()
	config.InitDatabase()

	router := routes.SetupRouter()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	log.Println("Server running at http://localhost:8081")
	log.Println("Swagger docs available at http://localhost:8081/swagger/index.html")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
