package handler

import (
	"context"

	"github.com/massivebugs/home-feature-server/internal/repeat"
	"github.com/massivebugs/home-feature-server/rest"
	"github.com/massivebugs/home-feature-server/rest/oapi"
)

type RepeatHandler struct {
	*rest.Handler
	repeat *repeat.Repeat
}

func NewRepeatHandler(cfg *rest.Config) *RepeatHandler {
	return &RepeatHandler{
		Handler: rest.NewHandler(cfg),
		repeat:  repeat.NewRepeat(),
	}
}

func (h *RepeatHandler) Repeat(ctx context.Context, req oapi.RepeatRequestObject) (oapi.RepeatResponseObject, error) {
	result := h.repeat.Run(ctx, req.Body.Message)

	return oapi.Repeat200JSONResponse{
		Message: result,
	}, nil
}
