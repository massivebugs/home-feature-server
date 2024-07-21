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
	Description string
	Balance     float64
	Currency    string
	Type        string
	OrderIndex  uint32
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   sql.NullTime
}
