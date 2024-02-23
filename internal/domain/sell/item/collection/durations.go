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

func NewDurations[T any](durations ...value.Duration[T]) (*Durations[T], error) {
	// startAtで昇順ソート
	sort.Slice(durations, func(i, j int) bool {
		return durations[i].StartAt().Before(durations[j].StartAt())
	})

	// 追加する期間が重複している場合は、追加する期間を優先して既存の期間を調整する
	buffer := make([]value.Duration[T], len(durations))
	for _, v := range durations {
		adjusting, err := Adjust(v, v)
		if err != nil {
			return nil, err
		}
		buffer = append(buffer, adjusting...)
	}

	// startAtで昇順ソートする
	sort.Slice(buffer, func(i, j int) bool {
		return buffer[i].StartAt().Before(buffer[j].StartAt())
	})

	durations2 := &Durations[T]{durations: buffer}
	return durations2, nil
}

func (d Durations[T]) Durations() []value.Duration[T] {
	return d.durations
}

// 重複している場合は、追加する期間を優先して既存の期間を調整する
func Adjust[T any](target, point value.Duration[T]) ([]value.Duration[T], error) {
	// 追加期間と重複せず、前方に位置するため、そのまま返す
	if point.StartAt().Compare(target.EndAt()) > 0 {
		return []value.Duration[T]{target}, nil
	}

	// 追加期間と重複せず、後方に位置するため、そのまま返す
	if point.EndAt().Compare(target.StartAt()) < 0 {
		return []value.Duration[T]{target}, nil
	}

	// 追加期間が包含するため、空で返す
	if point.StartAt().Compare(target.StartAt()) <= 0 && point.EndAt().Compare(target.EndAt()) >= 0 {
		return []value.Duration[T]{}, nil
	}

	// 追加期間を包含するため、追加期間を中心に前方と後方に分割して返す
	if point.StartAt().Compare(target.StartAt()) > 0 && point.EndAt().Compare(target.EndAt()) < 0 {

		foward, err := value.NewDuration(target.Value(), target.StartAt(), point.StartAt().Add(-1*time.Second))
		if err != nil {
			return nil, errors.Join(ErrDurationsUnexpection, err)
		}

		backward, err := value.NewDuration[T](target.Value(), point.EndAt().Add(1*time.Second), target.EndAt())
		if err != nil {
			return nil, errors.Join(ErrDurationsUnexpection, err)
		}

		return []value.Duration[T]{foward, backward}, nil
	}

	// 追加期間の前方と重複するため、既存の終了日時を調整して返す
	if point.StartAt().Compare(target.EndAt()) < 0 && point.EndAt().Compare(target.EndAt()) > 0 {
		foward, err := value.NewDuration[T](target.Value(), target.StartAt(), point.StartAt().Add(-1*time.Second))
		if err != nil {
			return nil, errors.Join(ErrDurationsUnexpection, err)
		}
		return []value.Duration[T]{foward}, nil
	}

	// 追加期間の後方と重複するため、既存の開始日時を調整して返す
	if point.StartAt().Compare(target.StartAt()) < 0 && point.EndAt().Compare(target.EndAt()) < 0 {
		backward, err := value.NewDuration[T](target.Value(), point.EndAt().Add(1*time.Second), target.EndAt())
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
