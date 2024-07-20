package handler

import (
	"github.com/labstack/echo/v4"
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
	req := new(repeat.RepeatRequestDTO)

	if err := ctx.Bind(req); err != nil {
		return api.NewAPIResponse(err, "")
	}

	if err := ctx.Validate(req); err != nil {
		return api.NewAPIResponse(err, "")
	}

	result := h.repeat.Run(ctx.Request().Context(), req)

	return api.NewAPIResponse(nil, result)
}
