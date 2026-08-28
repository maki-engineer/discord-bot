package repository

import (
	"context"
	"discord-bot/src/domain/member"
	"discord-bot/src/infrastructure/model"

	"gorm.io/gorm"
)

type MemberRepository struct {
	db *gorm.DB
}

func NewMemberRepository(db *gorm.DB) *MemberRepository {
	return &MemberRepository{
		db: db,
	}
}

func (r *MemberRepository) GetMembersByBirthdayMonth(ctx context.Context, birthdayMonth member.Month) ([]*member.MemberBirthday, error) {
	var members []*member.MemberBirthday
	err := r.db.
		WithContext(ctx).
		Model(&model.Member{}).
		Select("name, month, date").
		Where("month = ?", birthdayMonth).
		Order("date ASC").
		Find(&members).Error

	if err != nil {
		return nil, err
	}

	return members, nil
}
