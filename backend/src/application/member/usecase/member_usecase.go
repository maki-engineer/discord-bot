package usecase

import (
	"context"
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

func (uc *MemberUseCase) GetMembersByBirthdayMonth(ctx context.Context, birthdayMonth int) ([]MemberBirthdayOutputData, error) {
	month, err := member.NewMonth(birthdayMonth)
	if err != nil {
		return nil, err
	}

	members, err := uc.repo.GetMembersByBirthdayMonth(ctx, month)
	if err != nil {
		return nil, err
	}

	outputData := make([]MemberBirthdayOutputData, len(members))
	for i, member := range members {
		outputData[i] = MemberBirthdayOutputData{
			Name:  member.Name,
			Month: member.Month,
			Date:  member.Date,
		}
	}

	return outputData, nil
}
