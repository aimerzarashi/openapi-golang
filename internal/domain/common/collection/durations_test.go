package collection_test

import (
	"errors"
	"openapi/internal/domain/common/collection"
	"openapi/internal/domain/common/value"
	"reflect"
	"testing"
	"time"
)

func TestNewDurations(t *testing.T) {
	// Setup
	t.Parallel()

	currentAt := time.Now()

	var validDurations []value.Duration[string]
	for i := 0; i < 3; i++ {
		startAt := currentAt.Add(time.Duration(i) * time.Hour)
		endAt := currentAt.Add(time.Duration(i+1)*time.Hour - 1*time.Second)
		duration, err := value.NewDuration("a", startAt, endAt)
		if err != nil {
			t.Fatal(err)
		}
		validDurations = append(validDurations, duration)
		//		fmt.Printf("got: %v %v\n", duration.StartAt().Format(time.RFC3339), duration.EndAt().Format(time.RFC3339))
	}

	var invalidDurations []value.Duration[string]
	for i := 0; i < 3; i++ {
		startAt := currentAt.Add(time.Duration(i) * time.Hour)
		endAt := currentAt.Add(time.Duration(i+1)*time.Hour + 1*time.Second)
		duration, err := value.NewDuration("a", startAt, endAt)
		if err != nil {
			t.Fatal(err)
		}
		invalidDurations = append(invalidDurations, duration)
		//		fmt.Printf("got: %v %v\n", duration.StartAt().Format(time.RFC3339), duration.EndAt().Format(time.RFC3339))
	}

	type args struct {
		durations []value.Duration[string]
	}
	type want struct {
		durations []value.Duration[string]
	}
	tests := []struct {
		name    string
		args    args
		want    want
		wantErr bool
		errType error
	}{
		{
			name: "success",
			args: args{
				durations: validDurations,
			},
			want: want{
				durations: validDurations,
			},
			wantErr: false,
			errType: nil,
		},
		{
			name: "fail/invalid",
			args: args{
				durations: invalidDurations,
			},
			want:    want{},
			wantErr: true,
			errType: collection.ErrDurationsInvalid,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			// When
			got, err := collection.NewDurations(tt.args.durations)

			// Then
			if !tt.wantErr {
				if err != nil {
					t.Errorf("NewDurations() error = %v, wantErr %v", err, tt.wantErr)
				}
				if !reflect.DeepEqual(got.Durations(), tt.want.durations) {
					t.Errorf("NewDurations() = %v, want %v", got.Durations(), tt.want.durations)
				}
				return
			}

			if !errors.Is(err, tt.errType) {
				t.Errorf("NewDurations() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func NewDuration[T any](arg T, startAt, endAt time.Time) value.Duration[T] {
	duration, err := value.NewDuration(arg, startAt, endAt)
	if err != nil {
		panic(err)
	}
	return duration
}

func NewDurations[T any](arg T, startAt, endAt time.Time) ([]value.Duration[T], error) {
	duration, err := value.NewDuration(arg, startAt, endAt)
	if err != nil {
		return nil, err
	}
	var durations []value.Duration[T]
	durations = append(durations, duration)
	return durations, nil
}

func TestDurations_Merge(t *testing.T) {
	// Setup
	t.Parallel()

	type args struct {
		adding   value.Duration[string]
		existing value.Duration[string]
	}
	type want struct {
		durations []value.Duration[string]
	}
	tests := []struct {
		name    string
		args    args
		want    want
		wantErr bool
		errType error
	}{
		{
			name: "1",
			args: args{
				adding:   NewDuration("a", time.Date(2024, 1, 1, 9, 0, 0, 0, time.UTC), time.Date(2024, 1, 1, 9, 59, 59, 0, time.UTC)),
				existing: NewDuration("b", time.Date(2024, 1, 1, 8, 0, 0, 0, time.UTC), time.Date(2024, 1, 1, 8, 59, 59, 0, time.UTC)),
			},
			want: want{
				durations: []value.Duration[string]{
					NewDuration("b", time.Date(2024, 1, 1, 8, 0, 0, 0, time.UTC), time.Date(2024, 1, 1, 8, 59, 59, 0, time.UTC)),
					NewDuration("a", time.Date(2024, 1, 1, 9, 0, 0, 0, time.UTC), time.Date(2024, 1, 1, 9, 59, 59, 0, time.UTC)),
				},
			},
			wantErr: false,
			errType: nil,
		},
		{
			name: "2",
			args: args{
				adding:   NewDuration("a", time.Date(2024, 1, 1, 9, 0, 0, 0, time.UTC), time.Date(2024, 1, 1, 9, 59, 59, 0, time.UTC)),
				existing: NewDuration("b", time.Date(2024, 1, 1, 10, 0, 0, 0, time.UTC), time.Date(2024, 1, 1, 10, 59, 59, 0, time.UTC)),
			},
			want: want{
				durations: []value.Duration[string]{
					NewDuration("a", time.Date(2024, 1, 1, 9, 0, 0, 0, time.UTC), time.Date(2024, 1, 1, 9, 59, 59, 0, time.UTC)),
					NewDuration("b", time.Date(2024, 1, 1, 10, 0, 0, 0, time.UTC), time.Date(2024, 1, 1, 10, 59, 59, 0, time.UTC)),
				},
			},
			wantErr: false,
			errType: nil,
		},
		{
			name: "3",
			args: args{
				adding:   NewDuration("a", time.Date(2024, 1, 1, 9, 0, 0, 0, time.UTC), time.Date(2024, 1, 1, 9, 59, 59, 0, time.UTC)),
				existing: NewDuration("b", time.Date(2024, 1, 1, 9, 0, 0, 0, time.UTC), time.Date(2024, 1, 1, 9, 59, 59, 0, time.UTC)),
			},
			want: want{
				durations: []value.Duration[string]{
					NewDuration("a", time.Date(2024, 1, 1, 9, 0, 0, 0, time.UTC), time.Date(2024, 1, 1, 9, 59, 59, 0, time.UTC)),
				},
			},
			wantErr: false,
			errType: nil,
		},
		{
			name: "4",
			args: args{
				adding:   NewDuration("a", time.Date(2024, 1, 1, 9, 0, 0, 0, time.UTC), time.Date(2024, 1, 1, 9, 59, 59, 0, time.UTC)),
				existing: NewDuration("b", time.Date(2024, 1, 1, 8, 0, 0, 0, time.UTC), time.Date(2024, 1, 1, 10, 59, 59, 0, time.UTC)),
			},
			want: want{
				durations: []value.Duration[string]{
					NewDuration("b", time.Date(2024, 1, 1, 8, 0, 0, 0, time.UTC), time.Date(2024, 1, 1, 8, 59, 59, 0, time.UTC)),
					NewDuration("a", time.Date(2024, 1, 1, 9, 0, 0, 0, time.UTC), time.Date(2024, 1, 1, 9, 59, 59, 0, time.UTC)),
					NewDuration("b", time.Date(2024, 1, 1, 10, 0, 0, 0, time.UTC), time.Date(2024, 1, 1, 10, 59, 59, 0, time.UTC)),
				},
			},
			wantErr: false,
			errType: nil,
		},
		{
			name: "5",
			args: args{
				adding:   NewDuration("a", time.Date(2024, 1, 1, 9, 0, 0, 0, time.UTC), time.Date(2024, 1, 1, 9, 59, 59, 0, time.UTC)),
				existing: NewDuration("b", time.Date(2024, 1, 1, 8, 0, 0, 0, time.UTC), time.Date(2024, 1, 1, 9, 29, 59, 0, time.UTC)),
			},
			want: want{
				durations: []value.Duration[string]{
					NewDuration("b", time.Date(2024, 1, 1, 8, 0, 0, 0, time.UTC), time.Date(2024, 1, 1, 8, 59, 59, 0, time.UTC)),
					NewDuration("a", time.Date(2024, 1, 1, 9, 0, 0, 0, time.UTC), time.Date(2024, 1, 1, 9, 59, 59, 0, time.UTC)),
				},
			},
			wantErr: false,
			errType: nil,
		},
		{
			name: "6",
			args: args{
				adding:   NewDuration("a", time.Date(2024, 1, 1, 9, 0, 0, 0, time.UTC), time.Date(2024, 1, 1, 9, 59, 59, 0, time.UTC)),
				existing: NewDuration("b", time.Date(2024, 1, 1, 9, 30, 0, 0, time.UTC), time.Date(2024, 1, 1, 10, 59, 59, 0, time.UTC)),
			},
			want: want{
				durations: []value.Duration[string]{
					NewDuration("a", time.Date(2024, 1, 1, 9, 0, 0, 0, time.UTC), time.Date(2024, 1, 1, 9, 59, 59, 0, time.UTC)),
					NewDuration("b", time.Date(2024, 1, 1, 10, 00, 0, 0, time.UTC), time.Date(2024, 1, 1, 10, 59, 59, 0, time.UTC)),
				},
			},
			wantErr: false,
			errType: nil,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			// Given
			durations := []value.Duration[string]{tt.args.existing}
			got, err := collection.NewDurations(durations)
			if err != nil {
				t.Fatal(err)
			}

			// When
			err = got.Merge(tt.args.adding)

			// Then
			if !tt.wantErr {
				if err != nil {
					t.Errorf("NewDurations() error = %v, wantErr %v", err, tt.wantErr)
				}
				for i, v := range got.Durations() {
					if !v.StartAt().Equal(tt.want.durations[i].StartAt()) || !v.EndAt().Equal(tt.want.durations[i].EndAt()) {
						t.Errorf("NewDurations() value = %v, want %v", v, tt.want.durations[i])
					}
					if v.Value() != tt.want.durations[i].Value() {
						t.Errorf("NewDurations() value = %v, want %v", v.Value(), tt.want.durations[i].Value())
					}
				}
				return
			}

			if !errors.Is(err, tt.errType) {
				t.Errorf("NewDurations() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
