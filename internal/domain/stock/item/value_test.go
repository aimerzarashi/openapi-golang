package item_test

import (
	"openapi/internal/domain/stock/item"
	"testing"
)

func TestNewName(t *testing.T) {
	t.Parallel()

	//Given
	type args struct {
		v string
	}

	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
		errType error
	}{
		{
			name: "success",
			args: args{
				v: "test",
			},
			want:    "test",
			wantErr: false,
			errType: nil,
		},
		{
			name: "fail",
			args: args{
				v: "",
			},
			want:    "",
			wantErr: true,
			errType: item.ErrNameEmpty,
		},
	}

	//When & Then
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := item.NewName(tt.args.v)
			if (err != nil) != tt.wantErr && tt.errType != tt.errType {
				t.Errorf("NewName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got.String() != tt.want {
				t.Errorf("NewName() = %v, want %v", got, tt.want)
			}
		})
	}
}
