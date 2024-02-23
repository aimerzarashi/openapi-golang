package collection

import (
	"errors"
	"openapi/internal/domain/common/value"
	"reflect"
	"testing"
	"time"
)

func NewDuration[T any](arg T, startAt, endAt time.Time) value.Duration[T] {
	duration, err := value.NewDuration(arg, startAt, endAt)
	if err != nil {
		panic(err)
	}
	return duration
}

func TestAdjust(t *testing.T) {
	type args struct {
		target value.Duration[string]
		point  value.Duration[string]
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
				target: NewDuration("target", time.Date(2024, 1, 1, 9, 0, 0, 0, time.UTC), time.Date(2024, 1, 1, 9, 59, 59, 0, time.UTC)),
				point:  NewDuration("point", time.Date(2024, 1, 1, 8, 0, 0, 0, time.UTC), time.Date(2024, 1, 1, 8, 59, 59, 0, time.UTC)),
			},
			want: want{
				durations: []value.Duration[string]{
					NewDuration("target", time.Date(2024, 1, 1, 9, 0, 0, 0, time.UTC), time.Date(2024, 1, 1, 9, 59, 59, 0, time.UTC)),
				},
				err: nil,
			},
			wantErr: false,
		},
		{
			name: "success/2",
			args: args{
				target: NewDuration("target", time.Date(2024, 1, 1, 9, 0, 0, 0, time.UTC), time.Date(2024, 1, 1, 9, 59, 59, 0, time.UTC)),
				point:  NewDuration("point", time.Date(2024, 1, 1, 10, 0, 0, 0, time.UTC), time.Date(2024, 1, 1, 10, 59, 59, 0, time.UTC)),
			},
			want: want{
				durations: []value.Duration[string]{
					NewDuration("target", time.Date(2024, 1, 1, 9, 0, 0, 0, time.UTC), time.Date(2024, 1, 1, 9, 59, 59, 0, time.UTC)),
				},
				err: nil,
			},
			wantErr: false,
		},
		{
			name: "success/3",
			args: args{
				target: NewDuration("target", time.Date(2024, 1, 1, 9, 0, 0, 0, time.UTC), time.Date(2024, 1, 1, 9, 59, 59, 0, time.UTC)),
				point:  NewDuration("point", time.Date(2024, 1, 1, 9, 0, 0, 0, time.UTC), time.Date(2024, 1, 1, 9, 59, 59, 0, time.UTC)),
			},
			want: want{
				durations: []value.Duration[string]{},
				err: nil,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Adjust(tt.args.target, tt.args.point)
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
