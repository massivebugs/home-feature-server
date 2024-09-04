package handler

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/massivebugs/home-feature-server/api/config"
	"github.com/massivebugs/home-feature-server/api/response"
	"github.com/massivebugs/home-feature-server/db/service/auth_repository"
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
			auth_repository.New(),
		),
	}
}

func (h *AuthHandler) CreateUser(ctx echo.Context) *api.APIResponse {
	req := new(auth.CreateUserRequestDTO)

	if err := ctx.Bind(req); err != nil {
		return api.NewAPIResponse(err, nil)
	}

	if err := ctx.Validate(req); err != nil {
		return api.NewAPIResponse(err, nil)
	}

	err := h.auth.CreateAuthUser(ctx.Request().Context(), req)

	return api.NewAPIResponse(err, nil)
}

func (h *AuthHandler) CreateJWTToken(ctx echo.Context) *api.APIResponse {
	req := new(auth.UserAuthRequestDTO)

	if err := ctx.Bind(req); err != nil {
		return api.NewAPIResponse(err, nil)
	}

	if err := ctx.Validate(req); err != nil {
		return api.NewAPIResponse(err, nil)
	}

	now := time.Now()
	token, refreshToken, err := h.auth.CreateJWTToken(
		ctx.Request().Context(),
		now,
		h.cfg.AuthJWTSigningMethod,
		h.cfg.AuthJWTSecret,
		h.cfg.AuthJWTExpireSeconds,
		h.cfg.RefreshJWTSigningMethod,
		h.cfg.RefreshJWTSecret,
		h.cfg.RefreshJWTExpireSeconds,
		req,
	)

	if err != nil {
		return api.NewAPIResponse(err, nil)
	}

	ctx.SetCookie(&http.Cookie{
		Name:     h.cfg.AuthJWTCookieName,
		Value:    token,
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
		MaxAge:   h.cfg.AuthJWTExpireSeconds,
		Expires:  now.Add(time.Second * time.Duration(h.cfg.AuthJWTExpireSeconds)),
	})

	ctx.SetCookie(&http.Cookie{
		Name:     h.cfg.RefreshJWTCookieName,
		Value:    refreshToken,
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
		MaxAge:   h.cfg.RefreshJWTExpireSeconds,
		Expires:  now.Add(time.Second * time.Duration(h.cfg.RefreshJWTExpireSeconds)),
	})

	result := response.NewCreateJWTTokenResponseDTO(token, refreshToken)

	return api.NewAPIResponse(err, result)
}

func (h *AuthHandler) RefreshJWTToken(ctx echo.Context) *api.APIResponse {
	oldRefreshToken := ctx.Get("user").(*jwt.Token)
	claims := oldRefreshToken.Claims.(*auth.JWTClaims)

	now := time.Now()
	token, refreshToken, err := h.auth.RefreshJWTToken(
		ctx.Request().Context(),
		now,
		h.cfg.AuthJWTSigningMethod,
		h.cfg.AuthJWTSecret,
		h.cfg.AuthJWTExpireSeconds,
		h.cfg.RefreshJWTSigningMethod,
		h.cfg.RefreshJWTSecret,
		h.cfg.RefreshJWTExpireSeconds,
		claims,
	)

	ctx.SetCookie(&http.Cookie{
		Name:     h.cfg.AuthJWTCookieName,
		Value:    token,
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
		MaxAge:   h.cfg.AuthJWTExpireSeconds,
		Expires:  now.Add(time.Second * time.Duration(h.cfg.AuthJWTExpireSeconds)),
	})

	ctx.SetCookie(&http.Cookie{
		Name:     h.cfg.AuthJWTCookieName + "_refresh",
		Value:    refreshToken,
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
		MaxAge:   h.cfg.RefreshJWTExpireSeconds,
		Expires:  now.Add(time.Second * time.Duration(h.cfg.RefreshJWTExpireSeconds)),
	})

	result := response.NewCreateJWTTokenResponseDTO(token, refreshToken)

	return api.NewAPIResponse(err, result)
}

func (h *AuthHandler) GetAuthUser(ctx echo.Context) *api.APIResponse {
	token := ctx.Get("user").(*jwt.Token)
	claims := token.Claims.(*auth.JWTClaims)

	result, err := h.auth.GetAuthUser(ctx.Request().Context(), claims)

	return api.NewAPIResponse(err, response.NewAuthUserResponseDTO(result))
}
