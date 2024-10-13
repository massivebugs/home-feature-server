package db

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func OpenMySQLDatabase(user string, password string, host string, port string, database string) (*Handle, error) {
	h, err := sql.Open(
		"mysql",
		fmt.Sprintf(
			"%s:%s@tcp(%s:%s)/%s?multiStatements=true&parseTime=true",
			user,
			password,
			host,
			port,
			database,
		),
	)
	if err != nil {
		return nil, err
	}

	return NewDB(h), nil
}

// A common interface for both *sql.DB and *sql.Tx to abstract usage in transactions
// https://github.com/golang/go/issues/14468
type DB interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

type Handle struct {
	*sql.DB
}

func NewDB(h *sql.DB) *Handle {
	return &Handle{
		h,
	}
}

func (db *Handle) WithTx(ctx context.Context, fn func(tx DB) error) error {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return nil
	}

	err = fn(tx)
	if err != nil {
		if err := tx.Rollback(); err != nil {
			return err
		}
		return err
	}

	return tx.Commit()
}
