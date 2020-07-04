package UUtils

import (
	"strconv"
	"strings"
	"unicode"
)

// 字符串是否为空白 空白的定义如下:
// 1.为null
// 2.为不可见字符 (如空格)
// 3.""
func IsBlank(in string) bool {
	if len(in) == 0 || in == "" {
		return true
	}
	inRune := []rune(in)
	for _, r := range inRune {
		// 是否为空白符号
		if unicode.IsSpace(r) {
			return true
		}
	}
	return false
}

// 如果对象是字符串是否为空白 空白的定义如下:
// 1.为null
// 2.为不可见字符 (如空格)
// 3.""
func IsBlankIfStr(in interface{}) bool {
	_, isFlag := in.(string) // in2, isFlag := in.(string)
	if isFlag {
		return IsBlank(in.(string)) // return IsBlank(in2)
	}
	return false
}

// IsBlankIfStr 写法2 方法效果一样
func IsBlankIfStr2(in interface{}) bool {
	in2, isFlag := in.(string)
	if isFlag {
		return IsBlank(in2)
	}
	return false
}

// 字符串是否为非空白 非空白的定义如下:
// 1.不为null
// 2.不为不可见字符 (如空格)
// 3.不为""
func IsNotBlank(in string) bool {
	return false == IsBlank(in)
}

// 如果对象是字符串是否为非空白 非空白的定义如下:
// 1.不为null
// 2.不为不可见字符 (如空格)
// 3.不为""
func IsNotBlankIfStr(in interface{}) bool {
	_, isFlag := in.(string)
	if isFlag {
		return IsNotBlank(in.(string))
	}
	return false
}

// 字符串是否为空 空的定义如下:
// 1.为null
// 2.""
func IsEmpty(in string) bool {
	return len(in) == 0 || in == ""
}

// 如果对象是字符串是否为空字符串 空的定义如下:
// 1.为null
// 2.""
func IsEmptyIfStr(in interface{}) bool {
	_, isFlag := in.(string)
	if isFlag {
		return IsEmpty(in.(string))
	}
	return false
}

// 字符串是否为非空 非空的定义如下:
// 1.不为null
// 2.不为""
func IsNotEmpty(in string) bool {
	return false == IsEmpty(in)
}

// 如果对象是字符串是否为非空字符串 非空的定义如下:
// 1.不为null
// 2.不为""
func IsNotEmptyIfStr(in interface{}) bool {
	_, isFlag := in.(string)
	if isFlag {
		return IsNotEmpty(in.(string))
	}
	return false
}

// snake string, XxYy to xx_yy , XxYY to xx_yy
func SnakeString(s string) string {
	data := make([]byte, 0, len(s)*2)
	j := false
	num := len(s)
	for i := 0; i < num; i++ {
		d := s[i]
		if i > 0 && d >= 'A' && d <= 'Z' && j {
			data = append(data, '_')
		}
		if d != '_' {
			j = true
		}
		data = append(data, d)
	}
	return strings.ToLower(string(data[:]))
}

// camel string, xx_yy to XxYy
func CamelString(s string) string {
	data := make([]byte, 0, len(s))
	j := false
	k := false
	num := len(s) - 1
	for i := 0; i <= num; i++ {
		d := s[i]
		if k == false && d >= 'A' && d <= 'Z' {
			k = true
		}
		if d >= 'a' && d <= 'z' && (j || k == false) {
			d = d - 32
			j = false
			k = true
		}
		if k && d == '_' && num > i && s[i+1] >= 'a' && s[i+1] <= 'z' {
			j = true
			continue
		}
		data = append(data, d)
	}
	return string(data[:])
}

func CamelStringFirstLower(s string) string {
	data := make([]byte, 0, len(s))
	j := false
	k := false
	num := len(s) - 1
	for i := 0; i <= num; i++ {
		d := s[i]
		if k == false && d >= 'A' && d <= 'Z' {
			k = true
		}
		if d >= 'a' && d <= 'z' && (j || k == false) {
			d = d - 32
			j = false
			k = true
		}
		if k && d == '_' && num > i && s[i+1] >= 'a' && s[i+1] <= 'z' {
			j = true
			continue
		}
		if i == 0 && d >= 'A' && d <= 'Z' {
			d = d + 32
		}
		data = append(data, d)
	}
	return string(data[:])
}

// 字符串去除空格并将所有字母大写
func ToUpper(oldData string) string {
	return strings.ToUpper(strings.Replace(oldData, " ", "", -1))
}

// 字符串去除空格并将所有字母小写
func ToLower(oldData string) string {
	return strings.ToLower(strings.Replace(oldData, " ", "", -1))
}

// int转string
func IntToString(in int) string {
	return strconv.Itoa(in)
}

// string转int
func StringToInt(in string) int {
	out, err := strconv.Atoi(in)
	if err != nil {
		panic(err)
	}
	return out
}

// int64转string
func Int64ToString(in int64) string {
	return strconv.FormatInt(in, 10)
}

// string转int64
func StringToInt64(in string) int64 {
	out, err := strconv.ParseInt(in, 10, 64)
	if err != nil {
		panic(err)
	}
	return out
}
