package auth

import (
	"time"

	"github.com/massivebugs/home-feature-server/db/queries"
)

type AuthUser struct {
	Id         uint32
	Name       string
	LoggedInAt time.Time
	CreatedAt  time.Time
}

func NewAuthUserFromQueries(u *queries.User) *AuthUser {
	user := &AuthUser{
		Id:        u.ID,
		Name:      u.Name,
		CreatedAt: u.CreatedAt,
	}

	return user
}
