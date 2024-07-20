// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package user

import (
	"context"
	"database/sql"
)

type Querier interface {
	CreateUser(ctx context.Context, db DBTX, name string) (sql.Result, error)
	DeleteUser(ctx context.Context, db DBTX, id uint32) error
	GetUser(ctx context.Context, db DBTX, id uint32) (*User, error)
	UpdateUser(ctx context.Context, db DBTX, arg UpdateUserParams) error
}

var _ Querier = (*Queries)(nil)
