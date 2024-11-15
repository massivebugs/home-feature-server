package handler

import (
	"context"

	"github.com/massivebugs/home-feature-server/internal/ping"
	"github.com/massivebugs/home-feature-server/rest"
	"github.com/massivebugs/home-feature-server/rest/oapi"
)

type PingHandler struct {
	*rest.Handler
	ping *ping.Ping
}

func NewPingHandler(cfg *rest.Config) *PingHandler {
	return &PingHandler{
		Handler: rest.NewHandler(cfg),
		ping:    ping.NewPing(),
	}
}

func (h *PingHandler) Ping(ctx context.Context, req oapi.PingRequestObject) (oapi.PingResponseObject, error) {
	result := h.ping.Run(ctx)

	return oapi.Ping200JSONResponse{Message: result}, nil
}
