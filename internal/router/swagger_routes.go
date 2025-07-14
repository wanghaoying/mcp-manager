package router

import (
	"mcp-manager/internal/controller"
	"mcp-manager/internal/service"

	"github.com/gin-gonic/gin"
)

// RegisterSwaggerHandlers 注册所有 Swagger 相关的 HTTP 路由
// 包括接口解析、管理、测试、校验等能力
// 依赖 service.GetSwaggerService() 注入业务实现
func RegisterSwaggerHandlers(r *gin.Engine) {
	handler := controller.NewSwaggerServiceHandler(service.NewSwaggerService())

	// 业务接口相关
	r.POST("/api/swagger/parse", handler.ParseAndSave)               // 解析并保存 swagger 接口
	r.GET("/api/swagger/endpoints", handler.ListAPIEndpoints)        // 查询指定 swaggerID 下所有接口
	r.GET("/api/swagger/endpoint/:id", handler.GetAPIEndpointByID)   // 查询单个接口详情
	r.DELETE("/api/swagger/endpoint/:id", handler.DeleteAPIEndpoint) // 删除接口
	r.PUT("/api/swagger/endpoint", handler.UpdateAPIEndpoint)        // 更新接口
	r.POST("/api/swagger/endpoint/test", handler.TestAPIEndpoint)    // 测试接口

	// swagger 校验相关
	r.POST("/api/swagger/validate/file", controller.ValidateSwaggerByFile) // 文件上传校验
	r.POST("/api/swagger/validate/text", controller.ValidateSwaggerByText) // 文本内容校验
}
