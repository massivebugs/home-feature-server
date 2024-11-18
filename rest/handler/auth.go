package handler

import (
	"context"
	"net/http"
	"time"

	"github.com/massivebugs/home-feature-server/db"
	"github.com/massivebugs/home-feature-server/db/queries"
	"github.com/massivebugs/home-feature-server/internal/auth"
	"github.com/massivebugs/home-feature-server/rest"
	"github.com/massivebugs/home-feature-server/rest/oapi"
)

type AuthHandler struct {
	*rest.Handler
	auth *auth.Auth
}

func NewAuthHandler(cfg *rest.Config, db *db.Handle, querier queries.Querier) *AuthHandler {
	return &AuthHandler{
		Handler: rest.NewHandler(cfg),
		auth: auth.NewAuth(
			db,
			auth.NewUserRepository(querier),
			auth.NewUserPasswordRepository(querier),
			auth.NewUserRefreshTokenRepository(querier),
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
		h.Config.AuthJWTSigningMethod,
		h.Config.AuthJWTSecret,
		h.Config.AuthJWTExpireSeconds,
		request.Body.Username,
		request.Body.Password,
	)
	if err != nil {
		return nil, err
	}

	cookie := http.Cookie{
		Name:     h.Config.AuthJWTCookieName,
		Value:    result,
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
		MaxAge:   h.Config.AuthJWTExpireSeconds,
		Expires:  now.Add(time.Second * time.Duration(h.Config.AuthJWTExpireSeconds)),
	}

	return oapi.CreateJWTToken200Response{
		Headers: oapi.CreateJWTToken200ResponseHeaders{
			SetCookie: cookie.String(),
		},
	}, nil
}

func (h *AuthHandler) DeleteJWTToken(ctx context.Context, request oapi.DeleteJWTTokenRequestObject) (oapi.DeleteJWTTokenResponseObject, error) {
	cookie := http.Cookie{
		Name:     h.Config.AuthJWTCookieName,
		Value:    "", // Empty token value for good measure
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
		MaxAge:   0,
		Expires:  time.Unix(0, 0),
	}

	return oapi.DeleteJWTToken204Response{
		Headers: oapi.DeleteJWTToken204ResponseHeaders{
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
		h.Config.RefreshJWTSigningMethod,
		h.Config.RefreshJWTSecret,
		h.Config.RefreshJWTExpireSeconds,
		claims.UserID,
	)
	if err != nil {
		return nil, err
	}

	cookie := http.Cookie{
		Name:     h.Config.RefreshJWTCookieName,
		Value:    result,
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
		MaxAge:   h.Config.RefreshJWTExpireSeconds,
		Expires:  now.Add(time.Second * time.Duration(h.Config.RefreshJWTExpireSeconds)),
	}

	return oapi.CreateJWTRefreshToken200Response{
		Headers: oapi.CreateJWTRefreshToken200ResponseHeaders{
			SetCookie: cookie.String(),
		},
	}, nil
}

func (h *AuthHandler) DeleteJWTRefreshToken(ctx context.Context, request oapi.DeleteJWTRefreshTokenRequestObject) (oapi.DeleteJWTRefreshTokenResponseObject, error) {
	claims := h.GetClaims(ctx)

	err := h.auth.DeleteJWTRefreshToken(ctx, claims.UserID, claims.TokenID)
	if err != nil {
		return nil, err
	}

	cookie := http.Cookie{
		Name:     h.Config.RefreshJWTCookieName,
		Value:    "", // Empty token value for good measure
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
		MaxAge:   0,
		Expires:  time.Unix(0, 0),
	}

	return oapi.DeleteJWTRefreshToken204Response{
		Headers: oapi.DeleteJWTRefreshToken204ResponseHeaders{
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
		h.Config.AuthJWTSigningMethod,
		h.Config.AuthJWTSecret,
		h.Config.AuthJWTExpireSeconds,
		h.Config.RefreshJWTSigningMethod,
		h.Config.RefreshJWTSecret,
		h.Config.RefreshJWTExpireSeconds,
		claims.UserID,
		claims.TokenID,
	)
	if err != nil {
		return nil, err
	}

	cookie := http.Cookie{
		Name:     h.Config.AuthJWTCookieName,
		Value:    result,
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
		MaxAge:   h.Config.AuthJWTExpireSeconds,
		Expires:  now.Add(time.Second * time.Duration(h.Config.AuthJWTExpireSeconds)),
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
	res.User.CreatedAt = u.CreatedAt.Format(h.Config.APIDateTimeFormat)
	res.User.LoggedInAt = u.CreatedAt.Format(h.Config.APIDateTimeFormat)

	return res, nil
}
