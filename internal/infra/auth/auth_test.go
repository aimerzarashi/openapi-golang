package auth_test

import (
	domain "openapi/internal/domain/auth"
	infra "openapi/internal/infra/auth"
	"testing"

	"github.com/google/uuid"
)


func TestDecode(t *testing.T) {
	t.Parallel()

	type args struct {
		secret        string	
		unsignedToken string
	}
	tests := []struct {
		name    string
		args    args
		want    domain.AccessToken
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				secret: "your-256-bit-secret",
				unsignedToken: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwidXNlcl9pZCI6IjU1MGU4NDAwLWUyOWItNDFkNC1hNzE2LTQ0NjY1NTQ0MDAwMCIsImlhdCI6MTUxNjIzOTAyMn0.vJJHY6_UUYwOTcqsm2eyYiz96gcALF2BsZtuOIiVMcA",
			},
			want : domain.AccessToken{
				UserId: uuid.MustParse("550e8400-e29b-41d4-a716-446655440000"),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			// Given
			token, err := infra.NewToken()
			if err != nil {
				t.Fatal(err)
			}

			// When
			got, err := token.Decode(tt.args.secret, tt.args.unsignedToken)

			// Then
			if !tt.wantErr {
				if err != nil {
					t.Errorf("Decode() error = %v, wantErr %v", err, tt.wantErr)
				}
				if got.UserId != tt.want.UserId {
					t.Errorf("Decode() = %v, want %v", got.UserId, tt.want.UserId)
				}
				return
			}

			if err == nil {
				t.Errorf("Decode() = %v, want %v", got.UserId, tt.want.UserId)
			}
		})
	}
}
