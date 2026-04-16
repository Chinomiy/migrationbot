package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBURL      string
	TgToken    string
	AdminToken string
}

func MustLoad() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	token := os.Getenv("TG_TOKEN")
	if token == "" {
		log.Fatal("TG_TOKEN env variable not set")
	}
	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("DB_URL env variable not set")
	}
	adminToken := os.Getenv("ADMIN_TOKEN")
	if adminToken == "" {
		log.Fatal("ADMIN_TOKEN env variable not set")
	}
	return &Config{
		TgToken:    token,
		DBURL:      dbURL,
		AdminToken: adminToken,
	}, nil
}
