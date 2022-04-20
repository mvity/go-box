// Copyright © 2021 - 2022 vity <vityme@icloud.com>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package k

import "encoding/json"

// ToJSONString 转为JSON字符串
func ToJSONString(value any) string {
	jsonVal, err := json.Marshal(value)
	if err != nil {
		return "{}"
	}
	return string(jsonVal)
}
