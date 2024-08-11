// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package cashbunny_repository

import (
	"context"
	"database/sql"
)

type Querier interface {
	CreateAccount(ctx context.Context, db DBTX, arg CreateAccountParams) (sql.Result, error)
	CreateTransaction(ctx context.Context, db DBTX, arg CreateTransactionParams) (sql.Result, error)
	CreateUserCurrency(ctx context.Context, db DBTX, arg CreateUserCurrencyParams) (sql.Result, error)
	CreateUserPreferences(ctx context.Context, db DBTX, userID uint32) (sql.Result, error)
	DeleteAccount(ctx context.Context, db DBTX, arg DeleteAccountParams) error
	DeleteTransaction(ctx context.Context, db DBTX, arg DeleteTransactionParams) error
	DeleteTransactionsByAccountID(ctx context.Context, db DBTX, arg DeleteTransactionsByAccountIDParams) error
	GetAccountByID(ctx context.Context, db DBTX, arg GetAccountByIDParams) (*CashbunnyAccount, error)
	GetTransactionByID(ctx context.Context, db DBTX, arg GetTransactionByIDParams) (*CashbunnyTransaction, error)
	GetUserPreferenceByUserID(ctx context.Context, db DBTX, userID uint32) (*CashbunnyUserPreference, error)
	IncrementIndex(ctx context.Context, db DBTX, arg IncrementIndexParams) error
	ListAccounts(ctx context.Context, db DBTX, userID uint32) ([]*CashbunnyAccount, error)
	ListAccountsByIDs(ctx context.Context, db DBTX, arg ListAccountsByIDsParams) ([]*CashbunnyAccount, error)
	ListTransactions(ctx context.Context, db DBTX, userID uint32) ([]*CashbunnyTransaction, error)
	ListUserCurrencies(ctx context.Context, db DBTX, userID uint32) ([]*CashbunnyUserCurrency, error)
	UpdateAccountBalance(ctx context.Context, db DBTX, arg UpdateAccountBalanceParams) error
}

var _ Querier = (*Queries)(nil)