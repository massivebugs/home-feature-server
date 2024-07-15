package echo

import (
	"github.com/labstack/echo/v4"
	"github.com/massivebugs/home-feature-server/internal/api"
	"github.com/massivebugs/home-feature-server/internal/echo/dto"
)

type EchoHandler struct{}

func NewEchoHandler() *EchoHandler {
	return &EchoHandler{}
}

func (*EchoHandler) Echo(ctx echo.Context) error {
	v := new(dto.EchoRequestDTO)

	if err := ctx.Bind(v); err != nil {
		return api.NewAPIResponse(ctx, err, "")
	}

	if err := ctx.Validate(v); err != nil {
		return api.NewAPIResponse(ctx, err, "")
	}

	return api.NewAPIResponse(ctx, nil, v.Message)
}
