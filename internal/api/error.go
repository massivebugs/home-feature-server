package api

import (
	"errors"
	"net/http"

	validation "github.com/go-ozzo/ozzo-validation"
)

type APIErrorCode string

const (
	CodeBadRequest          APIErrorCode = "bad_request"
	CodeUnauthorized        APIErrorCode = "unauthorized"
	CodeForbidden           APIErrorCode = "forbidden"
	CodeNotFound            APIErrorCode = "not_found"
	CodeInternalServerError APIErrorCode = "internal_server_error"
	CodeValidationFailed    APIErrorCode = "validation_failed"
)

var (
	errValidation = errors.New("there were some problems with the data you provided")
)

type APIError struct {
	Code             APIErrorCode      `json:"code"`
	Message          string            `json:"message"`
	ValidationErrors validation.Errors `json:"validation_errors"`
}

func NewAPIError(code APIErrorCode, err error) *APIError {
	return &APIError{
		Code:    code,
		Message: err.Error(),
	}
}

func NewAPIValidationError(code APIErrorCode, valErrs validation.Errors) *APIError {
	return &APIError{
		Code:             code,
		Message:          errValidation.Error(),
		ValidationErrors: valErrs,
	}
}

func (err *APIError) Error() string {
	return err.Message
}

func (err *APIError) GetHTTPStatusCode() int {
	if err == nil {
		return http.StatusOK
	}

	switch err.Code {
	case CodeBadRequest, CodeValidationFailed:
		return http.StatusBadRequest
	case CodeUnauthorized:
		return http.StatusUnauthorized
	case CodeForbidden:
		return http.StatusForbidden
	case CodeNotFound:
		return http.StatusNotFound
	default:
		return http.StatusInternalServerError
	}
}
