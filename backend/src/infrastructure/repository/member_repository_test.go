package repository

import (
	"context"
	"reflect"
	"testing"

	"discord-bot/src/config"
	"discord-bot/src/domain/member"
	"discord-bot/src/infrastructure/db"
	"discord-bot/src/infrastructure/model"
)

func TestMemberRepository_GetMembersByBirthdayMonth(t *testing.T) {
	config := config.LoadConfig()
	db, err := db.NewDB(config)
	if err != nil {
		t.Fatalf("Failed to connect to the database: %v", err)
	}

	defer db.Exec("TRUNCATE TABLE birthday_for_235_members RESTART IDENTITY CASCADE")

	members := []model.Member{
		{Name: "Bob", UserID: "23456", Month: 5, Date: 20, SpeakerID: 2},
		{Name: "Alice", UserID: "12345", Month: 5, Date: 15, SpeakerID: 1},
		{Name: "Charlie", UserID: "34567", Month: 6, Date: 10, SpeakerID: 3},
	}

	if err := db.Create(&members).Error; err != nil {
		t.Fatalf("Failed to insert test data: %v", err)
	}

	repo := NewMemberRepository(db)

	expected := []*member.MemberBirthday{
		{Name: "Alice", Month: 5, Date: 15},
		{Name: "Bob", Month: 5, Date: 20},
	}

	tests := []struct {
		name            string
		birthdayMonth   int
		expectedCount   int
		expectedMembers []*member.MemberBirthday
	}{
		{
			name:            "5月が誕生日のメンバーが昇順に取得できること",
			birthdayMonth:   5,
			expectedCount:   2,
			expectedMembers: expected,
		},
		{
			name:            "指定された月の誕生日メンバーが存在しない場合、空のスライスが返ること",
			birthdayMonth:   7,
			expectedCount:   0,
			expectedMembers: []*member.MemberBirthday{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			month, err := member.NewMonth(tt.birthdayMonth)
			result, err := repo.GetMembersByBirthdayMonth(context.Background(), month)

			if err != nil {
				t.Fatalf("Failed to get members by birthday month: %v", err)
			}

			if len(result) != tt.expectedCount {
				t.Errorf("Expected %d members, but got %d", tt.expectedCount, len(result))
			}

			if !reflect.DeepEqual(result, tt.expectedMembers) {
				t.Errorf("Expected %v, but got %v", tt.expectedMembers, result)
			}
		})
	}
}
