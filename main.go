package main

import (
	"fmt"
	"github.com/ahviplc/GoJustToolc/UConfig"
	"github.com/ahviplc/GoJustToolc/UConsole"
	"github.com/ahviplc/GoJustToolc/UMode"
	"github.com/ahviplc/GoJustToolc/UUtils/UDateTimeUtil"
	"github.com/ahviplc/GoJustToolc/UUtils/UStringUtil"
	"github.com/tidwall/gjson"
	"time"
)

// some test main
func main() {
	// ToUpper ToLower
	fmt.Println(UStringUtil.ToUpper("ah vip lc"))
	fmt.Println(UStringUtil.ToLower("SH vip LC"))
	UConsole.Log(UStringUtil.ToUpper("ah vip lc"))
	UConsole.Log(UStringUtil.ToLower("SH vip LC"))

	// UConsole
	UConsole.Log("我就是一个长长的字符串")
	UConsole.Log("string1", "string2")
	UConsole.Log(110, 120)
	UConsole.Log("我是字符串", 2)
	UConsole.Log(1.1, 2.2, 3.3)
	UConsole.Log(1.11, 2.22, 3.33)
	UConsole.Log(1.11, 2.222, 3.3333)

	// gjson
	const json = `{"name":{"first":"ahviplc","last":"shviplc"},"age":18}`
	value := gjson.Get(json, "name.first")
	UConsole.Log(value.String())
	UConsole.Log(gjson.Get(json, "name.last").String())
	UConsole.Log(gjson.Get(json, "age").String())

	// UConsole.AUTHORS UConsole.VERSION
	UConsole.Log(UConsole.AUTHORS)
	UConsole.Log(UConsole.VERSION)

	// 输出当前日期时间
	UConsole.Log(UDateTimeUtil.NowDateTime())
	UConsole.Log(UDateTimeUtil.NowDateTime(UDateTimeUtil.DefaultLayout))
	UConsole.Log(UDateTimeUtil.NowDateTime("20060102150405"))
	UConsole.PrintAStraightLine()

	// time.Now()
	UConsole.PrintTypeAndValue(time.Now()) // 类型(Type):time.Time  值(Value):2020-07-03 23:16:10.98863 +0800 CST m=+0.000319369
	// UTC
	UConsole.PrintTypeAndValue(UDateTimeUtil.GetNowDateTimeUTC())
	UConsole.PrintTypeAndValue(UDateTimeUtil.StrTimeToTime(UDateTimeUtil.NowDateTime(), UDateTimeUtil.DefaultLayout).UTC()) // 类型(Type):time.Time  值(Value):2020-07-03 15:11:05 +0000 UTC
	// CST 比UTC多8小时(也就是多28800秒)
	UConsole.PrintTypeAndValue(UDateTimeUtil.GetNowDateTimeCST())
	UConsole.PrintTypeAndValue(UDateTimeUtil.StrTimeToTime(UDateTimeUtil.NowDateTime(), UDateTimeUtil.DefaultLayout)) // 类型(Type):time.Time  值(Value):2020-07-03 23:11:05 +0800 CST

	// 打印出当前Mode
	UConsole.Log(UMode.UMode) // 0 默认是0为dubug模式
	// 改变Mode为 release.
	UMode.SetMode(UMode.ReleaseMode)
	UConsole.Log(UMode.UMode) // 1 改成了1为release模式

	// UConfig
	UConsole.PrintAStraightLine()
	v := UConfig.InitUConfig("UConfig", "demo", "json")
	UConsole.Log(UConfig.GetJson(v, "author", false))            // LC<ahlc@sina.cn>
	UConsole.Log(UConfig.GetJson(v, "GoToolcUrl", false))        // https://github.com/ahviplc/GoJustToolc
	UConsole.Log(UConfig.GetJson(v, "message", false))           // I am root
	UConsole.Log(UConfig.GetJson(v, "contents.post1", false))    // 1
	UConsole.Log(UConfig.GetJson(v, "contents2.#.post1", false)) // ["3-1","3-2"]
	UConsole.PrintAStraightLine()
}
