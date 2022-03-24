// Copyright © 2021 - 2022 vity <vityme@icloud.com>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package gbox

import "log"

var (
	DEBUG = true
)

// SetDebug Set whether to enable Debug mode
func SetDebug(debug bool) {
	DEBUG = debug
}

// WARN Output console log, only available in Debug mode
func WARN(format string, v ...interface{}) {
	if DEBUG {
		log.Printf("WARN: "+format, v...)
	}
}
