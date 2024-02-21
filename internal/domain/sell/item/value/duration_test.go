package value_test

import (
	"errors"
	"openapi/internal/domain/sell/item/value"
	"testing"
	"time"
)

func TestNewDuration(t *testing.T) {
	// Setup
	t.Parallel()

	type args struct {
		startAt time.Time
		endAt   time.Time
	}
	type want struct {
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
				startAt: time.Now(),
				endAt:   time.Now().Add(1 * time.Hour),
			},
			want: want{
				startAt: time.Now(),
				endAt:   time.Now().Add(1 * time.Hour),
			},
			wantErr: false,
			errType: nil,
		},
		{
			name: "fail/startAtEmpty",
			args: args{
				startAt: time.Time{},
				endAt:   time.Now().Add(1 * time.Hour),
			},
			want:    want{},
			wantErr: true,
			errType: value.ErrDurationStartAtEmpty,
		},
		{
			name: "fail/endAtEmpty",
			args: args{
				startAt: time.Now(),
				endAt:   time.Time{},
			},
			want:    want{},
			wantErr: true,
			errType: value.ErrDurationEndAtEmpty,
		},
		{
			name: "fail/invalid",
			args: args{
				startAt: time.Now(),
				endAt:   time.Now().Add(-1 * time.Hour),
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

			got, err := value.NewDuration(tt.args.startAt, tt.args.endAt)

			if !tt.wantErr {
				if err != nil {
					t.Errorf("NewDuration() error = %v, wantErr %v", err, tt.wantErr)
				}
				if got.StartAt().Equal(tt.want.startAt) {
					t.Errorf("NewDuration() got = %v, want %v", got.StartAt(), tt.want.startAt)
				}
				if got.EndAt().Equal(tt.want.endAt) {
					t.Errorf("NewDuration() got = %v, want %v", got.EndAt(), tt.want.endAt)
				}
				return
			}

			if !errors.Is(err, tt.errType) {
				t.Errorf("NewDuration() error = %v, wantErr %v", err, tt.errType)
				return
			}
		})
	}
}
