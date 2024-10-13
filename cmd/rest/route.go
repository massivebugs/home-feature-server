package main

import (
	"github.com/labstack/echo/v4"
	"github.com/massivebugs/home-feature-server/db"
	"github.com/massivebugs/home-feature-server/db/queries"
	"github.com/massivebugs/home-feature-server/rest"
)

type handlers struct {
	PingHandler                  *rest.PingHandler
	RepeatHandler                *rest.RepeatHandler
	AuthHandler                  *rest.AuthHandler
	UserSystemPreferencesHandler *rest.SystemPreferencesHandler
	CashbunnyHandler             *rest.CashbunnyHandler
}

func newAPIHandlers(cfg *rest.Config, db *db.Handle) *handlers {
	// TODO: Move this where??
	querier := queries.New()

	return &handlers{
		PingHandler:                  rest.NewPingHandler(),
		RepeatHandler:                rest.NewRepeatHandler(),
		AuthHandler:                  rest.NewAuthHandler(cfg, db, querier),
		UserSystemPreferencesHandler: rest.NewSystemPreferencesHandler(cfg, db, querier),
		CashbunnyHandler:             rest.NewCashbunnyHandler(db, querier),
	}
}

func registerRoutes(e *echo.Echo, cfg *rest.Config, m apiMiddlewares, db *db.Handle) {
	api := e.Group("/api")
	h := newAPIHandlers(cfg, db)
	registerV1Routes(api, cfg, m, h)
}

func registerV1Routes(e *echo.Group, cfg *rest.Config, m apiMiddlewares, h *handlers) {
	v1 := e.Group("/v1", m.CSRF)

	// Ping
	v1.GET("/ping", rest.CreateEchoHandlerFunc(cfg, h.PingHandler.Ping))

	// Repeat
	v1.POST("/repeat", rest.CreateEchoHandlerFunc(cfg, h.RepeatHandler.Repeat))

	// Auth
	v1.POST("/auth", rest.CreateEchoHandlerFunc(cfg, h.AuthHandler.CreateUser))
	v1.POST("/auth/token", rest.CreateEchoHandlerFunc(cfg, h.AuthHandler.CreateJWTToken))
	v1.POST("/auth/refresh", rest.CreateEchoHandlerFunc(cfg, h.AuthHandler.RefreshJWTToken), m.JWTRefresh)

	// Authenticated routes
	v1Secure := v1.Group("/secure", m.JWT)

	user := v1Secure.Group("/user")
	user.GET("", rest.CreateEchoHandlerFunc(cfg, h.AuthHandler.GetAuthUser))

	systemPreferences := user.Group("/system_preferences")
	systemPreferences.GET("", rest.CreateEchoHandlerFunc(cfg, h.UserSystemPreferencesHandler.GetUserSystemPreferences))
	systemPreferences.POST("", rest.CreateEchoHandlerFunc(cfg, h.UserSystemPreferencesHandler.CreateDefaultUserSystemPreferences))
	systemPreferences.PUT("", rest.CreateEchoHandlerFunc(cfg, h.UserSystemPreferencesHandler.UpdateDefaultUserSystemPreferences))

	cashbunny := v1Secure.Group("/cashbunny")
	cashbunny.GET("/overview", rest.CreateEchoHandlerFunc(cfg, h.CashbunnyHandler.GetOverview))
	cashbunny.GET("/planner", rest.CreateEchoHandlerFunc(cfg, h.CashbunnyHandler.GetPlan))
	cashbunny.GET("/planner/parameters", rest.CreateEchoHandlerFunc(cfg, h.CashbunnyHandler.GetPlannerParameters))
	cashbunny.POST("/planner/parameters", rest.CreateEchoHandlerFunc(cfg, h.CashbunnyHandler.SavePlannerParameters))
	cashbunny.GET("/currencies", rest.CreateEchoHandlerFunc(cfg, h.CashbunnyHandler.GetCurrencies))
	cashbunny.GET("/user_preferences", rest.CreateEchoHandlerFunc(cfg, h.CashbunnyHandler.GetUserPreferences))
	cashbunny.POST("/user_preferences", rest.CreateEchoHandlerFunc(cfg, h.CashbunnyHandler.CreateDefaultUserPreferences))
	cashbunny.POST("/accounts", rest.CreateEchoHandlerFunc(cfg, h.CashbunnyHandler.CreateAccount))
	cashbunny.GET("/accounts", rest.CreateEchoHandlerFunc(cfg, h.CashbunnyHandler.ListAccounts))
	cashbunny.DELETE("/accounts/:id", rest.CreateEchoHandlerFunc(cfg, h.CashbunnyHandler.DeleteAccount))
	cashbunny.POST("/transactions", rest.CreateEchoHandlerFunc(cfg, h.CashbunnyHandler.CreateTransaction))
	cashbunny.GET("/transactions", rest.CreateEchoHandlerFunc(cfg, h.CashbunnyHandler.ListTransactions))
	cashbunny.DELETE("/transactions/:id", rest.CreateEchoHandlerFunc(cfg, h.CashbunnyHandler.DeleteTransaction))
}
