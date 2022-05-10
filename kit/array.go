// Copyright © 2021 - 2022 vity <vityme@icloud.com>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package k

import "reflect"

// ArrayEmpty  返回一个空 array
func ArrayEmpty[T any]() []T {
	return make([]T, 0)
}

// ArrayContains 检查指定元素是否包含在 array 内
func ArrayContains[T any](value T, slice []T) bool {
	for _, item := range slice {
		if reflect.DeepEqual(value, item) {
			return true
		}
	}
	return false
}
