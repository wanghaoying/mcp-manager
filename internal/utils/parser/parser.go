// Package parser provides utilities for parsing and manipulating paths.
package parser

import (
	"github.com/getkin/kin-openapi/openapi2"
	"github.com/getkin/kin-openapi/openapi3"
)

// Parser defines the interface for parsing OpenAPI documents.
type Parser[T interface{ *openapi3.T } | interface{ *openapi2.T }] interface {
	// Parse parses the given path and returns a structured representation.
	Parse(path string) (T, error)
	// ParseFromData parses the given byte data and returns a structured representation.
	ParseFromData(data []byte) (T, error)
	// Validate validates the structured representation of the path.
	Validate(doc T) error
}
