package api

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

type APIResponse struct {
	Error *APIError   `json:"error"`
	Data  interface{} `json:"data"`
}

func NewAPIResponse(err error, data interface{}) *APIResponse {
	apiErr, ok := err.(*APIError)
	if !ok {
		if valErrs, ok := err.(validation.Errors); ok {
			apiErr = NewAPIValidationError(CodeValidationFailed, valErrs)
		} else if err != nil {
			apiErr = NewAPIError(CodeInternalServerError, err)
		}
	}

	return &APIResponse{
		Error: apiErr,
		Data:  data,
	}
}
