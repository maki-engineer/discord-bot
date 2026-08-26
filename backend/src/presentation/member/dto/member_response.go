package dto

type MemberBirthday struct {
	Name  string `json:"name"`
	Month int    `json:"month"`
	Date  int    `json:"date"`
}

type GetMembersByBirthdayMonthResponse struct {
	Result  string           `json:"result"`
	Members []MemberBirthday `json:"members"`
}
