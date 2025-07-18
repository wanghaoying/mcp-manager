package service

import (
	"context"
	"errors"
	"io"
	"testing"

	"mcp-manager/internal/model"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

// MockSwaggerParser 模拟 SwaggerParser
type MockSwaggerParser struct {
	mock.Mock
}

func (m *MockSwaggerParser) ParseFromData(data []byte) (*openapi3.T, error) {
	args := m.Called(data)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*openapi3.T), args.Error(1)
}

func (m *MockSwaggerParser) Validate(doc *openapi3.T) error {
	args := m.Called(doc)
	return args.Error(0)
}

func (m *MockSwaggerParser) Parse(path string) (*openapi3.T, error) {
	args := m.Called(path)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*openapi3.T), args.Error(1)
}

func (m *MockSwaggerParser) ExtractAPIEndpoints(doc interface{}) []model.APIEndpoint {
	args := m.Called(doc)
	return args.Get(0).([]model.APIEndpoint)
}

// MockAPIEndpointDAO 模拟 APIEndpointDAO
type MockAPIEndpointDAO struct {
	mock.Mock
}

func (m *MockAPIEndpointDAO) Create(ctx context.Context, endpoint *model.APIEndpoint) error {
	args := m.Called(ctx, endpoint)
	return args.Error(0)
}

func (m *MockAPIEndpointDAO) Delete(ctx context.Context, id uint) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *MockAPIEndpointDAO) Update(ctx context.Context, endpoint *model.APIEndpoint) error {
	args := m.Called(ctx, endpoint)
	return args.Error(0)
}

func (m *MockAPIEndpointDAO) GetByID(ctx context.Context, id uint) (*model.APIEndpoint, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.APIEndpoint), args.Error(1)
}

func (m *MockAPIEndpointDAO) List(ctx context.Context, swaggerID uint) ([]model.APIEndpoint, error) {
	args := m.Called(ctx, swaggerID)
	return args.Get(0).([]model.APIEndpoint), args.Error(1)
}

// MockHTTPClient 模拟 HTTPClient
type MockHTTPClient struct {
	mock.Mock
}

func (m *MockHTTPClient) DoRequest(ctx context.Context, method, url string, body io.Reader) (string, error) {
	args := m.Called(ctx, method, url, body)
	return args.String(0), args.Error(1)
}

// 测试数据
var sampleEndpoint = &model.APIEndpoint{
	ID:          1,
	SwaggerID:   1,
	Path:        "/test/{id}",
	Method:      "GET",
	Summary:     "Test endpoint",
	Description: "Test description",
	OperationID: "testOp",
	Tags:        "test",
	Parameters: []model.APIParameter{
		{Name: "id", In: "path", Required: true, Type: "string", Value: "123"},
		{Name: "param1", In: "query", Required: false, Type: "string", Value: "value1"},
		{Name: "Authorization", In: "header", Required: false, Type: "string", Value: "Bearer token"},
	},
	Headers: model.StringMap{"Content-Type": "application/json"},
	Body:    "",
}

// TestMain 设置测试环境
func TestMain(m *testing.M) {
	// 运行测试
	m.Run()
}

func TestSwaggerService_ParseAndSave(t *testing.T) {
	mockParser := new(MockSwaggerParser)
	mockDAO := new(MockAPIEndpointDAO)
	mockHTTPClient := new(MockHTTPClient)

	service := &swaggerService{
		parser:     mockParser,
		dao:        mockDAO,
		httpClient: mockHTTPClient,
	}

	ctx := context.Background()
	swaggerContent := []byte(`{"openapi": "3.0.0"}`)
	expectedDoc := &openapi3.T{}
	expectedEndpoints := []model.APIEndpoint{*sampleEndpoint}

	// Mock expectations
	mockParser.On("ParseFromData", swaggerContent).Return(expectedDoc, nil)
	mockParser.On("Validate", expectedDoc).Return(nil)
	mockParser.On("ExtractAPIEndpoints", expectedDoc).Return(expectedEndpoints)
	mockDAO.On("Create", ctx, mock.AnythingOfType("*model.APIEndpoint")).Return(nil)

	// Execute
	result, err := service.ParseAndSave(ctx, swaggerContent)

	// Assertions
	assert.NoError(t, err)
	assert.Equal(t, expectedEndpoints, result)
	mockParser.AssertExpectations(t)
	mockDAO.AssertExpectations(t)
}

