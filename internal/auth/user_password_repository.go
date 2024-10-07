package auth

import (
	"context"

	"github.com/massivebugs/home-feature-server/db"
)

type CreateUserPasswordParams struct {
	UserID       uint32
	PasswordHash string
}

type UpdateUserPasswordParams struct {
	PasswordHash string
	ID           uint32
}

type IUserPasswordRepository interface {
	CreateUserPassword(ctx context.Context, db db.DB, arg CreateUserPasswordParams) error
	GetUserPasswordByUserID(ctx context.Context, db db.DB, userID uint32) (string, error)
	UpdateUserPassword(ctx context.Context, db db.DB, arg UpdateUserPasswordParams) error
}
