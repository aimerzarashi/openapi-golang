package collection_test

import (
	"errors"
	"openapi/internal/domain/common/collection"
	"openapi/internal/domain/common/value"
	"reflect"
	"testing"
	"time"
)

func NewDuration[T any](arg *T, startAt, endAt time.Time) value.Duration[T] {
	duration, err := value.NewDuration(arg, startAt, endAt)
	if err != nil {
		panic(err)
	}
	return duration
}

func TestAdjust(t *testing.T) {
	// Setup
	t.Parallel()

	existing := "existing"
	adding := "adding"

	type args struct {
		existing value.Duration[string]
		adding   value.Duration[string]
	}
	type want struct {
		durations []value.Duration[string]
		err       error
	}
	tests := []struct {
		name    string
		args    args
		want    want
		wantErr bool
	}{
		{
			name: "success/1",
			args: args{
				existing: NewDuration(&existing, time.Date(2024, 1, 1, 9, 0, 0, 0, time.UTC), time.Date(2024, 1, 1, 9, 59, 59, 0, time.UTC)),
				adding:   NewDuration(&adding, time.Date(2024, 1, 1, 8, 0, 0, 0, time.UTC), time.Date(2024, 1, 1, 8, 59, 59, 0, time.UTC)),
			},
			want: want{
				durations: []value.Duration[string]{
					NewDuration(&existing, time.Date(2024, 1, 1, 9, 0, 0, 0, time.UTC), time.Date(2024, 1, 1, 9, 59, 59, 0, time.UTC)),
				},
				err: nil,
			},
			wantErr: false,
		},
		{
			name: "success/2",
			args: args{
				existing: NewDuration(&existing, time.Date(2024, 1, 1, 9, 0, 0, 0, time.UTC), time.Date(2024, 1, 1, 9, 59, 59, 0, time.UTC)),
				adding:   NewDuration(&adding, time.Date(2024, 1, 1, 10, 0, 0, 0, time.UTC), time.Date(2024, 1, 1, 10, 59, 59, 0, time.UTC)),
			},
			want: want{
				durations: []value.Duration[string]{
					NewDuration(&existing, time.Date(2024, 1, 1, 9, 0, 0, 0, time.UTC), time.Date(2024, 1, 1, 9, 59, 59, 0, time.UTC)),
				},
				err: nil,
			},
			wantErr: false,
		},
		{
			name: "success/3",
			args: args{
				existing: NewDuration(&existing, time.Date(2024, 1, 1, 9, 0, 0, 0, time.UTC), time.Date(2024, 1, 1, 9, 59, 59, 0, time.UTC)),
				adding:   NewDuration(&adding, time.Date(2024, 1, 1, 9, 0, 0, 0, time.UTC), time.Date(2024, 1, 1, 9, 59, 59, 0, time.UTC)),
			},
			want: want{
				durations: []value.Duration[string]{},
				err:       nil,
			},
			wantErr: false,
		},
		{
			name: "success/4",
			args: args{
				existing: NewDuration(&existing, time.Date(2024, 1, 1, 8, 0, 0, 0, time.UTC), time.Date(2024, 1, 1, 10, 59, 59, 0, time.UTC)),
				adding:   NewDuration(&adding, time.Date(2024, 1, 1, 9, 0, 0, 0, time.UTC), time.Date(2024, 1, 1, 9, 59, 59, 0, time.UTC)),
			},
			want: want{
				durations: []value.Duration[string]{
					NewDuration(&existing, time.Date(2024, 1, 1, 8, 0, 0, 0, time.UTC), time.Date(2024, 1, 1, 8, 59, 59, 0, time.UTC)),
					NewDuration(&existing, time.Date(2024, 1, 1, 10, 0, 0, 0, time.UTC), time.Date(2024, 1, 1, 10, 59, 59, 0, time.UTC)),
				},
				err: nil,
			},
			wantErr: false,
		},
		{
			name: "success/5",
			args: args{
				existing: NewDuration(&existing, time.Date(2024, 1, 1, 8, 0, 0, 0, time.UTC), time.Date(2024, 1, 1, 9, 29, 59, 0, time.UTC)),
				adding:   NewDuration(&adding, time.Date(2024, 1, 1, 9, 0, 0, 0, time.UTC), time.Date(2024, 1, 1, 9, 59, 59, 0, time.UTC)),
			},
			want: want{
				durations: []value.Duration[string]{
					NewDuration(&existing, time.Date(2024, 1, 1, 8, 0, 0, 0, time.UTC), time.Date(2024, 1, 1, 8, 59, 59, 0, time.UTC)),
				},
				err: nil,
			},
			wantErr: false,
		},
		{
			name: "success/6",
			args: args{
				existing: NewDuration(&existing, time.Date(2024, 1, 1, 9, 30, 0, 0, time.UTC), time.Date(2024, 1, 1, 10, 59, 59, 0, time.UTC)),
				adding:   NewDuration(&adding, time.Date(2024, 1, 1, 9, 0, 0, 0, time.UTC), time.Date(2024, 1, 1, 9, 59, 59, 0, time.UTC)),
			},
			want: want{
				durations: []value.Duration[string]{
					NewDuration(&existing, time.Date(2024, 1, 1, 10, 00, 0, 0, time.UTC), time.Date(2024, 1, 1, 10, 59, 59, 0, time.UTC)),
				},
				err: nil,
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			// When
			got, err := collection.Adjust(tt.args.existing, tt.args.adding)

			// Then
			if !tt.wantErr {
				if err != nil {
					t.Errorf("Adjust() error = %v, wantErr %v", err, tt.wantErr)
				}

				if !reflect.DeepEqual(got, tt.want.durations) {
					t.Errorf("Adjust() = %v, want %v", got, tt.want.durations)
				}
				return
			}

			if err != nil {
				if !errors.Is(err, tt.want.err) {
					t.Errorf("Adjust() error = %v, wantErr %v", err, tt.want.err)
				}
				return
			}
		})
	}
}

