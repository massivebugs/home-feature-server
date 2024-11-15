package rest

import (
	"context"
	"net/http"
	"time"

	"github.com/massivebugs/home-feature-server/db"
	"github.com/massivebugs/home-feature-server/db/queries"
	"github.com/massivebugs/home-feature-server/internal/auth"
	"github.com/massivebugs/home-feature-server/internal/repository"
	"github.com/massivebugs/home-feature-server/rest/oapi"
)

type AuthHandler struct {
	*Handler
	auth *auth.Auth
}

func NewAuthHandler(cfg *Config, db *db.Handle, querier queries.Querier) *AuthHandler {
	return &AuthHandler{
		Handler: &Handler{
			cfg: cfg,
		},
		auth: auth.NewAuth(
			db,
			repository.NewUserRepository(querier),
			repository.NewUserPasswordRepository(querier),
			repository.NewUserRefreshTokenRepository(querier),
		),
	}
}

func (h *AuthHandler) CreateUser(ctx context.Context, request oapi.CreateUserRequestObject) (oapi.CreateUserResponseObject, error) {
	err := h.auth.CreateUser(ctx, request.Body.Username, request.Body.Email, request.Body.Password)
	if err != nil {
		return nil, err
	}

	return oapi.CreateUser202Response{}, nil
}

func (h *AuthHandler) CreateJWTToken(ctx context.Context, request oapi.CreateJWTTokenRequestObject) (oapi.CreateJWTTokenResponseObject, error) {
	now := time.Now()
	result, err := h.auth.CreateJWTToken(
		ctx,
		now,
		h.cfg.AuthJWTSigningMethod,
		h.cfg.AuthJWTSecret,
		h.cfg.AuthJWTExpireSeconds,
		request.Body.Username,
		request.Body.Password,
	)
	if err != nil {
		return nil, err
	}

	cookie := http.Cookie{
		Name:     h.cfg.AuthJWTCookieName,
		Value:    result,
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
		MaxAge:   h.cfg.AuthJWTExpireSeconds,
		Expires:  now.Add(time.Second * time.Duration(h.cfg.AuthJWTExpireSeconds)),
	}

	return oapi.CreateJWTToken200Response{
		Headers: oapi.CreateJWTToken200ResponseHeaders{
			SetCookie: cookie.String(),
		},
	}, nil
}

func (h *AuthHandler) CreateJWTRefreshToken(ctx context.Context, request oapi.CreateJWTRefreshTokenRequestObject) (oapi.CreateJWTRefreshTokenResponseObject, error) {
	claims := h.GetClaims(ctx)
	now := time.Now()

	result, err := h.auth.CreateJWTRefreshToken(
		ctx,
		now,
		h.cfg.RefreshJWTSigningMethod,
		h.cfg.RefreshJWTSecret,
		h.cfg.RefreshJWTExpireSeconds,
		claims.UserID,
	)
	if err != nil {
		return nil, err
	}

	cookie := http.Cookie{
		Name:     h.cfg.RefreshJWTCookieName,
		Value:    result,
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
		MaxAge:   h.cfg.RefreshJWTExpireSeconds,
		Expires:  now.Add(time.Second * time.Duration(h.cfg.RefreshJWTExpireSeconds)),
	}

	return oapi.CreateJWTRefreshToken200Response{
		Headers: oapi.CreateJWTRefreshToken200ResponseHeaders{
			SetCookie: cookie.String(),
		},
	}, nil
}

func (h *AuthHandler) RefreshJWTToken(ctx context.Context, request oapi.RefreshJWTTokenRequestObject) (oapi.RefreshJWTTokenResponseObject, error) {
	claims := h.GetClaims(ctx)

	now := time.Now()
	result, err := h.auth.RefreshJWTToken(
		ctx,
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
		return nil, err
	}

	cookie := http.Cookie{
		Name:     h.cfg.AuthJWTCookieName,
		Value:    result,
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
		MaxAge:   h.cfg.AuthJWTExpireSeconds,
		Expires:  now.Add(time.Second * time.Duration(h.cfg.AuthJWTExpireSeconds)),
	}

	return oapi.RefreshJWTToken200Response{
		Headers: oapi.RefreshJWTToken200ResponseHeaders{
			SetCookie: cookie.String(),
		},
	}, nil
}

func (h *AuthHandler) GetUser(ctx context.Context, request oapi.GetUserRequestObject) (oapi.GetUserResponseObject, error) {
	claims := h.GetClaims(ctx)

	u, err := h.auth.GetAuthUser(ctx, claims.UserID, time.Now())
	if err != nil {
		return nil, err
	}

	res := oapi.GetUser200JSONResponse{}
	res.User.Id = u.Id
	res.User.Name = u.Name
	res.User.CreatedAt = u.CreatedAt.String()
	res.User.LoggedInAt = u.CreatedAt.String()

	return res, nil
}
