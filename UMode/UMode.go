package UMode

import (
	"io"
	"os"
)

// EnvGoJustToolcMode indicates environment name for GoJustToolc mode.
const EnvGoJustToolcMode = "GoJustToolc_MODE"

// 定义UMode全部模式
const (
	// DebugMode indicates GoJustToolc mode is debug.
	DebugMode = "debug"
	// ReleaseMode indicates GoJustToolc mode is release.
	ReleaseMode = "release"
	// TestMode indicates GoJustToolc mode is test.
	TestMode = "test"
)

// 定义UMode全部模式代号Code
const (
	DebugCode   = iota // 0
	ReleaseCode        // 1
	TestCode           // 2
)

// DefaultWriter is the default io.Writer used by GoJustToolc for debug output and
var DefaultWriter io.Writer = os.Stdout

// DefaultErrorWriter is the default io.Writer used by GoJustToolc to debug errors
var DefaultErrorWriter io.Writer = os.Stderr

var UMode = DebugCode
var modeName = DebugMode

// init函数
//func init() {
//	mode := os.Getenv(EnvGoJustToolcMode)
//	SetMode(mode)
//}

// 设置Mode方法 改变GoJustToolc模式
// SetMode sets GoJustToolc mode according to input string.
func SetMode(value string) {
	switch value {
	case DebugMode, "":
		UMode = DebugCode
	case ReleaseMode:
		UMode = ReleaseCode
	case TestMode:
		UMode = TestCode
	default:
		panic("GoJustToolc mode unknown: " + value)
	}
	if value == "" {
		value = DebugMode
	}
	modeName = value
}

// Mode returns currently GoJustToolc mode.
func Mode() string {
	return modeName
}
