package middleware

import (
	"context"
	"discord-bot/src/domain/auth"
)

type SessionService interface {
	GetSessionByID(ctx context.Context, sessionID string) (*auth.Session, error)
}
