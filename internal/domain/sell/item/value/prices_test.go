package value

import (
	"reflect"
	"testing"
)

func TestNewPrice(t *testing.T) {
	// Setup
	t.Parallel()

	type args struct {
		ammount  float64
		currency string
	}
	tests := []struct {
		name    string
		args    args
		want    Price
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				ammount:  100,
				currency: "JPY",
			},
			want: Price{
				ammount:  100,
				currency: "JPY",
			},
			wantErr: false,
		},
		{
			name: "fail/ammount",
			args: args{
				ammount:  0,
				currency: "JPY",
			},
			want:    Price{},
			wantErr: true,
		},
		{
			name: "fail/currency",
			args: args{
				ammount:  100,
				currency: "",
			},
			want:    Price{},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			// When
			got, err := NewPrice(tt.args.ammount, tt.args.currency)

			// Then
			if (err != nil) != tt.wantErr {
				t.Errorf("NewPrice() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPrice() = %v, want %v", got, tt.want)
			}
		})
	}
}