func TestSwaggerService_ParseAndSave_ParseError(t *testing.T) {
	mockParser := new(MockSwaggerParser)
	mockDAO := new(MockAPIEndpointDAO)
	mockHTTPClient := new(MockHTTPClient)

	service := &swaggerService{
		parser:     mockParser,
		dao:        mockDAO,
		httpClient: mockHTTPClient,
	}

	ctx := context.Background()
	swaggerContent := []byte(`invalid json`)

	// Mock expectations
	mockParser.On("ParseFromData", swaggerContent).Return(nil, errors.New("parse error"))

	// Execute
	result, err := service.ParseAndSave(ctx, swaggerContent)

	// Assertions
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "parse error")
	mockParser.AssertExpectations(t)
}

func TestSwaggerService_ListAPIEndpoints(t *testing.T) {
	mockParser := new(MockSwaggerParser)
	mockDAO := new(MockAPIEndpointDAO)
	mockHTTPClient := new(MockHTTPClient)

	service := &swaggerService{
		parser:     mockParser,
		dao:        mockDAO,
		httpClient: mockHTTPClient,
	}

	ctx := context.Background()
	swaggerID := uint(1)
	expectedEndpoints := []model.APIEndpoint{*sampleEndpoint}

	// Mock expectations
	mockDAO.On("List", ctx, swaggerID).Return(expectedEndpoints, nil)

	// Execute
	result, err := service.ListAPIEndpoints(ctx, swaggerID)

	// Assertions
	assert.NoError(t, err)
	assert.Equal(t, expectedEndpoints, result)
	mockDAO.AssertExpectations(t)
}

func TestSwaggerService_GetAPIEndpointByID(t *testing.T) {
	mockParser := new(MockSwaggerParser)
	mockDAO := new(MockAPIEndpointDAO)
	mockHTTPClient := new(MockHTTPClient)

	service := &swaggerService{
		parser:     mockParser,
		dao:        mockDAO,
		httpClient: mockHTTPClient,
	}

	ctx := context.Background()
	endpointID := uint(1)

	// Mock expectations
	mockDAO.On("GetByID", ctx, endpointID).Return(sampleEndpoint, nil)

	// Execute
	result, err := service.GetAPIEndpointByID(ctx, endpointID)

	// Assertions
	assert.NoError(t, err)
	assert.Equal(t, sampleEndpoint, result)
	mockDAO.AssertExpectations(t)
}

func TestSwaggerService_GetAPIEndpointByID_NotFound(t *testing.T) {
	mockParser := new(MockSwaggerParser)
	mockDAO := new(MockAPIEndpointDAO)
	mockHTTPClient := new(MockHTTPClient)

	service := &swaggerService{
		parser:     mockParser,
		dao:        mockDAO,
		httpClient: mockHTTPClient,
	}

	ctx := context.Background()
	endpointID := uint(999)

	// Mock expectations
	mockDAO.On("GetByID", ctx, endpointID).Return(nil, gorm.ErrRecordNotFound)

	// Execute
	result, err := service.GetAPIEndpointByID(ctx, endpointID)

	// Assertions
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Equal(t, gorm.ErrRecordNotFound, err)
	mockDAO.AssertExpectations(t)
}

func TestSwaggerService_DeleteAPIEndpoint(t *testing.T) {
	mockParser := new(MockSwaggerParser)
	mockDAO := new(MockAPIEndpointDAO)
	mockHTTPClient := new(MockHTTPClient)

	service := &swaggerService{
		parser:     mockParser,
		dao:        mockDAO,
		httpClient: mockHTTPClient,
	}

	ctx := context.Background()
	endpointID := uint(1)

	// Mock expectations
	mockDAO.On("Delete", ctx, endpointID).Return(nil)

	// Execute
	err := service.DeleteAPIEndpoint(ctx, endpointID)

	// Assertions
	assert.NoError(t, err)
	mockDAO.AssertExpectations(t)
}

func TestSwaggerService_UpdateAPIEndpoint(t *testing.T) {
	mockParser := new(MockSwaggerParser)
	mockDAO := new(MockAPIEndpointDAO)
	mockHTTPClient := new(MockHTTPClient)

	service := &swaggerService{
		parser:     mockParser,
		dao:        mockDAO,
		httpClient: mockHTTPClient,
	}

	ctx := context.Background()

	// Mock expectations
	mockDAO.On("Update", ctx, sampleEndpoint).Return(nil)

	// Execute
	err := service.UpdateAPIEndpoint(ctx, sampleEndpoint)

	// Assertions
	assert.NoError(t, err)
	mockDAO.AssertExpectations(t)
}

