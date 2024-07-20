package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/massivebugs/home-feature-server/api/dto"
	"github.com/massivebugs/home-feature-server/internal/api"
	"github.com/massivebugs/home-feature-server/internal/repeat"
)

type RepeatHandler struct {
	repeat *repeat.Repeat
}

func NewRepeatHandler() *RepeatHandler {
	return &RepeatHandler{
		repeat: repeat.NewRepeat(),
	}
}

func (h *RepeatHandler) Repeat(ctx echo.Context) *api.APIResponse {
	req := new(dto.RepeatRequestDTO)

	if err := ctx.Bind(req); err != nil {
		return api.NewAPIResponse(ctx, err, "")
	}

	if err := ctx.Validate(req); err != nil {
		return api.NewAPIResponse(ctx, err, "")
	}

	result := h.repeat.Run(ctx.Request().Context(), req)

	return api.NewAPIResponse(ctx, nil, result)
}
