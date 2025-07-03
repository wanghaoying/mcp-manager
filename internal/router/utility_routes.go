package router

import (
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// RegisterUtilityRoutes 注册所有工具性路由
func RegisterUtilityRoutes(r *gin.Engine) {
	// 注册Swagger文档路由
	registerSwaggerRoutes(r)

	// 注册pprof性能分析路由
	registerPprofRoutes(r)

	// 注册ping路由组
	registerPingRoutes(r)

}

// registerSwaggerRoutes 注册Swagger文档路由
func registerSwaggerRoutes(r *gin.Engine) {
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

// registerPprofRoutes 注册pprof性能分析路由
func registerPprofRoutes(r *gin.Engine) {
	pprof.Register(r, "debug/pprof")
}

// registerPingRoutes 注册ping路由组
func registerPingRoutes(r *gin.Engine) {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})
}
