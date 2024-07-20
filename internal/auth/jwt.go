package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTClaims struct {
	UserID uint32 `json:"user_id"`
	jwt.RegisteredClaims
}

type JWTBuilder struct {
	claims JWTClaims
}

func NewJWTBuilder(now time.Time, expireHours int, userID uint32) JWTBuilder {
	return JWTBuilder{
		claims: JWTClaims{
			uint32(userID),
			jwt.RegisteredClaims{
				// TODO: Add info for other claims?
				ExpiresAt: jwt.NewNumericDate(
					now.Add(time.Hour * time.Duration(expireHours)),
				),
				IssuedAt: jwt.NewNumericDate(now),
			},
		},
	}
}

func (c *JWTBuilder) CreateAndSignToken(m *jwt.SigningMethodHMAC, secret string) (string, error) {
	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c.claims)

	// Generate encoded token and send it as response.
	tokenStr, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return tokenStr, nil
}
