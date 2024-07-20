package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTClaims struct {
	jwt.RegisteredClaims
	userID uint32
}

func NewJWTClaims(now time.Time, expireHours int, userID uint32) JWTClaims {
	return JWTClaims{
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(
				now.Add(time.Hour * time.Duration(expireHours)),
			),
		},
		uint32(userID),
	}
}
