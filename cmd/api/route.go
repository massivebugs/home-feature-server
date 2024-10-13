package main

import (
	"github.com/labstack/echo/v4"
	"github.com/massivebugs/home-feature-server/db"
	"github.com/massivebugs/home-feature-server/db/queries"
	"github.com/massivebugs/home-feature-server/http"
)

type handlers struct {
	PingHandler                  *http.PingHandler
	RepeatHandler                *http.RepeatHandler
	AuthHandler                  *http.AuthHandler
	UserSystemPreferencesHandler *http.SystemPreferencesHandler
	CashbunnyHandler             *http.CashbunnyHandler
}

func newAPIHandlers(cfg *http.Config, db *db.Handle) *handlers {
	// TODO: Move this where??
	querier := queries.New()

	return &handlers{
		PingHandler:                  http.NewPingHandler(),
		RepeatHandler:                http.NewRepeatHandler(),
		AuthHandler:                  http.NewAuthHandler(cfg, db, querier),
		UserSystemPreferencesHandler: http.NewSystemPreferencesHandler(cfg, db, querier),
		CashbunnyHandler:             http.NewCashbunnyHandler(db, querier),
	}
}

func registerRoutes(e *echo.Echo, cfg *http.Config, m apiMiddlewares, db *db.Handle) {
	api := e.Group("/api")
	h := newAPIHandlers(cfg, db)
	registerV1Routes(api, cfg, m, h)
}

func registerV1Routes(e *echo.Group, cfg *http.Config, m apiMiddlewares, h *handlers) {
	v1 := e.Group("/v1", m.CSRF)

	// Ping
	v1.GET("/ping", http.CreateEchoHandlerFunc(cfg, h.PingHandler.Ping))

	// Repeat
	v1.POST("/repeat", http.CreateEchoHandlerFunc(cfg, h.RepeatHandler.Repeat))

	// Auth
	v1.POST("/auth", http.CreateEchoHandlerFunc(cfg, h.AuthHandler.CreateUser))
	v1.POST("/auth/token", http.CreateEchoHandlerFunc(cfg, h.AuthHandler.CreateJWTToken))
	v1.POST("/auth/refresh", http.CreateEchoHandlerFunc(cfg, h.AuthHandler.RefreshJWTToken), m.JWTRefresh)

	// Authenticated routes
	v1Secure := v1.Group("/secure", m.JWT)

	user := v1Secure.Group("/user")
	user.GET("", http.CreateEchoHandlerFunc(cfg, h.AuthHandler.GetAuthUser))

	systemPreferences := user.Group("/system_preferences")
	systemPreferences.GET("", http.CreateEchoHandlerFunc(cfg, h.UserSystemPreferencesHandler.GetUserSystemPreferences))
	systemPreferences.POST("", http.CreateEchoHandlerFunc(cfg, h.UserSystemPreferencesHandler.CreateDefaultUserSystemPreferences))
	systemPreferences.PUT("", http.CreateEchoHandlerFunc(cfg, h.UserSystemPreferencesHandler.UpdateDefaultUserSystemPreferences))

	cashbunny := v1Secure.Group("/cashbunny")
	cashbunny.GET("/overview", http.CreateEchoHandlerFunc(cfg, h.CashbunnyHandler.GetOverview))
	cashbunny.GET("/planner", http.CreateEchoHandlerFunc(cfg, h.CashbunnyHandler.GetPlan))
	cashbunny.GET("/planner/parameters", http.CreateEchoHandlerFunc(cfg, h.CashbunnyHandler.GetPlannerParameters))
	cashbunny.POST("/planner/parameters", http.CreateEchoHandlerFunc(cfg, h.CashbunnyHandler.SavePlannerParameters))
	cashbunny.GET("/currencies", http.CreateEchoHandlerFunc(cfg, h.CashbunnyHandler.GetCurrencies))
	cashbunny.GET("/user_preferences", http.CreateEchoHandlerFunc(cfg, h.CashbunnyHandler.GetUserPreferences))
	cashbunny.POST("/user_preferences", http.CreateEchoHandlerFunc(cfg, h.CashbunnyHandler.CreateDefaultUserPreferences))
	cashbunny.POST("/accounts", http.CreateEchoHandlerFunc(cfg, h.CashbunnyHandler.CreateAccount))
	cashbunny.GET("/accounts", http.CreateEchoHandlerFunc(cfg, h.CashbunnyHandler.ListAccounts))
	cashbunny.DELETE("/accounts/:id", http.CreateEchoHandlerFunc(cfg, h.CashbunnyHandler.DeleteAccount))
	cashbunny.POST("/transactions", http.CreateEchoHandlerFunc(cfg, h.CashbunnyHandler.CreateTransaction))
	cashbunny.GET("/transactions", http.CreateEchoHandlerFunc(cfg, h.CashbunnyHandler.ListTransactions))
	cashbunny.DELETE("/transactions/:id", http.CreateEchoHandlerFunc(cfg, h.CashbunnyHandler.DeleteTransaction))
}
