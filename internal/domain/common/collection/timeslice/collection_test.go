package timeslice_test

import (
	"errors"
	"testing"
	"time"

	"openapi/internal/domain/common/collection/timeslice"
)

func NewItem[T any](value *T, startAt, endAt time.Time) *timeslice.Item[T] {
	item, err := timeslice.NewItem(value, startAt, endAt)
	if err != nil {
		panic(err)
	}
	return item
}

func TestNewCollection(t *testing.T) {
	// Setup
	t.Parallel()

	type T = string

	value1 := "value1"
	value2 := "value2"
	value3 := "value3"

	type args struct {
		items []*timeslice.Item[T]
	}
	type want struct {
		err error
	}
	tests := []struct {
		name string
		args args
		want want
		wantErr bool
	}{
		{
			name: "success/empty",
			args: args{
				items: []*timeslice.Item[T]{},
			},
			want: want{
				err: nil,
			},
			wantErr: false,
		},
		{
			name: "success/not empty",
			args: args{
				items: []*timeslice.Item[T]{
					NewItem(&value1, time.Date(2024, 1, 1, 9, 0, 0, 0, time.UTC), time.Date(2024, 1, 1, 9, 59, 59, 0, time.UTC)),
					NewItem(&value2, time.Date(2024, 1, 1, 10, 0, 0, 0, time.UTC), time.Date(2024, 1, 1, 10, 59, 59, 0, time.UTC)),
					NewItem(&value3, time.Date(2024, 1, 1, 11, 0, 0, 0, time.UTC), time.Date(2024, 1, 1, 11, 59, 59, 0, time.UTC)),
				},
			},
			want: want{
				err: nil,
			},
			wantErr: false,
		},
		{
			name: "fail",
			args: args{
				items: []*timeslice.Item[T]{
					NewItem(&value1, time.Date(2024, 1, 1, 9, 0, 0, 0, time.UTC), time.Date(2024, 1, 1, 9, 59, 59, 0, time.UTC)),
					NewItem(&value2, time.Date(2024, 1, 1, 9, 59, 59, 0, time.UTC), time.Date(2024, 1, 1, 10, 59, 59, 0, time.UTC)),
					NewItem(&value3, time.Date(2024, 1, 1, 11, 0, 0, 0, time.UTC), time.Date(2024, 1, 1, 11, 59, 59, 0, time.UTC)),
				},
			},
			want: want{
				err: timeslice.ErrCollectionInvalid,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			// When
			got, err := timeslice.NewCollection(tt.args.items...)

			// Then
			if !tt.wantErr {
				if err != nil {
					t.Errorf("NewCollection() error = %v, wantErr %v", err, tt.wantErr)
				}
				for i, item := range got.Items() {
					if item != tt.args.items[i] {
						t.Errorf("NewCollection() = %+v, want %+v", item, tt.args.items[i])
					}
				}
				return
			}

			if !errors.Is(err, tt.want.err) {
				t.Errorf("NewCollection() error = %v, wantErr %v", err, tt.want.err)
			}
		})
	}
}

func TestCollection_Find(t *testing.T) {
	// Setup
	t.Parallel()

	type T = string

	value1 := "value1"
	value2 := "value2"
	value3 := "value3"

	items := []*timeslice.Item[T]{
		NewItem(&value1, time.Date(2024, 1, 1, 9, 0, 0, 0, time.UTC), time.Date(2024, 1, 1, 9, 59, 59, 0, time.UTC)),
		NewItem(&value2, time.Date(2024, 1, 1, 10, 0, 0, 0, time.UTC), time.Date(2024, 1, 1, 10, 59, 59, 0, time.UTC)),
		NewItem(&value3, time.Date(2024, 1, 1, 11, 0, 0, 0, time.UTC), time.Date(2024, 1, 1, 11, 59, 59, 0, time.UTC)),
	}

	type args struct {
		criteria time.Time
	}
	type want struct {
		item *timeslice.Item[T]
		err error
	}
	tests := []struct {
		name string
		args args
		want want
		wantErr bool
	}{
		{
			name: "success/1",
			args: args{
				criteria: time.Date(2024, 1, 1, 9, 0, 0, 0, time.UTC),
			},
			want: want{
				item: items[0],
				err: nil,
			},
			wantErr: false,
		},
		{
			name: "success/2",
			args: args{
				criteria: time.Date(2024, 1, 1, 9, 59, 59, 0, time.UTC),
			},
			want: want{
				item: items[0],
				err: nil,
			},
			wantErr: false,
		},
		{
			name: "success/3",
			args: args{
				criteria: time.Date(2024, 1, 1, 10, 0, 0, 0, time.UTC),
			},
			want: want{
				item: items[1],
				err: nil,
			},
			wantErr: false,
		},
		{
			name: "success/4",
			args: args{
				criteria: time.Date(2024, 1, 1, 10, 59, 59, 0, time.UTC),
			},
			want: want{
				item: items[1],
				err: nil,
			},
			wantErr: false,
		},
		{
			name: "success/5",
			args: args{
				criteria: time.Date(2024, 1, 1, 11, 0, 0, 0, time.UTC),
			},
			want: want{
				item: items[2],
				err: nil,
			},
			wantErr: false,
		},
		{
			name: "success/6",
			args: args{
				criteria: time.Date(2024, 1, 1, 11, 59, 59, 0, time.UTC),
			},
			want: want{
				item: items[2],
				err: nil,
			},
			wantErr: false,
		},
		{
			name: "fail/1",
			args: args{
				criteria: time.Date(2024, 1, 1, 8, 59, 59, 0, time.UTC),
			},
			want: want{
				item: nil,
				err: timeslice.ErrCollectionNotFound,
			},
			wantErr: true,
		},
		{
			name: "fail/2",
			args: args{
				criteria: time.Date(2024, 1, 1, 12, 0, 0, 0, time.UTC),
			},
			want: want{
				item: nil,
				err: timeslice.ErrCollectionNotFound,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			// Given
			timeSlice, err := timeslice.NewCollection(items...)
			if err != nil {
				t.Fatal(err)
			}

			// When
			got, err := timeSlice.Find(tt.args.criteria)

			// Then
			if !tt.wantErr {
				if err != nil {
					t.Errorf("NewCollection() error = %v, wantErr %v", err, tt.wantErr)
				}
				if got.Value() != tt.want.item.Value() {
					t.Errorf("NewCollection() = %v, want %v", got, tt.want.item)
				}
				return
			}

			if !errors.Is(err, tt.want.err) {
				t.Errorf("NewCollection() error = %v, wantErr %v", err, tt.want.err)
			}
		})
	}
}