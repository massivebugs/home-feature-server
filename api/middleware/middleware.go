package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/massivebugs/home-feature-server/api/config"
)

type APIMiddleware struct {
	CSRF       echo.MiddlewareFunc
	CORS       echo.MiddlewareFunc
	JWT        echo.MiddlewareFunc
	JWTRefresh echo.MiddlewareFunc
}

func NewAPIMiddleware(cfg *config.Config) APIMiddleware {
	return APIMiddleware{
		CSRF:       NewCSRFMiddleware(cfg),
		CORS:       NewCORSMiddleware(cfg),
		JWT:        NewJWTMiddleware(cfg),
		JWTRefresh: NewJWTRefreshMiddleware(cfg),
	}
}
