package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/massivebugs/home-feature-server/internal/api"
	"github.com/massivebugs/home-feature-server/internal/ping"
)

type PingHandler struct {
	ping *ping.Ping
}

func NewPingHandler() *PingHandler {
	return &PingHandler{
		ping: ping.NewPing(),
	}
}

func (h *PingHandler) Ping(ctx echo.Context) error {
	result := h.ping.Run(ctx.Request().Context())

	return api.NewAPIResponse(ctx, nil, result)
}
