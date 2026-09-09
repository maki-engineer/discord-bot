package repository

import (
	"context"
	"discord-bot/src/config"
	"discord-bot/src/domain/auth"
	"discord-bot/src/infrastructure/db"
	"discord-bot/src/infrastructure/model"
	"reflect"
	"testing"
	"time"

	"gorm.io/gorm"
)

func TestSessionRepository_GetSessionByID(t *testing.T) {
	config := config.LoadConfig()
	db, err := db.NewDB(config)
	if err != nil {
		t.Fatalf("Failed to connect to the database: %v", err)
	}

	defer db.Exec("TRUNCATE TABLE sessions RESTART IDENTITY CASCADE")

	sessionID := "j32k2iei3ji23"
	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		t.Fatal(err)
	}
	expiresAt := time.Date(2026, 10, 1, 10, 0, 0, 0, jst)

	session := model.Session{
		SessionID: sessionID,
		ExpiresAt: expiresAt,
	}

	if err := db.Create(&session).Error; err != nil {
		t.Fatalf("Failed to insert test data: %v", err)
	}

	repo := NewSessionRepository(db)

	expected := &auth.Session{
		SessionID: sessionID,
		ExpiresAt: expiresAt,
	}

	result, err := repo.GetSessionByID(context.Background(), sessionID)
	if err == gorm.ErrRecordNotFound {
		t.Fatalf("Failed to get session by session_id: %v", err)
	}

	if !reflect.DeepEqual(result.SessionID, expected.SessionID) {
		t.Errorf("Expected %v, but got %v", expected.SessionID, result.SessionID)
	}

	if !result.ExpiresAt.Equal(expected.ExpiresAt) {
		t.Errorf("Expected %v, but got %v", expected.ExpiresAt, result.ExpiresAt)
	}
}
