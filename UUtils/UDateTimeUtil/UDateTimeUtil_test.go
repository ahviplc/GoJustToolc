package UDateTimeUtil

import (
	"fmt"
	"github.com/ahviplc/GoJustToolc/UConsole"
	"testing"
	"time"
)

// 输出当前日期时间
// Test UDateTimeUtil.NowDateTime()
func TestNowDateTime(t *testing.T) {
	UConsole.PrintAStraightLine()

	// 默认输出layout yyyy-MM-dd HH:mm:ss
	UConsole.Log(NowDateTime())              // 2020-07-03 14:22:38
	UConsole.Log(NowDateTime(DefaultLayout)) // 2020-07-03 14:22:38

	// 设置输出layout yyyy年MM月dd日 HH:mm:ss
	layout := "2006年01月02日 15:04:05"
	UConsole.Log(NowDateTime(layout)) // 2020年07月03日 14:22:38

	// 设置输出layout yyyyMMddHHmmss
	layout2 := "20060102150405"
	UConsole.Log(NowDateTime(layout2)) // 20200703142238

	// 设置输出layout yyMMddHHmmss 输入多个layout
	layout3 := "060102150405"
	UConsole.Log(NowDateTime(layout3, layout2)) // 200703142238

	// 设置输出layout yyyy-MM-dd 年月日
	layout4 := "2006-01-02"
	UConsole.Log(NowDateTime(layout4)) // 2020-07-03

	// 设置输出layout HHmmss 时分秒
	layout5 := "150405"
	UConsole.Log(NowDateTime(layout5)) // 142238

	UConsole.PrintAStraightLine()
}

// Test UDateTimeUtil.GetNowDateTimeCST()
func TestGetNowDateTimeCST(t *testing.T) {
	outTime := GetNowDateTimeCST()
	UConsole.PrintTypeAndValue(outTime)
	// 类型(Type):time.Time  值(Value):2020-07-03 23:38:00 +0800 CST
	// 比下面测试方法 UTC + 8h = CST
	UConsole.PrintAStraightLine()
}

// Test UDateTimeUtil.GetNowDateTimeUTC()
func TestGetNowDateTimeUTC(t *testing.T) {
	outTime := GetNowDateTimeUTC()
	UConsole.PrintTypeAndValue(outTime)
	// 类型(Type):time.Time  值(Value):2020-07-03 15:38:00 +0000 UTC
	UConsole.PrintAStraightLine()
}

// Test UDateTimeUtil.GetGivenTime()
func TestGetGivenTime(t *testing.T) {
	outTime := GetGivenTime("20200703142238", "20060102150405")
	fmt.Printf("类型:%T  时间:%v", outTime, outTime) // 类型:time.Time  时间:2020-07-03 14:22:38 +0000 CST
	fmt.Println()
	UConsole.PrintAStraightLine()
}

// Test UDateTimeUtil.GetYearMonthDay()
func TestGetYearMonthDay(t *testing.T) {
	year, month, day := GetYearMonthDay("20200703142238", "20060102150405")
	UConsole.PrintTypeAndValue(year, month, day)
	// 类型(Type):int  值(Value):2020
	// 类型(Type):time.Month  值(Value):July
	// 类型(Type):int  值(Value):3
	UConsole.PrintAStraightLine()
}

// Test UDateTimeUtil.GetHourMinSec()
func TestGetHourMinSec(t *testing.T) {
	UConsole.PrintAStraightLine()
	hour, min, sec := GetHourMinSec("20200703142238", "20060102150405")
	UConsole.PrintTypeAndValue(hour, min, sec)
	// 类型(Type):int  值(Value):14
	// 类型(Type):int  值(Value):22
	// 类型(Type):int  值(Value):38
	UConsole.PrintAStraightLine()
}

// Test UDateTimeUtil.GetDaysPassedThisYear()
func TestGetDaysPassedThisYear(t *testing.T) {
	UConsole.PrintAStraightLine()
	days := GetDaysPassedThisYear()
	UConsole.PrintTypeAndValue(days)
	UConsole.PrintAStraightLine()
}

// Test UDateTimeUtil.GetDaysLeftThisYear()
func TestGetDaysLeftThisYear(t *testing.T) {
	UConsole.PrintAStraightLine()
	days := GetDaysLeftThisYear()
	UConsole.PrintTypeAndValue(days)
	UConsole.PrintAStraightLine()
}

// Test UDateTimeUtil.IsLeapYear()
func TestIsLeapYear(t *testing.T) {
	UConsole.Log(IsLeapYear(2019)) // false
	UConsole.Log(IsLeapYear(2020)) // true
}

