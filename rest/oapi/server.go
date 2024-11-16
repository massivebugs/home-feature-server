// Package oapi provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.4.1 DO NOT EDIT.
package oapi

import (
	"bytes"
	"compress/gzip"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
	strictecho "github.com/oapi-codegen/runtime/strictmiddleware/echo"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {

	// (POST /api/v1/auth)
	CreateUser(ctx echo.Context) error

	// (POST /api/v1/auth/token)
	CreateJWTToken(ctx echo.Context) error

	// (PUT /api/v1/auth/token)
	RefreshJWTToken(ctx echo.Context) error

	// (GET /api/v1/ping)
	Ping(ctx echo.Context) error

	// (POST /api/v1/repeat)
	Repeat(ctx echo.Context) error

	// (POST /api/v1/secure/auth/token)
	CreateJWTRefreshToken(ctx echo.Context) error

	// (GET /api/v1/secure/system_preferences)
	GetUserSystemPreference(ctx echo.Context) error

	// (POST /api/v1/secure/system_preferences)
	CreateDefaultUserSystemPreference(ctx echo.Context) error

	// (PUT /api/v1/secure/system_preferences)
	UpdateUserSystemPreference(ctx echo.Context) error

	// (GET /api/v1/secure/user)
	GetUser(ctx echo.Context) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// CreateUser converts echo context to params.
func (w *ServerInterfaceWrapper) CreateUser(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.CreateUser(ctx)
	return err
}

// CreateJWTToken converts echo context to params.
func (w *ServerInterfaceWrapper) CreateJWTToken(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.CreateJWTToken(ctx)
	return err
}

// RefreshJWTToken converts echo context to params.
func (w *ServerInterfaceWrapper) RefreshJWTToken(ctx echo.Context) error {
	var err error

	ctx.Set(BearerAuthScopes, []string{})

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.RefreshJWTToken(ctx)
	return err
}

// Ping converts echo context to params.
func (w *ServerInterfaceWrapper) Ping(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.Ping(ctx)
	return err
}

// Repeat converts echo context to params.
func (w *ServerInterfaceWrapper) Repeat(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.Repeat(ctx)
	return err
}

// CreateJWTRefreshToken converts echo context to params.
func (w *ServerInterfaceWrapper) CreateJWTRefreshToken(ctx echo.Context) error {
	var err error

	ctx.Set(BearerAuthScopes, []string{})

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.CreateJWTRefreshToken(ctx)
	return err
}

// GetUserSystemPreference converts echo context to params.
func (w *ServerInterfaceWrapper) GetUserSystemPreference(ctx echo.Context) error {
	var err error

	ctx.Set(BearerAuthScopes, []string{})

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetUserSystemPreference(ctx)
	return err
}

// CreateDefaultUserSystemPreference converts echo context to params.
func (w *ServerInterfaceWrapper) CreateDefaultUserSystemPreference(ctx echo.Context) error {
	var err error

	ctx.Set(BearerAuthScopes, []string{})

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.CreateDefaultUserSystemPreference(ctx)
	return err
}

// UpdateUserSystemPreference converts echo context to params.
func (w *ServerInterfaceWrapper) UpdateUserSystemPreference(ctx echo.Context) error {
	var err error

	ctx.Set(BearerAuthScopes, []string{})

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.UpdateUserSystemPreference(ctx)
	return err
}

// GetUser converts echo context to params.
func (w *ServerInterfaceWrapper) GetUser(ctx echo.Context) error {
	var err error

	ctx.Set(BearerAuthScopes, []string{})

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetUser(ctx)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface) {
	RegisterHandlersWithBaseURL(router, si, "")
}

// Registers handlers, and prepends BaseURL to the paths, so that the paths
// can be served under a prefix.
func RegisterHandlersWithBaseURL(router EchoRouter, si ServerInterface, baseURL string) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.POST(baseURL+"/api/v1/auth", wrapper.CreateUser)
	router.POST(baseURL+"/api/v1/auth/token", wrapper.CreateJWTToken)
	router.PUT(baseURL+"/api/v1/auth/token", wrapper.RefreshJWTToken)
	router.GET(baseURL+"/api/v1/ping", wrapper.Ping)
	router.POST(baseURL+"/api/v1/repeat", wrapper.Repeat)
	router.POST(baseURL+"/api/v1/secure/auth/token", wrapper.CreateJWTRefreshToken)
	router.GET(baseURL+"/api/v1/secure/system_preferences", wrapper.GetUserSystemPreference)
	router.POST(baseURL+"/api/v1/secure/system_preferences", wrapper.CreateDefaultUserSystemPreference)
	router.PUT(baseURL+"/api/v1/secure/system_preferences", wrapper.UpdateUserSystemPreference)
	router.GET(baseURL+"/api/v1/secure/user", wrapper.GetUser)

}

type CreateUserRequestObject struct {
	Body *CreateUserJSONRequestBody
}

type CreateUserResponseObject interface {
	VisitCreateUserResponse(w http.ResponseWriter) error
}

type CreateUser202Response struct {
}

func (response CreateUser202Response) VisitCreateUserResponse(w http.ResponseWriter) error {
	w.WriteHeader(202)
	return nil
}

type CreateUser400JSONResponse Error

func (response CreateUser400JSONResponse) VisitCreateUserResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(400)

	return json.NewEncoder(w).Encode(response)
}

