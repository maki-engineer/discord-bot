package service

import (
	"context"
	"discord-bot/src/domain/auth"
	"reflect"
	"testing"
	"time"
)

type MockSessionRepository struct {
	Session *auth.Session
	err     error
}

func (m *MockSessionRepository) GetSessionByID(ctx context.Context, sessionID string) (*auth.Session, error) {
	return m.Session, m.err
}

func TestSessionService_GetSessionByID(t *testing.T) {
	sessionID := "342k5jk352kj"
	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		t.Fatal(err)
	}
	expiresAt := time.Date(2026, 10, 1, 0, 0, 0, 0, jst)

	repo := &MockSessionRepository{
		Session: &auth.Session{
			SessionID: sessionID,
			ExpiresAt: expiresAt,
		},
		err: nil,
	}

	expected := &auth.Session{
		SessionID: sessionID,
		ExpiresAt: expiresAt,
	}

	s := NewSessionService(repo)
	result, err := s.GetSessionByID(context.Background(), sessionID)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if !reflect.DeepEqual(result.SessionID, expected.SessionID) {
		t.Fatalf("Expected %v, but got %v", expected.SessionID, result.SessionID)
	}

	if !result.ExpiresAt.Equal(expected.ExpiresAt) {
		t.Fatalf("Expected %v, but got %v", expected.ExpiresAt, result.ExpiresAt)
	}
}
