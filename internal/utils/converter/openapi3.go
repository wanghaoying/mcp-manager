package converter

import (
	"github.com/getkin/kin-openapi/openapi3"
	"mcp-manager/internal/model"
)

// openapi3Converter implements the APIENdpointConverter interface for OpenAPI 3.0 documents.
type openapi3Converter struct{}

// NewOpenAPI3Converter creates a new instance of openapi3Converter.
func NewOpenAPI3Converter() APIEndpointConverter[*openapi3.T] {
	return &openapi3Converter{}
}

// ConvertToAPIEndpoint converts the OpenAPI 3.0 document to a slice of APIEndpoint models.
func (p *openapi3Converter) ConvertToAPIEndpoint(openapiDoc *openapi3.T) []model.APIEndpoint {
	var endpoints []model.APIEndpoint
	for path, pathItem := range openapiDoc.Paths.Map() {
		for method, operation := range pathItem.Operations() {
			endpoints = append(endpoints, model.APIEndpoint{
				Path:        path,
				Method:      method,
				Summary:     operation.Summary,
				Description: operation.Description,
				OperationID: operation.OperationID,
			})
		}
	}
	return endpoints
}
