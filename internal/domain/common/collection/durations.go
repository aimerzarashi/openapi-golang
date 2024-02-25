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
	ErrDurationNotFound     = errors.New("Durations: not found")
)

func NewDurations[T any](durations ...value.Duration[T]) (*Durations[T], error) {
	// startAtで昇順ソート
	sort.Slice(durations, func(i, j int) bool {
		return durations[i].StartAt().Before(durations[j].StartAt())
	})

	// 期間が重複していないか確認
	for i := 0; i < len(durations)-1; i++ {
		if durations[i].EndAt().Compare(durations[i+1].StartAt()) >= 0 {
			return nil, ErrDurationsInvalid
		}
	}

	return &Durations[T]{durations: durations}, nil
}

func (d Durations[T]) Durations() []value.Duration[T] {
	return d.durations
}

// 既存期間と追加期間が重複している場合は、追加期間を優先して既存期間を調整する
func Adjust[T any](existing, adding value.Duration[T]) ([]value.Duration[T], error) {
	// 追加期間に対し、既存期間は重複しない前方に位置するため、そのまま返す
	if adding.StartAt().Compare(existing.EndAt()) > 0 {
		return []value.Duration[T]{existing}, nil
	}

	// 追加期間に対し、既存期間は重複しない後方に位置するため、そのまま返す
	if adding.EndAt().Compare(existing.StartAt()) < 0 {
		return []value.Duration[T]{existing}, nil
	}

	// 追加期間が既存期間を包含するため、空で返す
	if adding.StartAt().Compare(existing.StartAt()) <= 0 && adding.EndAt().Compare(existing.EndAt()) >= 0 {
		return []value.Duration[T]{}, nil
	}

	// 追加期間が既存期間に包含されるため、追加期間を優先して前方と後方に分割して返す
	if adding.StartAt().Compare(existing.StartAt()) > 0 && adding.EndAt().Compare(existing.EndAt()) < 0 {

		// 分割した既存期間の前方は、開始日時を調整して返す
		foward, err := value.NewDuration(existing.Value(), existing.StartAt(), adding.StartAt().Add(-1*time.Second))
		if err != nil {
			return nil, errors.Join(ErrDurationsUnexpection, err)
		}

		// 分割した既存期間の後方は、終了日時を調整して返す
		backward, err := value.NewDuration[T](existing.Value(), adding.EndAt().Add(1*time.Second), existing.EndAt())
		if err != nil {
			return nil, errors.Join(ErrDurationsUnexpection, err)
		}

		return []value.Duration[T]{foward, backward}, nil
	}

	// 追加期間に対し、既存期間の終了日時が重複するため、既存期間の終了日時を調整して返す
	if adding.StartAt().Compare(existing.EndAt()) < 0 && adding.EndAt().Compare(existing.EndAt()) > 0 {
		foward, err := value.NewDuration[T](existing.Value(), existing.StartAt(), adding.StartAt().Add(-1*time.Second))
		if err != nil {
			return nil, errors.Join(ErrDurationsUnexpection, err)
		}
		return []value.Duration[T]{foward}, nil
	}

	// 追加期間に対し、既存期間の開始日時が重複するため、既存期間の開始日時を調整して返す
	if adding.StartAt().Compare(existing.StartAt()) < 0 && adding.EndAt().Compare(existing.EndAt()) < 0 {
		backward, err := value.NewDuration[T](existing.Value(), adding.EndAt().Add(1*time.Second), existing.EndAt())
		if err != nil {
			return nil, errors.Join(ErrDurationsUnexpection, err)
		}
		return []value.Duration[T]{backward}, nil
	}

	return nil, ErrDurationsUnexpection
}

func (d *Durations[T]) Merge(adding value.Duration[T]) error {
	buffer := make([]value.Duration[T], 0)
	buffer = append(buffer, adding)

	// 追加する期間が重複している場合は、追加する期間を優先して既存の期間を調整する
	for _, v := range d.durations {
		adjusting, err := Adjust(v, adding)
		if err != nil {
			return err
		}
		buffer = append(buffer, adjusting...)
	}

	// startAtで昇順ソートする
	sort.Slice(buffer, func(i, j int) bool {
		return buffer[i].StartAt().Before(buffer[j].StartAt())
	})

	// 調整済みの期間に置き換える
	d.durations = buffer
	return nil
}

func (d *Durations[T]) Find(criteria time.Time) (value.Duration[T], error) {
	for _, v := range d.durations {
		if v.Contains(criteria) {
			return v, nil
		}
	}

	return value.Duration[T]{}, ErrDurationNotFound
}
