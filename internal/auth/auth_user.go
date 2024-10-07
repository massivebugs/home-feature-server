package auth

import (
	"time"

	"github.com/massivebugs/home-feature-server/db/queries"
)

type AuthUser struct {
	ID         uint32
	Name       string
	LoggedInAt time.Time
	CreatedAt  time.Time
}

func NewAuthUserFromDBGateway(u *queries.User) *AuthUser {
	user := &AuthUser{
		ID:        u.ID,
		Name:      u.Name,
		CreatedAt: u.CreatedAt,
	}

	return user
}

func (u *AuthUser) SetLoginTime(claims *JWTClaims) {
	u.LoggedInAt = claims.IssuedAt.Time
}
