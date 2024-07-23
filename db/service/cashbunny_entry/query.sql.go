// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: query.sql

package cashbunny_entry

import (
	"context"
	"database/sql"
	"time"
)

const createEntry = `-- name: CreateEntry :execresult
INSERT INTO
  cashbunny_entries (
    user_id,
    account_id,
    transaction_id,
    amount,
    currency,
    transacted_at
  )
VALUES
  (?, ?, ?, ?, ?, ?)
`

type CreateEntryParams struct {
	UserID        uint32
	AccountID     uint32
	TransactionID uint32
	Amount        float64
	Currency      string
	TransactedAt  time.Time
}

func (q *Queries) CreateEntry(ctx context.Context, db DBTX, arg CreateEntryParams) (sql.Result, error) {
	return db.ExecContext(ctx, createEntry,
		arg.UserID,
		arg.AccountID,
		arg.TransactionID,
		arg.Amount,
		arg.Currency,
		arg.TransactedAt,
	)
}

const listEntriesByTransactionID = `-- name: ListEntriesByTransactionID :many
SELECT
  id, user_id, account_id, transaction_id, amount, currency, transacted_at, created_at, updated_at, deleted_at
FROM
  cashbunny_entries
WHERE
  user_id = ?
  AND transaction_id = ?
`

type ListEntriesByTransactionIDParams struct {
	UserID        uint32
	TransactionID uint32
}

func (q *Queries) ListEntriesByTransactionID(ctx context.Context, db DBTX, arg ListEntriesByTransactionIDParams) ([]*CashbunnyEntry, error) {
	rows, err := db.QueryContext(ctx, listEntriesByTransactionID, arg.UserID, arg.TransactionID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []*CashbunnyEntry{}
	for rows.Next() {
		var i CashbunnyEntry
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.AccountID,
			&i.TransactionID,
			&i.Amount,
			&i.Currency,
			&i.TransactedAt,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.DeletedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}