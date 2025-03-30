package config

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	DB       *mongo.Database
	BookColl *mongo.Collection
)

func InitDatabase() {
	err := godotenv.Load()
	if err != nil {
		log.Println("не удалось загрузить .env, используем системные переменные")
	}

	mongoURI := os.Getenv("MONGO_URI")
	dbName := os.Getenv("MONGO_DB_NAME")
	collectionName := os.Getenv("MONGO_BOOKS_COLLECTION")

	clientOptions := options.Client().ApplyURI(mongoURI)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal("Ошибка подключения к MongoDB:", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("Ошибка: MongoDB не отвечает:", err)
	}

	DB = client.Database(dbName)
	BookColl = DB.Collection(collectionName)

	log.Println("Подключение к MongoDB установлено!")
}
