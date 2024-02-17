package item_test

import (
	"openapi/internal/domain/stock/item"
	"testing"

	"github.com/google/uuid"
)

func TestNewId(t *testing.T) {
	t.Parallel()

	// Given
	id := uuid.New()

	type args struct {
		v uuid.UUID
	}
	tests := []struct {
		name    string
		args    args
		want    uuid.UUID
		wantErr bool
		errType error
	}{
		{
			name: "success",
			args: args{
				v: id,
			},
			want:    id,
			wantErr: false,
			errType: nil,
		},
		{
			name: "fail",
			args: args{
				v: uuid.Nil,
			},
			want:    uuid.Nil,
			wantErr: true,
			errType: item.ErrIdNil,
		},
	}

	// When & Then
	for _, tt := range tests {
		tt := tt // capture range variable
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := item.NewId(tt.args.v)
			if (err != nil) != tt.wantErr && tt.errType != tt.errType {
				t.Errorf("NewId() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got.UUID() != tt.want {
				t.Errorf("NewId() = %v, want %v", got, tt.want)
			}
			if got.String() != tt.want.String() {
				t.Errorf("NewId() = %v, want %v", got, tt.want)
			}
		})
	}
}