type CreateJWTTokenRequestObject struct {
	Body *CreateJWTTokenJSONRequestBody
}

type CreateJWTTokenResponseObject interface {
	VisitCreateJWTTokenResponse(w http.ResponseWriter) error
}

type CreateJWTToken200ResponseHeaders struct {
	SetCookie string
}

type CreateJWTToken200Response struct {
	Headers CreateJWTToken200ResponseHeaders
}

func (response CreateJWTToken200Response) VisitCreateJWTTokenResponse(w http.ResponseWriter) error {
	w.Header().Set("Set-Cookie", fmt.Sprint(response.Headers.SetCookie))
	w.WriteHeader(200)
	return nil
}

type CreateJWTToken400JSONResponse Error

func (response CreateJWTToken400JSONResponse) VisitCreateJWTTokenResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(400)

	return json.NewEncoder(w).Encode(response)
}

type CreateJWTToken403JSONResponse Error

func (response CreateJWTToken403JSONResponse) VisitCreateJWTTokenResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(403)

	return json.NewEncoder(w).Encode(response)
}

type RefreshJWTTokenRequestObject struct {
}

type RefreshJWTTokenResponseObject interface {
	VisitRefreshJWTTokenResponse(w http.ResponseWriter) error
}

type RefreshJWTToken200ResponseHeaders struct {
	SetCookie string
}

type RefreshJWTToken200Response struct {
	Headers RefreshJWTToken200ResponseHeaders
}

func (response RefreshJWTToken200Response) VisitRefreshJWTTokenResponse(w http.ResponseWriter) error {
	w.Header().Set("Set-Cookie", fmt.Sprint(response.Headers.SetCookie))
	w.WriteHeader(200)
	return nil
}

type PingRequestObject struct {
}

type PingResponseObject interface {
	VisitPingResponse(w http.ResponseWriter) error
}

type Ping200JSONResponse struct {
	// Message Pong!
	Message string `json:"message"`
}

func (response Ping200JSONResponse) VisitPingResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type RepeatRequestObject struct {
	Body *RepeatJSONRequestBody
}

type RepeatResponseObject interface {
	VisitRepeatResponse(w http.ResponseWriter) error
}

type Repeat200JSONResponse struct {
	// Message The message you sent
	Message string `json:"message"`
}

func (response Repeat200JSONResponse) VisitRepeatResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type CreateJWTRefreshTokenRequestObject struct {
}

type CreateJWTRefreshTokenResponseObject interface {
	VisitCreateJWTRefreshTokenResponse(w http.ResponseWriter) error
}

type CreateJWTRefreshToken200ResponseHeaders struct {
	SetCookie string
}

type CreateJWTRefreshToken200Response struct {
	Headers CreateJWTRefreshToken200ResponseHeaders
}

func (response CreateJWTRefreshToken200Response) VisitCreateJWTRefreshTokenResponse(w http.ResponseWriter) error {
	w.Header().Set("Set-Cookie", fmt.Sprint(response.Headers.SetCookie))
	w.WriteHeader(200)
	return nil
}

type CreateJWTRefreshToken403JSONResponse Error

func (response CreateJWTRefreshToken403JSONResponse) VisitCreateJWTRefreshTokenResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(403)

	return json.NewEncoder(w).Encode(response)
}

type GetUserSystemPreferenceRequestObject struct {
}

type GetUserSystemPreferenceResponseObject interface {
	VisitGetUserSystemPreferenceResponse(w http.ResponseWriter) error
}

type GetUserSystemPreference200JSONResponse struct {
	// UserSystemPreference Model defining user system preferences such as language, time zone etc.
	UserSystemPreference UserSystemPreference `json:"user_system_preference"`
}

