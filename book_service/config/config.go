package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	AuthServiceURL string
)

func InitConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Println("не удалось загрузить .env, используем системные переменные")
	}

	AuthServiceURL = os.Getenv("AUTH_SERVICE_URL")
	if AuthServiceURL == "" {
		log.Fatal("Ошибка: переменная AUTH_SERVICE_URL не задана в .env")
	}
}