func TestNewDurations(t *testing.T) {
	// Setup
	t.Parallel()

	adding := "adding"

	type args struct {
		adding []value.Duration[string]
	}
	type want struct {
		durations []value.Duration[string]
		err       error
	}
	tests := []struct {
		name    string
		args    args
		want    want
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				adding: []value.Duration[string]{
					NewDuration(&adding, time.Date(2024, 1, 1, 9, 0, 0, 0, time.UTC), time.Date(2024, 1, 1, 9, 59, 59, 0, time.UTC)),
					NewDuration(&adding, time.Date(2024, 1, 1, 10, 0, 0, 0, time.UTC), time.Date(2024, 1, 1, 10, 59, 59, 0, time.UTC)),
					NewDuration(&adding, time.Date(2024, 1, 1, 11, 0, 0, 0, time.UTC), time.Date(2024, 1, 1, 11, 59, 59, 0, time.UTC)),
				},
			},
			want: want{
				durations: []value.Duration[string]{
					NewDuration(&adding, time.Date(2024, 1, 1, 9, 0, 0, 0, time.UTC), time.Date(2024, 1, 1, 9, 59, 59, 0, time.UTC)),
					NewDuration(&adding, time.Date(2024, 1, 1, 10, 0, 0, 0, time.UTC), time.Date(2024, 1, 1, 10, 59, 59, 0, time.UTC)),
					NewDuration(&adding, time.Date(2024, 1, 1, 11, 0, 0, 0, time.UTC), time.Date(2024, 1, 1, 11, 59, 59, 0, time.UTC)),
				},
				err: nil,
			},
			wantErr: false,
		},
		{
			name: "fail",
			args: args{
				adding: []value.Duration[string]{
					NewDuration(&adding, time.Date(2024, 1, 1, 9, 0, 0, 0, time.UTC), time.Date(2024, 1, 1, 9, 59, 59, 0, time.UTC)),
					NewDuration(&adding, time.Date(2024, 1, 1, 9, 59, 59, 0, time.UTC), time.Date(2024, 1, 1, 10, 59, 59, 0, time.UTC)),
					NewDuration(&adding, time.Date(2024, 1, 1, 11, 0, 0, 0, time.UTC), time.Date(2024, 1, 1, 11, 59, 59, 0, time.UTC)),
				},
			},
			want: want{
				durations: nil,
				err:       collection.ErrDurationsInvalid,
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			// When
			got, err := collection.NewDurations(tt.args.adding...)

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
			if !errors.Is(err, tt.want.err) {
				t.Errorf("NewDurations() error = %v, wantErr %v", err, tt.want.err)
			}
		})
	}
}

