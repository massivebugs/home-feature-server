package auth

import (
	"github.com/labstack/echo/v4"
	"github.com/massivebugs/home-feature-server/internal/api"
	"github.com/massivebugs/home-feature-server/internal/auth/dto"
)

type AuthHandler struct{}

func NewAuthHandler() *AuthHandler {
	return &AuthHandler{}
}

func (*AuthHandler) CreateUser(ctx echo.Context) error {
	v := new(dto.CreateUserRequestDTO)

	if err := ctx.Bind(v); err != nil {
		return api.NewAPIResponse(ctx, err, "")
	}

	if err := ctx.Validate(v); err != nil {
		return api.NewAPIResponse(ctx, err, "")
	}

	return api.NewAPIResponse(ctx, nil, "Validation test OK!")
}

func (*AuthHandler) LogIn(ctx echo.Context) error {
	return api.NewAPIResponse(ctx, nil, "Pong!")
}
