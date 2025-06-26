package router

import (
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// RegisterRoutes 注册所有接口路由
func RegisterRoutes(r *gin.Engine) {
	// 注册示例路由组
	RegisterExampleRoutes(r)

	// Swagger 文档路由
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	pprof.Register(r, "debug/pprof")

	// 其他路由组注册
}
