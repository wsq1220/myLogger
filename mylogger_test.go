package mylogger

import "testing"

// 单元测试
func TestConstLevel(t *testing.T) {
	t.Logf("%v  %T\n", DebugLevel, DebugLevel)
	t.Logf("%v  %T\n", FatalLevel, FatalLevel)
}
