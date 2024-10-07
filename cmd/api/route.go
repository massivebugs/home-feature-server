package main

import (
	"github.com/labstack/echo/v4"
	"github.com/massivebugs/home-feature-server/api"
	"github.com/massivebugs/home-feature-server/api/handler"
	"github.com/massivebugs/home-feature-server/db"
	"github.com/massivebugs/home-feature-server/db/queries"
)

type handlers struct {
	PingHandler                  *handler.PingHandler
	RepeatHandler                *handler.RepeatHandler
	AuthHandler                  *handler.AuthHandler
	UserSystemPreferencesHandler *handler.SystemPreferencesHandler
	CashbunnyHandler             *handler.CashbunnyHandler
}

func newAPIHandlers(cfg *api.Config, db *db.Handle) *handlers {
	// TODO: Move this where??
	querier := queries.New()

	return &handlers{
		PingHandler:                  handler.NewPingHandler(),
		RepeatHandler:                handler.NewRepeatHandler(),
		AuthHandler:                  handler.NewAuthHandler(cfg, db, querier),
		UserSystemPreferencesHandler: handler.NewSystemPreferencesHandler(cfg, db, querier),
		CashbunnyHandler:             handler.NewCashbunnyHandler(db, querier),
	}
}

func registerRoutes(e *echo.Echo, cfg *api.Config, m apiMiddlewares, db *db.Handle) {
	api := e.Group("/api")
	h := newAPIHandlers(cfg, db)
	registerV1Routes(api, cfg, m, h)
}

func registerV1Routes(e *echo.Group, cfg *api.Config, m apiMiddlewares, h *handlers) {
	v1 := e.Group("/v1", m.CSRF)

	// Ping
	v1.GET("/ping", api.CreateEchoHandlerFunc(cfg, h.PingHandler.Ping))

	// Repeat
	v1.POST("/repeat", api.CreateEchoHandlerFunc(cfg, h.RepeatHandler.Repeat))

	// Auth
	v1.POST("/auth", api.CreateEchoHandlerFunc(cfg, h.AuthHandler.CreateUser))
	v1.POST("/auth/token", api.CreateEchoHandlerFunc(cfg, h.AuthHandler.CreateJWTToken))
	v1.POST("/auth/refresh", api.CreateEchoHandlerFunc(cfg, h.AuthHandler.RefreshJWTToken), m.JWTRefresh)

	// Authenticated routes
	v1Secure := v1.Group("/secure", m.JWT)

	user := v1Secure.Group("/user")
	user.GET("", api.CreateEchoHandlerFunc(cfg, h.AuthHandler.GetAuthUser))

	systemPreferences := user.Group("/system_preferences")
	systemPreferences.GET("", api.CreateEchoHandlerFunc(cfg, h.UserSystemPreferencesHandler.GetUserSystemPreferences))
	systemPreferences.POST("", api.CreateEchoHandlerFunc(cfg, h.UserSystemPreferencesHandler.CreateDefaultUserSystemPreferences))
	systemPreferences.PUT("", api.CreateEchoHandlerFunc(cfg, h.UserSystemPreferencesHandler.UpdateDefaultUserSystemPreferences))

	cashbunny := v1Secure.Group("/cashbunny")
	cashbunny.GET("/overview", api.CreateEchoHandlerFunc(cfg, h.CashbunnyHandler.GetOverview))
	cashbunny.GET("/planner", api.CreateEchoHandlerFunc(cfg, h.CashbunnyHandler.GetPlan))
	cashbunny.GET("/planner/parameters", api.CreateEchoHandlerFunc(cfg, h.CashbunnyHandler.GetPlannerParameters))
	cashbunny.POST("/planner/parameters", api.CreateEchoHandlerFunc(cfg, h.CashbunnyHandler.SavePlannerParameters))
	cashbunny.GET("/currencies", api.CreateEchoHandlerFunc(cfg, h.CashbunnyHandler.GetCurrencies))
	cashbunny.GET("/user_preferences", api.CreateEchoHandlerFunc(cfg, h.CashbunnyHandler.GetUserPreferences))
	cashbunny.POST("/user_preferences", api.CreateEchoHandlerFunc(cfg, h.CashbunnyHandler.CreateDefaultUserPreferences))
	cashbunny.POST("/accounts", api.CreateEchoHandlerFunc(cfg, h.CashbunnyHandler.CreateAccount))
	cashbunny.GET("/accounts", api.CreateEchoHandlerFunc(cfg, h.CashbunnyHandler.ListAccounts))
	cashbunny.DELETE("/accounts/:id", api.CreateEchoHandlerFunc(cfg, h.CashbunnyHandler.DeleteAccount))
	cashbunny.POST("/transactions", api.CreateEchoHandlerFunc(cfg, h.CashbunnyHandler.CreateTransaction))
	cashbunny.GET("/transactions", api.CreateEchoHandlerFunc(cfg, h.CashbunnyHandler.ListTransactions))
	cashbunny.DELETE("/transactions/:id", api.CreateEchoHandlerFunc(cfg, h.CashbunnyHandler.DeleteTransaction))
}
