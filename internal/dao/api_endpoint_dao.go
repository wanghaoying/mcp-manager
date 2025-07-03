package dao

import (
	"context"
	"gorm.io/gorm"
	"mcp-manager/internal/model"
)

// APIEndpointDAO 定义对 api_endpoints 表的基本操作
// 推荐通过依赖注入传递 *gorm.DB

type APIEndpointDAO interface {
	Create(ctx context.Context, endpoint *model.APIEndpoint) error
	Delete(ctx context.Context, id uint) error
	Update(ctx context.Context, endpoint *model.APIEndpoint) error
	GetByID(ctx context.Context, id uint) (*model.APIEndpoint, error)
	List(ctx context.Context, swaggerID uint) ([]model.APIEndpoint, error)
}

type apiEndpointDAO struct {
	db *gorm.DB
}

func NewAPIEndpointDAO(db *gorm.DB) APIEndpointDAO {
	return &apiEndpointDAO{db: db}
}

func (d *apiEndpointDAO) Create(ctx context.Context, endpoint *model.APIEndpoint) error {
	return d.db.WithContext(ctx).Create(endpoint).Error
}

func (d *apiEndpointDAO) Delete(ctx context.Context, id uint) error {
	return d.db.WithContext(ctx).Delete(&model.APIEndpoint{}, id).Error
}

func (d *apiEndpointDAO) Update(ctx context.Context, endpoint *model.APIEndpoint) error {
	return d.db.WithContext(ctx).Save(endpoint).Error
}

func (d *apiEndpointDAO) GetByID(ctx context.Context, id uint) (*model.APIEndpoint, error) {
	var endpoint model.APIEndpoint
	err := d.db.WithContext(ctx).First(&endpoint, id).Error
	if err != nil {
		return nil, err
	}
	return &endpoint, nil
}

func (d *apiEndpointDAO) List(ctx context.Context, swaggerID uint) ([]model.APIEndpoint, error) {
	var endpoints []model.APIEndpoint
	err := d.db.WithContext(ctx).Where("swagger_id = ?", swaggerID).Find(&endpoints).Error
	return endpoints, err
}
