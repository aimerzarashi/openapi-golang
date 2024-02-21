package value

import (
	"errors"
	"fmt"
	"time"
)

type (
	Duration struct {
		startAt time.Time
		endAt   time.Time
	}
)

var (
	ErrDurationStartAtEmpty = errors.New("Duration: startAt cannot be empty")
	ErrDurationEndAtEmpty   = errors.New("Duration: endAt cannot be empty")
	ErrDurationInvalid = errors.New("Duration: invalid")
)

func NewDuration (startAt, endAt time.Time) (Duration, error) {
	if startAt.IsZero() {
		return Duration{}, ErrDurationStartAtEmpty
	}
	if endAt.IsZero() {
		return Duration{}, ErrDurationEndAtEmpty
	}
	if startAt.Compare(endAt) > 0 {
		return Duration{}, fmt.Errorf("NewDuration: startAt (%v) cannot be after endAt (%v)", startAt.Format(time.RFC3339), endAt.Format(time.RFC3339))
	}
	return Duration{
		startAt: startAt,
		endAt:   endAt,
	}, nil
}

func (v Duration) StartAt() time.Time {
	return v.startAt
}

func (v Duration) EndAt() time.Time {
	return v.endAt
}