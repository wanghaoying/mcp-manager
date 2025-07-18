package router

import (
	"github.com/gin-gonic/gin"
	"mcp-manager/internal/middleware"
)

// RegisterRoutes 注册所有接口路由
func RegisterRoutes(r *gin.Engine) {
	// 注册跨域中间件
	r.Use(middleware.CORSMiddleware())

	// 注册工具性路由（Swagger、pprof、ping）
	RegisterUtilityRoutes(r)

	// 注册Swagger相关路由
	RegisterSwaggerHandlers(r)
}
