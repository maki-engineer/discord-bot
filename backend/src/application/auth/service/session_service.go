package service

import (
	"context"
	"discord-bot/src/domain/auth"
	"fmt"
	"time"
)

type SessionService struct {
	repo auth.SessionRepository
}

func NewSessionService(repo auth.SessionRepository) *SessionService {
	return &SessionService{
		repo: repo,
	}
}

func (s *SessionService) GetSessionByID(ctx context.Context, sessionID string) (*auth.Session, error) {
	session, err := s.repo.GetSessionByID(ctx, sessionID)
	if err != nil {
		return nil, err
	}

	if session.ExpiresAt.Before(time.Now()) {
		return nil, fmt.Errorf("session has expired")
	}

	return session, nil
}
