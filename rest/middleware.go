package rest

import (
	"context"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/massivebugs/home-feature-server/internal/auth"
)

type CustomCtxKey string

const (
	CtxClaimsKey CustomCtxKey = "claims"
)

func NewCORSMiddleware(cfg *Config) echo.MiddlewareFunc {
	return middleware.CORSWithConfig(
		middleware.CORSConfig{
			AllowOrigins:     cfg.AllowedOrigins,
			AllowCredentials: true,
			AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		},
	)
}

func NewCSRFMiddleware(cfg *Config) echo.MiddlewareFunc {
	return middleware.CSRFWithConfig(
		middleware.CSRFConfig{
			TokenLookup:    "header:X-CSRF-Token,cookie:_csrf",
			CookiePath:     "/",
			CookieSecure:   true,
			CookieHTTPOnly: true,
			CookieSameSite: http.SameSiteStrictMode,
		},
	)
}

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
				return strings.Contains(c.Path(), "/secure/") || !(strings.Contains(c.Path(), "/auth/token") && c.Request().Method == "PUT")
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
			SigningKey:  []byte(cfg.RefreshJWTSecret),
			TokenLookup: "header:Authorization:Bearer ,cookie:" + cfg.RefreshJWTCookieName,
		},
	)
}
