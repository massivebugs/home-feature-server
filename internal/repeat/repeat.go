package repeat

import (
	"context"

	"github.com/massivebugs/home-feature-server/api/dto"
)

type Repeat struct{}

func NewRepeat() *Repeat {
	return &Repeat{}
}

func (*Repeat) Run(ctx context.Context, req *dto.RepeatRequestDTO) string {
	return req.Message
}
