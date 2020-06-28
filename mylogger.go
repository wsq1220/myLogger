package mylogger

import "strings"

// 我的日志库文件

// Level 是一个自定义的类型 代表日志级别
type Level uint16

// 定义具体的日志级别常量
const (
	DebugLevel Level = iota
	InfoLevel
	WarningLevel
	ErrorLevel
	FatalLevel
)

// Logger 定义一个logger接口
type Logger interface {
	Debug(format string, args ...interface{})
	Info(format string, args ...interface{})
	Warn(format string, args ...interface{})
	Error(format string, args ...interface{})
	Fatal(format string, args ...interface{})
	Close()
}

// 写一个根据传进来的Level 获取对应的字符串
func getLevelStr(level Level) string {
	switch level {
	case DebugLevel:
		return "DEBUG"
	case InfoLevel:
		return "INFO"
	case WarningLevel:
		return "WARN"
	case ErrorLevel:
		return "Error"
	case FatalLevel:
		return "FATAL"
	default:
		return "DEBUG"
	}
}

// 根据用户传入的字符串类型的日志级别，解析出对应的Level
func parseLogLevel(levelStr string) Level {
	levelStr = strings.ToLower(levelStr) //将字符串转换成全小写
	switch levelStr {
	case "debug":
		return DebugLevel
	case "info":
		return InfoLevel
	case "warn":
		return WarningLevel
	case "error":
		return ErrorLevel
	case "fatal":
		return FatalLevel
	default:
		return DebugLevel
	}
}
