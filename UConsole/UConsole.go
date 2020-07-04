package UConsole

import (
	"fmt"
	"github.com/ahviplc/GoJustToolc/UUtils"
	"strconv"
)

// 打印控制台日志
func Log(args ...interface{}) {
	logData := ""
	for _, arg := range args {
		switch arg.(type) {
		case string:
			logData = logData + arg.(string)
		case int:
			logData = logData + UUtils.IntToString(arg.(int))
		case int64:
			logData = logData + UUtils.Int64ToString(arg.(int64))
		case float32:
			logData = logData + strconv.FormatFloat(float64(arg.(float32)), 'f', -1, 64)
		case float64:
			logData = logData + strconv.FormatFloat(arg.(float64), 'f', -1, 64)
		case bool:
			logData = logData + strconv.FormatBool(arg.(bool))
		}
	}
	fmt.Println(logData)
}

// 打印一条直线并换行 类似: ----------- * 10
func PrintAStraightLine() {
	Log("-----------------------------------------------------------------------------------------------------")
}

// 打印类型和值
func PrintTypeAndValue(in ...interface{}) {
	// 备注: 这里len(in) == 0 则会传进来 nil 也是可以range操作的 下面的代码不写也可
	// if len(in) == 0 {
	// 	return
	// }

	// 遍历 in
	for _, v := range in {
		fmt.Printf("类型(Type):%T  值(Value):%v", v, v)
		Log()
	}
}
