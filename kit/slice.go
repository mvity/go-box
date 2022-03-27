// Copyright © 2021 - 2022 vity <vityme@icloud.com>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package k

// InStringSlice Checks if slice contains the specified element
func InStringSlice(value string, slice []string) bool {
	for _, val := range slice {
		if val == value {
			return true
		}
	}
	return false
}
