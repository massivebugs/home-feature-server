package rest

import (
	"context"

	"github.com/massivebugs/home-feature-server/internal/repeat"
)

type RepeatHandler struct {
	*Handler
	repeat *repeat.Repeat
}

func NewRepeatHandler(cfg *Config) *RepeatHandler {
	return &RepeatHandler{
		Handler: &Handler{
			cfg: cfg,
		},
		repeat: repeat.NewRepeat(),
	}
}

func (h *RepeatHandler) Repeat(ctx context.Context, req RepeatRequestObject) (RepeatResponseObject, error) {
	// req := new(repeat.RepeatRequest)

	// err := h.Validate(c, req)
	// if err != nil {
	// 	return h.CreateErrorResponse(c, err)
	// }

	result := h.repeat.Run(ctx, req.Body.Message)

	return Repeat200JSONResponse(result), nil
}
