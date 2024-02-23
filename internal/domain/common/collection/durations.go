package collection

import (
	"errors"
	"openapi/internal/domain/common/value"
	"sort"
	"time"
)

type (
	Durations[T any] struct {
		durations []value.Duration[T]
	}
)

var (
	ErrDurationsInvalid     = errors.New("Durations: invalid")
	ErrDurationsUnexpection = errors.New("Durations: unexpection")
)

func NewDurations[T any](durations []value.Duration[T]) (*Durations[T], error) {
	// startAtで昇順ソート
	sort.Slice(durations, func(i, j int) bool {
		return durations[i].StartAt().Before(durations[j].StartAt())
	})

	// 該当の終了時刻と次の開始時刻が重複していないか確認
	for i := 0; i < len(durations)-1; i++ {
		if durations[i].EndAt().Compare(durations[i+1].StartAt()) > 0 {
			return &Durations[T]{}, ErrDurationsInvalid
		}
	}

	return &Durations[T]{
		durations: durations,
	}, nil
}

func (d Durations[T]) Durations() []value.Duration[T] {
	return d.durations
}

func (d *Durations[T]) Merge(a value.Duration[T]) error {
	buffer := make([]value.Duration[T], 0)
	buffer = append(buffer, a)

	// 追加する期間が重複している場合は、追加する期間を優先して既存の期間を調整する
	for _, v := range d.durations {
		if a.StartAt().Compare(v.EndAt()) > 0 {
			buffer = append(buffer, v)
			continue
		}
		if a.EndAt().Compare(v.StartAt()) < 0 {
			buffer = append(buffer, v)
			continue
		}
		if a.StartAt().Compare(v.StartAt()) <= 0 && a.EndAt().Compare(v.EndAt()) >= 0 {
			continue
		}
		if a.StartAt().Compare(v.StartAt()) > 0 && a.EndAt().Compare(v.EndAt()) < 0 {
			b, err := value.NewDuration(v.Value(), v.StartAt(), a.StartAt().Add(-1*time.Second))
			if err != nil {
				return errors.Join(ErrDurationsUnexpection, err)
			}
			buffer = append(buffer, b)
			c, err := value.NewDuration[T](v.Value(), a.EndAt().Add(1*time.Second), v.EndAt())
			if err != nil {
				return errors.Join(ErrDurationsUnexpection, err)
			}
			buffer = append(buffer, c)
			continue
		}
		if a.StartAt().Compare(v.EndAt()) < 0 && a.EndAt().Compare(v.EndAt()) > 0 {
			b, err := value.NewDuration[T](v.Value(), v.StartAt(), a.StartAt().Add(-1*time.Second))
			if err != nil {
				return errors.Join(ErrDurationsUnexpection, err)
			}
			buffer = append(buffer, b)
			continue
		}
		if a.StartAt().Compare(v.StartAt()) < 0 && a.EndAt().Compare(v.EndAt()) < 0 {
			b, err := value.NewDuration[T](v.Value(), a.EndAt().Add(1*time.Second), v.EndAt())
			if err != nil {
				return errors.Join(ErrDurationsUnexpection, err)
			}
			buffer = append(buffer, b)
			continue
		}
		return ErrDurationsUnexpection
	}

	// startAtで昇順ソートする
	sort.Slice(buffer, func(i, j int) bool {
		return buffer[i].StartAt().Before(buffer[j].StartAt())
	})

	// 調整済みの期間に置き換える
	d.durations = buffer
	return nil
}
