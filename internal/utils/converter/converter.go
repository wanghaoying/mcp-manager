package converter

import (
	"github.com/getkin/kin-openapi/openapi2"
	"github.com/getkin/kin-openapi/openapi3"
	"mcp-manager/internal/model"
)

// APIEndpointConverter defines an interface for converting OpenAPI documents to APIEndpoint models.
type APIEndpointConverter[T interface{ *openapi2.T } | interface{ *openapi3.T }] interface {
	// ConvertToAPIEndpoint converts the given data to an APIEndpoint model.
	ConvertToAPIEndpoint(data T) []model.APIEndpoint
}
