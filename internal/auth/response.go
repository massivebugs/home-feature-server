package auth

type jwtTokenResponse struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}

func newJWTTokenResponse(token string, refreshToken string) jwtTokenResponse {
	return jwtTokenResponse{
		Token:        token,
		RefreshToken: refreshToken,
	}
}
