package usecase

import (
	"context"
	"errors"
	"testing"
	"time"

	"discord-bot/src/domain/auth"
)

type mockDiscordGateway struct {
	accessToken   string
	user          *auth.DiscordUser
	err           error
	membershipErr error
	verified      bool
}

func (m *mockDiscordGateway) ExchangeCode(ctx context.Context, code string) (string, error) {
	return m.accessToken, m.err
}

func (m *mockDiscordGateway) GetUser(ctx context.Context, accessToken string) (*auth.DiscordUser, error) {
	return m.user, m.err
}

func (m *mockDiscordGateway) VerifyGuildMember(ctx context.Context, userID string) error {
	m.verified = true
	return m.membershipErr
}

type mockSessionRepository struct {
	session *auth.Session
	err     error
}

func (m *mockSessionRepository) GetSessionByID(ctx context.Context, sessionID string) (*auth.Session, error) {
	return m.session, m.err
}

func (m *mockSessionRepository) CreateSession(ctx context.Context, session *auth.Session) error {
	m.session = session
	return m.err
}

func TestDiscordAuthUseCase_LoginCreatesSessionForGuildMember(t *testing.T) {
	gateway := &mockDiscordGateway{
		accessToken: "access-token",
		user:        &auth.DiscordUser{ID: "user-id"},
	}
	repository := &mockSessionRepository{}
	useCase := NewDiscordAuthUseCase(gateway, repository)

	sessionID, err := useCase.Login(context.Background(), "authorization-code")
	if err != nil {
		t.Fatalf("Login returned an error: %v", err)
	}
	if sessionID == "" || repository.session == nil {
		t.Fatal("expected a session to be created")
	}
	if !gateway.verified {
		t.Fatal("expected guild membership to be verified")
	}
	if !repository.session.ExpiresAt.After(time.Now()) {
		t.Fatal("expected the session to expire in the future")
	}
}

func TestDiscordAuthUseCase_LoginRejectsNonMember(t *testing.T) {
	gateway := &mockDiscordGateway{
		accessToken:   "access-token",
		user:          &auth.DiscordUser{ID: "user-id"},
		membershipErr: auth.ErrNotGuildMember,
	}
	repository := &mockSessionRepository{}
	useCase := NewDiscordAuthUseCase(gateway, repository)

	_, err := useCase.Login(context.Background(), "authorization-code")
	if !errors.Is(err, auth.ErrNotGuildMember) {
		t.Fatalf("expected ErrNotGuildMember, got %v", err)
	}
	if repository.session != nil {
		t.Fatal("expected no session to be created")
	}
}
