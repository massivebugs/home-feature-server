package auth

import (
	"context"
	"database/sql"

	"github.com/massivebugs/home-feature-server/db"
)

// ### User ###

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

// ### User Password ###

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

// ### User Refresh Token ###

type CreateUserRefreshTokenParams struct {
	UserID    uint32
	Value     string
	ExpiresAt sql.NullTime
}

type DeleteUserRefreshTokenParams struct {
	UserID uint32
	Value  string
}

type GetUserRefreshTokenExistsByValueParams struct {
	UserID uint32
	Value  string
}

type IUserRefreshTokenRepository interface {
	CreateUserRefreshToken(ctx context.Context, db db.DB, arg CreateUserRefreshTokenParams) error
	GetUserRefreshTokenExistsByValue(ctx context.Context, db db.DB, arg GetUserRefreshTokenExistsByValueParams) (bool, error)
}
