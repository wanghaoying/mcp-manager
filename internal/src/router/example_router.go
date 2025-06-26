package router

import "github.com/gin-gonic/gin"

// RegisterExampleRoutes 注册示例路由组
func RegisterExampleRoutes(r *gin.Engine) {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})
}
