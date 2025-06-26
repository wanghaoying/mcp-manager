// Package logger
package logger

// @Title  logger.go
// @Description  日志包
// @Author  socketwang  2025/6/25 13:16
// @Update  socketwang  2025/6/25 13:16

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/pkg/errors"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

// SetDebugLogLevel sets the level of debug log
func SetDebugLogLevel(level string) error {
	logLevel, err := logrus.ParseLevel(level)
	if err != nil {
		return err
	}

	oldLevel := logrus.GetLevel()
	if logLevel == oldLevel {
		return nil
	}

	// 修改日志级别
	logrus.SetLevel(logLevel)
	return nil
}

// InitDebugLogger inits debug logger
func InitDebugLogger(level string, path string) error {
	// 初始化打文件的logrus
	file, err := rotatelogs.New(
		path,
		rotatelogs.WithMaxAge(7*24*time.Hour),     // 文件最大保存时间
		rotatelogs.WithRotationTime(24*time.Hour), // 日志切割时间间隔
	)

	if err != nil {
		return err
	}

	logLevel, err := logrus.ParseLevel(level)
	if err != nil {
		logLevel = logrus.WarnLevel
	}

	logrus.SetLevel(logLevel)
	logrus.SetReportCaller(true)
	lfHookDebug := lfshook.NewHook(lfshook.WriterMap{
		logrus.DebugLevel: file, // 为不同级别设置不同的输出目的
		logrus.InfoLevel:  file,
		logrus.WarnLevel:  file,
		logrus.ErrorLevel: file,
		logrus.FatalLevel: file,
		logrus.PanicLevel: file,
		logrus.TraceLevel: file,
	}, &IdataxFormatter{})
	if lfHookDebug == nil {
		return errors.New("new debug log hook failed")
	}

	logrus.AddHook(&RequestIdHook{})
	logrus.AddHook(lfHookDebug)
	logrus.SetOutput(ioutil.Discard)
	return nil
}

// IdataxFormatter TODO
type IdataxFormatter struct{}

// Format TODO
func (f *IdataxFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}
	requestID := entry.Data[requestID]
	if requestID == nil {
		requestID = ""
	}

	timestamp := entry.Time.Format("2006-01-02 15:04:05.999")
	var newLog string
	if entry.HasCaller() {
		newLog = fmt.Sprintf("[%s] [%s] [%s] [%s:%d] %s\n",
			timestamp, requestID, entry.Level, entry.Caller.File, entry.Caller.Line, entry.Message)
	} else {
		newLog = fmt.Sprintf("[%s] [%s] [%s] %s\n", timestamp, requestID, entry.Level, entry.Message)
	}

	b.WriteString(newLog)
	return b.Bytes(), nil
}

// PlainTextFormatter defines a plain text formatter
type PlainTextFormatter struct {
	logrus.TextFormatter
}

// Format formats a log
func (f *PlainTextFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	return []byte(fmt.Sprintf("%+v", entry.Message)), nil
}
