package dto

type GetMembersByBirthdayMonthRequest struct {
	BirthdayMonth int `form:"birthday_month" binding:"required,numeric,min=1,max=12" example:"5" description:"誕生日の月（1～12）"`
}
