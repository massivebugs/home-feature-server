// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package cashbunny_account

import (
	"database/sql"
	"time"
)

type CashbunnyAccount struct {
	ID          uint32
	UserID      uint32
	Name        string
	Description sql.NullString
	Balance     float64
	Currency    string
	Type        string
	OrderIndex  int32
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   sql.NullTime
}
