package UFileUtil

import (
	"github.com/ahviplc/GoJustToolc/UConsole"
	"testing"
)

// Test UFileUtil.GetCurrentDirectory()
func TestUFileUtilGetCurrentDirectory(t *testing.T) {
	UConsole.Log(GetCurrentDirectory()) // 【/private/var/folders/lc/jvxkwmxn4xz60k5qqp426j9w0000gn/T】【win -> C:/Users/Administrator/AppData/Local/Temp】
}

// Test UFileUtil.GetRootDir()
func TestGetRootDir(t *testing.T) {
	UConsole.Log(GetRootDir()) // 【C:\Users\Administrator\AppData\Local\Temp\】
}

// Test UFileUtil.GetExecFilePath()
func TestGetExecFilePath(t *testing.T) {
	UConsole.Log(GetExecFilePath()) // 【C:\Users\Administrator\AppData\Local\Temp\___TestGetExecFilePath_in_github_com_ahviplc_GoJustToolc_UUtils_UFileUtil.exe】
}

// Test UFileUtil.CheckFileIsExist()
func TestUFileUtilCheckFileIsExist(t *testing.T) {
	UConsole.Log(CheckFileIsExist("UFileUtil.go"))        // true
	UConsole.Log(CheckFileIsExist("UFileUtilNoExist.go")) // false
}
