// Copyright © 2021 - 2022 vity <vityme@icloud.com>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package r

import (
	"math/rand"
	"time"
)

// RandomFloat64 生成随机Float64数字
func RandomFloat64() float64 {
	rand.Seed(time.Now().UnixMilli())
	return rand.Float64()
}
