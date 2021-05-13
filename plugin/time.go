package plugin

import (
	"time"
)

const (
	TIMEFORMAT       = "20060102150405"
	SHORTTIMESTRING  = "20060102"
	NORMALTIMEFORMAT = "2006-01-02 15:04:05"
)

// 当前时间
func GetTime() time.Time {
	return time.Now()
}

// 格式化为:20060102150405
func GetTimeString(t time.Time) string {
	return t.Format(TIMEFORMAT)
}

// 格式化为:2006-01-02 15:04:05
func GetNormalTimeString(t time.Time) string {
	return t.Format(NORMALTIMEFORMAT)
}

// 转为时间戳->秒数
func GetTimeUnix(t time.Time) int64 {
	return t.Unix()
}

// 转为时间戳->毫秒数
func GetTimeMills(t time.Time) int64 {
	return t.UnixNano() / 1e6
}

// 时间戳转时间
func GetTimeByInt(t1 int64) time.Time {
	return time.Unix(t1, 0)
}

// 字符串转时间
func GetTimeByString(timestring string) (time.Time, error) {
	if timestring == "" {
		return time.Time{}, nil
	}
	return time.ParseInLocation(TIMEFORMAT, timestring, time.Local)
}

// 标准字符串转时间
func GetTimeByNormalString(timestring string) (time.Time, error) {
	if timestring == "" {
		return time.Time{}, nil
	}
	return time.ParseInLocation(NORMALTIMEFORMAT, timestring, time.Local)
}

// 比较两个时间大小
func CompareTime(t1, t2 time.Time) bool {
	return t1.Before(t2)
}

// n小时后的时间字符串
func GetNextHourTime(s string, n int64) string {
	t2, _ := time.ParseInLocation(TIMEFORMAT, s, time.Local)
	t1 := t2.Add(time.Hour * time.Duration(n))
	return GetTimeString(t1)
}

// 计算俩个时间差多少小时
func GetHourDiffer(start_time, end_time string) float32 {
	var hour float32
	t1, err := time.ParseInLocation(TIMEFORMAT, start_time, time.Local)
	t2, err := time.ParseInLocation(TIMEFORMAT, end_time, time.Local)
	if err == nil && CompareTime(t1, t2) {
		diff := GetTimeUnix(t2) - GetTimeUnix(t1)
		hour = float32(diff) / 3600
		return hour
	}
	return hour
}

// 判断当前时间是否是整点
func Checkhours() bool {
	_, m, s := GetTime().Clock()
	if m == s && m == 0 && s == 0 {
		return true
	}
	return false
}

// 时间字符串转为标准字符串
func StringToNormalString(t string) string {
	if !(len(TIMEFORMAT) == len(t) || len(SHORTTIMESTRING) == len(t)) {
		return t
	}
	if len(SHORTTIMESTRING) == len(t) {
		t += "000000"
	}
	if len(TIMEFORMAT) == len(t) {
		t1, err := GetTimeByString(t)
		if err != nil {
			return t
		}
		t = GetTimeString(t1)
	}
	return t
}
