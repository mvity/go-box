// Copyright © 2021 - 2022 vity <vityme@icloud.com>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package k

// MapEmpty  返回一个空Map
func MapEmpty[T any]() map[string]T {
	return make(map[string]T)
}
