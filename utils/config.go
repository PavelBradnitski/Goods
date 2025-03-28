package utils

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func GetMongoClientOptions(uri string) *options.ClientOptions {
	clientOptions := options.Client().ApplyURI(uri)

	// Try to connect to MongoDB
	client, err := GetMongoClient(clientOptions)
	if err != nil {
		log.Fatalf("Error creating MongoDB client: %v", err)
	}

	// Ping the MongoDB server to check the connection
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatalf("Error pinging MongoDB: %v", err)
	}

	fmt.Println("Successfully connected to MongoDB!")

	return clientOptions
}

func GetMongoClient(clientOptions *options.ClientOptions) (*mongo.Client, error) {
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		return nil, fmt.Errorf("error creating MongoDB client: %w", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		return nil, fmt.Errorf("error connecting to MongoDB: %w", err)
	}

	return client, nil
}

// Константы JWT (лучше хранить в переменных окружения)
var (
	JWTSecretKey   string             // Example "my-secret-key"
	JWTIssuer      string             // Example "my-auth-service"
	AccessTokenTTL = time.Second * 10 // Время жизни access token
	// RefreshTokenTTL = time.Hour * 24 * 7 // Время жизни refresh token (неделя)
	RefreshTokenTTL = time.Hour
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	JWTSecretKey = os.Getenv("JWT_SECRET")
	JWTIssuer = os.Getenv("JWT_ISSUER")
	if JWTSecretKey == "" {
		log.Fatal("JWT_SECRET must be set in .env")
	}
	if JWTIssuer == "" {
		log.Fatal("JWT_ISSUER must be set in .env")
	}
}
