package UDateTimeUtil

import (
	"fmt"
	"github.com/ahviplc/GoJustToolc/UConsole"
	"github.com/ahviplc/GoJustToolc/UUtils/UStringUtil"
	"time"
)

// 日期时间工具类

// time包:
// 1年=365天，day
// 1天=24小时，hour
// 1小时=60分钟， minute
// 1分钟=60秒， second
// 1秒钟=1000毫秒， millisecond
// 1毫秒=1000微秒， microsecond->us
// 1微秒=1000纳秒， nanosecond->ns
// 1纳秒=1000皮秒， picosecond->ps

// 设置默认时间模板layout
const DefaultLayout = "2006-01-02 15:04:05"

// 输出当前时间 layout:输出时间模板
// 如果layout传入多参数 只默认使用第一个模板layout
func NowDateTime(layoutArgs ...string) string {
	if len(layoutArgs) > 0 {

		// 下面遍历layoutArgs代码
		// for _, vLayout := range layoutArgs {
		//	return fmt.Sprint(time.Now().Format(vLayout))
		// }

		// 使用下标为0的layout
		return fmt.Sprint(time.Now().Format(layoutArgs[0]))
	}
	return fmt.Sprint(time.Now().Format(DefaultLayout))
}

// 获取当前时间(CST UT+8)
func GetNowDateTimeCST() time.Time {
	return StrTimeToTime(NowDateTime(), DefaultLayout)
}

// 获取当前时间(UTC)
func GetNowDateTimeUTC() time.Time {
	return StrTimeToTime(NowDateTime(), DefaultLayout).UTC()
}

// 获取指定时间
// in 输入的字符串类型时间 传入
// layout 时间模板
// 使用 t := time.Date(2020, 9, 9, 9, 9, 9, 0, time.Local) 也可
func GetGivenTime(in string, layout string) time.Time {
	return StrTimeToTime(in, layout)
}

// 根据输入时间,获取指定内容 年月日
// in 输入的字符串类型时间 传入
// layout 时间模板
func GetYearMonthDay(in string, layout string) (year int, month time.Month, day int) {
	return StrTimeToTime(in, layout).Date()
}

// 根据输入时间,获取指定内容 时分秒
// in 输入的字符串类型时间 传入
// layout 时间模板
func GetHourMinSec(in string, layout string) (hour, min, sec int) {
	return StrTimeToTime(in, layout).Clock()
}

// 获取今年过去了多少天
func GetDaysPassedThisYear() (days int) {
	year := time.Now().Year()
	// 因为 YearDay() returns the day of the year specified by t, in the range [1,365] for non-leap years,返回的结果是针对非闰年的
	// 判断是否是闰年
	if IsLeapYear(year) {
		// 是闰年 YearDay() - 1
		return time.Now().YearDay() - 1
	}
	// 非闰年 直接YearDay()即可
	return time.Now().YearDay()
}

// 获取今年还剩下多少天
func GetDaysLeftThisYear() (days int) {
	year := time.Now().Year()
	// 判断是否是闰年
	if IsLeapYear(year) {
		// 是闰年 366 - 已过去天数
		return 366 - GetDaysPassedThisYear()
	}
	// 非闰年 365 - 已过去天数
	return 365 - GetDaysPassedThisYear()
}

// 闰年判断
// year 输入的年
func IsLeapYear(year int) bool {
	if (year%4 == 0 && year%100 != 0) || (year%400 == 0) || (year%1000 == 0) {
		return true
	}
	return false
}

// 获取当前时间的时间戳 秒
func GetNowSecondTimeStamp() int64 {
	return time.Now().Unix()
}

// 获取当前时间的时间戳 毫秒
func GetNowMilliSecondTimeStamp() int64 {
	return time.Now().UnixNano() / 1e6
}

// 获取当前时间的时间戳 纳秒
func GetNowNanoTimeStamp() int64 {
	return time.Now().UnixNano()
}

