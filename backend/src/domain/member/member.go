package member

import "fmt"

var ErrInvalidBirthday = fmt.Errorf("invalid birthday")

type Birthday struct {
	month int
	date  int
}

func NewBirthday(month, date int) (Birthday, error) {
	if month < 1 || month > 12 {
		return Birthday{}, fmt.Errorf("%w: month=%d", ErrInvalidBirthday, month)
	}
	if date < 1 || date > 31 {
		return Birthday{}, fmt.Errorf("%w: date=%d", ErrInvalidBirthday, date)
	}
	return Birthday{month: month, date: date}, nil
}

func (b Birthday) Month() int { return b.month }
func (b Birthday) Date() int  { return b.date }
