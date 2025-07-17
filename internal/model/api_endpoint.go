package model

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"
)

// APIEndpoint represents an API endpoint in the system.
type APIEndpoint struct {
	ID          uint          `gorm:"primaryKey;column:id" json:"id"`                           // Unique identifier for the endpoint
	SwaggerID   uint          `gorm:"column:swagger_id" json:"swagger_id"`                      // ID from the Swagger/OpenAPI specification
	Path        string        `gorm:"column:path;type:varchar(255)" json:"path"`                // URL path of the endpoint
	Method      string        `gorm:"column:method;type:varchar(16)" json:"method"`             // HTTP method (GET, POST, etc.)
	Summary     string        `gorm:"column:summary;type:varchar(255)" json:"summary"`          // Brief summary of the endpoint
	Description string        `gorm:"column:description;type:text" json:"description"`          // Detailed description of the endpoint
	OperationID string        `gorm:"column:operation_id;type:varchar(64)" json:"operation_id"` // Unique operation ID
	Tags        string        `gorm:"column:tags;type:varchar(255)" json:"tags"`                // Tags associated with the endpoint
	Parameters  APIParameters `gorm:"type:json;column:parameters" json:"parameters"`            // List of parameters for the endpoint
	Responses   string        `gorm:"column:responses;type:json" json:"responses"`              // Responses returned by the endpoint
	Headers     StringMap     `gorm:"type:json;column:headers" json:"headers"`                  // Headers associated with the endpoint
	Body        string        `gorm:"column:body;type:text" json:"body"`                        // Request body for the endpoint
	CreatedAt   time.Time     `gorm:"column:created_at;autoCreateTime" json:"created_at"`       // Timestamp when the endpoint was created
	UpdatedAt   time.Time     `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`       // Timestamp when the endpoint was last updated
}

// APIParameter represents a single parameter in an API endpoint.
type APIParameter struct {
	Name     string `json:"name"`     // Name of the parameter
	In       string `json:"in"`       // Location of the parameter (e.g., query, path, header)
	Required bool   `json:"required"` // Whether the parameter is required
	Type     string `json:"type"`     // Data type of the parameter
	Value    string `json:"value"`    // Default value of the parameter
}

// APIParameters is a slice of APIParameter.
type APIParameters []APIParameter

// Value converts APIParameters to a database-compatible format.
func (a APIParameters) Value() (driver.Value, error) {
	b, err := json.Marshal(a)
	if err != nil {
		return nil, err
	}
	return string(b), nil
}

// Scan converts a database value back to APIParameters.
func (a *APIParameters) Scan(value interface{}) error {
	var bytes []byte
	switch v := value.(type) {
	case string:
		bytes = []byte(v)
	case []byte:
		bytes = v
	default:
		return fmt.Errorf("unsupported type: %T", value)
	}
	return json.Unmarshal(bytes, a)
}

// StringMap is a map of string key-value pairs.
type StringMap map[string]string

// Value converts StringMap to a database-compatible format.
func (m StringMap) Value() (driver.Value, error) {
	b, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}
	return string(b), nil
}

// Scan converts a database value back to StringMap.
func (m *StringMap) Scan(value interface{}) error {
	var bytes []byte
	switch v := value.(type) {
	case string:
		bytes = []byte(v)
	case []byte:
		bytes = v
	default:
		return fmt.Errorf("unsupported type: %T", value)
	}
	return json.Unmarshal(bytes, m)
}
