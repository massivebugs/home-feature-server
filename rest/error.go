package rest

import (
	"github.com/labstack/echo/v4"
	"github.com/massivebugs/home-feature-server/internal/app"
	"github.com/massivebugs/home-feature-server/rest/oapi"
)

const (
	validationErrMessage string = "there were some problems with the data you provided"
	http500ErrMessage    string = "something went wrong with your request"
)

func NewHTTPErrorHandler(cfg *Config) func(err error, c echo.Context) {
	return func(err error, c echo.Context) {
		if c.Response().Committed {
			return
		}

		resCode := 500
		resErr := oapi.Error{}
		if valErr, ok := err.(*validationError); ok {
			resCode = 400
			resErr.Message = validationErrMessage
			resErr.ValidationMessages = valErr.Messages
		} else if appErr, ok := err.(*app.AppError); ok {
			resCode = appErr.GetHTTPStatusCode()
			resErr.Message = appErr.Error()
		} else if httpErr, ok := err.(*echo.HTTPError); ok {
			msg, ok := httpErr.Message.(string)
			if !ok {
				msg = httpErr.Error()
			}

			resCode = httpErr.Code
			resErr.Message = msg
		} else {
			resErr.Message = err.Error()
		}

		// If we are in production environment, we don't want to expose error messages to the client.
		if resCode == 500 && cfg.Environment == EnvironmentProduction {
			resErr.Message = http500ErrMessage
		}

		c.JSON(resCode, resErr)
	}
}
