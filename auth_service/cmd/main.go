package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/kamva/mgm/v3"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/PavelBradnitski/Goods/auth_service/docs" // Импорт сгенерированной документации Swagger

	"github.com/PavelBradnitski/Goods/auth_service/api"
	"github.com/PavelBradnitski/Goods/auth_service/db/migrations"
	"github.com/PavelBradnitski/Goods/auth_service/middleware"
	"github.com/PavelBradnitski/Goods/auth_service/utils"
)

// @title Auth API
// @version 1.0
// @description Microservice for authentication using JWT.
// @termsOfService http://swagger.io/terms/
// @host localhost:8080
// @BasePath /api/v1
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	mongoURI := os.Getenv("MONGO_URI")
	dbName := os.Getenv("MONGO_DB_NAME")
	if mongoURI == "" || dbName == "" {
		log.Fatal("MONGO_URI and MONGO_DB_NAME must be set in .env")
	}

	err = mgm.SetDefaultConfig(nil, dbName, utils.GetMongoClientOptions(mongoURI))
	if err != nil {
		log.Fatalf("Error connecting to MongoDB: %v", err)
	}

	if err := migrations.Up(); err != nil {
		log.Fatalf("Ошибка при выполнении миграции: %v", err)
	}
	router := gin.Default()

	v1 := router.Group("/api/v1")
	{
		auth := v1.Group("/auth")
		{
			auth.POST("/register", api.Register)
			auth.POST("/login", api.Login)
			auth.POST("/refresh", api.RefreshToken)
		}

		users := v1.Group("/users")
		users.Use(middleware.AuthMiddleware())
		{
			users.GET("/me", api.GetCurrentUser)
		}
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Printf("Server listening on port %s\n", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
