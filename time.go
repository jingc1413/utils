package utils

import (
	"time"
)

// 下一分钟, 对齐时间, 0 秒
func WaitNextMinute() {
	now := time.Now()
	<-time.After(Get0Second(now.Add(time.Minute)).Sub(now))
}

// 当天 0 点
func Get0Hour(t time.Time) time.Time {
	y, m, d := t.Date()
	return time.Date(y, m, d, 0, 0, 0, 0, t.Location())
}

// 昨天 0 点
func Get0Yesterday(t time.Time) time.Time {
	return Get0Hour(t.AddDate(0, 0, -1))
}

// 昨天 0 点
func Get0Tomorrow(t time.Time) time.Time {
	return Get0Hour(t.AddDate(0, 0, 1))
}

// 0 分
func Get0Minute(t time.Time) time.Time {
	y, m, d := t.Date()
	return time.Date(y, m, d, t.Hour(), 0, 0, 0, t.Location())
}

// 0 秒
func Get0Second(t time.Time) time.Time {
	t.Truncate(time.Minute)
	y, m, d := t.Date()
	return time.Date(y, m, d, t.Hour(), t.Minute(), 0, 0, t.Location())
}

// 本周一 0 点
func Get0Week(t time.Time) time.Time {
	offset := int(time.Monday - t.Weekday())
	if offset > 0 {
		offset = -6
	}

	return Get0Hour(t).AddDate(0, 0, offset)
}

// 上周一 0 点
func Get0LastWeek(t time.Time) time.Time {
	return Get0Week(t.AddDate(0, 0, -7))
}

// 下周一 0 点
func Get0NextWeek(t time.Time) time.Time {
	return Get0Week(t.AddDate(0, 0, 7))
}

// 当月第一天 0 点
func Get0Month(t time.Time) time.Time {
	y, m, _ := t.Date()
	return time.Date(y, m, 1, 0, 0, 0, 0, t.Location())
}

// 上月第一天 0 点
func Get0LastMonth(t time.Time) time.Time {
	return Get0Month(t.AddDate(0, -1, 0))
}

// 下月第一天 0 点
func Get0NextMonth(t time.Time) time.Time {
	return Get0Month(t.AddDate(0, 1, 0))
}

// 当月天数
func GetMonthDays(t time.Time) int {
	return int(Get0NextMonth(t).Sub(Get0Month(t)).Hours() / 24)
}
