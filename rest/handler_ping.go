package rest

import (
	"context"

	"github.com/massivebugs/home-feature-server/internal/ping"
	"github.com/massivebugs/home-feature-server/rest/oapi"
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

func (h *PingHandler) Ping(ctx context.Context, req oapi.PingRequestObject) (oapi.PingResponseObject, error) {
	result := h.ping.Run(ctx)

	return oapi.Ping200JSONResponse{Message: result}, nil
}
