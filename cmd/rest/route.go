package main

// type handlers struct {
// 	PingHandler                  *rest.PingHandler
// 	RepeatHandler                *rest.RepeatHandler
// 	AuthHandler                  *rest.AuthHandler
// 	UserSystemPreferencesHandler *rest.SystemPreferencesHandler
// 	CashbunnyHandler             *rest.CashbunnyHandler
// }

// func newAPIHandlers(cfg *rest.Config, db *db.Handle) *handlers {
// 	// TODO: Move this where??
// 	querier := queries.New()

// 	return &handlers{
// 		PingHandler:                  rest.NewPingHandler(cfg),
// 		RepeatHandler:                rest.NewRepeatHandler(cfg),
// 		AuthHandler:                  rest.NewAuthHandler(cfg, db, querier),
// 		UserSystemPreferencesHandler: rest.NewSystemPreferencesHandler(cfg, db, querier),
// 		CashbunnyHandler:             rest.NewCashbunnyHandler(cfg, db, querier),
// 	}
// }

// func registerRoutes(e *echo.Echo, cfg *rest.Config, m apiMiddlewares, db *db.Handle) {
// 	api := e.Group("/api")
// 	h := newAPIHandlers(cfg, db)
// 	registerV1Routes(api, m, h)
// }

// func registerV1Routes(e *echo.Group, m apiMiddlewares, h *handlers) {
// 	v1 := e.Group("/v1", m.CSRF)

// 	// Ping
// 	v1.GET("/ping", h.PingHandler.Ping)

// 	// Repeat
// 	v1.POST("/repeat", h.RepeatHandler.Repeat)

// 	// Auth
// 	v1.POST("/auth", h.AuthHandler.CreateUser)
// 	v1.POST("/auth/token", h.AuthHandler.CreateJWTToken)
// 	v1.POST("/auth/refresh", h.AuthHandler.RefreshJWTToken, m.JWTRefresh)

// 	// Authenticated routes
// 	v1Secure := v1.Group("/secure", m.JWT)

// 	user := v1Secure.Group("/user")
// 	user.GET("", h.AuthHandler.GetAuthUser)

// 	systemPreferences := user.Group("/system_preferences")
// 	systemPreferences.GET("", h.UserSystemPreferencesHandler.GetUserSystemPreferences)
// 	systemPreferences.POST("", h.UserSystemPreferencesHandler.CreateDefaultUserSystemPreferences)
// 	systemPreferences.PUT("", h.UserSystemPreferencesHandler.UpdateDefaultUserSystemPreferences)

// 	cashbunny := v1Secure.Group("/cashbunny")
// 	cashbunny.GET("/overview", h.CashbunnyHandler.GetOverview)
// 	cashbunny.GET("/planner", h.CashbunnyHandler.GetPlan)
// 	cashbunny.GET("/planner/parameters", h.CashbunnyHandler.GetPlannerParameters)
// 	cashbunny.POST("/planner/parameters", h.CashbunnyHandler.SavePlannerParameters)
// 	cashbunny.GET("/currencies", h.CashbunnyHandler.GetCurrencies)
// 	cashbunny.GET("/user_preferences", h.CashbunnyHandler.GetUserPreferences)
// 	cashbunny.POST("/user_preferences", h.CashbunnyHandler.CreateDefaultUserPreferences)
// 	cashbunny.POST("/accounts", h.CashbunnyHandler.CreateAccount)
// 	cashbunny.GET("/accounts", h.CashbunnyHandler.ListAccounts)
// 	cashbunny.DELETE("/accounts/:id", h.CashbunnyHandler.DeleteAccount)
// 	cashbunny.POST("/transactions", h.CashbunnyHandler.CreateTransaction)
// 	cashbunny.GET("/transactions", h.CashbunnyHandler.ListTransactions)
// 	cashbunny.DELETE("/transactions/:id", h.CashbunnyHandler.DeleteTransaction)
// }
