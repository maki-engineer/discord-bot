package db

import (
	"fmt"

	"discord-bot/src/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDB(config config.Config) (*gorm.DB, error) {
	var dsn string

	if config.URL != "" {
		dsn = config.URL
	} else {
		dsn = fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=Asia/Tokyo",
			config.Host,
			config.User,
			config.Password,
			config.DBName,
		)
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
