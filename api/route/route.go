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

	// Echo
	v1.POST("/echo", h.Echohandler.Echo)
}
