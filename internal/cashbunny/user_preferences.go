package cashbunny

import (
	"time"

	"github.com/massivebugs/home-feature-server/db/service/cashbunny_repository"
)

type UserPreferences struct {
	UserCurrencies []*cashbunny_repository.CashbunnyUserCurrency
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

func NewUserPreferences(
	up *cashbunny_repository.CashbunnyUserPreference,
	uc []*cashbunny_repository.CashbunnyUserCurrency,
) *UserPreferences {
	return &UserPreferences{
		UserCurrencies: uc,
		CreatedAt:      up.CreatedAt,
		UpdatedAt:      up.UpdatedAt,
	}
}
