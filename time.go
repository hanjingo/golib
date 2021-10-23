package golib

import (
	"strconv"
	"time"
)

const FMT_TIME_STR string = "2006-01-02 15:04:05"

func StrToCstTime(timeStr string) time.Time {
	t, _ := time.ParseInLocation(FMT_TIME_STR, timeStr, time.Local)
	return t
}

func TimeToStdStr(t time.Time) string {
	return t.Format(FMT_TIME_STR)
}

func UnixMs(t time.Time) int64 {
	return t.UnixNano() / 1e6
}

func DurToMs(dur time.Duration) int64 {
	return dur.Nanoseconds() / 1e6
}

func TimeToStamp(t time.Time) string {
	return strconv.FormatInt(UnixMs(t), 10)
}

func StampToTime(TimeStamp string) time.Time {
	ms, _ := strconv.ParseInt(TimeStamp, 10, 64)
	second := ms / 1000
	ns := (ms % 1000) * 1000000
	return time.Unix(second, ns)
}
