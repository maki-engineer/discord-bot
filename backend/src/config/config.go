package config

import (
	"discord-bot/src/infrastructure/db"
	"os"
)

func LoadConfig() db.Config {
	env := os.Getenv("APP_ENV")

	if env == "development" {
		return db.Config{
			URL: os.Getenv("POSTGRES_URL"),
		}
	}

	if env == "unittest" {
		return db.Config{
			User:     os.Getenv("POSTGRES_USER_UNITTEST"),
			Password: os.Getenv("POSTGRES_PASSWORD_UNITTEST"),
			DBName:   os.Getenv("POSTGRES_DB_UNITTEST"),
		}
	}

	panic("unknown APP_ENV")
}
