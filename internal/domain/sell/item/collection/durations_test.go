package collection_test

import (
	"errors"
	"fmt"
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

func TestDurations_Add(t *testing.T) {
	// Setup
	t.Parallel()

	currentAt := time.Now()

	var validDurations []value.Duration
	for i := 0; i < 5; i++ {
		startAt := currentAt.Add(time.Duration(i-2)*time.Hour)
		endAt := currentAt.Add(time.Duration(i+1-2)*time.Hour-1*time.Second)
		duration, err := value.NewDuration(startAt, endAt)
		if err != nil {
			t.Fatal(err)
		}
		validDurations = append(validDurations, duration)
		fmt.Printf("got: %v %v\n", duration.StartAt().Format(time.RFC3339), duration.EndAt().Format(time.RFC3339))
	}
	fmt.Println("-----------------------")

	startAt := currentAt.Add(0*time.Hour+20*time.Minute)
	endAt := currentAt.Add(1*time.Hour+30*time.Minute-1*time.Second)
	addDuration, err := value.NewDuration(startAt, endAt)
	if err != nil {
		t.Fatal(err)
	}

	type args struct {
		duration value.Duration
	}
	type want struct {}
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
				duration: addDuration,
			},
			want: want{},
			wantErr: false,
			errType: nil,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			// Given
			got, err := collection.NewDurations(validDurations)
			if err != nil {
				t.Fatal(err)
			}

			// When
			err = got.Add(tt.args.duration)

			// Then
			if !tt.wantErr {
				if err != nil {
					t.Errorf("NewDurations() error = %v, wantErr %v", err, tt.wantErr)
				}
				for _, v := range got.Durations() {
					fmt.Printf("got: %v %v\n", v.StartAt().Format(time.RFC3339), v.EndAt().Format(time.RFC3339))
				}
				return
			}

			if !errors.Is(err, tt.errType) {
				t.Errorf("NewDurations() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
