package response

import "github.com/massivebugs/home-feature-server/internal/system_preference"

type UserSystemPreferencesDTO struct {
	Language *string `json:"language"`
}

func NewUserSystemPreferenceDTO(usp *system_preference.UserSystemPreference) UserSystemPreferencesDTO {
	return UserSystemPreferencesDTO{
		Language: usp.Language,
	}
}
