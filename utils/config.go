package utils

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)

var (
	JWTSecretKey    string
	JWTIssuer       string
	AccessTokenTTL  time.Duration
	RefreshTokenTTL time.Duration
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	JWTSecretKey = os.Getenv("JWT_SECRET")
	JWTIssuer = os.Getenv("JWT_ISSUER")
	var parseErr error
	AccessTokenTTLStr := os.Getenv("ACCESS_TOKEN_TTL")
	if AccessTokenTTLStr == "" {
		AccessTokenTTLStr = "5s"
	}
	AccessTokenTTL, parseErr = time.ParseDuration(AccessTokenTTLStr)
	if parseErr != nil {
		log.Fatalf("Ошибка: некорректное значение ACCESS_TOKEN_TTL: %v", parseErr)
	}
	RefreshTokenTTLStr := os.Getenv("REFRESH_TOKEN_TTL")
	if RefreshTokenTTLStr == "" {
		RefreshTokenTTLStr = "1h"
	}
	RefreshTokenTTL, parseErr = time.ParseDuration(RefreshTokenTTLStr)
	if parseErr != nil {
		log.Fatalf("Ошибка: некорректное значение REFRESH_TOKEN_TTL: %v", parseErr)
	}

	if JWTSecretKey == "" {
		log.Fatal("JWT_SECRET must be set in .env")
	}
	if JWTIssuer == "" {
		log.Fatal("JWT_ISSUER must be set in .env")
	}
}
