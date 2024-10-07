package cashbunny

import (
	"time"

	"github.com/massivebugs/home-feature-server/db/queries"
)

type UserPreferences struct {
	UserCurrencies []string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

func NewUserPreferencesFromDBGateway(
	up *queries.CashbunnyUserPreference,
	ucs []string,
) *UserPreferences {
	return &UserPreferences{
		UserCurrencies: ucs,
		CreatedAt:      up.CreatedAt,
		UpdatedAt:      up.UpdatedAt,
	}
}
