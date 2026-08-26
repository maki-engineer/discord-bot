package usecase

import (
	"discord-bot/src/domain/member"
)

type MemberUseCase struct {
	repo member.MemberRepository
}

func NewMemberUseCase(repo member.MemberRepository) *MemberUseCase {
	return &MemberUseCase{
		repo: repo,
	}
}

func (uc *MemberUseCase) GetMembersByBirthdayMonth(birthdayMonth int) ([]*member.MemberBirthday, error) {
	return uc.repo.GetMembersByBirthdayMonth(birthdayMonth)
}
