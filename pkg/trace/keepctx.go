// Package trace 链路跟踪
package trace

import (
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/timandy/routine"
)

var ginCtxMap sync.Map

// SetGinCtx 设置gin上下文
func SetGinCtx(c *gin.Context) {
	ginCtxMap.Store(routine.Goid(), c)
}

// GetGinCtx 获取gin上下文
func GetGinCtx() (c *gin.Context) {
	val, _ := ginCtxMap.Load(routine.Goid())
	c, _ = val.(*gin.Context)
	return
}

// DelGinCtx 删除gin上下文
func DelGinCtx() {
	ginCtxMap.Delete(routine.Goid())
}
