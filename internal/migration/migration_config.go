package migration

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	DBHost     string
	DBUser     string
	DBPassword string
	DBName     string
	DBPort     int
}

func LoadConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: .env file not found, relying on environment variables.")
	}

	idType := os.Getenv("MIGRATION_ID_TYPE")
	if idType == "" {
		idType = "serial"
	}

	port, _ := strconv.Atoi(os.Getenv("DB_PORT"))
	return &Config{
		DBHost:     os.Getenv("DB_HOST"),
		DBUser:     os.Getenv("DB_USERNAME"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBName:     os.Getenv("DB_DATABASE"),
		DBPort:     port,
	}, nil
}
