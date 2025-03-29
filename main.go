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

	_ "github.com/PavelBradnitski/Goods/docs" // Импорт сгенерированной документации Swagger

	"github.com/PavelBradnitski/Goods/api"
	"github.com/PavelBradnitski/Goods/db/migrations"
	"github.com/PavelBradnitski/Goods/middleware"
	"github.com/PavelBradnitski/Goods/utils"
)

// @title Auth API
// @version 1.0
// @description Microservice for authentication using JWT.
// @termsOfService http://swagger.io/terms/
// @host localhost:8080
// @BasePath /api/v1
func main() {
	// Загрузка конфигурации из .env
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	// Инициализация MongoDB
	mongoURI := os.Getenv("MONGO_URI")
	dbName := os.Getenv("MONGO_DB_NAME")
	if mongoURI == "" || dbName == "" {
		log.Fatal("MONGO_URI and MONGO_DB_NAME must be set in .env")
	}

	err = mgm.SetDefaultConfig(nil, dbName, utils.GetMongoClientOptions(mongoURI))
	if err != nil {
		log.Fatalf("Error connecting to MongoDB: %v", err)
	}

	// Запуск миграции
	if err := migrations.Up(); err != nil {
		log.Fatalf("Ошибка при выполнении миграции: %v", err)
	}
	// Создание роутера Gin
	router := gin.Default()

	// Настройка маршрутов
	v1 := router.Group("/api/v1")
	{
		// Авторизация
		auth := v1.Group("/auth")
		{
			auth.POST("/register", api.Register)
			auth.POST("/login", api.Login)
			auth.POST("/refresh", api.RefreshToken)
		}

		// Пользователи (требуется авторизация)
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
