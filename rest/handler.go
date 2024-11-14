package rest

import (
	"context"
	"net/http"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/golang-jwt/jwt/v5"
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
	cfg *Config
}

// Retrieve JWT Token from request headers
func (h *Handler) GetTokenClaims(c echo.Context) *auth.JWTClaims {
	token := c.Get("user").(*jwt.Token)
	return token.Claims.(*auth.JWTClaims)
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

func (h *Handler) CreateEchoHandlerFunc(c echo.Context, res *APIResponse) error {
	code := res.Error.GetHTTPStatusCode()

	// TODO: Move to middleware
	// Hide unhandled errors on production
	if code == http.StatusInternalServerError && h.cfg.Environment == EnvironmentProduction {
		// res.Error.Message = ""
	}

	return c.JSON(code, res)
}

// TODO: Move to StrictServer options
func (h *Handler) CreateResponse(c echo.Context, err error, data interface{}) error {
	apiErr, ok := err.(*app.AppError)
	if !ok {
		if _, ok := err.(validation.Errors); ok {
			// apiErr = app.NewAppValidationError(app.CodeValidationFailed, valErrs)
		} else if err != nil {
			apiErr = app.NewAppError(app.CodeInternalServerError, err)
		}
	}

	return h.CreateEchoHandlerFunc(c, &APIResponse{
		Error: apiErr,
		Data:  data,
	})
}

func (h *Handler) CreateErrorResponse(c echo.Context, err error) error {
	return h.CreateResponse(c, err, nil)
}
