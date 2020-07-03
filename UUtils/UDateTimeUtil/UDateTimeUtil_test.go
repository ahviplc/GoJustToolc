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
	outTime := StrTimeToTime(NowDateTime(), DefaultLayout)
	fmt.Printf("类型:%T  时间:%v", outTime, outTime) // 类型:time.Time  时间:2020-07-03 15:09:40 +0000 UTC
	fmt.Println()
	UConsole.PrintAStraightLine()
}

// Test UDateTimeUtil.GetGivenTime()
func TestGetGivenTime(t *testing.T) {
	outTime := GetGivenTime("20200703142238", "20060102150405")
	fmt.Printf("类型:%T  时间:%v", outTime, outTime) // 类型:time.Time  时间:2020-07-03 14:22:38 +0000 UTC
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

// Test UDateTimeUtil.GetTimeStamp() UDateTimeUtil.GetTimeStampPlus8H()
func TestGetTimeStamp(t *testing.T) {
	UConsole.PrintTypeAndValue(GetTimeStamp(NowDateTime(), DefaultLayout))
	UConsole.PrintTypeAndValue(GetTimeStampPlus8H(NowDateTime(), DefaultLayout)) // 返回的是比传入时间多8h(也就是多28800秒)的秒时间戳
	// 类型(Type):int64  值(Value):1593768863
	// 二者相差 8h (也就是28800秒)
	// 类型(Type):int64  值(Value):1593797663
}
