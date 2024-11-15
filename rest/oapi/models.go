// Package oapi provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.4.1 DO NOT EDIT.
package oapi

const (
	BearerAuthScopes = "bearerAuth.Scopes"
)

// Error Error code and underlying errors
type Error struct {
	// Error A useful message describing the error
	Error              string            `json:"error"`
	ValidationMessages map[string]string `json:"validation_messages"`
}

// CreateJWTTokenJSONBody defines parameters for CreateJWTToken.
type CreateJWTTokenJSONBody struct {
	// Password must be between 8 to 72 characters, and contain a letter, number and a special character
	Password string `json:"password"`

	// Username must be between 3 to 50 alphanumerical characters
	Username string `json:"username" validate:"alphanum,max=50"`
}

// RepeatJSONBody defines parameters for Repeat.
type RepeatJSONBody struct {
	// Message Some message you want to be sent back
	Message string `json:"message"`
}

// CreateJWTTokenJSONRequestBody defines body for CreateJWTToken for application/json ContentType.
type CreateJWTTokenJSONRequestBody CreateJWTTokenJSONBody

// RepeatJSONRequestBody defines body for Repeat for application/json ContentType.
type RepeatJSONRequestBody RepeatJSONBody