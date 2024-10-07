package auth

import (
	"context"
	"database/sql"

	"github.com/massivebugs/home-feature-server/db"
)

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
