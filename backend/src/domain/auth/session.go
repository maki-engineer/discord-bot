package auth

import "time"

type Session struct {
	SessionID string
	ExpiresAt time.Time
}
