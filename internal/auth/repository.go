package auth

import (
	"context"
	"time"

	"github.com/massivebugs/home-feature-server/db"
)

// ### User ###

type GetUsernameOrEmailExistsParams struct {
	Name  string
	Email string
}

type CreateUserParams struct {
	Name  string
	Email string
}

type UpdateUserParams struct {
	Name  string
	Email string
	ID    uint32
}

type IUserRepository interface {
	GetUsernameOrEmailExists(ctx context.Context, db db.DB, arg GetUsernameOrEmailExistsParams) (bool, error)
	CreateUser(ctx context.Context, db db.DB, arg CreateUserParams) (uint32, error)
	DeleteUser(ctx context.Context, db db.DB, id uint32) error
	GetUser(ctx context.Context, db db.DB, id uint32) (*User, error)
	GetUserByName(ctx context.Context, db db.DB, name string) (*User, error)
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
	ExpiresAt time.Time
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
