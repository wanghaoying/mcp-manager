package model

import (
	"time"
)

// APIEndpoint 表示 swagger 解析出的接口信息
// 对应数据库表 api_endpoints
// 参数和响应字段建议用 JSON 存储
// 你可以在 example/api_endpoints.sql 找到对应的 MySQL 建表语句
type APIEndpoint struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	SwaggerID   uint      `json:"swagger_id"` // 可选，关联 swagger 文档
	Path        string    `json:"path"`
	Method      string    `json:"method"`
	Summary     string    `json:"summary"`
	Description string    `json:"description"`
	Tags        string    `json:"tags"`       // 多标签用逗号分隔或 json
	Parameters  string    `json:"parameters"` // 建议存 json
	Responses   string    `json:"responses"`  // 建议存 json
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
