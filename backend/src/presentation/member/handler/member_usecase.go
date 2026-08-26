package handler

import "discord-bot/src/application/member/usecase"

type MemberUseCase interface {
	GetMembersByBirthdayMonth(birthdayMonth int) ([]usecase.MemberBirthdayOutputData, error)
}
