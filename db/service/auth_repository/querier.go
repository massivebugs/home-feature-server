// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package auth_repository

import (
	"context"
	"database/sql"
)

type Querier interface {
	CreateUser(ctx context.Context, db DBTX, name string) (sql.Result, error)
	CreateUserPassword(ctx context.Context, db DBTX, arg CreateUserPasswordParams) (sql.Result, error)
	CreateUserRefreshToken(ctx context.Context, db DBTX, arg CreateUserRefreshTokenParams) (sql.Result, error)
	DeleteUser(ctx context.Context, db DBTX, id uint32) error
	DeleteUserRefreshToken(ctx context.Context, db DBTX, arg DeleteUserRefreshTokenParams) error
	GetUser(ctx context.Context, db DBTX, id uint32) (*User, error)
	GetUserByName(ctx context.Context, db DBTX, name string) (*User, error)
	GetUserPasswordByUserID(ctx context.Context, db DBTX, userID uint32) (*UserPassword, error)
	GetUserRefreshTokenByValue(ctx context.Context, db DBTX, arg GetUserRefreshTokenByValueParams) (*UserRefreshToken, error)
	UpdateUser(ctx context.Context, db DBTX, arg UpdateUserParams) error
	UpdateUserPassword(ctx context.Context, db DBTX, arg UpdateUserPasswordParams) error
}

var _ Querier = (*Queries)(nil)
