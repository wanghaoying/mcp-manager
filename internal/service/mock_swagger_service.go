package service

import (
	"context"
	"mcp-manager/internal/model"
)

// mockSwaggerService 是 SwaggerService 的 mock 实现，便于单元测试和开发

type mockSwaggerService struct{}

func NewMockSwaggerService() SwaggerService {
	return &mockSwaggerService{}
}

func (m *mockSwaggerService) ParseAndSave(ctx context.Context, swaggerContent []byte) ([]model.APIEndpoint, error) {
	return []model.APIEndpoint{}, nil
}
func (m *mockSwaggerService) ListAPIEndpoints(ctx context.Context, swaggerID uint) ([]model.APIEndpoint, error) {
	return []model.APIEndpoint{}, nil
}
func (m *mockSwaggerService) GetAPIEndpointByID(ctx context.Context, id uint) (*model.APIEndpoint, error) {
	return &model.APIEndpoint{}, nil
}
func (m *mockSwaggerService) DeleteAPIEndpoint(ctx context.Context, id uint) error {
	return nil
}
func (m *mockSwaggerService) UpdateAPIEndpoint(ctx context.Context, endpoint *model.APIEndpoint) error {
	return nil
}
