/*
 * Copyright © 2021 - 2025 vity <vityme@icloud.com>.
 *
 * Use of this source code is governed by an MIT-style
 * license that can be found in the LICENSE file.
 */

package x

// CastToPtr 将值类型转换为指针类型
func CastToPtr[T any](v T) *T {
	return &v
}

// CastFromPtr 将指针类型转换为值类型（若指针为nil，返回T的零值）
func CastFromPtr[T any](p *T) T {
	if p == nil {
		return *new(T) // 返回T的零值（等价于 T{}）
	}
	return *p
}
