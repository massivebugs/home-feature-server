# /internal/repeat

This is an example module for testing purposes.

Send a `POST` request to `/api/v1/repeat` with a JSON body like the following:

```json
{
  "message": "Repeat this message!"
}
```

Request body is validated using `github.com/go-ozzo/ozzo-validation` via the defined DTO struct (`internal/api/repeat/dto/request.go`).

See OpenAPI specs(`/doc/openapi/repeat.yaml`) for more details!
