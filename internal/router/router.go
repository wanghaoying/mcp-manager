package router

import "github.com/gin-gonic/gin"

// RegisterRoutes 注册所有接口路由
func RegisterRoutes(r *gin.Engine) {
	// 注册工具性路由（Swagger、pprof、ping）
	RegisterUtilityRoutes(r)

	// 注册Swagger相关路由
	RegisterSwaggerHandlers(r)
}
