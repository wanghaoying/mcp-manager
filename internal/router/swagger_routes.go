package router

import (
	"github.com/gin-gonic/gin"
	"mcp-manager/internal/controller"
)

// RegisterSwaggerHandlers 注册swagger相关handler
func RegisterSwaggerHandlers(r *gin.Engine) {
	r.POST("/api/swagger/validate/file", controller.ValidateSwaggerByFile)
	r.POST("/api/swagger/validate/text", controller.ValidateSwaggerByText)
}
