package config

import (
	"migtationbot/logger"
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
		logger.Info("Error loading .env file")
	}
	token := os.Getenv("TG_TOKEN")
	if token == "" {
		logger.Info("TG_TOKEN env variable not set")
	}
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		logger.Info("DATABASE_URL env variable not set")
	}
	adminToken := os.Getenv("ADMIN_TOKEN")
	if adminToken == "" {
		logger.Info("ADMIN_TOKEN env variable not set")
	}
	return &Config{
		TgToken:    token,
		DBURL:      dbURL,
		AdminToken: adminToken,
	}, nil
}
