package UConsole

import (
	"errors"
	"fmt"
	"github.com/ahviplc/GoJustToolc/UMode"
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

	// -----------------------------------------------------------------------------------------------------
	// 我是UConsole.Log()的测试输出
	// 110
	// 1.1
	// true
	// bool
	// string
	// false
	//-----------------------------------------------------------------------------------------------------

	// 默认Mode为 debug 下面可以输出.
	DebugPrint("Loaded This UConsoleLog (%d): \n%s\n", len("123"), "456")
	PrintAStraightLine()
	DebugPrint(`[WARNING] DebugPrint.`)
	PrintAStraightLine()

	// 改变Mode为 release 下面不可以输出.
	UMode.SetMode(UMode.ReleaseMode) // release

	err := errors.New("panic error.")
	DebugPrintError(err)
	PrintAStraightLine()

	// 改变Mode为debug 下面可以输出.
	UMode.SetMode(UMode.DebugMode) // debug
	DebugPrint(`[WARNING] DebugPrint2.`)
	PrintAStraightLine()

	// 改变Mode为test 下面不可以输出.
	UMode.SetMode(UMode.TestMode) // test
	DebugPrint(`[WARNING] DebugPrint3.`)
	PrintAStraightLine()

	// 改变Mode为debug 下面可以输出.
	UMode.SetMode(UMode.DebugMode) // debug
	DebugPrint(`[WARNING] DebugPrint2.1.`)
	PrintAStraightLine()

	err2 := errors.New("panic error 2.")
	DebugPrintError(err2)
	PrintAStraightLine()

	// -----------------------------------------------------------------------------------------------------
	// [GoJustToolc-debug] Loaded This UConsoleLog (3):
	// 456
	// -----------------------------------------------------------------------------------------------------
	// [GoJustToolc-debug] [WARNING] DebugPrint.
	// -----------------------------------------------------------------------------------------------------
	// -----------------------------------------------------------------------------------------------------
	// [GoJustToolc-debug] [WARNING] DebugPrint2.
	// -----------------------------------------------------------------------------------------------------
	// -----------------------------------------------------------------------------------------------------
	// [GoJustToolc-debug] [WARNING] DebugPrint2.1.
	// -----------------------------------------------------------------------------------------------------
	// [GoJustToolc-debug] [ERROR] panic error 2.
	// -----------------------------------------------------------------------------------------------------

	// pass
	t.Log("TestUConsoleLog pass")

}
