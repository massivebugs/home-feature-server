// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package cashbunny_account

import (
	"context"
	"database/sql"
)

type Querier interface {
	CreateAccount(ctx context.Context, db DBTX, arg CreateAccountParams) (sql.Result, error)
	ListAccounts(ctx context.Context, db DBTX, userID uint32) ([]*CashbunnyAccount, error)
}

var _ Querier = (*Queries)(nil)
