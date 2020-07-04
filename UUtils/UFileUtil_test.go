package UUtils

import (
	"github.com/ahviplc/GoJustToolc/UConsole"
	"testing"
)

// Test UFileUtil.GetCurrentDirectory()
func TestUFileUtilGetCurrentDirectory(t *testing.T) {
	UConsole.Log(GetCurrentDirectory()) // 【/private/var/folders/lc/jvxkwmxn4xz60k5qqp426j9w0000gn/T】
}

// Test UFileUtil.CheckFileIsExist()
func TestUFileUtilCheckFileIsExist(t *testing.T) {
	UConsole.Log(CheckFileIsExist("UFileUtil.go"))        // true
	UConsole.Log(CheckFileIsExist("UFileUtilNoExist.go")) // false
}
