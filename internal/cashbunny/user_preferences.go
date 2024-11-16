package cashbunny

import (
	"time"

	"github.com/massivebugs/home-feature-server/db/queries"
)

type UserPreference struct {
	UserCurrencies []string
	createdAt      time.Time
	updatedAt      time.Time
}

func NewUserPreferenceFromQueries(
	up *queries.CashbunnyUserPreference,
	ucs []string,
) *UserPreference {
	return &UserPreference{
		UserCurrencies: ucs,
		createdAt:      up.CreatedAt,
		updatedAt:      up.UpdatedAt,
	}
}
