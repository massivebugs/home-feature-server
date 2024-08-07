package cashbunny

// import (
// 	"time"

// 	"github.com/Rhymond/go-money"
// 	"github.com/massivebugs/home-feature-server/db/service/cashbunny_entry"
// )

// type Entry struct {
// 	ID           uint32
// 	Amount       *money.Money
// 	Currency     string
// 	TransactedAt time.Time
// 	CreatedAt    time.Time

// 	Account *Account
// }

// func NewEntry(e *cashbunny_entry.CashbunnyEntry) Entry {
// 	return Entry{
// 		ID:           e.ID,
// 		Amount:       money.NewFromFloat(e.Amount, e.Currency),
// 		Currency:     e.Currency,
// 		TransactedAt: e.TransactedAt,
// 		CreatedAt:    e.CreatedAt,
// 	}
// }
