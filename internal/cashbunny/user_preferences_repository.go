package cashbunny

import (
	"context"

	"github.com/massivebugs/home-feature-server/db"
)

type CreateUserCurrencyParams struct {
	UserID       uint32
	CurrencyCode string
}

type IUserPreferencesRepository interface {
	GetUserPreferenceExistsByUserID(ctx context.Context, db db.DB, userID uint32) (bool, error)
	CreateUserPreferences(ctx context.Context, db db.DB, userID uint32) (uint32, error)
	GetUserPreferencesByUserID(ctx context.Context, db db.DB, userID uint32) (*UserPreferences, error)
	CreateUserCurrency(ctx context.Context, db db.DB, params CreateUserCurrencyParams) (uint32, error)
	ListUserCurrencies(ctx context.Context, db db.DB, userID uint32) ([]string, error)
}
