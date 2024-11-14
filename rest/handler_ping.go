package rest

import (
	"context"

	"github.com/massivebugs/home-feature-server/internal/ping"
)

type PingHandler struct {
	*Handler
	ping *ping.Ping
}

func NewPingHandler(cfg *Config) *PingHandler {
	return &PingHandler{
		Handler: &Handler{
			cfg: cfg,
		},
		ping: ping.NewPing(),
	}
}

func (h *PingHandler) Ping(ctx context.Context, req PingRequestObject) (PingResponseObject, error) {
	result := h.ping.Run(ctx)

	return Ping200JSONResponse{Message: result}, nil
}
