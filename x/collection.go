// Copyright © 2021 - 2023 vity <vityme@icloud.com>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package x

import "reflect"

// ArrayEmpty  返回一个空 array
func ArrayEmpty[T any]() []T {
	return make([]T, 0)
}

// ArrayContains 检查指定元素是否包含在 array 内
func ArrayContains[T any](array []T, value T) bool {
	for _, item := range array {
		if reflect.DeepEqual(item, value) {
			return true
		}
	}
	return false
}

// SliceEmpty  返回一个空 slice
func SliceEmpty[T any]() []T {
	return make([]T, 0)
}

// SliceContains 检查指定元素是否包含在 slice 内
func SliceContains[T any](slice []T, value T) bool {
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

// MapEmpty  返回一个空 map
func MapEmpty[T any]() map[string]T {
	return make(map[string]T)
}

// MapKeys 返回键集合
func MapKeys(m map[string]any) []string {
	var keys []string
	for key := range m {
		keys = append(keys, key)
	}
	return keys
}

// MapContainsKey 是否包含指定键
func MapContainsKey(key string, m map[string]any) bool {
	return SliceContains[string](MapKeys(m), key)
}
