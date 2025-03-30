package utils

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func GetMongoClientOptions(uri string) *options.ClientOptions {
	clientOptions := options.Client().ApplyURI(uri)
	// Ping the MongoDB server to check the connection
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	// Try to connect to MongoDB
	client, err := GetMongoClient(ctx, clientOptions)
	if err != nil {
		log.Fatalf("Error creating MongoDB client: %v", err)
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatalf("Error pinging MongoDB: %v", err)
	}

	fmt.Println("Successfully connected to MongoDB!")

	return clientOptions
}

func GetMongoClient(ctx context.Context, clientOptions *options.ClientOptions) (*mongo.Client, error) {
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, fmt.Errorf("error connecting to MongoDB: %w", err)
	}

	return client, nil
}
