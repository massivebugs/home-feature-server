// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package cashbunny_repository

import (
	"database/sql"
	"time"
)

type CashbunnyAccount struct {
	ID          uint32
	UserID      uint32
	Category    string
	Name        string
	Description string
	Balance     float64
	Currency    string
	OrderIndex  uint32
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   sql.NullTime
}

type CashbunnyRecurrenceRule struct {
	ID        uint32
	Freq      string
	Dtstart   time.Time
	Count     int32
	Interval  int32
	Until     time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}

type CashbunnyScheduledTransaction struct {
	ID            uint32
	UserID        uint32
	SrcAccountID  uint32
	DestAccountID uint32
	Description   string
	Amount        float64
	Currency      string
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     sql.NullTime
}

type CashbunnyTransaction struct {
	ID                     uint32
	UserID                 uint32
	ScheduledTransactionID sql.NullInt32
	SrcAccountID           uint32
	DestAccountID          uint32
	Description            string
	Amount                 float64
	Currency               string
	TransactedAt           time.Time
	CreatedAt              time.Time
	UpdatedAt              time.Time
	DeletedAt              sql.NullTime
}

type CashbunnyUserCurrency struct {
	ID           uint32
	UserID       uint32
	CurrencyCode string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type CashbunnyUserPreference struct {
	ID        uint32
	UserID    uint32
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime
}
