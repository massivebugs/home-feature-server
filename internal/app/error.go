package app

import (
	"errors"
	"net/http"

	validation "github.com/go-ozzo/ozzo-validation"
)

type AppErrorCode string

const (
	CodeBadRequest          AppErrorCode = "bad_request"
	CodeUnauthorized        AppErrorCode = "unauthorized"
	CodeForbidden           AppErrorCode = "forbidden"
	CodeNotFound            AppErrorCode = "not_found"
	CodeInternalServerError AppErrorCode = "internal_server_error"
	CodeValidationFailed    AppErrorCode = "validation_failed"
)

var (
	errValidation = errors.New("there were some problems with the data you provided")
)

type AppError struct {
	Code             AppErrorCode      `json:"code"`
	Message          string            `json:"message"`
	ValidationErrors validation.Errors `json:"validation_errors"`
}

func NewAppError(code AppErrorCode, err error) *AppError {
	return &AppError{
		Code:    code,
		Message: err.Error(),
	}
}

func NewAppValidationError(code AppErrorCode, valErrs validation.Errors) *AppError {
	return &AppError{
		Code:             code,
		Message:          errValidation.Error(),
		ValidationErrors: valErrs,
	}
}

func (err *AppError) Error() string {
	return err.Message
}

func (err *AppError) GetHTTPStatusCode() int {
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
