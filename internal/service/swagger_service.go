package service

import (
	"context"
	"fmt"
	"io"
	"mcp-manager/internal/dao"
	"mcp-manager/internal/model"
	http "mcp-manager/internal/utils/http"
	"mcp-manager/internal/utils/parser"
	"net/url"
	"strings"
)

// SwaggerService 定义 swagger 解析与 APIEndpoint 管理的业务接口
// 便于 mock 与扩展

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

// --- 新增：SwaggerParserWithExtract 接口，扩展 ExtractAPIEndpoints 能力 ---
type SwaggerParserWithExtract interface {
	parser.SwaggerParser
	ExtractAPIEndpoints(doc interface{}) []model.APIEndpoint
}

type swaggerService struct {
	parser     parser.SwaggerParser
	dao        dao.APIEndpointDAO
	httpClient http.HTTPClient
}

//
//func NewSwaggerService(parser parser.SwaggerParser, dao dao.APIEndpointDAO, httpClient utils.HTTPClient) SwaggerService {
//	return &swaggerService{parser: parser, dao: dao, httpClient: httpClient}
//}

func NewSwaggerService() SwaggerService {
	return &swaggerService{
		parser:     parser.NewSwaggerParser(),
		dao:        dao.NewAPIEndpointDAO(),
		httpClient: http.NewHTTPClient(),
	}
}

// ParseAndSave 解析 swagger 内容并保存所有接口到数据库
func (s *swaggerService) ParseAndSave(ctx context.Context, swaggerContent []byte) ([]model.APIEndpoint, error) {
	doc, err := s.parser.ParseFromData(swaggerContent)
	if err != nil {
		return nil, err
	}
	if err := s.parser.Validate(doc); err != nil {
		return nil, err
	}
	// 断言 parser 是否实现了 ExtractAPIEndpoints
	parserWithExtract, ok := s.parser.(SwaggerParserWithExtract)
	if !ok {
		return nil, fmt.Errorf("parser does not support ExtractAPIEndpoints")
	}
	endpoints := parserWithExtract.ExtractAPIEndpoints(doc)
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
