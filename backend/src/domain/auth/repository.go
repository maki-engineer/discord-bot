package auth

import "context"

type SessionRepository interface {
	GetSessionByID(ctx context.Context, sessionID string) (*Session, error)
}