func TestSwaggerService_TestAPIEndpoint(t *testing.T) {
	mockParser := new(MockSwaggerParser)
	mockDAO := new(MockAPIEndpointDAO)
	mockHTTPClient := new(MockHTTPClient)

	service := &swaggerService{
		parser:     mockParser,
		dao:        mockDAO,
		httpClient: mockHTTPClient,
	}

	ctx := context.Background()
	baseURL := "http://localhost:8080"
	expectedResponse := "success response"

	// Mock expectations - need to be more flexible with body matcher
	mockHTTPClient.On("DoRequest", ctx, "GET", "http://localhost:8080/test/123?param1=value1", mock.Anything).Return(expectedResponse, nil)

	// Execute
	result, err := service.TestAPIEndpoint(ctx, sampleEndpoint, baseURL)

	// Assertions
	assert.NoError(t, err)
	assert.Equal(t, expectedResponse, result)
	mockHTTPClient.AssertExpectations(t)
}

func TestSwaggerService_TestAPIEndpoint_WithBody(t *testing.T) {
	mockParser := new(MockSwaggerParser)
	mockDAO := new(MockAPIEndpointDAO)
	mockHTTPClient := new(MockHTTPClient)

	service := &swaggerService{
		parser:     mockParser,
		dao:        mockDAO,
		httpClient: mockHTTPClient,
	}

	ctx := context.Background()
	baseURL := "http://localhost:8080"
	expectedResponse := "success response"

	// Create endpoint with POST method and body
	postEndpoint := &model.APIEndpoint{
		ID:        1,
		SwaggerID: 1,
		Path:      "/test",
		Method:    "POST",
		Body:      `{"key": "value"}`,
		Headers:   model.StringMap{"Content-Type": "application/json"},
	}

	// Mock expectations
	mockHTTPClient.On("DoRequest", ctx, "POST", "http://localhost:8080/test", mock.Anything).Return(expectedResponse, nil)

	// Execute
	result, err := service.TestAPIEndpoint(ctx, postEndpoint, baseURL)

	// Assertions
	assert.NoError(t, err)
	assert.Equal(t, expectedResponse, result)
	mockHTTPClient.AssertExpectations(t)
}

func TestSwaggerService_TestAPIEndpoint_MissingRequiredParam(t *testing.T) {
	mockParser := new(MockSwaggerParser)
	mockDAO := new(MockAPIEndpointDAO)
	mockHTTPClient := new(MockHTTPClient)

	service := &swaggerService{
		parser:     mockParser,
		dao:        mockDAO,
		httpClient: mockHTTPClient,
	}

	ctx := context.Background()
	baseURL := "http://localhost:8080"

	// Create endpoint with missing required parameter
	endpointWithMissingParam := &model.APIEndpoint{
		ID:        1,
		SwaggerID: 1,
		Path:      "/test/{id}",
		Method:    "GET",
		Parameters: []model.APIParameter{
			{Name: "id", In: "path", Required: true, Type: "string", Value: ""}, // Missing value
		},
	}

	// Execute
	result, err := service.TestAPIEndpoint(ctx, endpointWithMissingParam, baseURL)

	// Assertions
	assert.Error(t, err)
	assert.Empty(t, result)
	assert.Contains(t, err.Error(), "missing required path parameter: id")
}

func TestSwaggerService_TestAPIEndpoint_HTTPError(t *testing.T) {
	mockParser := new(MockSwaggerParser)
	mockDAO := new(MockAPIEndpointDAO)
	mockHTTPClient := new(MockHTTPClient)

	service := &swaggerService{
		parser:     mockParser,
		dao:        mockDAO,
		httpClient: mockHTTPClient,
	}

	ctx := context.Background()
	baseURL := "http://localhost:8080"

	// Mock expectations
	mockHTTPClient.On("DoRequest", ctx, "GET", "http://localhost:8080/test/123?param1=value1", mock.Anything).Return("", errors.New("connection failed"))

	// Execute
	result, err := service.TestAPIEndpoint(ctx, sampleEndpoint, baseURL)

	// Assertions
	assert.Error(t, err)
	assert.Empty(t, result)
	assert.Contains(t, err.Error(), "connection failed")
	mockHTTPClient.AssertExpectations(t)
}
