package auth

import (
	"time"
)

type authUserResponse struct {
	ID         uint32    `json:"id"`
	Name       string    `json:"name"`
	LoggedInAt time.Time `json:"logged_in_at"`
	CreatedAt  time.Time `json:"created_at"`
}

func newAuthUserResponse(au *AuthUser) authUserResponse {
	return authUserResponse{
		ID:         au.id,
		Name:       au.name,
		LoggedInAt: au.loggedInAt,
		CreatedAt:  au.createdAt,
	}
}

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
