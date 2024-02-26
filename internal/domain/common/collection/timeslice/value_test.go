package timeslice_test

import (
	"errors"
	"testing"
	"time"

	"openapi/internal/domain/common/collection/timeslice"
)

func TestNewItem(t *testing.T) {
	// Setup
	t.Parallel()

	value := "value"

	type args struct {
		value   *string
		startAt time.Time
		endAt   time.Time
	}
	type want struct {
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
				value:   &value,
				startAt: time.Date(2024, 1, 1, 9, 0, 0, 0, time.UTC),
				endAt:   time.Date(2024, 1, 1, 10, 0, 0, 0, time.UTC),
			},
			want: want{
				err: nil,
			},
			wantErr: false,
		},
		{
			name: "fail/1",
			args: args{
				value:   &value,
				startAt: time.Time{},
				endAt:   time.Date(2024, 1, 1, 10, 0, 0, 0, time.UTC),
			},
			want: want{
				err: timeslice.ErrItemStartAtEmpty,
			},
			wantErr: true,
		},
		{
			name: "fail/2",
			args: args{
				value:   &value,
				startAt: time.Date(2024, 1, 1, 10, 0, 0, 0, time.UTC),
				endAt:   time.Time{},
			},
			want: want{
				err: timeslice.ErrItemEndAtEmpty,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			// When
			got, err := timeslice.NewItem(&tt.args.value, tt.args.startAt, tt.args.endAt)

			// Then
			if !tt.wantErr {
				if err != nil {
					t.Errorf("NewItem() error = %v, wantErr %v", err, tt.wantErr)
				}
				if got.Value() != tt.args.value {
					t.Errorf("NewItem() = %v, want %v", got.Value(), tt.args.value)
				}
				if got.StartAt() != tt.args.startAt {
					t.Errorf("NewItem() = %v, want %v", got.StartAt(), tt.args.startAt)
				}
				if got.EndAt() != tt.args.endAt {
					t.Errorf("NewItem() = %v, want %v", got.EndAt(), tt.args.endAt)
				}
				if !got.Contains(tt.args.startAt) {
					t.Errorf("NewItem() = %v, want %v", got.Contains(tt.args.startAt), true)
				}
				if !got.Contains(tt.args.endAt) {
					t.Errorf("NewItem() = %v, want %v", got.Contains(tt.args.endAt), true)
				}
				if got.Contains(tt.args.startAt.Add(-1 * time.Second)) {
					t.Errorf("NewItem() = %v, want %v", got.Contains(tt.args.startAt), false)
				}
				if got.Contains(tt.args.endAt.Add(1 * time.Second)) {
					t.Errorf("NewItem() = %v, want %v", got.Contains(tt.args.endAt), false)
				}
				return
			}

			if errors.Is(err, tt.want.err) {
				return
			}
			t.Errorf("NewItem() error = %v, wantErr %v", err, tt.wantErr)
		})
	}
}
