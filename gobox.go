// Copyright © 2021 - 2022 vity <vityme@icloud.com>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package gbox

import "log"

var (
	DEBUG = true
)

// SetDebug 设置GoBox是否为DEBUG模式
func SetDebug(debug bool) {
	DEBUG = debug
}

// WARN 输出警告信息，仅在DEBUG模式下有效
func WARN(format string, v ...any) {
	if DEBUG {
		log.Printf("WARN: "+format, v...)
	}
}
