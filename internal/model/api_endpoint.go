package model

import (
	"time"
)

// APIEndpoint 表示 swagger 解析出的接口信息
// 对应数据库表 api_endpoints
// 参数和响应字段建议用 JSON 存储
// 你可以在 example/api_endpoints.sql 找到对应的 MySQL 建表语句
type APIEndpoint struct {
	ID          uint              `gorm:"primaryKey" json:"id"`
	SwaggerID   uint              `json:"swagger_id"` // 可选，关联 swagger 文档
	Path        string            `json:"path"`
	Method      string            `json:"method"`
	Summary     string            `json:"summary"`
	Description string            `json:"description"`
	OperationID string            `json:"operation_id"` // 新增，唯一标识 operationId
	Tags        string            `json:"tags"`         // 多标签用逗号分隔或 json
	Parameters  []APIParameter    `json:"parameters"`   // 建议存 json
	Responses   string            `json:"responses"`    // 建议存 json
	Headers     map[string]string `json:"headers"`      // 新增，接口自定义 header
	Body        string            `json:"body"`         // 新增，接口自定义 body 内容
	CreatedAt   time.Time         `json:"created_at"`
	UpdatedAt   time.Time         `json:"updated_at"`
}

// APIParameter 表示单个参数
type APIParameter struct {
	Name     string `json:"name"`
	In       string `json:"in"` // path, query, header, body
	Required bool   `json:"required"`
	Type     string `json:"type"`
	Value    string `json:"value"` // 测试时传入的值
}
