package timeslice

import (
	"errors"
	"sort"
	"time"
)

type(
	Collection[T any] struct {
		items []*Item[T]
	}
)

var (
	ErrCollectionInvalid     = errors.New("Collection: invalid")
	ErrCollectionNotFound    = errors.New("Collection: not found")
	ErrCollectionUnexpection = errors.New("Collection: unexpection")
)

func NewCollection[T any](timeslices ...*Item[T]) (*Collection[T], error) {
	// startAtで昇順ソート
	sort.Slice(timeslices, func(i, j int) bool {
		return timeslices[i].StartAt().Before(timeslices[j].StartAt())
	})
	
	// 期間が重複していないか確認
	for i := 0; i < len(timeslices)-1; i++ {
		if timeslices[i].EndAt().Compare(timeslices[i+1].StartAt()) >= 0 {
			return nil, ErrCollectionInvalid
		}
	}

	var items []*Item[T]
	if len(timeslices) == 0 {
		items = make([]*Item[T], len(timeslices))
	} else {
		items = timeslices
	}

	return &Collection[T]{
		items: items,
	}, nil
}

func (c Collection[T]) Items() []*Item[T] {
	return c.items
}

func (d Collection[T]) Find(criteria time.Time) (Item[T], error) {
	for _, v := range d.items {
		if v.Contains(criteria) {
			return *v, nil
		}
	}

	return Item[T]{}, ErrCollectionNotFound
}