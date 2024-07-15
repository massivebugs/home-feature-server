package route

import (
	"github.com/massivebugs/home-feature-server/internal/echo"
	"github.com/massivebugs/home-feature-server/internal/ping"
)

type Handlers struct {
	PingHandler *ping.PingHandler
	Echohandler *echo.EchoHandler
}

func NewAPIHandlers() *Handlers {
	return &Handlers{
		PingHandler: ping.NewPingHandler(),
		Echohandler: echo.NewEchoHandler(),
	}
}
