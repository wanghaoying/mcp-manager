package logger

import (
	"github.com/sirupsen/logrus"
	"mcp-manager/pkg/common"
	"mcp-manager/pkg/trace"
)

const requestID = "request_id"

type RequestIdHook struct{}

func (h *RequestIdHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (h *RequestIdHook) Fire(e *logrus.Entry) error {
	ctx := trace.GetGinCtx()
	if ctx == nil {
		return nil
	}
	rid := ctx.Value(common.HeaderXRequestID)
	if rid != nil {
		e.Data[requestID] = rid
	}
	return nil
}
