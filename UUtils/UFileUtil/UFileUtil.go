package UFileUtil

import (
	"errors"
	"fmt"
	"github.com/ahviplc/GoJustToolc/UUtils/UStringUtil"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strconv"
	"strings"
)

// 获取指定目录下的所有文件，不进入下一级目录搜索，可以匹配后缀过滤。
func ListDirFile(dirPth string, suffix string) (files []string, err error) {
	files = make([]string, 0, 10)
	dir, err := ioutil.ReadDir(dirPth)
	if err != nil {
		return nil, err
	}
	PthSep := string(os.PathSeparator)
	suffix = strings.ToUpper(suffix) //忽略后缀匹配的大小写
	for _, fi := range dir {
		if fi.IsDir() { // 忽略目录
			continue
		}
		if strings.HasSuffix(strings.ToUpper(fi.Name()), suffix) { //匹配文件
			files = append(files, dirPth+PthSep+fi.Name())
		}
	}
	return files, nil
}

// 获取指定目录及所有子目录下的所有文件，可以匹配后缀过滤。
func WalkDirFile(dirPth, suffix string) (files []string, err error) {
	files = make([]string, 0, 30)
	suffix = strings.ToUpper(suffix)                                                     //忽略后缀匹配的大小写
	err = filepath.Walk(dirPth, func(filename string, fi os.FileInfo, err error) error { //遍历目录
		//if err != nil { //忽略错误
		// return err
		//}
		if fi.IsDir() { // 忽略目录
			return nil
		}
		if strings.HasSuffix(strings.ToUpper(fi.Name()), suffix) {
			files = append(files, filename)
		}
		return nil
	})
	return files, err
}

// 获取当前目录
// 获取当前执行文件的目录
func GetCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	// 返回绝对路径
	// filepath.Dir(os.Args[0])
	if err != nil {
		log.Fatal(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}

// 获取执行文件当前的目录 最后带/
func GetRootDir() string {
	// 文件不存在获取执行路径
	file, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		file = fmt.Sprintf(".%s", string(os.PathSeparator))
	} else {
		file = fmt.Sprintf("%s%s", file, string(os.PathSeparator))
	}
	return file
}

// 获取执行文件的路径
func GetExecFilePath() string {
	file, err := exec.LookPath(os.Args[0])
	if err != nil {
		file = fmt.Sprintf(".%s", string(os.PathSeparator))
	} else {
		file, err = filepath.Abs(file)
		if err != nil {
			log.Fatal(err.Error())
		}
	}
	return file
}

// 判断文件是否存在  存在返回 true 不存在返回false
func CheckFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

// SelfPath gets compiled executable file absolute path 等价于方法 GetExecFilePath()
func SelfPath() string {
	pathTemp, _ := filepath.Abs(os.Args[0])
	return pathTemp
}

// SelfDir gets compiled executable file directory 等价于方法 GetRootDir()
func SelfDir() string {
	return filepath.Dir(SelfPath())
}

// get filepath base name
// 获取文件名称
func Basename(file string) string {
	return path.Base(file)
}

// get filepath dir name
// 获取文件目录
func Dir(file string) string {
	return path.Dir(file)
}

// InsureDir
func InsureDir(path string) error {
	if IsExist(path) {
		return nil
	}
	return os.MkdirAll(path, os.ModePerm)
}

// Ext
// 获取文件类型后缀
func Ext(file string) string {
	return path.Ext(file)
}

// rename file name
// 文件重命名
func Rename(file string, to string) error {
	return os.Rename(file, to)
}

// delete file
// 删除文件
func Unlink(file string) error {
	return os.Remove(file)
}

// IsFile checks whether the path is a file,
// it returns false when it's a directory or does not exist.
func IsFile(filePath string) bool {
	f, e := os.Stat(filePath)
	if e != nil {
		return false
	}
	return !f.IsDir()
}

// IsDir checks whether the path is a dir,
// it returns false when it's a file or does not exist.
func IsDir(filePath string) bool {
	f, e := os.Stat(filePath)
	if e != nil {
		return false
	}
	return f.IsDir()
}

// IsExist checks whether a file or directory exists.
// It returns false when the file or directory does not exist.
func IsExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

