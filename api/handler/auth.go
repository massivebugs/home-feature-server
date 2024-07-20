package handler

import (
	"database/sql"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/massivebugs/home-feature-server/api/config"
	"github.com/massivebugs/home-feature-server/api/response"
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
	req := new(auth.UserAuthRequestDTO)

	if err := ctx.Bind(req); err != nil {
		return api.NewAPIResponse(err, "")
	}

	if err := ctx.Validate(req); err != nil {
		return api.NewAPIResponse(err, "")
	}

	err := h.auth.CreateAuthUser(ctx.Request().Context(), req)

	return api.NewAPIResponse(err, "")
}

func (h *AuthHandler) CreateJWTToken(ctx echo.Context) *api.APIResponse {
	req := new(auth.UserAuthRequestDTO)

	if err := ctx.Bind(req); err != nil {
		return api.NewAPIResponse(err, "")
	}

	if err := ctx.Validate(req); err != nil {
		return api.NewAPIResponse(err, "")
	}

	result, err := h.auth.CreateJWTToken(ctx.Request().Context(), h.cfg.JWTSecret, req)

	return api.NewAPIResponse(err, result)
}

func (h *AuthHandler) GetAuthUser(ctx echo.Context) *api.APIResponse {
	token := ctx.Get("user").(*jwt.Token)
	claims := token.Claims.(*auth.JWTClaims)

	result, err := h.auth.GetAuthUser(ctx.Request().Context(), claims)

	return api.NewAPIResponse(err, response.NewAuthUserResponseDTO(result))
}
