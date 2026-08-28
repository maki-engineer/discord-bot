package member

import "fmt"

var ErrInvalidBirthday = fmt.Errorf("invalid birthday")

const (
	minMonth = 1
	maxMonth = 12
)

type Month int

func NewMonth(value int) (Month, error) {
	if value < minMonth || value > maxMonth {
		return 0, fmt.Errorf("invalid argument: month must be between %d and %d (got %d)", minMonth, maxMonth, value)
	}

	return Month(value), nil
}

func (m Month) Int() int {
	return int(m)
}
