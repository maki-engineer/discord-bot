package auth

import (
	"context"
	"errors"
)

var ErrNotGuildMember = errors.New("user is not a member of the configured guild")

type DiscordUser struct {
	ID string
}

type DiscordGateway interface {
	ExchangeCode(ctx context.Context, code string) (string, error)
	GetUser(ctx context.Context, accessToken string) (*DiscordUser, error)
	VerifyGuildMember(ctx context.Context, userID string) error
}
