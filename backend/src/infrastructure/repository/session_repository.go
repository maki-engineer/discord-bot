package repository

import (
	"context"
	"discord-bot/src/domain/auth"
	"discord-bot/src/infrastructure/model"

	"gorm.io/gorm"
)

type SessionRepository struct {
	db *gorm.DB
}

func NewSessionRepository(db *gorm.DB) *SessionRepository {
	return &SessionRepository{
		db: db,
	}
}

func (r *SessionRepository) GetSessionByID(ctx context.Context, sessionId string) (*auth.Session, error) {
	var session model.Session
	err := r.db.WithContext(ctx).
		Model(&model.Session{}).
		Where("session_id = ?", sessionId).
		First(&session).Error

	if err != nil {
		return nil, err
	}

	return &auth.Session{
		SessionID: session.SessionID,
		ExpiresAt: session.ExpiresAt,
	}, nil
}
