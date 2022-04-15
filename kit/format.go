// Copyright © 2021 - 2022 vity <vityme@icloud.com>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package k

import (
	"fmt"
	"strings"
)

// FormatChinaMoney 格式化人民币，从分到元
func FormatChinaMoney(value int64) string {
	return fmt.Sprintf("%.2f", (float64(value))/100.0)
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
