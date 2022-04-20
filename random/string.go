// Copyright © 2021 - 2022 vity <vityme@icloud.com>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package r

import (
	"math/rand"
	"time"
	"unsafe"
)

const (
	rndStrBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	rndNumBytes = "0123456789"
	rndAllBytes = rndStrBytes + rndNumBytes
	rndIdxBits  = 6                 // 6 bits to represent a letter index
	rndIdxMask  = 1<<rndIdxBits - 1 // All 1-bits, as many as rndIdxBits
	rndIdxMax   = 63 / rndIdxBits   // # of letter indices fitting in 63 bits
)

var src = rand.NewSource(time.Now().UnixNano())

// RandomString 生成随机字符
func RandomString(size int, str bool, num bool) string {
	var chars string
	if str && num || !str && !num {
		chars = rndAllBytes
	} else if str {
		chars = rndStrBytes
	} else {
		chars = rndNumBytes
	}
	b := make([]byte, size)
	for i, cache, remain := size-1, src.Int63(), rndIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), rndIdxMax
		}
		if idx := int(cache & rndIdxMask); idx < len(chars) {
			b[i] = chars[idx]
			i--
		}
		cache >>= rndIdxBits
		remain--
	}
	return *(*string)(unsafe.Pointer(&b))
}
