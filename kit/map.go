// Copyright © 2021 - 2022 vity <vityme@icloud.com>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package k

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
	return SliceContains[string](key, MapKeys(m))
}