func TestDurations_Merge(t *testing.T) {
	t.Parallel()

	adding := "adding"
	type args struct {
		existing []value.Duration[string]
		adding   []value.Duration[string]
	}
	type want struct {
		durations []value.Duration[string]
		err       error
	}
	tests := []struct {
		name    string
		args    args
		want    want
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				existing: []value.Duration[string]{
					NewDuration(&adding, time.Date(2024, 1, 1, 9, 0, 0, 0, time.UTC), time.Date(2024, 1, 1, 9, 59, 59, 0, time.UTC)),
					NewDuration(&adding, time.Date(2024, 1, 1, 10, 0, 0, 0, time.UTC), time.Date(2024, 1, 1, 10, 59, 59, 0, time.UTC)),
					NewDuration(&adding, time.Date(2024, 1, 1, 11, 0, 0, 0, time.UTC), time.Date(2024, 1, 1, 11, 59, 59, 0, time.UTC)),
				},
				adding: []value.Duration[string]{
					NewDuration(&adding, time.Date(2024, 1, 1, 9, 20, 0, 0, time.UTC), time.Date(2024, 1, 1, 9, 39, 59, 0, time.UTC)),
				},
			},
			want: want{
				durations: []value.Duration[string]{
					NewDuration(&adding, time.Date(2024, 1, 1, 9, 0, 0, 0, time.UTC), time.Date(2024, 1, 1, 9, 19, 59, 0, time.UTC)),
					NewDuration(&adding, time.Date(2024, 1, 1, 9, 20, 0, 0, time.UTC), time.Date(2024, 1, 1, 9, 39, 59, 0, time.UTC)),
					NewDuration(&adding, time.Date(2024, 1, 1, 9, 40, 0, 0, time.UTC), time.Date(2024, 1, 1, 9, 59, 59, 0, time.UTC)),
					NewDuration(&adding, time.Date(2024, 1, 1, 10, 0, 0, 0, time.UTC), time.Date(2024, 1, 1, 10, 59, 59, 0, time.UTC)),
					NewDuration(&adding, time.Date(2024, 1, 1, 11, 0, 0, 0, time.UTC), time.Date(2024, 1, 1, 11, 59, 59, 0, time.UTC)),
				},
				err: nil,
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			// Given
			d, err := collection.NewDurations(tt.args.existing...)
			if err != nil {
				t.Fatal(err)
			}

			// When
			err = d.Merge(tt.args.adding[0])

			// Then
			if !tt.wantErr {
				if err != nil {
					t.Errorf("Merge() error = %v, wantErr %v", err, tt.wantErr)
				}
				if !reflect.DeepEqual(d.Durations(), tt.want.durations) {
					t.Errorf("Merge() = %v, want %v", d.Durations(), tt.want.durations)
				}
				return
			}

			if !errors.Is(err, tt.want.err) {
				t.Errorf("Merge() error = %v, wantErr %v", err, tt.want.err)
			}
		})
	}
}

