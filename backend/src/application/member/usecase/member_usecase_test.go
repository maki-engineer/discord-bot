package usecase

import (
	"discord-bot/src/domain/member"
	"errors"
	"reflect"
	"testing"
)

type MockMemberRepository struct {
	members []*member.MemberBirthday
	err     error
}

func (m *MockMemberRepository) GetMembersByBirthdayMonth(birthdayMonth int) ([]*member.MemberBirthday, error) {
	return m.members, m.err
}

func TestMemberUseCase_GetMembersByBirthdayMonth(t *testing.T) {
	tests := []struct {
		name            string
		members         []*member.MemberBirthday
		birthdayMonth   int
		expectedCount   int
		expectedMembers []MemberBirthdayOutputData
	}{
		{
			name: "5月が誕生日のメンバーが取得できること",
			members: []*member.MemberBirthday{
				{Name: "Alice", Month: 5, Date: 15},
				{Name: "Bob", Month: 5, Date: 20},
			},
			birthdayMonth: 5,
			expectedCount: 2,
			expectedMembers: []MemberBirthdayOutputData{
				{Name: "Alice", Month: 5, Date: 15},
				{Name: "Bob", Month: 5, Date: 20},
			},
		},
		{
			name:            "指定された月の誕生日メンバーが存在しない場合、空のスライスが返ること",
			members:         []*member.MemberBirthday{},
			birthdayMonth:   7,
			expectedCount:   0,
			expectedMembers: []MemberBirthdayOutputData{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &MockMemberRepository{
				members: tt.members,
				err:     nil,
			}

			uc := NewMemberUseCase(repo)
			result, err := uc.GetMembersByBirthdayMonth(tt.birthdayMonth)
			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}

			if len(result) != tt.expectedCount {
				t.Fatalf("Expected %d members, but got %d", tt.expectedCount, len(result))
			}

			if !reflect.DeepEqual(result, tt.expectedMembers) {
				t.Fatalf("Expected %v, but got %v", tt.expectedMembers, result)
			}
		})
	}
}

func TestMemberUseCase_GetMembersByBirthdayMonth_RepositoryError(t *testing.T) {
	expectedError := errors.New("repository error")

	repo := &MockMemberRepository{
		members: nil,
		err:     expectedError,
	}

	uc := NewMemberUseCase(repo)
	_, err := uc.GetMembersByBirthdayMonth(5)
	if err == nil {
		t.Fatalf("Expected error, but got nil")
	}

	if !errors.Is(err, expectedError) {
		t.Fatalf("Expected error %v, but got %v", expectedError, err)
	}
}
