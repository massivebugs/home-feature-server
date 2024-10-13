package http

import (
	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/massivebugs/home-feature-server/internal/auth"
)

func NewJWTMiddleware(cfg *Config) echo.MiddlewareFunc {
	return echojwt.WithConfig(
		echojwt.Config{
			NewClaimsFunc: func(c echo.Context) jwt.Claims {
				return new(auth.JWTClaims)
			},
			SigningKey:  []byte(cfg.AuthJWTSecret),
			TokenLookup: "header:Authorization:Bearer ,cookie:" + cfg.AuthJWTCookieName,
		},
	)
}

func NewJWTRefreshMiddleware(cfg *Config) echo.MiddlewareFunc {
	return echojwt.WithConfig(
		echojwt.Config{
			NewClaimsFunc: func(c echo.Context) jwt.Claims {
				return new(auth.JWTClaims)
			},
			SigningKey:  []byte(cfg.RefreshJWTSecret),
			TokenLookup: "header:Authorization:Bearer ,cookie:" + cfg.RefreshJWTCookieName,
		},
	)
}
