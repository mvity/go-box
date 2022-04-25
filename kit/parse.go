// Copyright © 2021 - 2022 vity <vityme@icloud.com>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package k

import (
	"strconv"
	"strings"
	"time"
)

// ParseDateTime 转换 2006-01-02 15:04:05 格式时间字符串
func ParseDateTime(value string) time.Time {
	if strings.TrimSpace(value) == "" {
		return time.Time{}
	}
	t, err := time.ParseInLocation("2006-01-02 15:04:05", strings.TrimSpace(value), time.Now().Location())
	if err != nil {
		panic(err)
	}
	return t
}

// ParseDate 转换 2006-01-02 格式日期字符串
func ParseDate(value string) time.Time {
	if strings.TrimSpace(value) == "" {
		return time.Time{}
	}
	t, err := time.ParseInLocation("2006-01-02", strings.TrimSpace(value), time.Now().Location())
	if err != nil {
		panic(err)
	}
	return t
}

// ParseTime 转换 15:04:05 格式时间字符串
func ParseTime(value string) time.Time {
	if strings.TrimSpace(value) == "" {
		return time.Time{}
	}
	t, err := time.ParseInLocation("15:04:05", strings.TrimSpace(value), time.Now().Location())
	if err != nil {
		panic(err)
	}
	return t
}

// ParseLocation 转换 lng,lat 格式经纬度字符串，例：113.617525,34.751782
func ParseLocation(value string) (float64, float64) {
	if strings.TrimSpace(value) == "" || strings.Contains(value, ",") {
		return 0, 0
	}
	locs := strings.Split(value, ",")
	lng, err := strconv.ParseFloat(locs[0], 64) // 经度
	if err != nil {
		panic(err)
	}
	lat, err := strconv.ParseFloat(locs[1], 64) // 纬度
	if err != nil {
		panic(err)
	}
	return lng, lat
}
