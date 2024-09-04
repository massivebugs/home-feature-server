package middleware

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/massivebugs/home-feature-server/api/config"
)

func NewCSRFMiddleware(cfg *config.Config) echo.MiddlewareFunc {
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
