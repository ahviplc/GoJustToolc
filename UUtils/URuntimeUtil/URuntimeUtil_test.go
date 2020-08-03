package URuntimeUtil

import (
	"fmt"
	"github.com/ahviplc/GoJustToolc/UConsole"
	"os/exec"
	"runtime"
	"syscall"
	"testing"
	"time"
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

// ----------------------------------------------------------------------------------------

// ip

// Test URuntimeUtil.IP2Long()
func TestIP2Long(t *testing.T) {
	type args struct {
		ip string
	}
	tests := []struct {
		name string
		args args
		want uint32
	}{
		{
			name: "t1",
			args: args{ip: "192.0.34.166"},
			want: 3221234342,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IP2Long(tt.args.ip); got != tt.want {
				t.Errorf("IP2Long() = %v, want %v", got, tt.want)
			}
		})
	}
}

// Test URuntimeUtil.Long2IP()
func TestLong2IP(t *testing.T) {
	type args struct {
		ip uint32
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "t1",
			args: args{ip: 3221234342},
			want: "192.0.34.166",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Long2IP(tt.args.ip); got != tt.want {
				t.Errorf("Long2IP() = %v, want %v", got, tt.want)
			}
		})
	}
}

// Test URuntimeUtil.LocalIP()
func TestLocalIP(t *testing.T) {
	ip, err := LocalIP()
	if err != nil {
		panic(err)
	}
	UConsole.Log(ip) // 192.168.0.16
}

// Test URuntimeUtil.LocalDnsName()
func TestLocalDnsName(t *testing.T) {
	dnsName, err := LocalDnsName()
	if err != nil {
		panic(err)
	}
	UConsole.Log(dnsName)
}

// Test URuntimeUtil.GrabEphemeralPort()
func TestGrabEphemeralPort(t *testing.T) {
	port, err := GrabEphemeralPort()
	if err != nil {
		panic(err)
	}
	fmt.Println(port) // 52769
	// 将uint16强转成int
	UConsole.Log(int(port)) // // 52769
}

// Test URuntimeUtil.IntranetIP()
func TestIntranetIP(t *testing.T) {
	ip, err := IntranetIP()
	if err != nil {
		panic(err)
	}
	fmt.Println(ip) // [192.168.0.16 192.168.112.1 192.168.160.1 192.168.192.126]
}

// Test URuntimeUtil.IsIntranetIP()
func TestIsIntranetIP(t *testing.T) {
	UConsole.Log(IsIntranetIP("192.168.0.16"))
}

// ----------------------------------------------------------------------------------------

// Test URuntimeUtil.CmdOut()
func TestCmdOut(t *testing.T) {
	command := "go"
	params := []string{"version"}
	//执行cmd命令: go version
	out, err := CmdOut(command, params...)
	if err != nil {
		panic(err)
	}
	fmt.Println(out) // go version go1.13 windows/amd64
}

// Test URuntimeUtil.CmdOutBytes()
func TestCmdOutBytes(t *testing.T) {
	command := "go"
	params := []string{"version"}
	//执行cmd命令: go version
	out, err := CmdOutBytes(command, params...)
	if err != nil {
		panic(err)
	}
	fmt.Println(out) // [103 111 32 118 101 114 115 105 111 110 32 103 111 49 46 49 51 32 119 105 110 100 111 119 115 47 97 109 100 54 52 10]
}

// Test URuntimeUtil.CmdOutNoLn()
func TestCmdOutNoLn(t *testing.T) {
	command := "go"
	params := []string{"version"}
	//执行cmd命令: go version
	out, err := CmdOutNoLn(command, params...)
	if err != nil {
		panic(err)
	}
	fmt.Println(out) // go version go1.13 windows/amd64
}

// Test URuntimeUtil.CmdRunWithTimeout()
// todo 待完善测试用例
func TestCmdRunWithTimeout(t *testing.T) {
	cmd := exec.Command("go", "env")
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: false}
	cmd.Start() // attention!
	CmdRunWithTimeout(cmd, time.Duration(10)*time.Second)
	// 睡眠
	// time.Sleep(time.Hour * 1)
}
