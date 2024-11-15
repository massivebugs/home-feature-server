package app

import (
	"net/http"
)

type AppErrorCode string

const (
	CodeBadRequest          AppErrorCode = "bad_request"
	CodeUnauthorized        AppErrorCode = "unauthorized"
	CodeForbidden           AppErrorCode = "forbidden"
	CodeNotFound            AppErrorCode = "not_found"
	CodeInternalServerError AppErrorCode = "internal_server_error"
)

type AppError struct {
	code AppErrorCode
	error
}

func NewAppError(code AppErrorCode, err error) *AppError {
	return &AppError{
		code:  code,
		error: err,
	}
}

func (e *AppError) GetHTTPStatusCode() int {
	if e == nil {
		return http.StatusOK
	}

	switch e.code {
	case CodeBadRequest:
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
