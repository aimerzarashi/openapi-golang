package collection_test

import (
	"errors"
	"openapi/internal/domain/sell/item/collection"
	"openapi/internal/domain/sell/item/value"
	"reflect"
	"testing"
	"time"
)

func TestNewDurations(t *testing.T) {
	// Setup
	t.Parallel()

	currentAt := time.Now()

	var validDurations []value.Duration
	for i := 0; i < 3; i++ {
		startAt := currentAt.Add(time.Duration(i)*time.Hour)
		endAt := currentAt.Add(time.Duration(i+1)*time.Hour-1*time.Second)
		duration, err := value.NewDuration(startAt, endAt)
		if err != nil {
			t.Fatal(err)
		}
		validDurations = append(validDurations, duration)
//		fmt.Printf("got: %v %v\n", duration.StartAt().Format(time.RFC3339), duration.EndAt().Format(time.RFC3339))
	}

	var invalidDurations []value.Duration
	for i := 0; i < 3; i++ {
		startAt := currentAt.Add(time.Duration(i)*time.Hour)
		endAt := currentAt.Add(time.Duration(i+1)*time.Hour+1*time.Second)
		duration, err := value.NewDuration(startAt, endAt)
		if err != nil {
			t.Fatal(err)
		}
		invalidDurations = append(invalidDurations, duration)
//		fmt.Printf("got: %v %v\n", duration.StartAt().Format(time.RFC3339), duration.EndAt().Format(time.RFC3339))
	}

	type args struct {
		durations []value.Duration
	}
	type want struct {
		durations []value.Duration
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

func NewDuration(startAt, endAt time.Time) value.Duration {
	duration, err := value.NewDuration(startAt, endAt)
	if err != nil {
		panic(err)
	}
	return duration
}

func NewDurations(startAt, endAt time.Time) ([]value.Duration) {
	duration, err := value.NewDuration(startAt, endAt)
	if err != nil {
		panic(err)
	}
	var durations []value.Duration
	durations = append(durations, duration)
	return durations
}

func TestDurations_Merge(t *testing.T) {
	// Setup
	t.Parallel()

	type args struct {
		adding value.Duration
		existing value.Duration
	}
	type want struct {
		durations []value.Duration
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
				adding: NewDuration(time.Date(2024, 1, 1, 9, 0, 0, 0, time.UTC), time.Date(2024, 1, 1, 9, 59, 59, 0, time.UTC)),
				existing: NewDuration(time.Date(2024, 1, 1, 8, 0, 0, 0, time.UTC), time.Date(2024, 1, 1, 8, 59, 59, 0, time.UTC)),
			},
			want: want{
				durations: []value.Duration{
					NewDuration(time.Date(2024, 1, 1, 8, 0, 0, 0, time.UTC), time.Date(2024, 1, 1, 8, 59, 59, 0, time.UTC)),
					NewDuration(time.Date(2024, 1, 1, 9, 0, 0, 0, time.UTC), time.Date(2024, 1, 1, 9, 59, 59, 0, time.UTC)),
				},
			},
			wantErr: false,
			errType: nil,
		},
		{
			name: "2",
			args: args{
				adding: NewDuration(time.Date(2024, 1, 1, 9, 0, 0, 0, time.UTC), time.Date(2024, 1, 1, 9, 59, 59, 0, time.UTC)),
				existing: NewDuration(time.Date(2024, 1, 1, 10, 0, 0, 0, time.UTC), time.Date(2024, 1, 1, 10, 59, 59, 0, time.UTC)),
			},
			want: want{
				durations: []value.Duration{
					NewDuration(time.Date(2024, 1, 1, 9, 0, 0, 0, time.UTC), time.Date(2024, 1, 1, 9, 59, 59, 0, time.UTC)),
					NewDuration(time.Date(2024, 1, 1, 10, 0, 0, 0, time.UTC), time.Date(2024, 1, 1, 10, 59, 59, 0, time.UTC)),
				},
			},
			wantErr: false,
			errType: nil,
		},
		{
			name: "3",
			args: args{
				adding: NewDuration(time.Date(2024, 1, 1, 9, 0, 0, 0, time.UTC), time.Date(2024, 1, 1, 9, 59, 59, 0, time.UTC)),
				existing: NewDuration(time.Date(2024, 1, 1, 9, 0, 0, 0, time.UTC), time.Date(2024, 1, 1, 9, 59, 59, 0, time.UTC)),
			},
			want: want{
				durations: []value.Duration{
					NewDuration(time.Date(2024, 1, 1, 9, 0, 0, 0, time.UTC), time.Date(2024, 1, 1, 9, 59, 59, 0, time.UTC)),
				},
			},
			wantErr: false,
			errType: nil,
		},
		{
			name: "4",
			args: args{
				adding: NewDuration(time.Date(2024, 1, 1, 9, 0, 0, 0, time.UTC), time.Date(2024, 1, 1, 9, 59, 59, 0, time.UTC)),
				existing: NewDuration(time.Date(2024, 1, 1, 8, 0, 0, 0, time.UTC), time.Date(2024, 1, 1, 10, 59, 59, 0, time.UTC)),
			},
			want: want{
				durations: []value.Duration{
					NewDuration(time.Date(2024, 1, 1, 8, 0, 0, 0, time.UTC), time.Date(2024, 1, 1, 8, 59, 59, 0, time.UTC)),
					NewDuration(time.Date(2024, 1, 1, 9, 0, 0, 0, time.UTC), time.Date(2024, 1, 1, 9, 59, 59, 0, time.UTC)),
					NewDuration(time.Date(2024, 1, 1, 10, 0, 0, 0, time.UTC), time.Date(2024, 1, 1, 10, 59, 59, 0, time.UTC)),
				},
			},
			wantErr: false,
			errType: nil,
		},
		{
			name: "5",
			args: args{
				adding: NewDuration(time.Date(2024, 1, 1, 9, 0, 0, 0, time.UTC), time.Date(2024, 1, 1, 9, 59, 59, 0, time.UTC)),
				existing: NewDuration(time.Date(2024, 1, 1, 8, 0, 0, 0, time.UTC), time.Date(2024, 1, 1, 9, 29, 59, 0, time.UTC)),
			},
			want: want{
				durations: []value.Duration{
					NewDuration(time.Date(2024, 1, 1, 8, 0, 0, 0, time.UTC), time.Date(2024, 1, 1, 8, 59, 59, 0, time.UTC)),
					NewDuration(time.Date(2024, 1, 1, 9, 0, 0, 0, time.UTC), time.Date(2024, 1, 1, 9, 59, 59, 0, time.UTC)),
				},
			},
			wantErr: false,
			errType: nil,
		},
		{
			name: "6",
			args: args{
				adding: NewDuration(time.Date(2024, 1, 1, 9, 0, 0, 0, time.UTC), time.Date(2024, 1, 1, 9, 59, 59, 0, time.UTC)),
				existing: NewDuration(time.Date(2024, 1, 1, 9, 30, 0, 0, time.UTC), time.Date(2024, 1, 1, 10, 59, 59, 0, time.UTC)),
			},
			want: want{
				durations: []value.Duration{
					NewDuration(time.Date(2024, 1, 1, 9, 0, 0, 0, time.UTC), time.Date(2024, 1, 1, 9, 59, 59, 0, time.UTC)),
					NewDuration(time.Date(2024, 1, 1, 10, 00, 0, 0, time.UTC), time.Date(2024, 1, 1, 10, 59, 59, 0, time.UTC)),
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
			durations := []value.Duration{tt.args.existing}
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
				}
				return
			}

			if !errors.Is(err, tt.errType) {
				t.Errorf("NewDurations() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
