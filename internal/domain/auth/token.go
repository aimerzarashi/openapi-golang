package auth

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type (
	IToken interface {
		Decode(secret string, unsignedToken string) (*AccessToken, error)
	}

	AccessToken struct {
		UserId uuid.UUID
		jwt.RegisteredClaims
	}
)
