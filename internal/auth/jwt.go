package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTCustomClaims struct {
	UserID  uint32 `json:"user_id"`
	TokenID string `json:"token_id"`
}

type JWTClaims struct {
	JWTCustomClaims
	jwt.RegisteredClaims
}

type JWTBuilder struct {
	claims JWTClaims
}

func NewJWTBuilder(now time.Time, expireSeconds int, customClaims JWTCustomClaims) JWTBuilder {
	return JWTBuilder{
		claims: JWTClaims{
			customClaims,
			jwt.RegisteredClaims{
				// TODO: Add info for other claims?
				ExpiresAt: jwt.NewNumericDate(
					now.Add(time.Second * time.Duration(expireSeconds)),
				),
				IssuedAt: jwt.NewNumericDate(now),
			},
		},
	}
}

func (c *JWTBuilder) CreateAndSignToken(m *jwt.SigningMethodHMAC, secret string) (string, error) {
	// Create token with claims
	token := jwt.NewWithClaims(m, c.claims)

	// Generate encoded token and send it as response.
	tokenStr, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return tokenStr, nil
}
