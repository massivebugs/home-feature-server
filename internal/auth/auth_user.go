package auth

import (
	"time"

	"github.com/massivebugs/home-feature-server/db/service/auth_repository"
)

type AuthUser struct {
	ID         uint32
	Name       string
	LoggedInAt time.Time
	CreatedAt  time.Time
}

func NewAuthUser(user *auth_repository.User, claims *JWTClaims) AuthUser {
	return AuthUser{
		ID:         user.ID,
		Name:       user.Name,
		LoggedInAt: claims.IssuedAt.Time,
		CreatedAt:  user.CreatedAt,
	}
}
