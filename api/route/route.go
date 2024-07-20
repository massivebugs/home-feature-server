package route

import (
	"database/sql"

	"github.com/labstack/echo/v4"
	"github.com/massivebugs/home-feature-server/api/config"
	"github.com/massivebugs/home-feature-server/api/handler"
)

func RegisterRoutes(e *echo.Echo, cfg *config.Config, db *sql.DB) {
	api := e.Group("/api")
	h := handler.NewAPIHandlers(db)
	registerV1Routes(api, cfg, h)
}

func registerV1Routes(e *echo.Group, cfg *config.Config, h *handler.Handlers) {
	v1 := e.Group("/v1")

	// Ping
	v1.GET("/ping", handler.CreateEchoHandlerFunc(cfg, h.PingHandler.Ping))

	// Repeat
	v1.POST("/repeat", handler.CreateEchoHandlerFunc(cfg, h.RepeatHandler.Repeat))

	// Auth
	v1.POST("/auth", handler.CreateEchoHandlerFunc(cfg, h.AuthHandler.CreateUser))
	v1.POST("/auth/token", handler.CreateEchoHandlerFunc(cfg, h.AuthHandler.CreateJWTToken))
}
