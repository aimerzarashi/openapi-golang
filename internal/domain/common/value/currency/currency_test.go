package currency_test

import (
	"reflect"
	"testing"

	"openapi/internal/domain/common/value/currency"
)

func TestNewCurrency(t *testing.T) {
	// Setup
	t.Parallel()

	type args struct {
		ammount float64
		unit    string
	}
	type want struct {
		currency struct {
			ammount float64
			unit    string
		}
		err error
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
				ammount: 100,
				unit:    "JPY",
			},
			want: want{
				currency: struct {
					ammount float64
					unit    string
				}{
					ammount: 100,
					unit:    "JPY",
				},
				err: nil,
			},
			wantErr: false,
		},
		{
			name: "fail/unit",
			args: args{
				ammount: 100,
				unit:    "",
			},
			want: want{
				currency: struct {
					ammount float64
					unit    string
				}{
					ammount: 0,
					unit:    "",
				},
				err: currency.ErrCurrencyUnitInvalid,
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			// When
			got, err := currency.NewCurrency(tt.args.ammount, tt.args.unit)

			// Then
			if (err != nil) != tt.wantErr {
				t.Errorf("NewCurrency() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCurrency() = %v, want %v", got, tt.want)
			}
		})
	}
}
