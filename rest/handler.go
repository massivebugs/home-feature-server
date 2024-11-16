package rest

import (
	"context"

	"github.com/labstack/echo/v4"
	"github.com/massivebugs/home-feature-server/internal/app"
	"github.com/massivebugs/home-feature-server/internal/auth"
)

type APIResponse struct {
	Error *app.AppError `json:"error"`
	Data  interface{}   `json:"data"`
}

// Default interface for API Handlers
// I'm not entirely sure if we need this?
type IAPIHandler interface {
	GetJWTTokenClaims(c echo.Context) *auth.JWTClaims
	Validate(c echo.Context, req interface{}) error
	CreateResponse(err error, data interface{}) *APIResponse
	CreateErrorResponse(err error) *APIResponse
}

type Handler struct {
	Config *Config
}

func NewHandler(cfg *Config) *Handler {
	return &Handler{
		Config: cfg,
	}
}

// Retrieve JWT Claims from request headers
func (h *Handler) GetClaims(ctx context.Context) *auth.JWTClaims {
	return ctx.Value(CtxClaimsKey).(*auth.JWTClaims)
}

// Binds body/params to an interface, and validates.
// Interface must be a pointer
func (h *Handler) Validate(c echo.Context, req interface{}) error {
	if err := c.Bind(req); err != nil {
		return err
	}

	if err := c.Validate(req); err != nil {
		return err
	}

	return nil
}
