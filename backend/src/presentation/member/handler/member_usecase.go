package handler

import (
	"context"
	"discord-bot/src/application/member/usecase"
)

type MemberUseCase interface {
	GetMembersByBirthdayMonth(ctx context.Context, birthdayMonth int) ([]usecase.MemberBirthdayOutputData, error)
}
