package handler

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/massivebugs/home-feature-server/api"
	"github.com/massivebugs/home-feature-server/api/response"
	"github.com/massivebugs/home-feature-server/db"
	"github.com/massivebugs/home-feature-server/db/queries"
	"github.com/massivebugs/home-feature-server/internal/auth"
	"github.com/massivebugs/home-feature-server/repository"
)

type AuthHandler struct {
	*api.Handler
	cfg  *api.Config
	auth *auth.Auth
}

func NewAuthHandler(cfg *api.Config, db *db.Handle, querier queries.Querier) *AuthHandler {
	return &AuthHandler{
		cfg: cfg,
		auth: auth.NewAuth(
			db,
			repository.NewUserDBRepository(querier),
			repository.NewUserPasswordDBRepository(querier),
			repository.NewUserRefreshTokenDBRepository(querier),
		),
	}
}

func (h *AuthHandler) CreateUser(c echo.Context) *api.APIResponse {
	req := new(auth.CreateUserDTO)

	err := h.Validate(c, req)
	if err != nil {
		return h.CreateErrorResponse(err)
	}

	err = h.auth.CreateAuthUser(c.Request().Context(), req)

	return h.CreateResponse(err, nil)
}

func (h *AuthHandler) CreateJWTToken(c echo.Context) *api.APIResponse {
	req := new(auth.UserAuthDTO)

	err := h.Validate(c, req)
	if err != nil {
		return h.CreateErrorResponse(err)
	}

	now := time.Now()
	token, refreshToken, err := h.auth.CreateJWTToken(
		c.Request().Context(),
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
		return h.CreateErrorResponse(err)
	}

	c.SetCookie(&http.Cookie{
		Name:     h.cfg.AuthJWTCookieName,
		Value:    token,
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
		MaxAge:   h.cfg.AuthJWTExpireSeconds,
		Expires:  now.Add(time.Second * time.Duration(h.cfg.AuthJWTExpireSeconds)),
	})

	c.SetCookie(&http.Cookie{
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

	return h.CreateResponse(nil, result)
}

func (h *AuthHandler) RefreshJWTToken(c echo.Context) *api.APIResponse {
	claims := h.GetTokenClaims(c)

	now := time.Now()
	token, refreshToken, err := h.auth.RefreshJWTToken(
		c.Request().Context(),
		now,
		h.cfg.AuthJWTSigningMethod,
		h.cfg.AuthJWTSecret,
		h.cfg.AuthJWTExpireSeconds,
		h.cfg.RefreshJWTSigningMethod,
		h.cfg.RefreshJWTSecret,
		h.cfg.RefreshJWTExpireSeconds,
		claims.UserID,
		claims.TokenID,
	)
	if err != nil {
		return h.CreateErrorResponse(err)
	}

	c.SetCookie(&http.Cookie{
		Name:     h.cfg.AuthJWTCookieName,
		Value:    token,
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
		MaxAge:   h.cfg.AuthJWTExpireSeconds,
		Expires:  now.Add(time.Second * time.Duration(h.cfg.AuthJWTExpireSeconds)),
	})

	c.SetCookie(&http.Cookie{
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

	return h.CreateResponse(nil, result)
}

func (h *AuthHandler) GetAuthUser(c echo.Context) *api.APIResponse {
	claims := h.GetTokenClaims(c)

	result, err := h.auth.GetAuthUser(c.Request().Context(), claims)

	return h.CreateResponse(err, response.NewAuthUserResponseDTO(result))
}
