package auth

import (
	"time"

	"github.com/massivebugs/home-feature-server/db/queries"
)

type AuthUser struct {
	id         uint32
	name       string
	loggedInAt time.Time
	createdAt  time.Time
}

func NewAuthUserFromQueries(u *queries.User) *AuthUser {
	user := &AuthUser{
		id:        u.ID,
		name:      u.Name,
		createdAt: u.CreatedAt,
	}

	return user
}

func (u *AuthUser) SetLoginTime(claims *JWTClaims) {
	u.loggedInAt = claims.IssuedAt.Time
}
