// Copyright © 2021 - 2022 vity <vityme@icloud.com>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package r

import (
	"math/rand"
	"time"
)

// RandomInt 随机生成指定区间的数字
func RandomInt(min int64, max int64) int64 {
	if min == max {
		return min
	}
	if min > max {
		min, max = max, min
	}
	rand.Seed(time.Now().UnixNano())
	return rand.Int63n(max-min) + min
}
