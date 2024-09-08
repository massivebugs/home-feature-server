package response

import (
	"time"

	"github.com/massivebugs/home-feature-server/internal/auth"
)

type AuthUserResponseDTO struct {
	ID         uint32    `json:"id"`
	Name       string    `json:"name"`
	LoggedInAt time.Time `json:"logged_in_at"`
	CreatedAt  time.Time `json:"created_at"`
}

func NewAuthUserResponseDTO(au auth.AuthUser) AuthUserResponseDTO {
	return AuthUserResponseDTO{
		ID:         au.ID,
		Name:       au.Name,
		LoggedInAt: au.LoggedInAt,
		CreatedAt:  au.CreatedAt,
	}
}

type CreateJWTTokenResponseDTO struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}

func NewCreateJWTTokenResponseDTO(token string, refreshToken string) CreateJWTTokenResponseDTO {
	return CreateJWTTokenResponseDTO{
		Token:        token,
		RefreshToken: refreshToken,
	}
}
