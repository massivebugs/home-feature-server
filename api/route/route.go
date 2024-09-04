package route

import (
	"database/sql"

	"github.com/labstack/echo/v4"
	"github.com/massivebugs/home-feature-server/api/config"
	"github.com/massivebugs/home-feature-server/api/handler"
	"github.com/massivebugs/home-feature-server/api/middleware"
)

func RegisterRoutes(e *echo.Echo, cfg *config.Config, m middleware.APIMiddleware, db *sql.DB) {
	api := e.Group("/api")
	h := handler.NewAPIHandlers(cfg, db)
	registerV1Routes(api, cfg, m, h)
}

func registerV1Routes(e *echo.Group, cfg *config.Config, m middleware.APIMiddleware, h *handler.Handlers) {
	v1 := e.Group("/v1", m.CSRF)

	// Ping
	v1.GET("/ping", handler.CreateEchoHandlerFunc(cfg, h.PingHandler.Ping))

	// Repeat
	v1.POST("/repeat", handler.CreateEchoHandlerFunc(cfg, h.RepeatHandler.Repeat))

	// Auth
	v1.POST("/auth", handler.CreateEchoHandlerFunc(cfg, h.AuthHandler.CreateUser))
	v1.POST("/auth/token", handler.CreateEchoHandlerFunc(cfg, h.AuthHandler.CreateJWTToken))
	v1.POST("/auth/refresh", handler.CreateEchoHandlerFunc(cfg, h.AuthHandler.RefreshJWTToken), m.JWTRefresh)

	// Authenticated routes
	v1Secure := v1.Group("/secure", m.JWT)
	v1Secure.GET("/auth", handler.CreateEchoHandlerFunc(cfg, h.AuthHandler.GetAuthUser))

	cashbunny := v1Secure.Group("/cashbunny")
	cashbunny.GET("/overview", handler.CreateEchoHandlerFunc(cfg, h.CashbunnyHandler.GetOverview))
	cashbunny.GET("/currencies", handler.CreateEchoHandlerFunc(cfg, h.CashbunnyHandler.GetCurrencies))
	cashbunny.GET("/user_preferences", handler.CreateEchoHandlerFunc(cfg, h.CashbunnyHandler.GetUserPreferences))
	cashbunny.POST("/user_preferences", handler.CreateEchoHandlerFunc(cfg, h.CashbunnyHandler.CreateDefaultUserPreferences))
	cashbunny.POST("/accounts", handler.CreateEchoHandlerFunc(cfg, h.CashbunnyHandler.CreateAccount))
	cashbunny.GET("/accounts", handler.CreateEchoHandlerFunc(cfg, h.CashbunnyHandler.ListAccounts))
	cashbunny.DELETE("/accounts/:id", handler.CreateEchoHandlerFunc(cfg, h.CashbunnyHandler.DeleteAccount))
	cashbunny.POST("/transactions", handler.CreateEchoHandlerFunc(cfg, h.CashbunnyHandler.CreateTransaction))
	cashbunny.GET("/transactions", handler.CreateEchoHandlerFunc(cfg, h.CashbunnyHandler.ListTransactions))
	cashbunny.DELETE("/transactions/:id", handler.CreateEchoHandlerFunc(cfg, h.CashbunnyHandler.DeleteTransaction))
}
