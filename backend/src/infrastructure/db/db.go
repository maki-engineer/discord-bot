package db

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	User     string
	Password string
	DBName   string
	URL      string
}

func NewDB(config Config) (*gorm.DB, error) {
	var dsn string

	if config.URL != "" {
		dsn = config.URL
	} else {
		dsn = fmt.Sprintf(
			"host=unittest.db user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=Asia/Tokyo",
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
