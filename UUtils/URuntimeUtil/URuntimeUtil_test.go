package URuntimeUtil

import (
	"github.com/ahviplc/GoJustToolc/UConsole"
	"runtime"
	"testing"
)

// Test URuntimeUtil.GetGOOS()
func TestGetGOOS(t *testing.T) {
	UConsole.Log(GetGOOS()) // darwin
}

// Test URuntimeUtil.GetGOARCH()
func TestGetGOARCH(t *testing.T) {
	UConsole.Log(GetGOARCH()) // amd64
}

// Test URuntimeUtil.GetNumCPU()
func TestGetNumCPU(t *testing.T) {
	UConsole.Log(GetNumCPU())           // 12
	UConsole.Log(runtime.GOMAXPROCS(0)) // 12
}

// Test URuntimeUtil.GetGOInfo()
func TestGetGOInfo(t *testing.T) {
	goRoot, goVersion := GetGOInfo()
	UConsole.Log(goRoot)    // 【/Volumes/MacOS-SSD-LCKu/DevelopSoftKu/go/GOROOT】
	UConsole.Log(goVersion) // go1.13.6
}

// Test URuntimeUtil.GetNumGoroutine()
func TestGetNumGoroutine(t *testing.T) {
	UConsole.Log(GetNumGoroutine()) // 2
}

// Test URuntimeUtil.IsMacOS()
func TestIsMacOS(t *testing.T) {
	UConsole.Log(IsMacOS()) // true
}

// Test URuntimeUtil.IsWinOS()
func TestIsWinOS(t *testing.T) {
	UConsole.Log(IsWinOS()) // false
}

// Test URuntimeUtil.IsLinuxOS()
func TestIsLinuxOS(t *testing.T) {
	UConsole.Log(IsLinuxOS()) // false
}
