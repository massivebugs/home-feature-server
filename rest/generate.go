//go:generate go run -modfile=../tools/go.mod github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen -config=../config/oapi-codegen.server.yaml rest.yaml
//go:generate go run -modfile=../tools/go.mod github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen -config=../config/oapi-codegen.models.yaml rest.yaml

package rest
