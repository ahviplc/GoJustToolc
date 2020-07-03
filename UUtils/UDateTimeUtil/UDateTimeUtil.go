package UDateTimeUtil

import (
	"fmt"
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
	out, err := time.Parse(layout, in) // 走完这步 time 是 世界协调时间 (UTC)了
	if err != nil {
		panic(err)
	}
	return out
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
	// time.Now().Unix() 直接就是 CST 的 秒时间戳 1593768863
	// 但是StrTimeToTime()转成的time走Unix()会比 CST 的 秒时间戳 多8小时(也就是多28800秒) 1593797663
	// 所以需要减去 8h
	h, _ := time.ParseDuration("-1h")
	outTime := StrTimeToTime(in, layout).Add(8 * h)
	return outTime.Unix()
}

// 获取指定时间的时间戳 秒 方法2
func GetTimeStamp2(in string, layout string) int64 {
	outTime, _ := time.ParseInLocation(layout, in, time.Local) // 返回的是 CST
	return outTime.Unix()
}

// 获取比指定时间多8h的时间戳 秒
func GetTimeStampPlus8H(in string, layout string) int64 {
	// time.Now().Unix() 直接就是 CST 的 秒时间戳 1593768863
	// 但是StrTimeToTime()转成的time走Unix()会比 CST 的 秒时间戳 多8小时(也就是多28800秒) 1593797663
	return StrTimeToTime(in, layout).Unix()
}
