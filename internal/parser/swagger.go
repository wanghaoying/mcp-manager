package parser

import (
	"context"
	"fmt"
	"io/ioutil"

	"github.com/getkin/kin-openapi/openapi3"
	"mcp-manager/internal/model" // 修改为实际的导入路径
)

// SwaggerParser 定义了 Swagger 解析器的接口
type SwaggerParser interface {
	// Parse 解析 Swagger 文档
	Parse(path string) (*openapi3.T, error)
	// ParseFromData 通过字节数据解析 Swagger 文档
	ParseFromData(data []byte) (*openapi3.T, error)
	// Validate 验证 Swagger 文档
	Validate(doc *openapi3.T) error
	// ExtractAPIEndpoints 从 openapi3.T 文档中提取所有 APIEndpoint
	ExtractAPIEndpoints(doc interface{}) []model.APIEndpoint
}

// DefaultSwaggerParser 是 Swagger 解析器的默认实现
type DefaultSwaggerParser struct{}

// NewSwaggerParser 创建一个新的 Swagger 解析器
func NewSwaggerParser() SwaggerParser {
	return &DefaultSwaggerParser{}
}

// Parse 解析 Swagger 文档
func (p *DefaultSwaggerParser) Parse(path string) (*openapi3.T, error) {
	// 读取文件内容
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read swagger file: %v", err)
	}

	// 解析 Swagger 文档
	loader := openapi3.NewLoader()
	doc, err := loader.LoadFromData(data)
	if err != nil {
		return nil, fmt.Errorf("failed to parse swagger file: %v", err)
	}

	// 验证文档
	if err := p.Validate(doc); err != nil {
		return nil, err
	}

	return doc, nil
}

// ParseFromData 通过字节数据解析 Swagger 文档
func (p *DefaultSwaggerParser) ParseFromData(data []byte) (*openapi3.T, error) {
	loader := openapi3.NewLoader()
	doc, err := loader.LoadFromData(data)
	if err != nil {
		return nil, fmt.Errorf("failed to parse swagger data: %v", err)
	}
	if err := p.Validate(doc); err != nil {
		return nil, err
	}
	return doc, nil
}

// Validate 验证 Swagger 文档
func (p *DefaultSwaggerParser) Validate(doc *openapi3.T) error {
	// 确保Components对象存在（处理空components情况）
	if doc.Components == nil {
		doc.Components = &openapi3.Components{}
	}
	// 兼容空 components 字段
	if doc.Components.Extensions == nil {
		doc.Components.Extensions = map[string]interface{}{}
	}
	if doc.Components.Schemas == nil {
		doc.Components.Schemas = openapi3.Schemas{}
	}
	if doc.Components.SecuritySchemes == nil {
		doc.Components.SecuritySchemes = openapi3.SecuritySchemes{}
	}
	if doc.Components.Parameters == nil {
		doc.Components.Parameters = openapi3.ParametersMap{}
	}
	if doc.Components.RequestBodies == nil {
		doc.Components.RequestBodies = openapi3.RequestBodies{}
	}
	if doc.Components.Responses == nil {
		doc.Components.Responses = openapi3.ResponseBodies{}
	}
	if doc.Components.Headers == nil {
		doc.Components.Headers = openapi3.Headers{}
	}
	if doc.Components.Examples == nil {
		doc.Components.Examples = openapi3.Examples{}
	}
	if doc.Components.Links == nil {
		doc.Components.Links = openapi3.Links{}
	}
	if doc.Components.Callbacks == nil {
		doc.Components.Callbacks = openapi3.Callbacks{}
	}
	// 删除所有 value 为 nil 的 key
	for k, v := range doc.Components.Schemas {
		if v == nil {
			delete(doc.Components.Schemas, k)
		}
	}
	for k, v := range doc.Components.SecuritySchemes {
		if v == nil {
			delete(doc.Components.SecuritySchemes, k)
		}
	}
	for k, v := range doc.Components.Parameters {
		if v == nil {
			delete(doc.Components.Parameters, k)
		}
	}
	for k, v := range doc.Components.RequestBodies {
		if v == nil {
			delete(doc.Components.RequestBodies, k)
		}
	}
	for k, v := range doc.Components.Responses {
		if v == nil {
			delete(doc.Components.Responses, k)
		}
	}
	for k, v := range doc.Components.Headers {
		if v == nil {
			delete(doc.Components.Headers, k)
		}
	}
	for k, v := range doc.Components.Examples {
		if v == nil {
			delete(doc.Components.Examples, k)
		}
	}
	for k, v := range doc.Components.Links {
		if v == nil {
			delete(doc.Components.Links, k)
		}
	}
	for k, v := range doc.Components.Callbacks {
		if v == nil {
			delete(doc.Components.Callbacks, k)
		}
	}
	// 验证文档
	if err := doc.Validate(context.Background()); err != nil {
		return fmt.Errorf("invalid swagger document: %v", err)
	}
	// 检查必要的字段及其内容
	if doc.Info == nil {
		return fmt.Errorf("swagger document missing info section")
	}
	if doc.Info.Title == "" {
		return fmt.Errorf("info.title is required")
	}
	if doc.Info.Version == "" {
		return fmt.Errorf("info.version is required")
	}
	if doc.Paths == nil {
		return fmt.Errorf("swagger document missing paths section")
	}
	// 确保Paths被正确初始化
	if doc.Paths.Map() == nil {
		doc.Paths = &openapi3.Paths{}
	}
	return nil
}

// ExtractAPIEndpoints 从 openapi3.T 文档中提取所有 APIEndpoint
func (p *DefaultSwaggerParser) ExtractAPIEndpoints(doc interface{}) []model.APIEndpoint {
	openapiDoc, ok := doc.(*openapi3.T)
	if !ok || openapiDoc == nil {
		return nil
	}
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
