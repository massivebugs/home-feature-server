package rest

import (
	"github.com/labstack/echo/v4"
	"github.com/massivebugs/home-feature-server/internal/app"
	"github.com/massivebugs/home-feature-server/rest/oapi"
)

const (
	validationErrMessage = "there were some problems with the data you provided"
)

func NewErrorFromAppError(appErr *app.AppError) oapi.Error {
	return oapi.Error{
		Error: appErr.Error(),
	}
}

func HTTPErrorHandler(err error, c echo.Context) {
	e := c.Echo()

	// If the response has already been written to the client, then the error might have been returned by one of the middlewares.
	// In that case, we don't have to do anything.
	// This idea is derived from echo.DefaultHTTPErrorHandler.
	if c.Response().Committed {
		e.Logger.Info("Response already committed. Ignoring the error - ", err)
		return
	}

	if valErr, ok := err.(*validationError); ok {
		c.JSON(400, oapi.Error{
			Error:              validationErrMessage,
			ValidationMessages: valErr.Messages,
		})
		return
	}

	if appErr, ok := err.(*app.AppError); ok {
		c.JSON(appErr.GetHTTPStatusCode(), NewErrorFromAppError(appErr))
		return
	}

	if httpErr, ok := err.(*echo.HTTPError); ok {
		msg, ok := httpErr.Message.(string)
		if !ok {
			msg = httpErr.Error()
		}

		c.JSON(httpErr.Code, oapi.Error{
			Error: msg,
		})
		return
	}

	c.JSON(500, oapi.Error{
		Error: err.Error(),
	})
}
