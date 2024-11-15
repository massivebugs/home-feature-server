package rest

import (
	"context"
	"net/http"
	"time"

	"github.com/massivebugs/home-feature-server/db"
	"github.com/massivebugs/home-feature-server/db/queries"
	"github.com/massivebugs/home-feature-server/internal/app"
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

// func (h *AuthHandler) CreateAuthUser(ctx context.Context, req CreateAuthUserRequestObject) (CreateAuthUserResponseObject, error) {
// 	err := h.auth.CreateAuthUser(ctx, req.Body.Username, req.Body.Password)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return CreateAuthUser200Response{}, nil
// }

func (h *AuthHandler) CreateJWTToken(ctx context.Context, req oapi.CreateJWTTokenRequestObject) (oapi.CreateJWTTokenResponseObject, error) {
	now := time.Now()
	result, err := h.auth.CreateJWTToken(
		ctx,
		now,
		h.cfg.AuthJWTSigningMethod,
		h.cfg.AuthJWTSecret,
		h.cfg.AuthJWTExpireSeconds,
		h.cfg.RefreshJWTSigningMethod,
		h.cfg.RefreshJWTSecret,
		h.cfg.RefreshJWTExpireSeconds,
		req.Body.Username,
		req.Body.Password,
	)
	if err != nil {
		if appErr, ok := err.(*app.AppError); ok {
			return oapi.CreateJWTToken401JSONResponse(NewErrorFromAppError(appErr)), nil
		}
		return nil, err
	}

	tokenCookie := http.Cookie{
		Name:     h.cfg.AuthJWTCookieName,
		Value:    result.Token,
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
		MaxAge:   h.cfg.AuthJWTExpireSeconds,
		Expires:  now.Add(time.Second * time.Duration(h.cfg.AuthJWTExpireSeconds)),
	}

	// refreshTokenCookie := http.Cookie{
	// 	Name:     h.cfg.RefreshJWTCookieName,
	// 	Value:    result.RefreshToken,
	// 	Path:     "/",
	// 	HttpOnly: true,
	// 	Secure:   true,
	// 	SameSite: http.SameSiteLaxMode,
	// 	MaxAge:   h.cfg.RefreshJWTExpireSeconds,
	// 	Expires:  now.Add(time.Second * time.Duration(h.cfg.RefreshJWTExpireSeconds)),
	// }

	return oapi.CreateJWTToken200Response{
		Headers: oapi.CreateJWTToken200ResponseHeaders{
			SetCookie: tokenCookie.String(),
		},
	}, nil
	// return h.CreateResponse(c, nil, result)
}

// func (h *AuthHandler) CreateJWTRefreshToken(ctx context.Context, req CreateJWTRefreshTokenRequestObject) (CreateJWTRefreshTokenResponseObject, error) {
// 	now := time.Now()
// 	result, err := h.auth.CreateJWTToken(
// 		ctx,
// 		now,
// 		h.cfg.AuthJWTSigningMethod,
// 		h.cfg.AuthJWTSecret,
// 		h.cfg.AuthJWTExpireSeconds,
// 		h.cfg.RefreshJWTSigningMethod,
// 		h.cfg.RefreshJWTSecret,
// 		h.cfg.RefreshJWTExpireSeconds,
// 		req.Body.Username,
// 		req.Body.Password,
// 	)
// 	if err != nil {
// 		return nil, err
// 	}

// 	tokenCookie := http.Cookie{
// 		Name:     h.cfg.AuthJWTCookieName,
// 		Value:    result.Token,
// 		Path:     "/",
// 		HttpOnly: true,
// 		Secure:   true,
// 		SameSite: http.SameSiteLaxMode,
// 		MaxAge:   h.cfg.AuthJWTExpireSeconds,
// 		Expires:  now.Add(time.Second * time.Duration(h.cfg.AuthJWTExpireSeconds)),
// 	}

// 	refreshTokenCookie := http.Cookie{
// 		Name:     h.cfg.RefreshJWTCookieName,
// 		Value:    result.RefreshToken,
// 		Path:     "/",
// 		HttpOnly: true,
// 		Secure:   true,
// 		SameSite: http.SameSiteLaxMode,
// 		MaxAge:   h.cfg.RefreshJWTExpireSeconds,
// 		Expires:  now.Add(time.Second * time.Duration(h.cfg.RefreshJWTExpireSeconds)),
// 	}

// 	return CreateJWTToken200Response{
// 		Headers: CreateJWTToken200ResponseHeaders{
// 			SetCookie: tokenCookie.String(),
// 		},
// 	}, nil
// 	// return h.CreateResponse(c, nil, result)
// }

// func (h *AuthHandler) RefreshJWTToken(c echo.Context) error {
// 	claims := h.GetTokenClaims(c)

// 	now := time.Now()
// 	result, err := h.auth.RefreshJWTToken(
// 		c.Request().Context(),
// 		now,
// 		h.cfg.AuthJWTSigningMethod,
// 		h.cfg.AuthJWTSecret,
// 		h.cfg.AuthJWTExpireSeconds,
// 		h.cfg.RefreshJWTSigningMethod,
// 		h.cfg.RefreshJWTSecret,
// 		h.cfg.RefreshJWTExpireSeconds,
// 		claims.UserID,
// 		claims.TokenID,
// 	)
// 	if err != nil {
// 		return h.CreateErrorResponse(c, err)
// 	}

// 	c.SetCookie(&http.Cookie{
// 		Name:     h.cfg.AuthJWTCookieName,
// 		Value:    result.Token,
// 		Path:     "/",
// 		HttpOnly: true,
// 		Secure:   true,
// 		SameSite: http.SameSiteLaxMode,
// 		MaxAge:   h.cfg.AuthJWTExpireSeconds,
// 		Expires:  now.Add(time.Second * time.Duration(h.cfg.AuthJWTExpireSeconds)),
// 	})

// 	c.SetCookie(&http.Cookie{
// 		Name:     h.cfg.AuthJWTCookieName + "_refresh",
// 		Value:    result.RefreshToken,
// 		Path:     "/",
// 		HttpOnly: true,
// 		Secure:   true,
// 		SameSite: http.SameSiteLaxMode,
// 		MaxAge:   h.cfg.RefreshJWTExpireSeconds,
// 		Expires:  now.Add(time.Second * time.Duration(h.cfg.RefreshJWTExpireSeconds)),
// 	})

// 	return h.CreateResponse(c, nil, result)
// }

func (h *AuthHandler) GetAuthUser(ctx context.Context, req oapi.GetAuthUserRequestObject) (oapi.GetAuthUserResponseObject, error) {
	claims := h.GetClaims(ctx)

	u, err := h.auth.GetAuthUser(ctx, claims.UserID, time.Now())
	if err != nil {
		return nil, err
	}

	res := oapi.GetAuthUser200JSONResponse{}
	res.User.Id = u.Id
	res.User.Name = u.Name
	res.User.CreatedAt = u.CreatedAt.String()
	res.User.LoggedInAt = u.CreatedAt.String()

	return res, nil
}
