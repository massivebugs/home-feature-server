package cashbunny

import (
	"time"

	"github.com/massivebugs/home-feature-server/db/queries"
)

type UserPreferences struct {
	userCurrencies []string
	createdAt      time.Time
	updatedAt      time.Time
}

func NewUserPreferencesFromQueries(
	up *queries.CashbunnyUserPreference,
	ucs []string,
) *UserPreferences {
	return &UserPreferences{
		userCurrencies: ucs,
		createdAt:      up.CreatedAt,
		updatedAt:      up.UpdatedAt,
	}
}
