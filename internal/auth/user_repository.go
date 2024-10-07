package auth

import (
	"context"

	"github.com/massivebugs/home-feature-server/db"
)

type UpdateUserParams struct {
	Name string
	ID   uint32
}

type IUserRepository interface {
	GetUsernameExists(ctx context.Context, db db.DB, name string) (bool, error)
	CreateUser(ctx context.Context, db db.DB, name string) (uint32, error)
	DeleteUser(ctx context.Context, db db.DB, id uint32) error
	GetUser(ctx context.Context, db db.DB, id uint32) (*AuthUser, error)
	GetUserByName(ctx context.Context, db db.DB, name string) (*AuthUser, error)
	UpdateUser(ctx context.Context, db db.DB, arg UpdateUserParams) error
}
