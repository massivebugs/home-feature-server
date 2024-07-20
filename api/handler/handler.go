package handler

import (
	"database/sql"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/massivebugs/home-feature-server/api/config"
	"github.com/massivebugs/home-feature-server/internal/api"
)

type Handlers struct {
	PingHandler   *PingHandler
	RepeatHandler *RepeatHandler
	AuthHandler   *AuthHandler
}

func NewAPIHandlers(db *sql.DB) *Handlers {
	return &Handlers{
		PingHandler:   NewPingHandler(),
		RepeatHandler: NewRepeatHandler(),
		AuthHandler:   NewAuthHandler(db),
	}
}

func CreateEchoHandlerFunc(cfg *config.Config, handlerFunc func(ctx echo.Context) *api.APIResponse) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		res := handlerFunc(ctx)
		code := res.Error.GetHTTPStatusCode()

		// Hide unhandled errors on production
		if code == http.StatusInternalServerError && cfg.Environment == config.EnvironmentProduction {
			res.Error.Message = ""
		}

		return ctx.JSON(code, res)
	}
}
