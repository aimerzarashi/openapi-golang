package value

import (
	"reflect"
	"testing"
)

func TestNewCurrency(t *testing.T) {
	// Setup
	t.Parallel()

	type args struct {
		ammount float64
		unit    string
	}
	tests := []struct {
		name    string
		args    args
		want    Currency
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				ammount: 100,
				unit:    "JPY",
			},
			want: Currency{
				ammount: 100,
				unit:    "JPY",
			},
			wantErr: false,
		},
		{
			name: "fail/unit",
			args: args{
				ammount: 100,
				unit:    "",
			},
			want:    Currency{},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			// When
			got, err := NewCurrency(tt.args.ammount, tt.args.unit)

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
