package http

import (
	"github.com/labstack/echo/v4"
	"github.com/massivebugs/home-feature-server/internal/repeat"
)

type RepeatHandler struct {
	*Handler
	repeat *repeat.Repeat
}

func NewRepeatHandler() *RepeatHandler {
	return &RepeatHandler{
		repeat: repeat.NewRepeat(),
	}
}

func (h *RepeatHandler) Repeat(c echo.Context) *APIResponse {
	req := new(repeat.RepeatRequest)

	err := h.Validate(c, req)
	if err != nil {
		return h.CreateErrorResponse(err)
	}

	result := h.repeat.Run(c.Request().Context(), req)

	return h.CreateResponse(nil, result)
}
