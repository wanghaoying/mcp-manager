package common

import (
	"mcp-manager/pkg/trace"

	"github.com/gin-gonic/gin"
)

// Response 通用响应结构体
// code: 业务码，0为成功，非0为失败
// message: 提示信息
// data: 返回数据
// trace_id: 链路追踪ID

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	TraceID string      `json:"trace_id,omitempty"`
}

// getTraceIDFromCtx 从 ginCtxMap 获取 trace_id
func getTraceIDFromCtx() string {
	ctx := trace.GetGinCtx()
	if ctx == nil {
		return ""
	}
	if v, ok := ctx.Get("trace_id"); ok {
		if s, ok := v.(string); ok {
			return s
		}
	}
	return ""
}

// Success 返回成功响应
func Success(c *gin.Context, data interface{}) {
	traceID := getTraceIDFromCtx()
	c.JSON(200, Response{
		Code:    0,
		Message: "success",
		Data:    data,
		TraceID: traceID,
	})
}

// Error 返回错误响应
func Error(c *gin.Context, code int, message string) {
	traceID := getTraceIDFromCtx()
	c.JSON(200, Response{
		Code:    code,
		Message: message,
		TraceID: traceID,
	})
}
