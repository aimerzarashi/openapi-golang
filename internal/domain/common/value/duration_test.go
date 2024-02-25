package value_test

import (
	"errors"
	"testing"
	"time"

	"openapi/internal/domain/common/value"
)

func TestNewDuration(t *testing.T) {
	// Setup
	t.Parallel()

	attribute := "value"
	startAt := time.Date(2024, 1, 1, 9, 0, 0, 0, time.UTC)
	endAt := time.Date(2024, 1, 1, 9, 59, 59, 0, time.UTC)

	type args struct {
		value   *string
		startAt time.Time
		endAt   time.Time
	}
	type want struct {
		value   *string
		startAt time.Time
		endAt   time.Time
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
				value:   &attribute,
				startAt: startAt,
				endAt:   endAt,
			},
			want: want{
				value:   &attribute,
				startAt: startAt,
				endAt:   endAt,
			},
			wantErr: false,
			errType: nil,
		},
		{
			name: "fail/startAtEmpty",
			args: args{
				startAt: time.Time{},
				endAt:   endAt,
			},
			want:    want{},
			wantErr: true,
			errType: value.ErrDurationStartAtEmpty,
		},
		{
			name: "fail/endAtEmpty",
			args: args{
				startAt: startAt,
				endAt:   time.Time{},
			},
			want:    want{},
			wantErr: true,
			errType: value.ErrDurationEndAtEmpty,
		},
		{
			name: "fail/invalid",
			args: args{
				startAt: endAt,
				endAt:   startAt,
			},
			want:    want{},
			wantErr: true,
			errType: value.ErrDurationInvalid,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := value.NewDuration(&tt.args.value, tt.args.startAt, tt.args.endAt)

			if !tt.wantErr {
				if err != nil {
					t.Errorf("NewDuration() error = %v, wantErr %v", err, tt.wantErr)
				}
				if *got.Value() != tt.want.value {
					t.Errorf("NewDuration() got = %v, want %v", got.Value(), tt.want.value)
				}
				if !got.StartAt().Equal(tt.want.startAt) {
					t.Errorf("NewDuration() got = %v, want %v", got.StartAt(), tt.want.startAt)
				}
				if !got.EndAt().Equal(tt.want.endAt) {
					t.Errorf("NewDuration() got = %v, want %v", got.EndAt(), tt.want.endAt)
				}
				if !got.Contains(startAt.Add(30 * time.Minute)) {
					t.Errorf("NewDuration() got = %v, want %v", true, got.Contains(startAt.Add(30 * time.Minute)))
				}
				if got.Contains(endAt.Add(30 * time.Minute)) {
					t.Errorf("NewDuration() got = %v, want %v", false, got.Contains(endAt.Add(30 * time.Minute)))
				}
				return
			}

			if !errors.Is(err, tt.errType) {
				t.Errorf("NewDuration() error = %v, wantErr %v", err, tt.errType)
			}
		})
	}
}
