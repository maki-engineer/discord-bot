package model

type Member struct {
	Name      string `gorm:"not null"`
	UserID    string `gorm:"not null;default:''"`
	Month     int    `gorm:"not null"`
	Date      int    `gorm:"not null"`
	SpeakerID int    `gorm:"not null;default:62"`
}

func (Member) TableName() string {
	return "birthday_for_235_members"
}
