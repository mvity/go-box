/*
 * Copyright © 2021 - 2023 vity <vityme@icloud.com>.
 *
 * Use of this source code is governed by an MIT-style
 * license that can be found in the LICENSE file.
 */

package x

// Ternary 三元表达式
func Ternary[T any](expr bool, first T, second T) T {
	if expr {
		return first
	} else {
		return second
	}
}