// Search a file in paths.
// this is often used in search config file in /etc ~/
func SearchFile(filename string, paths ...string) (fullPath string, err error) {
	for _, pathTemp := range paths {
		if fullPath = filepath.Join(pathTemp, filename); IsExist(fullPath) {
			return
		}
	}
	err = errors.New(fullPath + " not found in paths")
	return
}

// get absolute filepath, based on built executable file
// 获取文件真实路径
func RealPath(file string) (string, error) {
	if path.IsAbs(file) {
		return file, nil
	}
	wd, err := os.Getwd()
	return path.Join(wd, file), err
}

// get file modified time
// 获取文件修改时间戳
func FileMTime(file string) (int64, error) {
	f, e := os.Stat(file)
	if e != nil {
		return 0, e
	}
	return f.ModTime().Unix(), nil
}

// get file size as how many bytes
// 返回的是字节大小
func FileSize(file string) (int64, error) {
	f, e := os.Stat(file)
	if e != nil {
		return 0, e
	}
	return f.Size(), nil
}

// list dirs under dirPath
// 获取路径下所有的文件夹Dirs
func DirsUnder(dirPath string) ([]string, error) {
	if !IsExist(dirPath) {
		return []string{}, nil
	}
	fs, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return []string{}, err
	}
	sz := len(fs)
	if sz == 0 {
		return []string{}, nil
	}
	ret := []string{}
	for i := 0; i < sz; i++ {
		if fs[i].IsDir() {
			name := fs[i].Name()
			if name != "." && name != ".." {
				ret = append(ret, name)
			}
		}
	}
	return ret, nil
}

// list files under dirPath
// 获取路径下所有的文件Files
func FilesUnder(dirPath string) ([]string, error) {
	if !IsExist(dirPath) {
		return []string{}, nil
	}
	fs, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return []string{}, err
	}
	sz := len(fs)
	if sz == 0 {
		return []string{}, nil
	}
	ret := []string{}
	for i := 0; i < sz; i++ {
		if !fs[i].IsDir() {
			ret = append(ret, fs[i].Name())
		}
	}
	return ret, nil
}

// -------------------------------------------------

// ---------------FileToUint64 FileToInt64-----------------

func FileToUint64(file string) (uint64, error) {
	content, err := ReadFileToStringNoLn(file)
	if err != nil {
		return 0, err
	}

	var ret uint64
	if ret, err = strconv.ParseUint(content, 10, 64); err != nil {
		return 0, err
	}
	return ret, nil
}

func FileToInt64(file string) (int64, error) {
	content, err := ReadFileToStringNoLn(file)
	if err != nil {
		return 0, err
	}

	var ret int64
	if ret, err = strconv.ParseInt(content, 10, 64); err != nil {
		return 0, err
	}
	return ret, nil
}

// ---------------reader-----------------

// ReadFileToBytes reads data type '[]byte' from file by given path.
// It returns error when fail to finish operation.
func ReadFileToBytes(filePath string) ([]byte, error) {
	b, err := ioutil.ReadFile(filePath)
	if err != nil {
		return []byte(""), err
	}
	return b, nil
}

// ReadFileToString reads data type 'string' from file by given path.
// It returns error when fail to finish operation.
func ReadFileToString(filePath string) (string, error) {
	b, err := ReadFileToBytes(filePath)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func ReadFileToStringNoLn(filePath string) (string, error) {
	str, err := ReadFileToString(filePath)
	if err != nil {
		return "", err
	}
	return UStringUtil.TrimRightSpace(str), nil
}

// ---------------writer-----------------

// WriteBytesToFile saves content type '[]byte' to file by given path.
// It returns error when fail to finish operation.
func WriteBytesToFile(filePath string, b []byte) (int, error) {
	os.MkdirAll(path.Dir(filePath), os.ModePerm)
	fw, err := os.Create(filePath)
	if err != nil {
		return 0, err
	}
	defer fw.Close()
	return fw.Write(b)
}

// WriteStringFile saves content type 'string' to file by given path.
// It returns error when fail to finish operation.
func WriteStringToFile(filePath string, s string) (int, error) {
	return WriteBytesToFile(filePath, []byte(s))
}

// -------------------------------------------------
