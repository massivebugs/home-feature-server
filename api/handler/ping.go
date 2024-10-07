package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/massivebugs/home-feature-server/api"
	"github.com/massivebugs/home-feature-server/internal/ping"
)

type PingHandler struct {
	*api.Handler
	ping *ping.Ping
}

func NewPingHandler() *PingHandler {
	return &PingHandler{
		ping: ping.NewPing(),
	}
}

func (h *PingHandler) Ping(c echo.Context) *api.APIResponse {
	result := h.ping.Run(c.Request().Context())

	return h.CreateResponse(nil, result)
}
