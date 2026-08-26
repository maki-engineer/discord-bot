package repository

import (
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

func (r *MemberRepository) GetMembersByBirthdayMonth(birthdayMonth int) ([]*member.MemberBirthday, error) {
	var members []*member.MemberBirthday
	err := r.db.
		Model(&model.Member{}).
		Select("name, month, date").
		Where("month = ?", birthdayMonth).
		Find(&members).Error

	if err != nil {
		return nil, err
	}

	return members, nil
}
