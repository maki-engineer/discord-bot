package config

import (
	"os"
)

type Config struct {
	Host     string
	User     string
	Password string
	DBName   string
	URL      string
}

func LoadConfig() Config {
	env := os.Getenv("APP_ENV")

	if env == "development" {
		return Config{
			URL: os.Getenv("POSTGRES_URL"),
		}
	}

	if env == "unittest" {
		return Config{
			Host:     "unittest-db",
			User:     os.Getenv("POSTGRES_USER_UNITTEST"),
			Password: os.Getenv("POSTGRES_PASSWORD_UNITTEST"),
			DBName:   os.Getenv("POSTGRES_DB_UNITTEST"),
		}
	}

	panic("unknown APP_ENV")
}
