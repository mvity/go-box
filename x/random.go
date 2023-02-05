/*
 * Copyright © 2021 - 2023 vity <vityme@icloud.com>.
 *
 * Use of this source code is governed by an MIT-style
 * license that can be found in the LICENSE file.
 */

package x

import (
	"math/rand"
	"time"
)

// RandomString 生成随机字符
func RandomString(count int, letters bool, numbers bool) string {
	var runes []rune
	if letters && numbers || !letters && !numbers {
		runes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	} else if letters {
		runes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	} else {
		runes = []rune("0123456789")
	}
	b := make([]rune, count)
	for i := range b {
		b[i] = runes[rand.Intn(len(runes))]
	}
	return string(b)
}

// RandomInt 随机生成指定区间的数字
func RandomInt(minimum int64, maximum int64) int64 {
	if minimum == maximum {
		return minimum
	}
	if minimum > maximum {
		minimum, maximum = maximum, minimum
	}
	return rand.New(rand.NewSource(time.Now().UnixNano())).Int63n(maximum-minimum) + minimum
}
