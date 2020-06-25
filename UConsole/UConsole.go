package UConsole

import (
	"fmt"
	"strconv"
)

//打印控制台日志
func Log(args ...interface{}) {
	logData := ""
	for _, arg := range args {
		switch arg.(type) {
		case string:
			logData = logData + arg.(string)
		case int:
			logData = logData + strconv.Itoa(arg.(int))
		case int64:
			logData = logData + strconv.FormatInt(arg.(int64), 10)
		case float32:
			logData = logData + strconv.FormatFloat(float64(arg.(float32)), 'f', -1, 64)
		case float64:
			logData = logData + strconv.FormatFloat(arg.(float64), 'f', -1, 64)
		}
	}
	fmt.Println(logData)
}
