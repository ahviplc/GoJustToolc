package URuntimeUtil

import (
	"bytes"
	"fmt"
	"github.com/ahviplc/GoJustToolc/UUtils/UStringUtil"
	"math"
	"net"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"time"
)

// go运行时工具类 runtime cmd命令执行
// 系统工具类 system 含ip处理相关

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

// ----------------------------------------------------------------------------------------

// ip dns

// IP2Long converts a string containing an (IPv4) Internet Protocol dotted address into a long integer.
func IP2Long(ip string) uint32 {
	ipv4 := net.ParseIP(ip).To4()
	if ipv4 == nil {
		return 0
	}
	return uint32(ipv4[0])<<24 | uint32(ipv4[1])<<16 | uint32(ipv4[2])<<8 | uint32(ipv4[3])
}

// Long2IP converts an long integer address into a string in (IPv4) Internet standard dotted format.
func Long2IP(ip uint32) string {
	if ip > math.MaxUint32 {
		return ""
	}
	return net.IPv4(byte(ip>>24), byte(ip>>16), byte(ip>>8), byte(ip)).String()
}

// LocalIP() 返回本地ip
func LocalIP() (string, error) {
	addr, err := net.ResolveUDPAddr("udp", "1.2.3.4:1")
	if err != nil {
		return "", err
	}
	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		return "", err
	}
	defer conn.Close()
	host, _, err := net.SplitHostPort(conn.LocalAddr().String())
	if err != nil {
		return "", err
	}
	// host = "10.180.2.66"
	return host, nil
}

// LocalDnsName()
// 本地dns名称
// just use in linux and mac
func LocalDnsName() (hostname string, err error) {
	if IsWinOS() {
		err = fmt.Errorf("LocalDnsName() can not use in winOS,just use in linux and mac")
		return "", err
	}
	var ip string
	ip, err = LocalIP()
	if err != nil {
		return
	}
	cmd := exec.Command("host", ip)
	var out bytes.Buffer
	cmd.Stdout = &out
	err = cmd.Run()
	if err != nil {
		return
	}
	tmp := out.String()
	arr := strings.Split(tmp, ".\n")
	if len(arr) > 1 {
		content := arr[0]
		arr = strings.Split(content, " ")
		return arr[len(arr)-1], nil
	}
	err = fmt.Errorf("parse host %s fail", ip)
	return
}

// GrabEphemeralPort()
//【uint16 无符号 16 位整型 (0 到 65535)】【int16 有符号 16 位整型 (-32768 到 32767)】
func GrabEphemeralPort() (port uint16, err error) {
	var listener net.Listener
	var portStr string
	var p int
	listener, err = net.Listen("tcp", ":0")
	if err != nil {
		return
	}
	defer listener.Close()
	_, portStr, err = net.SplitHostPort(listener.Addr().String())
	if err != nil {
		return
	}
	p, err = strconv.Atoi(portStr)
	port = uint16(p)
	return
}

// IntranetIP() 获取所有局域网ip
func IntranetIP() (ips []string, err error) {
	ips = make([]string, 0)
	ifaces, e := net.Interfaces()
	if e != nil {
		return ips, e
	}
	for _, iface := range ifaces {
		if iface.Flags&net.FlagUp == 0 {
			continue // interface down
		}
		if iface.Flags&net.FlagLoopback != 0 {
			continue // loopback interface
		}
		if strings.HasPrefix(iface.Name, "docker") || strings.HasPrefix(iface.Name, "w-") {
			continue
		}
		addrs, e := iface.Addrs()
		if e != nil {
			return ips, e
		}
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			if ip == nil || ip.IsLoopback() {
				continue
			}
			ip = ip.To4()
			if ip == nil {
				continue // not an ipv4 address
			}
			ipStr := ip.String()
			if IsIntranetIP(ipStr) {
				ips = append(ips, ipStr)
			}
		}
	}
	return ips, nil
}

// IsIntranetIP() 是否是局域网ip
func IsIntranetIP(ipStr string) bool {
	if strings.HasPrefix(ipStr, "10.") || strings.HasPrefix(ipStr, "192.168.") {
		return true
	}
	if strings.HasPrefix(ipStr, "172.") {
		// 172.16.0.0-172.31.255.255
		arr := strings.Split(ipStr, ".")
		if len(arr) != 4 {
			return false
		}

		second, err := strconv.ParseInt(arr[1], 10, 64)
		if err != nil {
			return false
		}

		if second >= 16 && second <= 31 {
			return true
		}
	}
	return false
}

// ----------------------------------------------------------------------------------------

// command 执行命令

// CmdOut()
// 返回字符串执行结果
func CmdOut(name string, arg ...string) (string, error) {
	cmd := exec.Command(name, arg...)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	return out.String(), err
}

// CmdOutBytes()
// 返回字节切片执行结果
func CmdOutBytes(name string, arg ...string) ([]byte, error) {
	cmd := exec.Command(name, arg...)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	return out.Bytes(), err
}

// CmdOutNoLn()
// 返回字符串执行结果
// 输出结果 去掉换行,空格,回车,制表
func CmdOutNoLn(name string, arg ...string) (out string, err error) {
	out, err = CmdOut(name, arg...)
	if err != nil {
		return
	}
	return UStringUtil.TrimRightSpace(string(out)), nil
}

// CmdRunWithTimeout()
// 带有超时时间的执行cmd
func CmdRunWithTimeout(cmd *exec.Cmd, timeout time.Duration) (error, bool) {
	done := make(chan error)
	go func() {
		fmt.Println("CmdRunWithTimeout() run...")
		done <- cmd.Wait()
	}()

	var err error
	select {
	case <-time.After(timeout):
		//timeout
		if err = cmd.Process.Kill(); err != nil {
			fmt.Printf("failed to kill: %s, error: %s", cmd.Path, err)
		}
		go func() {
			<-done // allow goroutine to exit
		}()
		fmt.Printf("process:%s killed", cmd.Path)
		return err, true
	case err = <-done:
		return err, false
	}
}

// ----------------------------------------------------------------------------------------
