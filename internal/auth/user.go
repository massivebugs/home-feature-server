package auth

import (
	"time"

	"github.com/massivebugs/home-feature-server/db/queries"
)

type User struct {
	Id         uint32
	Name       string
	Email      string
	DisabledAt *time.Time
	LoggedInAt time.Time
	CreatedAt  time.Time
}

func (u *User) IsDisabled() bool {
	return u.DisabledAt != nil
}

func NewUserFromQueries(u *queries.User) *User {
	user := &User{
		Id:        u.ID,
		Name:      u.Name,
		Email:     u.Email,
		CreatedAt: u.CreatedAt,
	}

	if u.DisabledAt.Valid {
		user.DisabledAt = &u.DisabledAt.Time
	}

	return user
}
