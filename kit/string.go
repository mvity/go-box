// Copyright © 2021 - 2022 vity <vityme@icloud.com>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package k

import "strings"

// StringReverse 反转字符串
func StringReverse(value string) string {
	r := []rune(value)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

// StringIsAnyBlank 判断给定的字符串参数中是否存在空字符串
func StringIsAnyBlank(value ...string) bool {
	for _, s := range value {
		if strings.TrimSpace(s) == "" {
			return true
		}
	}
	return false
}

// StringIsAllBlank 判断给定的字符串参数中是否全部为空字符串
func StringIsAllBlank(value ...string) bool {
	for _, s := range value {
		if strings.TrimSpace(s) != "" {
			return false
		}
	}
	return true
}