// Test UDateTimeUtil.GetNowSecondTimeStamp()
func TestGetNowSecondTimeStamp(t *testing.T) {
	UConsole.PrintTypeAndValue(GetNowSecondTimeStamp())
	// 类型(Type):int64  值(Value):1593767272
}

// Test UDateTimeUtil.GetNowMilliSecondTimeStamp()
func TestGetNowMilliSecondTimeStamp(t *testing.T) {
	UConsole.PrintTypeAndValue(GetNowMilliSecondTimeStamp())
	// 类型(Type):int64  值(Value):1593767290011
}

// Test UDateTimeUtil.GetNowNanoTimeStamp()
func TestGetNowNanoTimeStamp(t *testing.T) {
	UConsole.PrintTypeAndValue(GetNowNanoTimeStamp())
	// 类型(Type):int64  值(Value):1593767333204879600
}

// Test UDateTimeUtil.GetTimeStamp()
func TestGetTimeStamp(t *testing.T) {
	UConsole.PrintTypeAndValue(GetTimeStamp(NowDateTime(), DefaultLayout))
	// 类型(Type):int64  值(Value):1593790426
}

// Test UDateTimeUtil.GetMilliSecondTimeStamp()
// 因为指定的时间只到秒 所以毫秒后面都是0补充
func TestGetMilliSecondTimeStamp(t *testing.T) {
	UConsole.PrintTypeAndValue(GetMilliSecondTimeStamp(NowDateTime(), DefaultLayout))
	// 类型(Type):int64  值(Value):1593790426000
}

// Test UDateTimeUtil.GetNanoTimeStamp()
// 因为指定的时间只到秒 所以纳秒后面都是0补充
func TestGetNanoTimeStamp(t *testing.T) {
	UConsole.PrintTypeAndValue(GetNanoTimeStamp(NowDateTime(), DefaultLayout))
	// 类型(Type):int64  值(Value):1593790426000000000
}

// Test UDateTimeUtil.AddDataTime()
func TestAddDataTime(t *testing.T) {
	UConsole.Log("当前时间:")
	thisTime := GetNowDateTimeCST()
	UConsole.PrintTypeAndValue(thisTime)
	UConsole.PrintAStraightLine()
	// 年
	UConsole.PrintTypeAndValue(AddDataTime("yy", +1, thisTime))
	UConsole.PrintTypeAndValue(AddDataTime("yy", -1, thisTime))
	UConsole.PrintAStraightLine()
	// 月
	UConsole.PrintTypeAndValue(AddDataTime("MM", +1, thisTime))
	UConsole.PrintTypeAndValue(AddDataTime("MM", -1, thisTime))
	UConsole.PrintAStraightLine()
	// 日
	UConsole.PrintTypeAndValue(AddDataTime("dd", +1, thisTime))
	UConsole.PrintTypeAndValue(AddDataTime("dd", -1, thisTime))
	UConsole.PrintAStraightLine()
	// 小时
	UConsole.PrintTypeAndValue(AddDataTime("HH", +1, thisTime))
	UConsole.PrintTypeAndValue(AddDataTime("HH", -1, thisTime))
	UConsole.PrintAStraightLine()
	// 分钟
	UConsole.PrintTypeAndValue(AddDataTime("mm", +5, thisTime))
	UConsole.PrintTypeAndValue(AddDataTime("mm", -5, thisTime))
	UConsole.PrintAStraightLine()
	// 秒
	UConsole.PrintTypeAndValue(AddDataTime("ss", +10, thisTime))
	UConsole.PrintTypeAndValue(AddDataTime("ss", -10, thisTime))
	UConsole.PrintAStraightLine()
	// 来个不合法的 不存在的datepart 来个"##"
	UConsole.PrintTypeAndValue(AddDataTime("##", +1, thisTime))
	UConsole.PrintAStraightLine()
	// 当前时间:
	// 类型(Type):time.Time  值(Value):2020-07-04 00:57:38 +0800 CST
	// -----------------------------------------------------------------------------------------------------
	// 类型(Type):time.Time  值(Value):2021-07-04 00:57:38 +0800 CST
	// 类型(Type):time.Time  值(Value):2019-07-04 00:57:38 +0800 CST
	// -----------------------------------------------------------------------------------------------------
	// 类型(Type):time.Time  值(Value):2020-08-04 00:57:38 +0800 CST
	// 类型(Type):time.Time  值(Value):2020-06-04 00:57:38 +0800 CST
	// -----------------------------------------------------------------------------------------------------
	// 类型(Type):time.Time  值(Value):2020-07-05 00:57:38 +0800 CST
	// 类型(Type):time.Time  值(Value):2020-07-03 00:57:38 +0800 CST
	// -----------------------------------------------------------------------------------------------------
	// 类型(Type):time.Time  值(Value):2020-07-04 01:57:38 +0800 CST
	// 类型(Type):time.Time  值(Value):2020-07-03 23:57:38 +0800 CST
	// -----------------------------------------------------------------------------------------------------
	// 类型(Type):time.Time  值(Value):2020-07-04 01:02:38 +0800 CST
	// 类型(Type):time.Time  值(Value):2020-07-04 00:52:38 +0800 CST
	// -----------------------------------------------------------------------------------------------------
	// 类型(Type):time.Time  值(Value):2020-07-04 00:57:48 +0800 CST
	// 类型(Type):time.Time  值(Value):2020-07-04 00:57:28 +0800 CST
	// -----------------------------------------------------------------------------------------------------
	// ##是不合法的datepart,代替输出当前时间->
	// 类型(Type):time.Time  值(Value):2020-07-04 00:57:38 +0800 CST
	// -----------------------------------------------------------------------------------------------------
}

