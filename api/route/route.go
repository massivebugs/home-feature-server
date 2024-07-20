package route

import (
	"database/sql"

	"github.com/labstack/echo/v4"
	"github.com/massivebugs/home-feature-server/api/handler"
)

func RegisterRoutes(e *echo.Echo, db *sql.DB) {
	api := e.Group("/api")
	h := handler.NewAPIHandlers(db)
	registerV1Routes(api, h)
}

func registerV1Routes(e *echo.Group, h *handler.Handlers) {
	v1 := e.Group("/v1")

	// Ping
	v1.GET("/ping", h.PingHandler.Ping)

	// Repeat
	v1.POST("/repeat", h.RepeatHandler.Repeat)

	// Auth
	v1.POST("/auth", h.AuthHandler.CreateUser)
	v1.POST("/auth/login", h.AuthHandler.LogIn)
}
