package auth

import (
	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	Uid uint `json:"uid"`
	jwt.RegisteredClaims
}

type Auth struct {
	signingKey string
}

type Authentication interface {
	GenerateJWT(uid uint) (string, error)
	ValidateToken(tokenString string) (*Claims, error)
}

func NewAuth(signingKey string) Authentication {
	return &Auth{
		signingKey: signingKey,
	}
}
