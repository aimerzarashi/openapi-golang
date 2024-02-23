package value

import (
	"errors"
	"fmt"
	"time"
)

type (
	Duration[T any] struct {
		startAt time.Time
		endAt   time.Time
		value   T
	}
)

var (
	ErrDurationStartAtEmpty = errors.New("Duration: startAt cannot be empty")
	ErrDurationEndAtEmpty   = errors.New("Duration: endAt cannot be empty")
	ErrDurationInvalid      = errors.New("Duration: invalid")
)

func NewDuration[T any](value T, startAt, endAt time.Time) (Duration[T], error) {
	if startAt.IsZero() {
		return Duration[T]{}, ErrDurationStartAtEmpty
	}
	if endAt.IsZero() {
		return Duration[T]{}, ErrDurationEndAtEmpty
	}
	if startAt.Compare(endAt) > 0 {
		return Duration[T]{}, errors.Join(ErrDurationInvalid, fmt.Errorf(" want startAt: %s <= endAt: %s", startAt.Format(time.RFC3339), endAt.Format(time.RFC3339)))
	}
	return Duration[T]{
		startAt: startAt,
		endAt:   endAt,
		value:   value,
	}, nil
}

func (v Duration[T]) StartAt() time.Time {
	return v.startAt
}

func (v Duration[T]) EndAt() time.Time {
	return v.endAt
}

func (v Duration[T]) Value() T {
	return v.value
}
