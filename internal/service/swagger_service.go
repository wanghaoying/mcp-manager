// Package service provides the business logic for handling Swagger/OpenAPI specifications and managing API endpoints.
package service

import (
	"context"
	"fmt"
	"github.com/getkin/kin-openapi/openapi2"
	"github.com/getkin/kin-openapi/openapi3"
	"io"
	"mcp-manager/internal/dao"
	"mcp-manager/internal/model"
	http "mcp-manager/internal/utils/http"
	"mcp-manager/internal/utils/parser"
	"net/url"
	"strings"
)

// SwaggerService 定义 swagger 解析与 APIEndpoint 管理的业务接口
type SwaggerService interface {
	// ParseAndSave 解析 swagger 内容并保存所有接口到数据库
	ParseAndSave(ctx context.Context, swaggerContent []byte) ([]model.APIEndpoint, error)
	// ListAPIEndpoints 查询指定 swaggerID 下的所有 APIEndpoint
	ListAPIEndpoints(ctx context.Context, swaggerID uint) ([]model.APIEndpoint, error)
	// GetAPIEndpointByID 根据 ID 查询 APIEndpoint
	GetAPIEndpointByID(ctx context.Context, id uint) (*model.APIEndpoint, error)
	// DeleteAPIEndpoint 删除指定的 APIEndpoint
	DeleteAPIEndpoint(ctx context.Context, id uint) error
	// UpdateAPIEndpoint 更新指定的 APIEndpoint
	UpdateAPIEndpoint(ctx context.Context, endpoint *model.APIEndpoint) error
	// TestAPIEndpoint 测试指定 APIEndpoint，返回响应内容
	TestAPIEndpoint(ctx context.Context, endpoint *model.APIEndpoint, baseURL string) (string, error)
}

// swaggerService 实现 SwaggerService 接口
type swaggerService struct {
	swagger2Parser parser.Parser[*openapi2.T]
	openapi3Parser parser.Parser[*openapi3.T]
	dao            dao.APIEndpointDAO
	httpClient     http.HTTPClient
}

// NewSwaggerService 创建一个新的 SwaggerService 实例
func NewSwaggerService() SwaggerService {
	return &swaggerService{
		swagger2Parser: parser.NewSwagger2Parser(),
		openapi3Parser: parser.NewOpenAPI3Parser(),
		dao:            dao.NewAPIEndpointDAO(nil),
		httpClient:     http.NewHTTPClient(),
	}
}

// ParseAndSave 解析 swagger 内容并保存所有接口到数据库
func (s *swaggerService) ParseAndSave(ctx context.Context, swaggerContent []byte) ([]model.APIEndpoint, error) {
	var (
		endpoints []model.APIEndpoint
	)

	contentStr := string(swaggerContent)
	isOpenAPI3 := strings.Contains(contentStr, "openapi") && strings.Contains(contentStr, "3.")
	isSwagger2 := strings.Contains(contentStr, "swagger") && strings.Contains(contentStr, "2.")

	if isOpenAPI3 {
		doc, err := s.openapi3Parser.ParseFromData(swaggerContent)
		if err != nil {
			return nil, err
		}
		if err := s.openapi3Parser.Validate(doc); err != nil {
			return nil, err
		}
		parserWithExtract, ok := s.openapi3Parser.(parser.SwaggerParserWithExtract[*openapi3.T])
		if !ok {
			return nil, fmt.Errorf("openapi3Parser does not support ExtractAPIEndpoints")
		}
		endpoints = parserWithExtract.ExtractAPIEndpoints(doc)
	} else if isSwagger2 {
		doc, err := s.swagger2Parser.ParseFromData(swaggerContent)
		if err != nil {
			return nil, err
		}
		if err := s.swagger2Parser.Validate(doc); err != nil {
			return nil, err
		}
		parserWithExtract, ok := s.swagger2Parser.(parser.SwaggerParserWithExtract[*openapi2.T])
		if !ok {
			return nil, fmt.Errorf("swagger2Parser does not support ExtractAPIEndpoints")
		}
		endpoints = parserWithExtract.ExtractAPIEndpoints(doc)
	} else {
		return nil, fmt.Errorf("unknown swagger/openapi version")
	}

	for i := range endpoints {
		err := s.dao.Create(ctx, &endpoints[i])
		if err != nil {
			return nil, err
		}
	}
	return endpoints, nil
}

func (s *swaggerService) ListAPIEndpoints(ctx context.Context, swaggerID uint) ([]model.APIEndpoint, error) {
	return s.dao.List(ctx, swaggerID)
}

func (s *swaggerService) GetAPIEndpointByID(ctx context.Context, id uint) (*model.APIEndpoint, error) {
	return s.dao.GetByID(ctx, id)
}

func (s *swaggerService) DeleteAPIEndpoint(ctx context.Context, id uint) error {
	return s.dao.Delete(ctx, id)
}

func (s *swaggerService) UpdateAPIEndpoint(ctx context.Context, endpoint *model.APIEndpoint) error {
	return s.dao.Update(ctx, endpoint)
}

func (s *swaggerService) TestAPIEndpoint(ctx context.Context, endpoint *model.APIEndpoint, baseURL string) (string, error) {
	// 1. 处理 path 参数
	accURL := baseURL + endpoint.Path
	for _, param := range endpoint.Parameters {
		if param.In == "path" {
			if param.Value == "" && param.Required {
				return "", fmt.Errorf("missing required path parameter: %s", param.Name)
			}
			accURL = strings.ReplaceAll(accURL, "{"+param.Name+"}", param.Value)
		}
	}

	// 2. 处理 query 参数
	query := url.Values{}
	for _, param := range endpoint.Parameters {
		if param.In == "query" && param.Value != "" {
			query.Add(param.Name, param.Value)
		}
	}
	if len(query) > 0 {
		accURL += "?" + query.Encode()
	}

	// 3. 处理 header
	headers := make(map[string]string)
	for _, param := range endpoint.Parameters {
		if param.In == "header" && param.Value != "" {
			headers[param.Name] = param.Value
		}
	}
	for k, v := range endpoint.Headers {
		headers[k] = v
	}

	// 4. 处理 body（支持 application/json）
	var bodyReader io.Reader
	if endpoint.Method == "POST" || endpoint.Method == "PUT" || endpoint.Method == "PATCH" {
		var bodyStr string
		for _, param := range endpoint.Parameters {
			if param.In == "body" {
				if param.Value == "" && param.Required {
					return "", fmt.Errorf("missing required body parameter: %s", param.Name)
				}
				bodyStr = param.Value
				break
			}
		}
		if bodyStr == "" && endpoint.Body != "" {
			bodyStr = endpoint.Body
		}
		if bodyStr != "" {
			bodyReader = strings.NewReader(bodyStr)
		}
	}

	// 5. 发起请求
	return s.httpClient.DoRequest(ctx, endpoint.Method, accURL, bodyReader)
}
