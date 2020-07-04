package UUtils

import (
	"github.com/ahviplc/GoJustToolc/UConsole"
	"testing"
)

// Test UMd5Util.Md5()
func TestUMd5UtilMd5(t *testing.T) {
	UConsole.Log(Md5("admin")) // 21232f297a57a5a743894a0e4a801fc3
}

// Test UMd5Util.GetGuid()
func TestUMd5UtilGetGuid(t *testing.T) {
	UConsole.Log(GetGuid()) // e70e70ead9c60eab4cfdfd51fb3adf8a
}
