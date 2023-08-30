/*
 * Copyright © 2021 - 2023 vity <vityme@icloud.com>.
 *
 * Use of this source code is governed by an MIT-style
 * license that can be found in the LICENSE file.
 */

package x

import (
	"time"
)

// GetZeroTime  获取指定时间零点时刻
func GetZeroTime(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}

// GetLastTime  获取指定时间最后时刻 23:59:59
func GetLastTime(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 0, t.Location())
}

// GetWeekdayOfDay 获取指定日期本周周几的日期
func GetWeekdayOfDay(t time.Time, weekday time.Weekday) time.Time {
	if t.Weekday() == weekday {
		return GetZeroTime(t)
	}
	offset := int(weekday - t.Weekday())
	if t.Weekday() == time.Sunday {
		offset = 0 - (7 - offset)
	}
	day := GetZeroTime(t.AddDate(0, 0, offset))
	if weekday == time.Sunday {
		day = day.AddDate(0, 0, 7)
	}
	return day
}
