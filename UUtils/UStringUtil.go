package UUtils

import "strings"

//字符串去除空格并将所有字母大写
func ToUpper(oldData string) string {
	return strings.ToUpper(strings.Replace(oldData, " ", "", -1))
}

//字符串去除空格并将所有字母小写
func ToLower(oldData string) string {
	return strings.ToLower(strings.Replace(oldData, " ", "", -1))
}
