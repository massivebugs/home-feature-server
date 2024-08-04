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
	CreateCategory(ctx context.Context, db DBTX, arg CreateCategoryParams) (sql.Result, error)
	GetCategoryByID(ctx context.Context, db DBTX, arg GetCategoryByIDParams) (*CashbunnyCategory, error)
	GetCategoryByName(ctx context.Context, db DBTX, arg GetCategoryByNameParams) (*CashbunnyCategory, error)
	IncrementIndex(ctx context.Context, db DBTX, arg IncrementIndexParams) error
	ListAccounts(ctx context.Context, db DBTX, userID uint32) ([]*CashbunnyAccount, error)
	ListAccountsAndCategories(ctx context.Context, db DBTX, userID uint32) ([]*ListAccountsAndCategoriesRow, error)
	ListCategoriesByUserID(ctx context.Context, db DBTX, userID uint32) ([]*CashbunnyCategory, error)
}

var _ Querier = (*Queries)(nil)
