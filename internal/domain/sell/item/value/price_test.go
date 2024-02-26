package value_test

import (
	"testing"

	"openapi/internal/domain/sell/item/value"
)

func TestNewPrice(t *testing.T) {
	// Setup
	t.Parallel()

	type arg struct {
		ammount float64
		unit string 
	}
	type want struct {
		ammount float64
		unit string 
	}
	tests := []struct {
		name string
		arg  arg
		want want
	}{
		{
			name: "success",
			arg: arg{
				ammount: 100,
				unit: "JPY",
			},
			want: want{
				ammount: 100,
				unit: "JPY",
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := value.NewPrice(tt.arg.ammount, tt.arg.unit)
			if err != nil {
				t.Fatal(err)
			}

			if got.Ammount() != tt.want.ammount {
				t.Errorf("want %v, but got %v", tt.want.ammount, got.Ammount())
			}
			if got.Unit() != tt.want.unit {
				t.Errorf("want %v, but got %v", tt.want.unit, got.Unit())
			}

		})
	}

}