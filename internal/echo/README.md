# Echo

Not to be confused with the Go web framework, this is an example module for testing purposes.

Send a `POST` request to `/api/v1/echo` with a JSON body like the following:
```json
{
    "message": "Repeat this message!"
}
```

Request body is validated using `github.com/go-ozzo/ozzo-validation` via the defined DTO struct (`internal/api/echo/dto/request.go`).

See OpenAPI specs(`/doc/openapi/echo.yaml`) for more details! 

