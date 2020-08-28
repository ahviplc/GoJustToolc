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

// 设置RFC3339时间模板layout
const DefaultLayoutRFC3339 = "2006-01-02T15:04:05+08:00" // time.RFC3339 ->【RFC3339     = "2006-01-02T15:04:05Z07:00"】

// 输出当前时间 支持按照指定的转换格式输出 layout:输出时间模板
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

// 获取当前时间(CST UT+8) RFC3339格式
func GetNowDateTimeCSTRFC3339() string {
	return StrTimeToTime(NowDateTime(), DefaultLayout).Format(time.RFC3339)
}

// 获取当前时间(UTC)
func GetNowDateTimeUTC() time.Time {
	return StrTimeToTime(NowDateTime(), DefaultLayout).UTC()
}

// 时间转换成RFC3339格式
// in 输入的字符串类型时间 传入
// inLayout 时间模板 输入的格式
func ToRFC3339(in string, inLayout string) string {
	loc, _ := time.LoadLocation("Local")
	tm, err := time.ParseInLocation(inLayout, in, loc)
	if err != nil {
		panic(err)
		// return ""
	}
	return tm.Format(time.RFC3339)
}

// 将RFC3339格式转成正常日期时间格式
// in 输入的RFC3339格式类型时间 传入
// outLayout 时间模板 输出的转成正常日期时间格式 输出的格式
func RFC3339To(in string, outLayout string) string {
	loc, _ := time.LoadLocation("Local")
	t, err := time.ParseInLocation(DefaultLayoutRFC3339, in, loc)
	if err != nil {
		panic(err)
		// return "0000-00-00 00:00:00"
	}
	return t.Format(outLayout)
}

// 2006-01-02T15:04:05Z07:00 RFC3339格式转成时间戳 秒
// in 输入的RFC3339格式类型时间 传入
// 输出时间戳 秒
func RFC3339ToUnix(in string) int64 { //转化所需模板
	loc, _ := time.LoadLocation("Local")
	tm, err := time.ParseInLocation(time.RFC3339, in, loc) //使用模板在对应时区转化为time.time类型
	if err != nil {
		panic(err)
		//return int64(0)
	}
	return tm.Unix()
}

// 万能的format转换格式函数 等价于方法 TimeToStrTime()
// Format time.Time struct to string
// MM - month - 01
// M - month - 1, single bit
// DD - day - 02
// D - day 2
// YYYY - year - 2006
// YY - year - 06
// HH - 24 hours - 03
// H - 24 hours - 3
// hh - 12 hours - 03
// h - 12 hours - 3
// mm - minute - 04
// m - minute - 4
// ss - second - 05
// s - second = 5
// 备注:自定义格式对应如下: YYYY-MM-DD HH:mm:ss 必须使用 2006-01-02 15:04:05 替代
// YYYY-MM-DD HH:mm:ss -> 2006-01-02 15:04:05
// YYYY-MM-DD -> 2006-01-02
// HH:mm:ss -> 15:04:05
// YYYY/MM/DD HH:mm:ss -> 2006/01/02 15:04:05
func UFormat(in time.Time, format string) string {
	return in.Format(format)
}

//获取UFormat格式化的数据 等价于方法 TimeToStrTime()
func GetUFormat(in time.Time, format string) string {
	return UFormat(in, format)
}

// 获取指定时间
// in 输入的字符串类型时间 传入
// inLayout 时间模板 输入的格式
// 使用 t := time.Date(2020, 9, 9, 9, 9, 9, 0, time.Local) 也可
func GetGivenTime(in string, inLayout string) time.Time {
	return StrTimeToTime(in, inLayout)
}

// 根据输入时间,获取指定内容 年月日
// in 输入的字符串类型时间 传入
// inLayout 时间模板 输入的格式
func GetYearMonthDay(in string, inLayout string) (year int, month time.Month, day int) {
	return StrTimeToTime(in, inLayout).Date()
}

