package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	SQLiteDBPath string
}

var AppConfig *Config

func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
	}

	dbPath := os.Getenv("SQLITE_DB_PATH")
	if dbPath == "" {
		log.Fatal("SQLITE_DB_PATH is not set in environment variables")
	}

	AppConfig = &Config{
		SQLiteDBPath: dbPath,
	}
}
