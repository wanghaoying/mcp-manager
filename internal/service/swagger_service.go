package service

import (
	"context"
	"fmt"
	"mcp-manager/internal/dao"
	"mcp-manager/internal/model"
	"mcp-manager/internal/parser"
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
}

// --- 新增：SwaggerParserWithExtract 接口，扩展 ExtractAPIEndpoints 能力 ---
type SwaggerParserWithExtract interface {
	parser.SwaggerParser
	ExtractAPIEndpoints(doc interface{}) []model.APIEndpoint
}

type swaggerService struct {
	parser parser.SwaggerParser
	dao    dao.APIEndpointDAO
}

func NewSwaggerService(parser parser.SwaggerParser, dao dao.APIEndpointDAO) SwaggerService {
	return &swaggerService{parser: parser, dao: dao}
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

// --- mock 实现已移除，如需 mock 请在 mock_swagger_service.go 中实现 ---
