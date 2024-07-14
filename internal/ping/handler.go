package ping

import (
	"github.com/labstack/echo/v4"
	"github.com/massivebugs/home-server/internal/api"
)

type PingHandler struct{}

func NewPingHandler() *PingHandler {
	return &PingHandler{}
}

func (*PingHandler) Ping(ctx echo.Context) error {
	return api.NewAPIResponse(ctx, nil, "Pong!")
}