func (response GetUserSystemPreference200JSONResponse) VisitGetUserSystemPreferenceResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type GetUserSystemPreference404Response struct {
}

func (response GetUserSystemPreference404Response) VisitGetUserSystemPreferenceResponse(w http.ResponseWriter) error {
	w.WriteHeader(404)
	return nil
}

type CreateDefaultUserSystemPreferenceRequestObject struct {
}

type CreateDefaultUserSystemPreferenceResponseObject interface {
	VisitCreateDefaultUserSystemPreferenceResponse(w http.ResponseWriter) error
}

type CreateDefaultUserSystemPreference200JSONResponse struct {
	// UserSystemPreference Model defining user system preferences such as language, time zone etc.
	UserSystemPreference UserSystemPreference `json:"user_system_preference"`
}

func (response CreateDefaultUserSystemPreference200JSONResponse) VisitCreateDefaultUserSystemPreferenceResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type UpdateUserSystemPreferenceRequestObject struct {
	Body *UpdateUserSystemPreferenceJSONRequestBody
}

type UpdateUserSystemPreferenceResponseObject interface {
	VisitUpdateUserSystemPreferenceResponse(w http.ResponseWriter) error
}

type UpdateUserSystemPreference200JSONResponse struct {
	// UserSystemPreference Model defining user system preferences such as language, time zone etc.
	UserSystemPreference UserSystemPreference `json:"user_system_preference"`
}

func (response UpdateUserSystemPreference200JSONResponse) VisitUpdateUserSystemPreferenceResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type GetUserRequestObject struct {
}

type GetUserResponseObject interface {
	VisitGetUserResponse(w http.ResponseWriter) error
}

type GetUser200JSONResponse struct {
	User User `json:"user"`
}

func (response GetUser200JSONResponse) VisitGetUserResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

// StrictServerInterface represents all server handlers.
type StrictServerInterface interface {

	// (POST /api/v1/auth)
	CreateUser(ctx context.Context, request CreateUserRequestObject) (CreateUserResponseObject, error)

	// (POST /api/v1/auth/token)
	CreateJWTToken(ctx context.Context, request CreateJWTTokenRequestObject) (CreateJWTTokenResponseObject, error)

	// (PUT /api/v1/auth/token)
	RefreshJWTToken(ctx context.Context, request RefreshJWTTokenRequestObject) (RefreshJWTTokenResponseObject, error)

	// (GET /api/v1/ping)
	Ping(ctx context.Context, request PingRequestObject) (PingResponseObject, error)

	// (POST /api/v1/repeat)
	Repeat(ctx context.Context, request RepeatRequestObject) (RepeatResponseObject, error)

	// (POST /api/v1/secure/auth/token)
	CreateJWTRefreshToken(ctx context.Context, request CreateJWTRefreshTokenRequestObject) (CreateJWTRefreshTokenResponseObject, error)

	// (GET /api/v1/secure/system_preferences)
	GetUserSystemPreference(ctx context.Context, request GetUserSystemPreferenceRequestObject) (GetUserSystemPreferenceResponseObject, error)

	// (POST /api/v1/secure/system_preferences)
	CreateDefaultUserSystemPreference(ctx context.Context, request CreateDefaultUserSystemPreferenceRequestObject) (CreateDefaultUserSystemPreferenceResponseObject, error)

	// (PUT /api/v1/secure/system_preferences)
	UpdateUserSystemPreference(ctx context.Context, request UpdateUserSystemPreferenceRequestObject) (UpdateUserSystemPreferenceResponseObject, error)

	// (GET /api/v1/secure/user)
	GetUser(ctx context.Context, request GetUserRequestObject) (GetUserResponseObject, error)
}

type StrictHandlerFunc = strictecho.StrictEchoHandlerFunc
type StrictMiddlewareFunc = strictecho.StrictEchoMiddlewareFunc

func NewStrictHandler(ssi StrictServerInterface, middlewares []StrictMiddlewareFunc) ServerInterface {
	return &strictHandler{ssi: ssi, middlewares: middlewares}
}

type strictHandler struct {
	ssi         StrictServerInterface
	middlewares []StrictMiddlewareFunc
}

