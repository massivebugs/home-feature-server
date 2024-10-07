package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/massivebugs/home-feature-server/api"
	"github.com/massivebugs/home-feature-server/internal/repeat"
)

type RepeatHandler struct {
	*api.Handler
	repeat *repeat.Repeat
}

func NewRepeatHandler() *RepeatHandler {
	return &RepeatHandler{
		repeat: repeat.NewRepeat(),
	}
}

func (h *RepeatHandler) Repeat(c echo.Context) *api.APIResponse {
	req := new(repeat.RepeatDTO)

	err := h.Validate(c, req)
	if err != nil {
		return h.CreateErrorResponse(err)
	}

	result := h.repeat.Run(c.Request().Context(), req)

	return h.CreateResponse(nil, result)
}
