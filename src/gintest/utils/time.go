package utils

import (
	"time"
)

// 一些与时间相关的通用方法

// GetLastWeekDate 获取上周一和本周一的日期
func GetLastWeekDate() (thisMonday, lastMonday time.Time) {
	now := time.Now()
	// 当前距离本周一的天数
	diffMonday := (int)(now.Weekday() + 7) % 7
	thisMonday = now.AddDate(0, 0, -diffMonday)
	lastMonday = thisMonday.AddDate(0, 0, -7)
	return
}
