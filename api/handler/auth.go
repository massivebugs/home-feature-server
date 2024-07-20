package handler

import (
	"database/sql"

	"github.com/labstack/echo/v4"
	"github.com/massivebugs/home-feature-server/api/config"
	"github.com/massivebugs/home-feature-server/api/dto"
	"github.com/massivebugs/home-feature-server/db/service/user"
	"github.com/massivebugs/home-feature-server/db/service/user_password"
	"github.com/massivebugs/home-feature-server/internal/api"
	"github.com/massivebugs/home-feature-server/internal/auth"
)

type AuthHandler struct {
	cfg  *config.Config
	auth *auth.Auth
}

func NewAuthHandler(cfg *config.Config, db *sql.DB) *AuthHandler {
	return &AuthHandler{
		cfg: cfg,
		auth: auth.NewAuth(
			db,
			user.New(),
			user_password.New(),
		),
	}
}

func (h *AuthHandler) CreateUser(ctx echo.Context) *api.APIResponse {
	req := new(dto.UserAuthRequestDTO)

	if err := ctx.Bind(req); err != nil {
		return api.NewAPIResponse(ctx, err, "")
	}

	if err := ctx.Validate(req); err != nil {
		return api.NewAPIResponse(ctx, err, "")
	}

	err := h.auth.CreateAuthUser(ctx.Request().Context(), req)

	return api.NewAPIResponse(ctx, err, "")
}

func (h *AuthHandler) CreateJWTToken(ctx echo.Context) *api.APIResponse {
	req := new(dto.UserAuthRequestDTO)

	if err := ctx.Bind(req); err != nil {
		return api.NewAPIResponse(ctx, err, "")
	}

	if err := ctx.Validate(req); err != nil {
		return api.NewAPIResponse(ctx, err, "")
	}

	result, err := h.auth.CreateJWTToken(ctx.Request().Context(), h.cfg.JWTSecret, req)

	return api.NewAPIResponse(ctx, err, result)
}
