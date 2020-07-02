package UConsole

import (
	"fmt"
	"strconv"
	"testing"
)

// Test UConsole.Log()
func TestUConsoleLog(t *testing.T) {
	// 打印一条直线
	PrintAStraightLine()

	// string
	Log("我是UConsole.Log()的测试输出")

	// int
	Log(110)

	// float64
	Log(1.1)

	// bool
	Log(true)
	fmt.Printf("%T", true)
	Log("")
	fmt.Printf("%T", strconv.FormatBool(true))
	Log("")
	Log(false)

	// 打印一条直线
	PrintAStraightLine()

	// pass
	t.Log("TestUConsoleLog pass")
}
