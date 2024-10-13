package system_preference

import "github.com/massivebugs/home-feature-server/db/queries"

type UserSystemPreference struct {
	ID       uint32
	UserID   uint32
	Language *string
}

func NewUserSystemPreferenceFromQueries(data *queries.UserSystemPreference) *UserSystemPreference {
	usp := &UserSystemPreference{
		ID:     data.ID,
		UserID: data.UserID,
	}

	if data.Language.Valid {
		usp.Language = &data.Language.String
	}

	return usp
}
