package system_preference

type userSystemPreferenceResponse struct {
	Language *string `json:"language"`
}

func newUserSystemPreferenceResponse(usp *UserSystemPreference) userSystemPreferenceResponse {
	return userSystemPreferenceResponse{
		Language: usp.Language,
	}
}
