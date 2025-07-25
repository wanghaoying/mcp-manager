package converter

import (
	"github.com/getkin/kin-openapi/openapi2"
	"mcp-manager/internal/model"
)

type swagger2Converter struct{}

// NewSwagger2Converter creates a new instance of swagger2Converter.
func NewSwagger2Converter() APIEndpointConverter[*openapi2.T] {
	return &swagger2Converter{}
}

// ConvertToAPIEndpoint converts the OpenAPI 2.0 document to a slice of APIEndpoint models.
func (p *swagger2Converter) ConvertToAPIEndpoint(swaggerDoc *openapi2.T) []model.APIEndpoint {
	var endpoints []model.APIEndpoint
	for path, pathItem := range swaggerDoc.Paths {
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
