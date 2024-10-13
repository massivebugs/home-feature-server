package ping

import "context"

type Ping struct{}

func NewPing() *Ping {
	return &Ping{}
}

func (*Ping) Run(ctx context.Context) string {
	return "Pong!"
}
