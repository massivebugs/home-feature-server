package repeat

import (
	"context"
)

type Repeat struct{}

func NewRepeat() *Repeat {
	return &Repeat{}
}

func (*Repeat) Run(ctx context.Context, message string) string {
	return message
}
