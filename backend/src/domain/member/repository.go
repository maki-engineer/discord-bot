package member

type MemberRepository interface {
	GetMembersByBirthdayMonth(birthdayMonth int) ([]*MemberBirthday, error)
}