func TestSubDataTime(t *testing.T) {
	tim1 := GetNowDateTimeCST()
	// 一天前
	tim2 := AddDataTime("dd", -1, tim1)
	UConsole.Log("当前时间:")
	UConsole.PrintTypeAndValue(tim1)
	UConsole.Log("一天前:")
	UConsole.PrintTypeAndValue(tim2)
	UConsole.PrintAStraightLine()
	// 日
	UConsole.PrintTypeAndValue(SubDataTime("dd", tim1, tim2))
	UConsole.PrintTypeAndValue(SubDataTime("dd", tim2, tim1))
	UConsole.PrintAStraightLine()
	// 小时
	UConsole.PrintTypeAndValue(SubDataTime("HH", tim1, tim2))
	UConsole.PrintTypeAndValue(SubDataTime("HH", tim2, tim1))
	UConsole.PrintAStraightLine()
	// 分钟
	UConsole.PrintTypeAndValue(SubDataTime("mm", tim1, tim2))
	UConsole.PrintTypeAndValue(SubDataTime("mm", tim2, tim1))
	UConsole.PrintAStraightLine()
	// 秒
	UConsole.PrintTypeAndValue(SubDataTime("ss", tim1, tim2))
	UConsole.PrintTypeAndValue(SubDataTime("ss", tim2, tim1))
	UConsole.PrintAStraightLine()
	// 来个不合法的 不存在的datepart 来个"**"
	UConsole.PrintTypeAndValue(SubDataTime("**", tim1, tim2))
	UConsole.PrintAStraightLine()

	// 当前时间:
	// 类型(Type):time.Time  值(Value):2020-07-04 11:13:40 +0800 CST
	// 一天前:
	// 类型(Type):time.Time  值(Value):2020-07-03 11:13:40 +0800 CST
	//-----------------------------------------------------------------------------------------------------
	//类型(Type):string  值(Value):1 天
	//
	//类型(Type):string  值(Value):-1 天
	//
	//-----------------------------------------------------------------------------------------------------
	//类型(Type):string  值(Value):24 小时
	//
	//类型(Type):string  值(Value):-24 小时
	//
	//-----------------------------------------------------------------------------------------------------
	//类型(Type):string  值(Value):1440 分钟
	//
	//类型(Type):string  值(Value):-1440 分钟
	//
	//-----------------------------------------------------------------------------------------------------
	//类型(Type):string  值(Value):86400 秒
	//
	//类型(Type):string  值(Value):-86400 秒
	//
	//-----------------------------------------------------------------------------------------------------
	//**是不合法的datepart,代替输出当前字符串时间->
	//类型(Type):string  值(Value):2020-07-04 11:13:40
	//-----------------------------------------------------------------------------------------------------
}

// Test UDateTimeUtil.TimeToStrTime()
func TestTimeToStrTime(t *testing.T) {
	// TimeToStrTime
	outTime := TimeToStrTime(time.Now(), DefaultLayout)
	fmt.Printf("类型:%T  时间:%v", outTime, outTime) // 类型:string  时间:2020-07-03 15:09:40
	fmt.Println()
	UConsole.PrintAStraightLine()
}

// Test UDateTimeUtil.StrTimeToTime()
func TestStrTimeToTime(t *testing.T) {
	// StrTimeToTime
	// 返回(CST UT+8)时间 类型:time.Time
	outTime := StrTimeToTime(NowDateTime(), DefaultLayout)
	fmt.Printf("类型:%T  时间:%v", outTime, outTime) // 类型:time.Time  时间:2020-07-03 15:09:40 +0000 CST
	fmt.Println()
	UConsole.PrintAStraightLine()
}
