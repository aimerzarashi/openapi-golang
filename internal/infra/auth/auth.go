package auth

import (
	"errors"
	"openapi/internal/domain/auth"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type (
	Token struct {
		auth.IToken
	}
)

func NewToken() (auth.IToken, error) {
	return &Token{}, nil
}

func (t *Token) Decode(secret string, unsignedToken string) (*auth.AccessToken, error) {
	token, err := jwt.Parse(unsignedToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(secret), nil
	})
	if err != nil {
		return nil, errors.Join(err, err)
	}

	return &auth.AccessToken{
		UserId: uuid.MustParse(token.Claims.(jwt.MapClaims)["user_id"].(string)),
	}, nil
}
