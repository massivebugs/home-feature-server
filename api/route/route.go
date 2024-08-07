package route

import (
	"database/sql"

	"github.com/labstack/echo/v4"
	"github.com/massivebugs/home-feature-server/api/config"
	"github.com/massivebugs/home-feature-server/api/handler"
)

func RegisterRoutes(e *echo.Echo, cfg *config.Config, jwtMiddleware echo.MiddlewareFunc, db *sql.DB) {
	api := e.Group("/api")
	h := handler.NewAPIHandlers(cfg, db)
	registerV1Routes(api, cfg, jwtMiddleware, h)
}

func registerV1Routes(e *echo.Group, cfg *config.Config, jwtMiddleware echo.MiddlewareFunc, h *handler.Handlers) {
	v1 := e.Group("/v1")

	// Ping
	v1.GET("/ping", handler.CreateEchoHandlerFunc(cfg, h.PingHandler.Ping))

	// Repeat
	v1.POST("/repeat", handler.CreateEchoHandlerFunc(cfg, h.RepeatHandler.Repeat))

	// Auth
	v1.POST("/auth", handler.CreateEchoHandlerFunc(cfg, h.AuthHandler.CreateUser))
	v1.POST("/auth/token", handler.CreateEchoHandlerFunc(cfg, h.AuthHandler.CreateJWTToken))

	// Authenticated routes
	v1Secure := v1.Group("/secure")
	v1Secure.Use(jwtMiddleware)
	v1Secure.GET("/auth", handler.CreateEchoHandlerFunc(cfg, h.AuthHandler.GetAuthUser))

	v1Secure.POST("/cashbunny/account_categories", handler.CreateEchoHandlerFunc(cfg, h.CashbunnyHandler.CreateAccountCategory))
	v1Secure.GET("/cashbunny/account_categories", handler.CreateEchoHandlerFunc(cfg, h.CashbunnyHandler.ListAccountCategories))
	v1Secure.POST("/cashbunny/accounts", handler.CreateEchoHandlerFunc(cfg, h.CashbunnyHandler.CreateAccount))
	v1Secure.GET("/cashbunny/accounts", handler.CreateEchoHandlerFunc(cfg, h.CashbunnyHandler.ListAccounts))
	v1Secure.DELETE("/cashbunny/accounts/:id", handler.CreateEchoHandlerFunc(cfg, h.CashbunnyHandler.DeleteAccount))
}
