package UConsole

import (
	"testing"
)

// Test UConsole.Log()
func TestUConsoleLog(t *testing.T) {
	Log("我是UConsole.Log()的测试输出")
	t.Log("TestUConsoleLog pass")
}