// 获取指定时间的时间戳 秒
func GetTimeStamp(in string, layout string) int64 {
	// time.Now().Unix()直接就是(CST UT+8)的秒时间戳 1593768863
	// 但是StrTimeToTime()转成UTC的time走Unix()会比(CST UT+8)的秒时间戳多8小时(也就是多28800秒) 1593797663
	// 所以需要减去 8h
	// 下面写法废弃
	// h, _ := time.ParseDuration("-1h")
	// outTime := StrTimeToTime(in, layout).Add(8 * h) // 使用这个写法 因为StrTimeToTime(in, layout)返回的是当前时间+8h的UTC 所以要减去8h得到当前时间
	// 现在StrTimeToTime()方法直接返回的就是
	return StrTimeToTime(in, layout).Unix() // 返回的是(CST UT+8)的time的时间戳 秒
}

// 获取指定时间的时间戳 毫秒
func GetMilliSecondTimeStamp(in string, layout string) int64 {
	return StrTimeToTime(in, layout).UnixNano() / 1e6 // 返回的是(CST UT+8)的time的时间戳 秒
}

// 获取指定时间的时间戳 纳秒
func GetNanoTimeStamp(in string, layout string) int64 {
	return StrTimeToTime(in, layout).UnixNano() // 返回的是(CST UT+8)的time的时间戳 秒
}

// 把日期时间加减年、月、日、时、分、秒后得到新的日期时间
// datepart yy-年、MM-月、dd-日、HH-时、mm-分、ss-秒
// plusSubSum   加减因子 例如:+1 -1
// thisTime 需要加减的日期时间
// 备注: ParseDuration介绍
// ParseDuration解析一个时间段字符串。一个时间段字符串是一个序列，每个片段包含可选的正负号,
// 十进制数、可选的小数部分和单位后缀，如"300ms"、"-1.5h"、"2h45m"。
// 合法的单位有"ns"纳秒,"us","µs"、"ms"毫秒、"s"秒、"m"分钟、"h"。
func AddDataTime(datepart string, plusSubSum int, thisTime time.Time) time.Time {
	switch datepart {
	case "yy":
		thatTime := thisTime.AddDate(plusSubSum, 0, 0)
		return thatTime
	case "MM":
		thatTime := thisTime.AddDate(0, plusSubSum, 0)
		return thatTime
	case "dd":
		thatTime := thisTime.AddDate(0, 0, plusSubSum)
		return thatTime
	case "HH":
		h, _ := time.ParseDuration(UStringUtil.IntToString(plusSubSum) + "h")
		thatTime := thisTime.Add(h)
		return thatTime
	case "mm":
		m, _ := time.ParseDuration(UStringUtil.IntToString(plusSubSum) + "m")
		thatTime := thisTime.Add(m)
		return thatTime
	case "ss":
		m, _ := time.ParseDuration(UStringUtil.IntToString(plusSubSum) + "s")
		thatTime := thisTime.Add(m)
		return thatTime
	}
	UConsole.Log(datepart, "是不合法的datepart,代替输出当前时间->")
	return GetNowDateTimeCST()
}

// 计算两个时间差
// datepart string dd-日、HH-时、mm-分、ss-秒
// time1
// time2
// 计算 time1-time2的时间差
func SubDataTime(datepart string, time1 time.Time, time2 time.Time) string {
	switch datepart {
	case "dd":
		return fmt.Sprintln(time1.Sub(time2).Hours()/24, "天")
	case "HH":
		return fmt.Sprintln(time1.Sub(time2).Hours(), "小时")
	case "mm":
		return fmt.Sprintln(time1.Sub(time2).Minutes(), "分钟")
	case "ss":
		return fmt.Sprintln(time1.Sub(time2).Seconds(), "秒")
	}
	UConsole.Log(datepart, "是不合法的datepart,代替输出当前字符串时间->")
	return NowDateTime()
}

// time.Time类型转成string字符串类型时间
// in 输入的time.Time类型时间
// inLayout 时间模板
func TimeToStrTime(in time.Time, layout string) string {
	return fmt.Sprint(in.Format(layout))
}

// string字符串类型时间转成time.Time类型
// in 输入的字符串类型时间
// layout 时间模板
func StrTimeToTime(in string, layout string) time.Time {
	// 走完这步 time是世界协调时间(UTC)了 不是(CST UT+8)
	// 就没有time.Now()返回(CST UT+8)的效果 所以废弃此方法
	// out, err := time.Parse(layout, in)
	out, err := time.ParseInLocation(layout, in, time.Local) // 此方法返回的time是(CST UT+8) 中国标准时间 China Standard Time UT+8:00
	if err != nil {
		panic(err)
	}
	return out
}
