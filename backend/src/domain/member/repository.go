package member

type MemberRepository interface {
	GetMembersByBirthdayMonth(BirthdayMonth int) ([]*MemberBirthday, error)
}
