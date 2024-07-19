package route

import "github.com/massivebugs/home-feature-server/api/handler"

type Handlers struct {
	PingHandler   *handler.PingHandler
	RepeatHandler *handler.RepeatHandler
	AuthHandler   *handler.AuthHandler
}

func NewAPIHandlers() *Handlers {
	return &Handlers{
		PingHandler:   handler.NewPingHandler(),
		RepeatHandler: handler.NewRepeatHandler(),
		AuthHandler:   handler.NewAuthHandler(),
	}
}
