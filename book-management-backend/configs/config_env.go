package config

import (
	"log"
	"os"

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

	// Check ENVIRONMENT
	log.Println("========================== ENVIRONMENT ==========================")
	log.Printf("🚀 Running with environment: %s \n", envFile)
	log.Printf("🚀 Running with DB_NAME: %s", os.Getenv("DB_NAME"))
	log.Printf("🚀 Running with DB_HOST: %s", os.Getenv("DB_HOST"))
	log.Println("=================================================================")

	return &Config{
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBName:     os.Getenv("DB_NAME"),
		DBSSLMode:  os.Getenv("DB_SSLMODE"),
		HTTPPort:   os.Getenv("HTTP_PORT"),
	}
}