// CreateUser operation middleware
func (sh *strictHandler) CreateUser(ctx echo.Context) error {
	var request CreateUserRequestObject

	var body CreateUserJSONRequestBody
	if err := ctx.Bind(&body); err != nil {
		return err
	}
	request.Body = &body

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.CreateUser(ctx.Request().Context(), request.(CreateUserRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "CreateUser")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(CreateUserResponseObject); ok {
		return validResponse.VisitCreateUserResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// CreateJWTToken operation middleware
func (sh *strictHandler) CreateJWTToken(ctx echo.Context) error {
	var request CreateJWTTokenRequestObject

	var body CreateJWTTokenJSONRequestBody
	if err := ctx.Bind(&body); err != nil {
		return err
	}
	request.Body = &body

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.CreateJWTToken(ctx.Request().Context(), request.(CreateJWTTokenRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "CreateJWTToken")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(CreateJWTTokenResponseObject); ok {
		return validResponse.VisitCreateJWTTokenResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// RefreshJWTToken operation middleware
func (sh *strictHandler) RefreshJWTToken(ctx echo.Context) error {
	var request RefreshJWTTokenRequestObject

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.RefreshJWTToken(ctx.Request().Context(), request.(RefreshJWTTokenRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "RefreshJWTToken")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(RefreshJWTTokenResponseObject); ok {
		return validResponse.VisitRefreshJWTTokenResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// Ping operation middleware
func (sh *strictHandler) Ping(ctx echo.Context) error {
	var request PingRequestObject

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.Ping(ctx.Request().Context(), request.(PingRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "Ping")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(PingResponseObject); ok {
		return validResponse.VisitPingResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// Repeat operation middleware
func (sh *strictHandler) Repeat(ctx echo.Context) error {
	var request RepeatRequestObject

	var body RepeatJSONRequestBody
	if err := ctx.Bind(&body); err != nil {
		return err
	}
	request.Body = &body

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.Repeat(ctx.Request().Context(), request.(RepeatRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "Repeat")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(RepeatResponseObject); ok {
		return validResponse.VisitRepeatResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// CreateJWTRefreshToken operation middleware
func (sh *strictHandler) CreateJWTRefreshToken(ctx echo.Context) error {
	var request CreateJWTRefreshTokenRequestObject

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.CreateJWTRefreshToken(ctx.Request().Context(), request.(CreateJWTRefreshTokenRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "CreateJWTRefreshToken")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(CreateJWTRefreshTokenResponseObject); ok {
		return validResponse.VisitCreateJWTRefreshTokenResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// GetUserSystemPreference operation middleware
func (sh *strictHandler) GetUserSystemPreference(ctx echo.Context) error {
	var request GetUserSystemPreferenceRequestObject

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.GetUserSystemPreference(ctx.Request().Context(), request.(GetUserSystemPreferenceRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetUserSystemPreference")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(GetUserSystemPreferenceResponseObject); ok {
		return validResponse.VisitGetUserSystemPreferenceResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// CreateDefaultUserSystemPreference operation middleware
func (sh *strictHandler) CreateDefaultUserSystemPreference(ctx echo.Context) error {
	var request CreateDefaultUserSystemPreferenceRequestObject

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.CreateDefaultUserSystemPreference(ctx.Request().Context(), request.(CreateDefaultUserSystemPreferenceRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "CreateDefaultUserSystemPreference")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(CreateDefaultUserSystemPreferenceResponseObject); ok {
		return validResponse.VisitCreateDefaultUserSystemPreferenceResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// UpdateUserSystemPreference operation middleware
func (sh *strictHandler) UpdateUserSystemPreference(ctx echo.Context) error {
	var request UpdateUserSystemPreferenceRequestObject

	var body UpdateUserSystemPreferenceJSONRequestBody
	if err := ctx.Bind(&body); err != nil {
		return err
	}
	request.Body = &body

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.UpdateUserSystemPreference(ctx.Request().Context(), request.(UpdateUserSystemPreferenceRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "UpdateUserSystemPreference")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(UpdateUserSystemPreferenceResponseObject); ok {
		return validResponse.VisitUpdateUserSystemPreferenceResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// GetUser operation middleware
func (sh *strictHandler) GetUser(ctx echo.Context) error {
	var request GetUserRequestObject

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.GetUser(ctx.Request().Context(), request.(GetUserRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetUser")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(GetUserResponseObject); ok {
		return validResponse.VisitGetUserResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/+RYb2/bthP+Klf+fi+V2IlbbDBQYG3Xrs0wLGgS9EUQGLR4tthIJMc/SbzA3304UrKt",
	"SE7S1F037E3gSLzj3XPPczzqluW6Mlqh8o6Nb5nLC6x4/InWaks/BLrcSuOlVmzM3tJjyLVA4EpAUAJt",
	"uZBqDtHAsYwZqw1aLzH6qdA5Pseup1cQHM5CCfUKSK+n5MsXmPyxjOENr0yJbMx8gRbhmv44XSEYq6cl",
	"Vg6upS+ijeCew0IHenUlBQqWMb8wZOy8lWrOlhm74qUUnKKY1FvHQLkQkh7y8riVQMd+FdAtozAmM4ml",
	"SOFBiWruC6iC8zBFmKK/RlRwEME6HJJ57U9PP2Pu2XKZMYt/BGlRsPH5Cq7+MC865hkLDmOd2rDnFrlH",
	"MeG+Nwcp6PFM24oWsCCVHx2uwZLK4xwtrSz1fI5iItU2V4pX2PPiTl6SShGX3nGZbYa6Lb+ThfNYHVuc",
	"oUWV97DpNy2wBIEzqYhAZAQuWoFZmTlwIS+AOyi5mgc+xwy8rBD+1AoBfb7foW+zMEpixcTPfO/omBIK",
	"Zcmn9MjbgFkPBt1qO8yDlX5xQlpLm0yRW7Svgi/W/71ranP06ZRlSZnkKb1dl6rw3rAlOZZqpmMhpI9B",
	"vieNvEPug0U4QXuFFj6+PTmFV8cfiF9oXcLuYH+4H6mpDSpuJBuz0f5wf0RgcF/EEAfcyMHVwYDXQRrt",
	"fLcKb2IpHXBQeJ2KEFxTD6p+BlhxWUY5GO7ctbZEDEI8Uv2DWHk5czFNIhE6/1qLReS1Vh5V3JobU8o8",
	"mg0+O9q/aWBdOcRdu/E2OuUQ1ZaCa/ccdJ6C/6l+tJ/rqtNUMnazp7mRe9QX56j28MZbvuf5PG5eK5kM",
	"0gZUrVX2W4NqmseP4DX8cAh5wS3PPVqXRfwICi4VcCjRe7QZqFBN0caXHJzBXPJybdbJa9LEMDk4HD37",
	"iqxWjmJmTa0fzmxEmb0YAi9NwVWo0Mp8M2LXW4qvCLTZKKukejnKKn7z8sUwqWezWTU0WGWyUa5ui2pb",
	"UyOID5zRyiX2HQ4Pu2D8/itJ7vlw+EW0/r/FGRuz/w3WB/egPrUH6cSM8bS3es0FfEw6otfLrCXngdeX",
	"qLaL+iP6YFUj6qNPp3BKBjDTFjjUFelT8NGn07hyZyre1Mw9VO479B8m5VYSZmm+qPiNrEJFfF29g1wH",
	"5b8dS7fwc3e8HPbyMmMFckH6G9+yE/R7b7S+lNguTdsqUWK9FpKHbiliWH8/62nP0bffkw4tkA6EdDQX",
	"iLTGhF5ZzSy6Al2cXPNgLSq/lldHU/X6lqi+fy03Bho2Pm+PMucXy4vNZmPIaHzL5uj7rgO5VleoJA1q",
	"gEoYLZWnAyIvML+M3SbXSmEeDe6ic5wk1gfJE5vN1ovLsVbzZy3NN0/uH4Ebh/1K7R4OrUZt0WAawO9v",
	"0sQlR0Sqd8sicNSUaAwzwRpN6HTJFd3vqlFvxe6EZtLmwkcXtWueijytw57y/LKF7XssS/214D6qDe46",
	"19OinSrlt/vUHuBNFCc+8ZyvW84XnPe1xT+pQ32/zv8FrbEuU7qxTjZurFv7ZVOt/ptup0K/oD/ru0jv",
	"VAoUy6STw0OQ9l7w+0aeHteP1QSR4HkXxLNe8KDgDpSmqRAV1J8nYIH+ETXN7r0Zg8AZD6VPVbunXGn5",
	"z2n1f7lwjwG8b7o6MzRMP4hzWrYV4KedhU/A7BufUf/i+nabZPPB8962SHMQHXqoPIGEIlEhfiPb0ht3",
	"L6nHANwL6A7hiyvsVTxez++i9U5b2LyTsowFW9ZfFd14MCh1zstCOz8+GB2O2PJi+VcAAAD//75QsJoz",
	"GAAA",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %w", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	res := make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	resolvePath := PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		pathToFile := url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
