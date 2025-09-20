package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	DBSSLMode  string
	HTTPPort   string
	JWTSecret  string
	Env        string

	// JWT TTL
	AccessTokenTTLMinutes int
	RefreshTokenTTLHours  int
}

func LoadConfig() *Config {
	// Load ENV
	env := os.Getenv("ENV")

	envFile := ".env.local" // default
	if env == "prod" {
		envFile = ".env.prod"
	}

	if err := godotenv.Load(envFile); err != nil {
		log.Printf("⚠️ %s not found, fallback to system environment", envFile)
	}

	// Convert TTL
	accessTTL, _ := strconv.Atoi(getEnv("ACCESS_TOKEN_TTL_MIN", "15"))
	refreshTTL, _ := strconv.Atoi(getEnv("REFRESH_TOKEN_TTL_HOUR", "168"))

	// Check ENVIRONMENT
	log.Println("========================== ENVIRONMENT ==========================")
	log.Printf("🚀 Running with environment: %s", envFile)
	log.Printf("🚀 Running with DB_NAME: %s", os.Getenv("DB_NAME"))
	log.Printf("🚀 Running with DB_HOST: %s", os.Getenv("DB_HOST"))
	log.Printf("🔑 JWT_SECRET loaded: %v", os.Getenv("JWT_SECRET") != "")
	log.Println("=================================================================")

	return &Config{
		DBHost:                os.Getenv("DB_HOST"),
		DBPort:                os.Getenv("DB_PORT"),
		DBUser:                os.Getenv("DB_USER"),
		DBPassword:            os.Getenv("DB_PASSWORD"),
		DBName:                os.Getenv("DB_NAME"),
		DBSSLMode:             os.Getenv("DB_SSLMODE"),
		HTTPPort:              os.Getenv("HTTP_PORT"),
		JWTSecret:             os.Getenv("JWT_SECRET"),
		Env:                   env,
		AccessTokenTTLMinutes: accessTTL,
		RefreshTokenTTLHours:  refreshTTL,
	}
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
