// Package parser provides utilities for parsing Swagger 2.0 specifications.
package parser

import (
	"encoding/json"
	"fmt"
	"github.com/getkin/kin-openapi/openapi2"
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

// Swagger2Parser 定义了 Swagger 2.0 解析器的接口实现
type Swagger2Parser struct{}

// NewSwagger2Parser 创建一个新的 Swagger2Parser 实例
func NewSwagger2Parser() Parser[*openapi2.T] {
	return &Swagger2Parser{}
}

// Parse 解析 Swagger2.0 文档
func (p *Swagger2Parser) Parse(path string) (*openapi2.T, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read swagger2 file: %v", err)
	}
	return p.ParseFromData(data)
}

// ParseFromData 通过字节数据解析 Swagger2.0 文档
func (p *Swagger2Parser) ParseFromData(data []byte) (*openapi2.T, error) {
	var doc openapi2.T
	// 先尝试 JSON
	if err := json.Unmarshal(data, &doc); err == nil {
		return &doc, nil
	}
	// 再尝试 YAML
	if err := yaml.Unmarshal(data, &doc); err == nil {
		return &doc, nil
	}
	return nil, fmt.Errorf("failed to parse swagger2 data as JSON or YAML")
}

// Validate 验证 Swagger2.0 文档
func (p *Swagger2Parser) Validate(doc *openapi2.T) error {
	if doc == nil {
		return fmt.Errorf("swagger2 document is nil")
	}
	if doc.Swagger != "2.0" {
		return fmt.Errorf("swagger version must be 2.0")
	}
	// openapi2.Info 可能为零值结构体而不是 nil
	if doc.Info.Title == "" {
		return fmt.Errorf("info.title is required")
	}
	if doc.Info.Version == "" {
		return fmt.Errorf("info.version is required")
	}
	if doc.Paths == nil || len(doc.Paths) == 0 {
		return fmt.Errorf("swagger2 document paths section is required and cannot be empty")
	}
	if doc.Host != "" {
		if doc.Host == "string" || len(doc.Host) > 253 {
			return fmt.Errorf("invalid host format")
		}
	}
	if doc.BasePath != "" && (doc.BasePath[0] != '/' || len(doc.BasePath) > 256) {
		return fmt.Errorf("basePath must start with '/' and be less than 256 characters")
	}
	return nil
}