func TestDurations_Find(t *testing.T) {
	// Setup
	t.Parallel()

	value1 := "1"
	value2 := "2"
	value3 := "3"

	type args struct {
		existing []value.Duration[string]
		criteria time.Time
	}
	type want struct {
		duration value.Duration[string]
		err      error
	}
	tests := []struct {
		name    string
		args    args
		want    want
		wantErr bool
	}{
		{
			name: "success/1",
			args: args{
				existing: []value.Duration[string]{
					NewDuration(&value1, time.Date(2024, 1, 1, 9, 0, 0, 0, time.UTC), time.Date(2024, 1, 1, 9, 59, 59, 0, time.UTC)),
					NewDuration(&value2, time.Date(2024, 1, 1, 10, 0, 0, 0, time.UTC), time.Date(2024, 1, 1, 10, 59, 59, 0, time.UTC)),
					NewDuration(&value3, time.Date(2024, 1, 1, 11, 0, 0, 0, time.UTC), time.Date(2024, 1, 1, 11, 59, 59, 0, time.UTC)),
				},
				criteria: time.Date(2024, 1, 1, 9, 59, 59, 0, time.UTC),
			},
			want: want{
				duration: NewDuration(&value1, time.Date(2024, 1, 1, 9, 0, 0, 0, time.UTC), time.Date(2024, 1, 1, 9, 59, 59, 0, time.UTC)),
				err:      nil,
			},
			wantErr: false,
		},
		{
			name: "success/2",
			args: args{
				existing: []value.Duration[string]{
					NewDuration(&value1, time.Date(2024, 1, 1, 9, 0, 0, 0, time.UTC), time.Date(2024, 1, 1, 9, 59, 59, 0, time.UTC)),
					NewDuration(&value2, time.Date(2024, 1, 1, 10, 0, 0, 0, time.UTC), time.Date(2024, 1, 1, 10, 59, 59, 0, time.UTC)),
					NewDuration(&value3, time.Date(2024, 1, 1, 11, 0, 0, 0, time.UTC), time.Date(2024, 1, 1, 11, 59, 59, 0, time.UTC)),
				},
				criteria: time.Date(2024, 1, 1, 10, 0, 0, 0, time.UTC),
			},
			want: want{
				duration: NewDuration(&value2, time.Date(2024, 1, 1, 10, 0, 0, 0, time.UTC), time.Date(2024, 1, 1, 10, 59, 59, 0, time.UTC)),
				err:      nil,
			},
			wantErr: false,
		},
		{
			name: "success/3",
			args: args{
				existing: []value.Duration[string]{
					NewDuration(&value1, time.Date(2024, 1, 1, 9, 0, 0, 0, time.UTC), time.Date(2024, 1, 1, 9, 59, 59, 0, time.UTC)),
					NewDuration(&value2, time.Date(2024, 1, 1, 10, 0, 0, 0, time.UTC), time.Date(2024, 1, 1, 10, 59, 59, 0, time.UTC)),
					NewDuration(&value3, time.Date(2024, 1, 1, 11, 0, 0, 0, time.UTC), time.Date(2024, 1, 1, 11, 59, 59, 0, time.UTC)),
				},
				criteria: time.Date(2024, 1, 1, 10, 59, 59, 0, time.UTC),
			},
			want: want{
				duration: NewDuration(&value2, time.Date(2024, 1, 1, 10, 0, 0, 0, time.UTC), time.Date(2024, 1, 1, 10, 59, 59, 0, time.UTC)),
				err:      nil,
			},
			wantErr: false,
		},
		{
			name: "success/4",
			args: args{
				existing: []value.Duration[string]{
					NewDuration(&value1, time.Date(2024, 1, 1, 9, 0, 0, 0, time.UTC), time.Date(2024, 1, 1, 9, 59, 59, 0, time.UTC)),
					NewDuration(&value2, time.Date(2024, 1, 1, 10, 0, 0, 0, time.UTC), time.Date(2024, 1, 1, 10, 59, 59, 0, time.UTC)),
					NewDuration(&value3, time.Date(2024, 1, 1, 11, 0, 0, 0, time.UTC), time.Date(2024, 1, 1, 11, 59, 59, 0, time.UTC)),
				},
				criteria: time.Date(2024, 1, 1, 11, 0, 0, 0, time.UTC),
			},
			want: want{
				duration: NewDuration(&value3, time.Date(2024, 1, 1, 11, 0, 0, 0, time.UTC), time.Date(2024, 1, 1, 11, 59, 59, 0, time.UTC)),
				err:      nil,
			},
			wantErr: false,
		},
		{
			name: "fail/1",
			args: args{
				existing: []value.Duration[string]{
					NewDuration(&value1, time.Date(2024, 1, 1, 9, 0, 0, 0, time.UTC), time.Date(2024, 1, 1, 9, 59, 59, 0, time.UTC)),
					NewDuration(&value2, time.Date(2024, 1, 1, 10, 0, 0, 0, time.UTC), time.Date(2024, 1, 1, 10, 59, 59, 0, time.UTC)),
					NewDuration(&value3, time.Date(2024, 1, 1, 11, 0, 0, 0, time.UTC), time.Date(2024, 1, 1, 11, 59, 59, 0, time.UTC)),
				},
				criteria: time.Date(2024, 1, 1, 8, 59, 59, 0, time.UTC),
			},
			want: want{
				duration: value.Duration[string]{},
				err:      collection.ErrDurationNotFound,
			},
			wantErr: true,
		},
		{
			name: "fail/2",
			args: args{
				existing: []value.Duration[string]{
					NewDuration(&value1, time.Date(2024, 1, 1, 9, 0, 0, 0, time.UTC), time.Date(2024, 1, 1, 9, 59, 59, 0, time.UTC)),
					NewDuration(&value2, time.Date(2024, 1, 1, 10, 0, 0, 0, time.UTC), time.Date(2024, 1, 1, 10, 59, 59, 0, time.UTC)),
					NewDuration(&value3, time.Date(2024, 1, 1, 11, 0, 0, 0, time.UTC), time.Date(2024, 1, 1, 11, 59, 59, 0, time.UTC)),
				},
				criteria: time.Date(2024, 1, 1, 12, 0, 0, 0, time.UTC),
			},
			want: want{
				duration: value.Duration[string]{},
				err:      collection.ErrDurationNotFound,
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			// Given
			d, err := collection.NewDurations(tt.args.existing...)
			if err != nil {
				t.Fatal(err)
			}

			// When
			duration, err := d.Find(tt.args.criteria)

			// Then
			if !tt.wantErr {
				if err != nil {
					t.Errorf("Find() error = %v, wantErr %v", err, tt.wantErr)
				}
				if !reflect.DeepEqual(duration, tt.want.duration) {
					t.Errorf("Find() = %+v, want %+v", duration, tt.want.duration)
				}
				return
			}

			if !errors.Is(err, tt.want.err) {
				t.Errorf("Find() error = %v, wantErr %v", err, tt.want.err)
			}
		})
	}
}
