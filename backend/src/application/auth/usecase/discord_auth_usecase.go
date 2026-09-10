package usecase

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"log"
	"time"

	"discord-bot/src/domain/auth"
)

const sessionDuration = 24 * time.Hour

type DiscordAuthUseCase struct {
	discordGateway  auth.DiscordGateway
	sessionRepo     auth.SessionRepository
	sessionDuration time.Duration
}

func NewDiscordAuthUseCase(discordGateway auth.DiscordGateway, sessionRepo auth.SessionRepository) *DiscordAuthUseCase {
	return &DiscordAuthUseCase{
		discordGateway:  discordGateway,
		sessionRepo:     sessionRepo,
		sessionDuration: sessionDuration,
	}
}

func (u *DiscordAuthUseCase) Login(ctx context.Context, code string) (string, error) {
	accessToken, err := u.discordGateway.ExchangeCode(ctx, code)
	if err != nil {
		return "", fmt.Errorf("exchange Discord authorization code: %w", err)
	}

	log.Println("discord login: ExchangeCode success")

	user, err := u.discordGateway.GetUser(ctx, accessToken)
	if err != nil {
		return "", fmt.Errorf("get Discord user: %w", err)
	}

	log.Println("discord login: GetUser success")

	if err := u.discordGateway.VerifyGuildMember(ctx, user.ID); err != nil {
		return "", fmt.Errorf("verify Discord guild membership: %w", err)
	}

	log.Println("discord login: verifyGuildMember success")

	sessionID, err := generateSessionID()
	if err != nil {
		return "", err
	}

	log.Println("discord login: generateSessionID success")

	if err := u.sessionRepo.CreateSession(ctx, &auth.Session{
		SessionID: sessionID,
		ExpiresAt: time.Now().Add(u.sessionDuration),
	}); err != nil {
		return "", fmt.Errorf("create session: %w", err)
	}

	log.Println("discord login: CreateSession success")

	return sessionID, nil
}

func generateSessionID() (string, error) {
	value := make([]byte, 32)
	if _, err := rand.Read(value); err != nil {
		return "", fmt.Errorf("generate session ID: %w", err)
	}
	return hex.EncodeToString(value), nil
}
