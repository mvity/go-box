/*
 * Copyright © 2021 - 2023 vity <vityme@icloud.com>.
 *
 * Use of this source code is governed by an MIT-style
 * license that can be found in the LICENSE file.
 */

package x

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// FormatChinaMoney 格式化人民币，从分到元
func FormatChinaMoney(value int64) string {
	return fmt.Sprintf("%.2f", (float64(value))/100.0)
}

// FormatChinaMoneyWord 格式化人民币大写格式
func FormatChinaMoneyWord(num float64) string {
	strnum := strconv.FormatFloat(num*100, 'f', 0, 64)
	sliceUnit := []string{"仟", "佰", "拾", "亿", "仟", "佰", "拾", "万", "仟", "佰", "拾", "元", "角", "分"}
	s := sliceUnit[len(sliceUnit)-len(strnum):]
	upperDigitUnit := map[string]string{"0": "零", "1": "壹", "2": "贰", "3": "叁", "4": "肆", "5": "伍", "6": "陆", "7": "柒", "8": "捌", "9": "玖"}
	str := ""
	for k, v := range strnum[:] {
		str = str + upperDigitUnit[string(v)] + s[k]
	}
	reg, err := regexp.Compile(`零角零分$`)
	str = reg.ReplaceAllString(str, "整")

	reg, err = regexp.Compile(`零角`)
	str = reg.ReplaceAllString(str, "零")

	reg, err = regexp.Compile(`零分$`)
	str = reg.ReplaceAllString(str, "整")

	reg, err = regexp.Compile(`零[仟佰拾]`)
	str = reg.ReplaceAllString(str, "零")

	reg, err = regexp.Compile(`零{2,}`)
	str = reg.ReplaceAllString(str, "零")

	reg, err = regexp.Compile(`零亿`)
	str = reg.ReplaceAllString(str, "亿")

	reg, err = regexp.Compile(`零万`)
	str = reg.ReplaceAllString(str, "万")

	reg, err = regexp.Compile(`零*元`)
	str = reg.ReplaceAllString(str, "元")

	reg, err = regexp.Compile(`亿零{0, 3}万`)
	str = reg.ReplaceAllString(str, "^元")

	reg, err = regexp.Compile(`零元`)
	str = reg.ReplaceAllString(str, "零")
	if err != nil {
		log.Fatal(err)
	}
	return str
}

// FormatKilometer 格式化公里，从米到千米
func FormatKilometer(value int64, scale uint8) string {
	strs := strings.Split(fmt.Sprintf("%.3f", (float64(value))/1000.0), ".")
	if scale == 0 {
		return strs[0]
	}
	str := strs[0] + "." + strs[1][:scale]
	return str
}

// FormatDateTime 格式化时间为 2006-01-02 15:04:05 格式
func FormatDateTime(dt time.Time) string {
	if dt.IsZero() {
		return ""
	}
	return dt.Format("2006-01-02 15:04:05")
}

// FormatDateTimeLayout 格式化指定格式的字符串
func FormatDateTimeLayout(dt time.Time, layout string) string {
	if dt.IsZero() {
		return ""
	}
	return dt.Format(layout)
}

// FormatDate 格式化时间为 2006-01-02 格式
func FormatDate(dt time.Time) string {
	if dt.IsZero() {
		return ""
	}
	return dt.Format("2006-01-02")
}

// FormatTime 格式化时间为 15:04:05 格式
func FormatTime(dt time.Time) string {
	if dt.IsZero() {
		return ""
	}
	return dt.Format("15:04:05")
}
