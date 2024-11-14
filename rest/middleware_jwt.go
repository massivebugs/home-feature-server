package rest

import (
	"context"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/massivebugs/home-feature-server/internal/auth"
)

type CustomCtxKey string

const (
	CtxClaimsKey CustomCtxKey = "claims"
)

func NewJWTMiddleware(cfg *Config) echo.MiddlewareFunc {
	return echojwt.WithConfig(
		echojwt.Config{
			Skipper: func(c echo.Context) bool {
				return !strings.Contains(c.Path(), "/secure/")
			},
			NewClaimsFunc: func(c echo.Context) jwt.Claims {
				return new(auth.JWTClaims)
			},
			SuccessHandler: func(c echo.Context) {
				token := c.Get("user").(*jwt.Token)
				ctxWithClaims := context.WithValue(c.Request().Context(), CtxClaimsKey, token.Claims)

				// Since a oapi-codegen Strict server only allows handlers to access the Request's context,
				// We can add the claims info directly to the request's context by making a shallow copy.
				c.SetRequest(c.Request().WithContext(ctxWithClaims))
			},
			SigningKey:  []byte(cfg.AuthJWTSecret),
			TokenLookup: "header:Authorization:Bearer ,cookie:" + cfg.AuthJWTCookieName,
		},
	)
}

func NewJWTRefreshMiddleware(cfg *Config) echo.MiddlewareFunc {
	return echojwt.WithConfig(
		echojwt.Config{
			Skipper: func(c echo.Context) bool {
				return !(strings.Contains(c.Path(), "/auth/refresh") && c.Request().Method == "PUT")
			},
			NewClaimsFunc: func(c echo.Context) jwt.Claims {
				return new(auth.JWTClaims)
			},
			SigningKey:  []byte(cfg.RefreshJWTSecret),
			TokenLookup: "header:Authorization:Bearer ,cookie:" + cfg.RefreshJWTCookieName,
		},
	)
}
