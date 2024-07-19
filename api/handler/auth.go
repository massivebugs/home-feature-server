package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/massivebugs/home-feature-server/api/dto"
	"github.com/massivebugs/home-feature-server/internal/api"
	"github.com/massivebugs/home-feature-server/internal/auth"
)

type AuthHandler struct {
	auth *auth.Auth
}

func NewAuthHandler() *AuthHandler {
	return &AuthHandler{
		auth: auth.NewAuth(),
	}
}

func (h *AuthHandler) CreateUser(ctx echo.Context) error {
	req := new(dto.CreateUserRequestDTO)

	if err := ctx.Bind(req); err != nil {
		return api.NewAPIResponse(ctx, err, "")
	}

	if err := ctx.Validate(req); err != nil {
		return api.NewAPIResponse(ctx, err, "")
	}

	result := h.auth.CreateUser(ctx.Request().Context(), req)

	return api.NewAPIResponse(ctx, nil, result)
}

func (*AuthHandler) LogIn(ctx echo.Context) error {
	return api.NewAPIResponse(ctx, nil, "Pong!")
}
