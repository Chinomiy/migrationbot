package config

import (
	"errors"
	"migtationbot/logger"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBURL      string
	TgToken    string
	AdminToken string
}

var (
	ErrTGTokenNotFound    = errors.New("TG token not found")
	ErrDBURLNotFound      = errors.New("DB URL not found")
	ErrAdminTokenNotFound = errors.New("Admin token not found")
)

func MustLoad() *Config {
	if err := godotenv.Load(); err != nil {
		logger.Info(".env not found (ok in prod)")
	}

	token := os.Getenv("TG_TOKEN")
	if token == "" {
		logger.Error(ErrTGTokenNotFound)
		os.Exit(1)
	}

	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		logger.Error(ErrDBURLNotFound)
		os.Exit(1)
	}

	adminToken := os.Getenv("ADMIN_TOKEN")
	if adminToken == "" {
		logger.Error(ErrAdminTokenNotFound)
		os.Exit(1)
	}

	return &Config{
		TgToken:    token,
		DBURL:      dbURL,
		AdminToken: adminToken,
	}
}
