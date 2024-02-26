package timeslice

import (
	"errors"
	"fmt"
	"time"
)

type(
	Item[T any] struct {
		value   *T
		startAt time.Time
		endAt   time.Time
	}
)

var (
	ErrItemStartAtEmpty = errors.New("Item: startAt cannot be empty")
	ErrItemEndAtEmpty   = errors.New("Item: endAt cannot be empty")
	ErrItemInvalid      = errors.New("Item: invalid")
)

func NewItem[T any](value *T, startAt, endAt time.Time) (*Item[T], error) {
	if startAt.IsZero() {
		return nil, ErrItemStartAtEmpty
	}
	if endAt.IsZero() {
		return nil, ErrItemEndAtEmpty
	}
	if startAt.Compare(endAt) > 0 {
		return nil, errors.Join(ErrItemInvalid, fmt.Errorf(" want startAt: %s <= endAt: %s", startAt.Format(time.RFC3339), endAt.Format(time.RFC3339)))
	}
	return &Item[T]{
		value:   value,
		startAt: startAt,
		endAt:   endAt,
	},nil
}

func (i Item[T]) Value() T {
	return *i.value
}

func (i Item[T]) StartAt() time.Time {
	return i.startAt
}

func (i Item[T]) EndAt() time.Time {
	return i.endAt
}

func (i Item[T]) Contains(t time.Time) bool {
	return i.StartAt().Compare(t) <= 0 && i.EndAt().Compare(t) >= 0
}