// 根据输入时间,获取指定内容 时分秒
// in 输入的字符串类型时间 传入
// inLayout 时间模板 输入的格式
func GetHourMinSec(in string, inLayout string) (hour, min, sec int) {
	return StrTimeToTime(in, inLayout).Clock()
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
// in 输入的指定时间
// inLayout 时间模板 输入的格式
func GetTimeStamp(in string, inLayout string) int64 {
	// time.Now().Unix()直接就是(CST UT+8)的秒时间戳 1593768863
	// 但是StrTimeToTime()转成UTC的time走Unix()会比(CST UT+8)的秒时间戳多8小时(也就是多28800秒) 1593797663
	// 所以需要减去 8h
	// 下面写法废弃
	// h, _ := time.ParseDuration("-1h")
	// outTime := StrTimeToTime(in, inLayout).Add(8 * h) // 使用这个写法 因为StrTimeToTime(in, inLayout)返回的是当前时间+8h的UTC 所以要减去8h得到当前时间
	// 现在StrTimeToTime()方法直接返回的就是
	return StrTimeToTime(in, inLayout).Unix() // 返回的是(CST UT+8)的time的时间戳 秒
}

// 获取指定时间的时间戳 毫秒
// in 输入的指定时间
// inLayout 时间模板 输入的格式
func GetMilliSecondTimeStamp(in string, inLayout string) int64 {
	return StrTimeToTime(in, inLayout).UnixNano() / 1e6 // 返回的是(CST UT+8)的time的时间戳 秒
}

// 获取指定时间的时间戳 纳秒
// in 输入的指定时间
// inLayout 时间模板 输入的格式
func GetNanoTimeStamp(in string, inLayout string) int64 {
	return StrTimeToTime(in, inLayout).UnixNano() // 返回的是(CST UT+8)的time的时间戳 秒
}

// 2006-01-02 15:04:05格式字符串类型时间 转 时间戳 秒 等价于方法 GetTimeStamp()
// in 输入的字符串类型时间 传入 必须是DefaultLayout格式 【2006-01-02 15:04:05】
// inLayout 输入的字符串日期时间格式 输入的格式
func ToUnix(in string, inLayout string) int64 {
	loc, _ := time.LoadLocation("Local")
	t, err := time.ParseInLocation(inLayout, in, loc)
	if err != nil {
		panic(err)
	}
	return t.Unix()
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
// outLayout 时间模板 输出的格式
func TimeToStrTime(in time.Time, outLayout string) string {
	return fmt.Sprint(in.Format(outLayout))
}

// string字符串类型时间转成time.Time类型
// in 输入的字符串类型时间
// inLayout 时间模板 输入的格式
func StrTimeToTime(in string, inLayout string) time.Time {
	// 走完这步 time是世界协调时间(UTC)了 不是(CST UT+8)
	// 就没有time.Now()返回(CST UT+8)的效果 所以废弃此方法
	// out, err := time.Parse(layout, in)
	out, err := time.ParseInLocation(inLayout, in, time.Local) // 此方法返回的time是(CST UT+8) 中国标准时间 China Standard Time UT+8:00
	if err != nil {
		panic(err)
	}
	return out
}

// 获取日期时间的的零点
// in 输入的字符串类型时间
// inLayout 时间模板 输入的格式
func GetDateTimeStart(in, inLayout string) (result string, err error) {
	loc, _ := time.LoadLocation("Local")
	tt, err := time.Parse(inLayout, in)
	if err != nil {
		//panic(err)
		return "", err
	}
	tt = time.Date(tt.Year(), tt.Month(), tt.Day(), 0, 0, 0, 0, loc)
	return tt.Format(DefaultLayout), nil
}

// 获取日期时间的的最后一刻
// in 输入的字符串类型时间
// inLayout 时间模板 输入的格式
func GetDateTimeEnd(in, fromFormat string) (result string, err error) {
	loc, _ := time.LoadLocation("Local")
	tt, err := time.Parse(fromFormat, in)
	if err != nil {
		//panic(err)
		return "", err
	}
	tt = time.Date(tt.Year(), tt.Month(), tt.Day(), 23, 59, 59, 0, loc)
	return tt.Format(DefaultLayout), nil
}

// 秒时间戳转成字符串时间
// in 输入的时间戳时间
// 输出2006-01-02 15:04:05格式字符串类型时间
func TimeStampToTimeStr(in int64) string {
	return time.Unix(in, 0).Format(DefaultLayout)
}

// WeekAround returns the date of monday and sunday for current week
// 输出当前时间所在周的周一和周日的日期 输出格式为 20060102
func WeekAround() (monday, sunday string) {
	now := time.Now()
	offset := int(time.Monday - now.Weekday())
	if offset > 0 {
		offset = -6
	}
	today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local)
	monday = today.AddDate(0, 0, offset).Format("20060102")
	offset = int(time.Sunday - now.Weekday())
	if offset < 0 {
		offset += 7
	}
	sunday = today.AddDate(0, 0, offset).Format("20060102")
	return
}
