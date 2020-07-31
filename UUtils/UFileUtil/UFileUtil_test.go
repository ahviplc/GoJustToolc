package UFileUtil

import (
	"fmt"
	"github.com/ahviplc/GoJustToolc/UConsole"
	"github.com/ahviplc/GoJustToolc/UUtils/UDateTimeUtil"
	"io/ioutil"
	"testing"
)

// Test UFileUtil.ListDirFile()
// 获取指定目录下的所有文件，不进入下一级目录搜索，可以匹配后缀过滤。
func TestListDirFile(t *testing.T) {
	list, err := ListDirFile(GetRootDir(), "ini")
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, v := range list {
		fmt.Println(v)
	}
	// 输出:
	// C:\Users\Administrator\AppData\Local\Temp\\SmartInfo.ini
	// C:\Users\Administrator\AppData\Local\Temp\\SogouPinyin.ini
	// C:\Users\Administrator\AppData\Local\Temp\\popup.ini
	// C:\Users\Administrator\AppData\Local\Temp\\version.ini
}

// Test UFileUtil.WalkDirFile()
// 获取指定目录及所有子目录下的所有文件，可以匹配后缀过滤。
func TestWalkDirFile(t *testing.T) {
	list, err := WalkDirFile(GetRootDir(), "txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, v := range list {
		fmt.Println(v)
	}
	// 输出:
	// C:\Users\Administrator\AppData\Local\Temp\VSCode Crashes\operation_log.txt
	// C:\Users\Administrator\AppData\Local\Temp\sgsi\LC.txt
}

// Test UFileUtil.GetCurrentDirectory()
func TestGetCurrentDirectory(t *testing.T) {
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
func TestCheckFileIsExist(t *testing.T) {
	UConsole.Log(CheckFileIsExist("UFileUtil.go"))        // true
	UConsole.Log(CheckFileIsExist("UFileUtilNoExist.go")) // false
}

// Test UFileUtil.SelfPath()
func TestSelfPath(t *testing.T) {
	out := SelfPath()
	UConsole.Log(out) // 【C:\Users\Administrator\AppData\Local\Temp\___TestSelfPath_in_github_com_ahviplc_GoJustToolc_UUtils_UFileUtil.exe】
}

// Test UFileUtil.SelfDir()
func TestSelfDir(t *testing.T) {
	out := SelfDir()
	UConsole.Log(out) // 【C:\Users\Administrator\AppData\Local\Temp】
}

// 定义一个文件路径
var filePath = "d:/LC.txt"

// init 写入文件内容
func init() {
	content := "1001"
	data := []byte(content)
	if ioutil.WriteFile(filePath, data, 0666) == nil {
		fmt.Println("写入成功")
	}
}

// Test UFileUtil.Basename()
// 获取文件名称
func TestBasename(t *testing.T) {
	UConsole.Log(Basename(filePath)) // LC.txt
}

// Test UFileUtil.Dir()
// 获取文件目录
func TestDir(t *testing.T) {
	UConsole.Log(Dir(filePath)) // c:
}

// Test UFileUtil.FileToInt64()
// 将文件内的数字转换成int64
func TestFileToInt64(t *testing.T) {
	rs, err := FileToInt64(filePath)
	if err != nil {
		t.Error("FileToInt64 err", err.Error())
	}
	fmt.Println(rs) // 1001
}

// Test UFileUtil.FileToUint64()
// 将文件内的数字转换成无符号的int64
func TestFileToUint64(t *testing.T) {
	rs, err := FileToUint64(filePath)
	if err != nil {
		t.Error("FileToUint64 err", err.Error())
	}
	fmt.Println(rs) // 1001
}

// Test UFileUtil.Ext()
// 获取文件类型后缀
func TestExt(t *testing.T) {
	out := Ext(filePath)
	fmt.Println(out) // .txt
}

// Test UFileUtil.Rename()
// 文件重命名
func TestRename(t *testing.T) {
	Rename(filePath, Dir(filePath)+"/"+"LC2.txt")
}

// Test UFileUtil.Unlink()
// 删除文件
func TestUnlink(t *testing.T) {
	Unlink(filePath)
}

// Test UFileUtil.IsFile()
// 获取一个路径是否是文件
func TestIsFile(t *testing.T) {
	out := IsFile(filePath)
	fmt.Println(out) // true
}

// Test UFileUtil.IsDir()
// 获取一个路径是否是目录
func TestIsDir(t *testing.T) {
	out := IsDir(filePath)
	fmt.Println(out) // false
}

// Test UFileUtil.RealPath()
// 获取文件真实路径
func TestRealPath(t *testing.T) {
	out, err := RealPath("UFileUtil.go")
	if err != nil {
		t.Errorf("RealPath err: %v \n", err.Error())
	}
	fmt.Println(out) // 【C:\_developSoftKu\ideaIU-2019.1.3.win\#GOPATHCodeKu\src\GoJustToolc\UUtils\UFileUtil/UFileUtil.go】
}

// Test UFileUtil.FileMTime()
// 获取文件修改时间戳
func TestFileMTime(t *testing.T) {
	rs, err := FileMTime(filePath)
	if err != nil {
		t.Errorf("FileMTime err: %v \n", err.Error())
	}
	fmt.Println(UDateTimeUtil.TimeStampToTimeStr(rs)) // 1596185828 -> 2020-07-31 16:57:08
}

// Test UFileUtil.FileSize()
// 获取文件大小
func TestFileSize(t *testing.T) {
	size, err := FileSize(filePath)
	if err != nil {
		t.Errorf("FileSize err: %v \n", err.Error())
	}
	fmt.Println(size) // 6 -> 6 bytes 代表 6字节
}

// Test UFileUtil.DirsUnder()
// 获取路径下所有的文件夹Dirs
func TestDirsUnder(t *testing.T) {
	rs, err := DirsUnder(GetRootDir())
	if err != nil {
		t.Error("DirsUnder: ", err.Error())
	}
	fmt.Println(rs)
}

// Test UFileUtil.FilesUnder()
// 获取路径下所有的文件Files
func TestFilesUnder(t *testing.T) {
	rs, err := FilesUnder(GetRootDir())
	if err != nil {
		t.Error("FilesUnder: ", err.Error())
	}
	fmt.Println(rs)
}
