package config

import (
	"os"
)

type Config struct {
	Host                string
	User                string
	Password            string
	DBName              string
	URL                 string
	AllowURL            string
	FrontendURL         string
	BasicUserName       string
	BasicPassword       string
	DiscordClientID     string
	DiscordClientSecret string
	DiscordRedirectURI  string
	DiscordGuildID      string
	DiscordBotToken     string
}

func LoadConfig() Config {
	env := os.Getenv("APP_ENV")

	if env == "production" {
		return Config{
			URL:                 os.Getenv("POSTGRES_URL"),
			AllowURL:            os.Getenv("VERCEL_URL"),
			FrontendURL:         os.Getenv("FRONTEND_URL"),
			BasicUserName:       os.Getenv("BASIC_AUTH_USERNAME"),
			BasicPassword:       os.Getenv("BASIC_AUTH_PASSWORD"),
			DiscordClientID:     os.Getenv("DISCORD_CLIENT_ID"),
			DiscordClientSecret: os.Getenv("DISCORD_CLIENT_SECRET"),
			DiscordRedirectURI:  os.Getenv("DISCORD_REDIRECT_URI"),
			DiscordGuildID:      os.Getenv("DISCORD_GUILD_ID"),
			DiscordBotToken:     os.Getenv("DISCORD_BOT_TOKEN"),
		}
	}

	if env == "development" {
		return Config{
			Host:                "development-db",
			User:                os.Getenv("POSTGRES_USER_DEVELOPMENT"),
			Password:            os.Getenv("POSTGRES_PASSWORD_DEVELOPMENT"),
			DBName:              os.Getenv("POSTGRES_DB_DEVELOPMENT"),
			AllowURL:            os.Getenv("VERCEL_URL"),
			FrontendURL:         os.Getenv("FRONTEND_URL"),
			BasicUserName:       os.Getenv("BASIC_AUTH_USERNAME"),
			BasicPassword:       os.Getenv("BASIC_AUTH_PASSWORD"),
			DiscordClientID:     os.Getenv("DISCORD_CLIENT_ID"),
			DiscordClientSecret: os.Getenv("DISCORD_CLIENT_SECRET"),
			DiscordRedirectURI:  os.Getenv("DISCORD_REDIRECT_URI"),
			DiscordGuildID:      os.Getenv("DISCORD_GUILD_ID"),
			DiscordBotToken:     os.Getenv("DISCORD_BOT_TOKEN"),
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
