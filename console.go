package mylogger

import (
	"fmt"
	"os"
	"time"
)

/*
   往终端打印日志
*/

// ConsoleLogger
type ConsoleLogger struct {
	level Level
}

// NewConsoleLogger 构造函数
func NewConsoleLogger(levelStr string) *ConsoleLogger {
	logLevel := parseLogLevel(levelStr)
	return &ConsoleLogger{
		level: logLevel,
	}

}

func (c *ConsoleLogger) log(level Level, format string, args ...interface{}) {
	if c.level > level {
		return
	}
	msg := fmt.Sprintf(format, args...) // 格式化
	// 日志格式：[时间][文件:行号][函数名][日志级别] 日志信息
	nowStr := time.Now().Format("2006-01-02 15:04:05.000")
	fileName, line, funcName := getCallerInfo(3)
	logLevelStr := getLevelStr(level)
	logMsg := fmt.Sprintf("[%s][%s:%d][%s][%s]%s",
		nowStr, fileName, line, funcName, logLevelStr, msg)
	fmt.Fprintln(os.Stdout, logMsg)
}

// Debug
func (c *ConsoleLogger) Debug(format string, args ...interface{}) {
	c.log(DebugLevel, format, args...)
}

// Info
func (c *ConsoleLogger) Info(format string, args ...interface{}) {
	c.log(InfoLevel, format, args...)
}

// Warn
func (c *ConsoleLogger) Warn(format string, args ...interface{}) {
	c.log(WarningLevel, format, args...)
}

// Error
func (c *ConsoleLogger) Error(format string, args ...interface{}) {
	c.log(ErrorLevel, format, args...)
}

// Fatal
func (c *ConsoleLogger) Fatal(format string, args ...interface{}) {
	c.log(FatalLevel, format, args...)
}

// console不需要关闭
func (c *ConsoleLogger) Close() {

}
