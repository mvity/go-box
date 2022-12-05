// Copyright © 2021 - 2022 vity <vityme@icloud.com>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

// Deprecated: The package move to x
package k

import "reflect"

// SliceEmpty  返回一个空 slice
func SliceEmpty[T any]() []T {
	return make([]T, 0)
}

// SliceContains 检查指定元素是否包含在 slice 内
func SliceContains[T any](value T, slice []T) bool {
	for _, item := range slice {
		if reflect.DeepEqual(value, item) {
			return true
		}
	}
	return false
}

// SliceInsert 插入元素到指定位置
func SliceInsert[T any](slice []T, index int, value T) []T {
	return append(slice[:index], append([]T{value}, slice[index:]...)...)
}
