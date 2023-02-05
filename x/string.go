// Copyright © 2021 - 2023 vity <vityme@icloud.com>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package x

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

// StringIsNotBlank 判断给定的字符串是否为空字符串
func StringIsNotBlank(value string) bool {
	return strings.TrimSpace(value) != ""
}

// StringIsBlank 判断给定的字符串是否为空字符串
func StringIsBlank(value string) bool {
	return strings.TrimSpace(value) == ""
}

// StringDefaultIfBlank 返回字符串，若给定的字符串为空，则返回默认字符串
func StringDefaultIfBlank(value, defaultString string) string {
	if StringIsBlank(value) {
		return defaultString
	}
	return value
}

// StringLeftPad 左补齐给定字符串
func StringLeftPad(value string, size int, padString string) string {
	padLen := len(padString)
	strLen := len(value)
	pads := size - strLen
	if pads <= 0 {
		return value
	}
	if padLen == 1 {
		return strings.Repeat(padString, pads) + value
	}
	if pads == padLen {
		return padString + value
	}
	if pads < padLen {
		return padString[:(pads)] + value
	}
	padding := make([]rune, pads)
	padRunes := []rune(padString)
	for i := 0; i < pads; i++ {
		padding[i] = padRunes[i%padLen]
	}
	return string(padding) + value
}

// StringRightPad 右补齐给定字符串
func StringRightPad(value string, size int, padString string) string {
	padLen := len(padString)
	strLen := len(value)
	pads := size - strLen
	if pads <= 0 {
		return value
	}
	if padLen == 1 {
		return value + strings.Repeat(padString, pads)
	}
	if pads == padLen {
		return value + padString
	}
	if pads < padLen {
		return value + padString[:(pads)]
	}
	padding := make([]rune, pads)
	padRunes := []rune(padString)
	for i := 0; i < pads; i++ {
		padding[i] = padRunes[i%padLen]
	}
	return value + string(padding)
}
