package route

import "github.com/labstack/echo/v4"

func RegisterRoutes(e *echo.Echo) {
	api := e.Group("/api")
	h := NewAPIHandlers()
	registerV1Routes(api, h)
}

func registerV1Routes(e *echo.Group, h *Handlers) {
	v1 := e.Group("/v1")

	// Ping
	v1.GET("/ping", h.PingHandler.Ping)

	// Repeat
	v1.POST("/repeat", h.RepeatHandler.Repeat)

	// Auth
	v1.POST("/auth", h.AuthHandler.CreateUser)
	v1.POST("/auth/login", h.AuthHandler.LogIn)
}
