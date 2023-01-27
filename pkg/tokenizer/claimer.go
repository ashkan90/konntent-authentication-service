package tokenizer

import (
	"github.com/golang-jwt/jwt/v4"
	"konntent-authentication-service/internal/app/authorize"
)

type Claims struct {
	authorize.JWTClaim
	jwt.RegisteredClaims
}

func NewClaim(
	jwtClaim authorize.JWTClaim) Claims {
	return Claims{
		JWTClaim: jwtClaim,
	}
}
