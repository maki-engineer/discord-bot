package model

import "time"

type Session struct {
	SessionID string    `gorm:"not null"`
	ExpiresAt time.Time `gorm:"not null"`
}

func (Session) TableName() string {
	return "sessions"
}
