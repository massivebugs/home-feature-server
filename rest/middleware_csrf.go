package rest

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

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
