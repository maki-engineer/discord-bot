package member

import "context"

type MemberRepository interface {
	GetMembersByBirthdayMonth(ctx context.Context, birthdayMonth Month) ([]*MemberBirthday, error)
}
