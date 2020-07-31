package URuntimeUtil

import "runtime"

// go运行时工具类

// 获取当前操作系统名称
// mac -> darwin
// win -> windows
// linux -> linux
func GetGOOS() string {
	return runtime.GOOS
}

// 获取当前操作系统架构
func GetGOARCH() string {
	return runtime.GOARCH
}

// 获取GOROOT 本机的GO路径 和 Go版本
func GetGOInfo() (string, string) {
	return runtime.GOROOT(), runtime.Version()
}

// 获取当前系统的CPU核数量
func GetNumCPU() int {
	return runtime.NumCPU()
	// 下面等价方法
	// return runtime.GOMAXPROCS(0)
}

// 获取正在执行和排队的任务总数
func GetNumGoroutine() int {
	return runtime.NumGoroutine()
}

// 是否是mac系统
func IsMacOS() bool {
	if GetGOOS() == "darwin" {
		return true
	}
	return false
}

// 是否是win系统
func IsWinOS() bool {
	if GetGOOS() == "windows" {
		return true
	}
	return false
}

// 是否是linux系统
func IsLinuxOS() bool {
	if GetGOOS() == "linux" {
		return true
	}
	return false
}
