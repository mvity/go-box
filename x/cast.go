/*
 * Copyright © 2021 - 2025 vity <vityme@icloud.com>.
 *
 * Use of this source code is governed by an MIT-style
 * license that can be found in the LICENSE file.
 */

package x

import (
	"math"
	"strconv"
	"strings"

	gbox "github.com/mvity/go-box"
)

// ToPtr 将值类型转换为指针类型
func ToPtr[T any](v T) *T {
	return &v
}

// FromPtr 将指针类型转换为值类型（若指针为nil，返回T的零值）
func FromPtr[T any](p *T) T {
	if p == nil {
		return *new(T) // 返回T的零值（等价于 T{}）
	}
	return *p
}

// ToBool 转换为 bool 类型
func ToBool(value any) bool {
	if value == nil {
		return false
	}
	switch value.(type) {
	case bool:
		return value.(bool)
	case string:
		val, _ := strconv.ParseBool(strings.TrimSpace(value.(string)))
		return val
	case int:
		return value.(int) > 0
	case int8:
		return value.(int8) > 0
	case int16:
		return value.(int16) > 0
	case int32:
		return value.(int32) > 0
	case int64:
		return value.(int64) > 0
	case uint:
		return value.(uint) > 0
	case uint8:
		return value.(uint8) > 0
	case uint16:
		return value.(uint16) > 0
	case uint32:
		return value.(uint32) > 0
	case uint64:
		return value.(uint64) > 0
	case float32:
		return value.(float32) > 0
	case float64:
		return value.(float64) > 0
	}
	return false
}

// ToString 转换为 string 类型
func ToString(value any) string {
	if value == nil {
		return ""
	}
	switch value.(type) {
	case bool:
		return strconv.FormatBool(value.(bool))
	case string:
		return value.(string)
	case int:
		return strconv.Itoa(value.(int))
	case int8:
		return strconv.Itoa(int(value.(int8)))
	case int16:
		return strconv.Itoa(int(value.(int16)))
	case int32:
		return strconv.Itoa(int(value.(int32)))
	case int64:
		return strconv.FormatInt(value.(int64), 10)
	case uint:
		return strconv.FormatUint(uint64(value.(uint)), 10)
	case uint8:
		return strconv.FormatUint(uint64(value.(uint8)), 10)
	case uint16:
		return strconv.FormatUint(uint64(value.(uint16)), 10)
	case uint32:
		return strconv.FormatUint(uint64(value.(uint32)), 10)
	case uint64:
		return strconv.FormatUint(value.(uint64), 10)
	case float32:
		return strconv.FormatFloat(float64(value.(float32)), 'g', 4, 64)
	case float64:
		return strconv.FormatFloat(value.(float64), 'g', 4, 64)
	}
	return ""
}

// ToInt 转换为 int 类型
func ToInt(value any) int {
	if value == nil {
		return 0
	}
	switch value.(type) {
	case bool:
		var val = 0
		if value.(bool) {
			val = 1
		}
		return val
	case string:
		sval := strings.TrimSpace(value.(string))
		if RegexpNumeric.MatchString(sval) {
			if strings.Contains(sval, ".") {
				val, _ := strconv.ParseFloat(sval, 64)
				if val > math.MaxInt {
					gbox.WARN("Data overflow during type conversion. value: %v", val)
				}
				return int(val)
			} else {
				val, _ := strconv.ParseInt(sval, 10, 64)
				if val > math.MaxInt {
					gbox.WARN("Data overflow during type conversion. value: %v", val)
				}
				return int(val)
			}
		}
	case int:
		return value.(int)
	case int8:
		return int(value.(int8))
	case int16:
		return int(value.(int16))
	case int32:
		return int(value.(int32))
	case int64:
		val := value.(int64)
		if val > math.MaxInt {
			gbox.WARN("Data overflow during type conversion. value: %v", val)
		}
		return int(val)
	case uint:
		val := value.(uint)
		if val > math.MaxInt {
			gbox.WARN("Data overflow during type conversion. value: %v", val)
		}
		return int(val)
	case uint8:
		return int(value.(uint8))
	case uint16:
		return int(value.(uint16))
	case uint32:
		return int(value.(uint32))
	case uint64:
		val := value.(uint64)
		if val > math.MaxInt {
			gbox.WARN("Data overflow during type conversion. value: %v", val)
		}
		return int(val)
	case float32:
		val := value.(float32)
		if val > math.MaxInt {
			gbox.WARN("Data overflow during type conversion. value: %v", val)
		}
		return int(val)
	case float64:
		val := value.(float64)
		if val > math.MaxInt {
			gbox.WARN("Data overflow during type conversion. value: %v", val)
		}
		return int(val)
	}
	return 0
}

// ToInt8 转换为 int8 类型
func ToInt8(value any) int8 {
	if value == nil {
		return 0
	}
	switch value.(type) {
	case bool:
		var val int8 = 0
		if value.(bool) {
			val = 1
		}
		return val
	case string:
		sval := strings.TrimSpace(value.(string))
		if RegexpNumeric.MatchString(sval) {
			if strings.Contains(sval, ".") {
				val, _ := strconv.ParseFloat(sval, 64)
				if val > math.MaxInt8 {
					gbox.WARN("Data overflow during type conversion. value: %v", val)
				}
				return int8(val)
			} else {
				val, _ := strconv.ParseInt(sval, 10, 64)
				if val > math.MaxInt8 {
					gbox.WARN("Data overflow during type conversion. value: %v", val)
				}
				return int8(val)
			}
		}
	case int:
		val := value.(int)
		if val > math.MaxInt8 {
			gbox.WARN("Data overflow during type conversion. value: %v", val)
		}
		return int8(val)
	case int8:
		return value.(int8)
	case int16:
		val := value.(int16)
		if val > math.MaxInt8 {
			gbox.WARN("Data overflow during type conversion. value: %v", val)
		}
		return int8(val)
	case int32:
		val := value.(int32)
		if val > math.MaxInt8 {
			gbox.WARN("Data overflow during type conversion. value: %v", val)
		}
		return int8(val)
	case int64:
		val := value.(int64)
		if val > math.MaxInt8 {
			gbox.WARN("Data overflow during type conversion. value: %v", val)
		}
		return int8(val)
	case uint:
		val := value.(uint)
		if val > math.MaxInt8 {
			gbox.WARN("Data overflow during type conversion. value: %v", val)
		}
		return int8(val)
	case uint8:
		val := value.(uint8)
		if val > math.MaxInt8 {
			gbox.WARN("Data overflow during type conversion. value: %v", val)
		}
		return int8(val)
	case uint16:
		val := value.(uint16)
		if val > math.MaxInt8 {
			gbox.WARN("Data overflow during type conversion. value: %v", val)
		}
		return int8(val)
	case uint32:
		val := value.(uint32)
		if val > math.MaxInt8 {
			gbox.WARN("Data overflow during type conversion. value: %v", val)
		}
		return int8(val)
	case uint64:
		val := value.(uint64)
		if val > math.MaxInt8 {
			gbox.WARN("Data overflow during type conversion. value: %v", val)
		}
		return int8(val)
	case float32:
		val := value.(float32)
		if val > math.MaxInt8 {
			gbox.WARN("Data overflow during type conversion. value: %v", val)
		}
		return int8(val)
	case float64:
		val := value.(float64)
		if val > math.MaxInt8 {
			gbox.WARN("Data overflow during type conversion. value: %v", val)
		}
		return int8(val)
	}
	return 0
}

// ToInt16 转换为 int16 类型
func ToInt16(value any) int16 {
	if value == nil {
		return 0
	}
	switch value.(type) {
	case bool:
		var val int16 = 0
		if value.(bool) {
			val = 1
		}
		return val
	case string:
		sval := strings.TrimSpace(value.(string))
		if RegexpNumeric.MatchString(sval) {
			if strings.Contains(sval, ".") {
				val, _ := strconv.ParseFloat(value.(string), 64)
				if val > math.MaxInt16 {
					gbox.WARN("Data overflow during type conversion. value: %v", val)
				}
				return int16(val)
			} else {
				val, _ := strconv.ParseInt(value.(string), 10, 64)
				if val > math.MaxInt16 {
					gbox.WARN("Data overflow during type conversion. value: %v", val)
				}
				return int16(val)
			}
		}
	case int:
		val := value.(int)
		if val > math.MaxInt16 {
			gbox.WARN("Data overflow during type conversion. value: %v", val)
		}
		return int16(val)
	case int8:
		return int16(value.(int8))
	case int16:
		return value.(int16)
	case int32:
		val := value.(int32)
		if val > math.MaxInt16 {
			gbox.WARN("Data overflow during type conversion. value: %v", val)
		}
		return int16(val)
	case int64:
		val := value.(int64)
		if val > math.MaxInt16 {
			gbox.WARN("Data overflow during type conversion. value: %v", val)
		}
		return int16(val)
	case uint:
		val := value.(uint)
		if val > math.MaxInt16 {
			gbox.WARN("Data overflow during type conversion. value: %v", val)
		}
		return int16(val)
	case uint8:
		return int16(value.(uint8))
	case uint16:
		val := value.(uint16)
		if val > math.MaxInt16 {
			gbox.WARN("Data overflow during type conversion. value: %v", val)
		}
		return int16(val)
	case uint32:
		val := value.(uint32)
		if val > math.MaxInt16 {
			gbox.WARN("Data overflow during type conversion. value: %v", val)
		}
		return int16(val)
	case uint64:
		val := value.(uint64)
		if val > math.MaxInt16 {
			gbox.WARN("Data overflow during type conversion. value: %v", val)
		}
		return int16(val)
	case float32:
		val := value.(float32)
		if val > math.MaxInt16 {
			gbox.WARN("Data overflow during type conversion. value: %v", val)
		}
		return int16(val)
	case float64:
		val := value.(float64)
		if val > math.MaxInt16 {
			gbox.WARN("Data overflow during type conversion. value: %v", val)
		}
		return int16(val)
	}
	return 0
}

// ToInt32 转换为 int32 类型
func ToInt32(value any) int32 {
	if value == nil {
		return 0
	}
	switch value.(type) {
	case bool:
		var val int32 = 0
		if value.(bool) {
			val = 1
		}
		return val
	case string:
		sval := strings.TrimSpace(value.(string))
		if RegexpNumeric.MatchString(sval) {
			if strings.Contains(sval, ".") {
				val, _ := strconv.ParseFloat(sval, 64)
				if val > math.MaxInt32 {
					gbox.WARN("Data overflow during type conversion. value: %v", val)
				}
				return int32(val)
			} else {
				val, _ := strconv.ParseInt(sval, 10, 64)
				if val > math.MaxInt32 {
					gbox.WARN("Data overflow during type conversion. value: %v", val)
				}
				return int32(val)
			}
		}
	case int:
		val := value.(int)
		if val > math.MaxInt32 {
			gbox.WARN("Data overflow during type conversion. value: %v", val)
		}
		return int32(val)
	case int8:
		return int32(value.(int8))
	case int16:
		return int32(value.(int16))
	case int32:
		return value.(int32)
	case int64:
		val := value.(int64)
		if val > math.MaxInt32 {
			gbox.WARN("Data overflow during type conversion. value: %v", val)
		}
		return int32(val)
	case uint:
		val := value.(uint)
		if val > math.MaxInt32 {
			gbox.WARN("Data overflow during type conversion. value: %v", val)
		}
		return int32(val)
	case uint8:
		return int32(value.(uint8))
	case uint16:
		return int32(value.(uint16))
	case uint32:
		val := value.(uint32)
		if val > math.MaxInt32 {
			gbox.WARN("Data overflow during type conversion. value: %v", val)
		}
		return int32(val)
	case uint64:
		val := value.(uint64)
		if val > math.MaxInt32 {
			gbox.WARN("Data overflow during type conversion. value: %v", val)
		}
		return int32(val)
	case float32:
		val := value.(float32)
		if val > math.MaxInt32 {
			gbox.WARN("Data overflow during type conversion. value: %v", val)
		}
		return int32(val)
	case float64:
		val := value.(float64)
		if val > math.MaxInt32 {
			gbox.WARN("Data overflow during type conversion. value: %v", val)
		}
		return int32(val)
	}
	return 0
}

// ToInt64 转换为 int64 类型
func ToInt64(value any) int64 {
	if value == nil {
		return 0
	}
	switch value.(type) {
	case bool:
		var val int64 = 0
		if value.(bool) {
			val = 1
		}
		return val
	case string:
		sval := strings.TrimSpace(value.(string))
		if RegexpNumeric.MatchString(sval) {
			if strings.Contains(sval, ".") {
				val, _ := strconv.ParseFloat(sval, 64)
				if val > math.MaxInt64 {
					gbox.WARN("Data overflow during type conversion. value: %v", val)
				}
				return int64(val)
			} else {
				val, _ := strconv.ParseInt(sval, 10, 64)
				return val
			}
		} else {
			value = 0
		}
	case int:
		return int64(value.(int))
	case int8:
		return int64(value.(int8))
	case int16:
		return int64(value.(int16))
	case int32:
		return int64(value.(int32))
	case int64:
		return value.(int64)
	case uint:
		val := value.(uint)
		if val > math.MaxInt64 {
			gbox.WARN("Data overflow during type conversion. value: %v", val)
		}
		return int64(val)
	case uint8:
		return int64(value.(uint8))
	case uint16:
		return int64(value.(uint16))
	case uint32:
		return int64(value.(uint32))
	case uint64:
		val := value.(uint64)
		if val > math.MaxInt64 {
			gbox.WARN("Data overflow during type conversion. value: %v", val)
		}
		return int64(val)
	case float32:
		val := value.(float32)
		if val > math.MaxInt64 {
			gbox.WARN("Data overflow during type conversion. value: %v", val)
		}
		return int64(val)
	case float64:
		val := value.(float64)
		if val > math.MaxInt64 {
			gbox.WARN("Data overflow during type conversion. value: %v", val)
		}
		return int64(val)
	}

	return 0
}

// ToUInt 转换为 uint 类型
func ToUInt(value any) uint {
	if value == nil {
		return 0
	}
	switch value.(type) {
	case bool:
		var val uint = 0
		if value.(bool) {
			val = 1
		}
		return val
	case string:
		sval := strings.TrimSpace(value.(string))
		if RegexpNumeric.MatchString(sval) {
			if strings.Contains(sval, ".") {
				val, _ := strconv.ParseFloat(sval, 64)
				if val < 0 || val > math.MaxUint {
					gbox.WARN("Data overflow during type conversion. value: %v", val)
				}
				return uint(val)
			} else if strings.Contains(sval, "-") {
				val, _ := strconv.ParseInt(sval, 10, 64)
				if val < 0 {
					gbox.WARN("Data overflow during type conversion. value: %v", val)
				}
				return uint(val)
			} else {
				val, _ := strconv.ParseUint(sval, 10, 64)
				if val < 0 || val > math.MaxUint {
					gbox.WARN("Data overflow during type conversion. value: %v", val)
				}
				return uint(val)
			}
		}
	case int:
		val := value.(int)
		if val < 0 {
			gbox.WARN("Data overflow during type conversion. value: %v", val)
		}
		return uint(val)
	case int8:
		val := value.(int8)
		if val < 0 {
			gbox.WARN("Data overflow during type conversion. value: %v", val)
		}
		return uint(val)
	case int16:
		val := value.(int16)
		if val < 0 {
			gbox.WARN("Data overflow during type conversion. value: %v", val)
		}
		return uint(val)
	case int32:
		val := value.(int32)
		if val < 0 {
			gbox.WARN("Data overflow during type conversion. value: %v", val)
		}
		return uint(val)
	case int64:
		val := value.(int64)
		if val < 0 {
			gbox.WARN("Data overflow during type conversion. value: %v", val)
		}
		return uint(val)
	case uint:
		return value.(uint)
	case uint8:
		return uint(value.(uint8))
	case uint16:
		return uint(value.(uint16))
	case uint32:
		return uint(value.(uint32))
	case uint64:
		val := value.(uint64)
		if val > math.MaxUint {
			gbox.WARN("Data overflow during type conversion. value: %v", val)
		}
		return uint(val)
	case float32:
		val := value.(float32)
		if val < 0 || val > math.MaxUint {
			gbox.WARN("Data overflow during type conversion. value: %v", val)
		}
		return uint(val)
	case float64:
		val := value.(float64)
		if val < 0 || val > math.MaxUint {
			gbox.WARN("Data overflow during type conversion. value: %v", val)
		}
		return uint(val)
	}

	return 0
}

// ToUInt8 转换为 uint8 类型
func ToUInt8(value any) uint8 {
	if value == nil {
		return 0
	}
	switch value.(type) {
	case bool:
		var val uint8 = 0
		if value.(bool) {
			val = 1
		}
		return val
	case string:
		sval := strings.TrimSpace(value.(string))
		if RegexpNumeric.MatchString(sval) {
			if strings.Contains(sval, ".") {
				val, _ := strconv.ParseFloat(sval, 64)
				if val < 0 || val > math.MaxUint8 {
					gbox.WARN("Data overflow during type conversion. value: %v", val)
				}
				return uint8(val)
			} else if strings.Contains(sval, "-") {
				val, _ := strconv.ParseInt(sval, 10, 64)
				if val < 0 || val > math.MaxUint8 {
					gbox.WARN("Data overflow during type conversion. value: %v", val)
				}
				return uint8(val)
			} else {
				val, _ := strconv.ParseUint(sval, 10, 64)
				if val < 0 || val > math.MaxUint8 {
					gbox.WARN("Data overflow during type conversion. value: %v", val)
				}
				return uint8(val)
			}
		}
	case int:
		val := value.(int)
		if val < 0 || val > math.MaxUint8 {
			gbox.WARN("Data overflow during type conversion. value: %v", val)
		}
		return uint8(val)
	case int8:
		val := value.(int8)
		if val < 0 {
			gbox.WARN("Data overflow during type conversion. value: %v", val)
		}
		return uint8(val)
	case int16:
		val := value.(int16)
		if val < 0 || val > math.MaxUint8 {
			gbox.WARN("Data overflow during type conversion. value: %v", val)
		}
		return uint8(val)
	case int32:
		val := value.(int32)
		if val < 0 || val > math.MaxUint8 {
			gbox.WARN("Data overflow during type conversion. value: %v", val)
		}
		return uint8(val)
	case int64:
		val := value.(int64)
		if val < 0 || val > math.MaxUint8 {
			gbox.WARN("Data overflow during type conversion. value: %v", val)
		}
		return uint8(val)
	case uint:
		val := value.(uint)
		if val > math.MaxUint8 {
			gbox.WARN("Data overflow during type conversion. value: %v", val)
		}
		return uint8(val)
	case uint8:
		return value.(uint8)
	case uint16:
		val := value.(uint16)
		if val > math.MaxUint8 {
			gbox.WARN("Data overflow during type conversion. value: %v", val)
		}
		return uint8(val)
	case uint32:
		val := value.(uint32)
		if val > math.MaxUint8 {
			gbox.WARN("Data overflow during type conversion. value: %v", val)
		}
		return uint8(val)
	case uint64:
		val := value.(uint64)
		if val > math.MaxUint8 {
			gbox.WARN("Data overflow during type conversion. value: %v", val)
		}
		return uint8(val)
	case float32:
		val := value.(float32)
		if val < 0 || val > math.MaxUint8 {
			gbox.WARN("Data overflow during type conversion. value: %v", val)
		}
		return uint8(val)
	case float64:
		val := value.(float64)
		if val < 0 || val > math.MaxUint8 {
			gbox.WARN("Data overflow during type conversion. value: %v", val)
		}
		return uint8(val)
	}
	return 0
}

// ToUInt16 转换为 uint16 类型
func ToUInt16(value any) uint16 {
	if value == nil {
		return 0
	}
	switch value.(type) {
	case bool:
		var val uint16 = 0
		if value.(bool) {
			val = 1
		}
		return val
	case string:
		sval := strings.TrimSpace(value.(string))
		if RegexpNumeric.MatchString(sval) {
			if strings.Contains(sval, ".") {
				val, _ := strconv.ParseFloat(sval, 64)
				if val < 0 || val > math.MaxUint16 {
					gbox.WARN("Data overflow during type conversion. value: %v", val)
				}
				return uint16(val)
			} else if strings.Contains(sval, "-") {
				val, _ := strconv.ParseInt(sval, 10, 64)
				if val < 0 || val > math.MaxUint16 {
					gbox.WARN("Data overflow during type conversion. value: %v", val)
				}
				return uint16(val)
			} else {
				val, _ := strconv.ParseUint(sval, 10, 64)
				if val < 0 || val > math.MaxUint16 {
					gbox.WARN("Data overflow during type conversion. value: %v", val)
				}
				return uint16(val)
			}
		}
	case int:
		val := value.(int)
		if val < 0 || val > math.MaxUint16 {
			gbox.WARN("Data overflow during type conversion. value: %v", val)
		}
		return uint16(val)
	case int8:
		val := value.(int8)
		if val < 0 {
			gbox.WARN("Data overflow during type conversion. value: %v", val)
		}
		return uint16(val)
	case int16:
		val := value.(int16)
		if val < 0 {
			gbox.WARN("Data overflow during type conversion. value: %v", val)
		}
		return uint16(val)
	case int32:
		val := value.(int32)
		if val < 0 || val > math.MaxUint16 {
			gbox.WARN("Data overflow during type conversion. value: %v", val)
		}
		return uint16(val)
	case int64:
		val := value.(int64)
		if val < 0 || val > math.MaxUint16 {
			gbox.WARN("Data overflow during type conversion. value: %v", val)
		}
		return uint16(val)
	case uint:
		val := value.(uint)
		if val > math.MaxUint16 {
			gbox.WARN("Data overflow during type conversion. value: %v", val)
		}
		return uint16(val)
	case uint8:
		return uint16(value.(uint8))
	case uint16:
		return value.(uint16)
	case uint32:
		val := value.(uint32)
		if val > math.MaxUint16 {
			gbox.WARN("Data overflow during type conversion. value: %v", val)
		}
		return uint16(val)
	case uint64:
		val := value.(uint64)
		if val > math.MaxUint16 {
			gbox.WARN("Data overflow during type conversion. value: %v", val)
		}
		return uint16(val)
	case float32:
		val := value.(float32)
		if val < 0 || val > math.MaxUint16 {
			gbox.WARN("Data overflow during type conversion. value: %v", val)
		}
		return uint16(val)
	case float64:
		val := value.(float64)
		if val < 0 || val > math.MaxUint16 {
			gbox.WARN("Data overflow during type conversion. value: %v", val)
		}
		return uint16(val)
	}

	return 0
}

// ToUInt32 转换为 uint32 类型
func ToUInt32(value any) uint32 {
	if value == nil {
		return 0
	}
	switch value.(type) {
	case bool:
		var val uint32 = 0
		if value.(bool) {
			val = 1
		}
		return val
	case string:
		sval := strings.TrimSpace(value.(string))
		if RegexpNumeric.MatchString(sval) {
			if strings.Contains(sval, ".") {
				val, _ := strconv.ParseFloat(sval, 64)
				if val < 0 || val > math.MaxUint32 {
					gbox.WARN("Data overflow during type conversion. value: %v", val)
				}
				return uint32(val)
			} else if strings.Contains(sval, "-") {
				val, _ := strconv.ParseInt(sval, 10, 64)
				if val < 0 || val > math.MaxUint32 {
					gbox.WARN("Data overflow during type conversion. value: %v", val)
				}
				return uint32(val)
			} else {
				val, _ := strconv.ParseUint(sval, 10, 64)
				if val < 0 || val > math.MaxUint32 {
					gbox.WARN("Data overflow during type conversion. value: %v", val)
				}
				return uint32(val)
			}
		}
	case int:
		val := value.(int)
		if val < 0 || val > math.MaxUint32 {
			gbox.WARN("Data overflow during type conversion. value: %v", val)
		}
		return uint32(val)
	case int8:
		val := value.(int8)
		if val < 0 {
			gbox.WARN("Data overflow during type conversion. value: %v", val)
		}
		return uint32(val)
	case int16:
		val := value.(int16)
		if val < 0 {
			gbox.WARN("Data overflow during type conversion. value: %v", val)
		}
		return uint32(val)
	case int32:
		val := value.(int32)
		if val < 0 {
			gbox.WARN("Data overflow during type conversion. value: %v", val)
		}
		return uint32(val)

	case int64:
		val := value.(int64)
		if val < 0 || val > math.MaxUint32 {
			gbox.WARN("Data overflow during type conversion. value: %v", val)
		}
		return uint32(val)
	case uint:
		val := value.(uint)
		if val > math.MaxUint32 {
			gbox.WARN("Data overflow during type conversion. value: %v", val)
		}
		return uint32(val)
	case uint8:
		return uint32(value.(uint8))
	case uint16:
		return uint32(value.(uint16))
	case uint32:
		return value.(uint32)
	case uint64:
		val := value.(uint64)
		if val > math.MaxUint32 {
			gbox.WARN("Data overflow during type conversion. value: %v", val)
		}
		return uint32(val)
	case float32:
		val := value.(float32)
		if val < 0 || val > math.MaxUint32 {
			gbox.WARN("Data overflow during type conversion. value: %v", val)
		}
		return uint32(val)
	case float64:
		val := value.(float64)
		if val < 0 || val > math.MaxUint32 {
			gbox.WARN("Data overflow during type conversion. value: %v", val)
		}
		return uint32(val)
	}
	return 0
}

// ToUInt64 转换为 uint64 类型
func ToUInt64(value any) uint64 {
	if value == nil {
		return 0
	}
	switch value.(type) {
	case bool:
		var val uint64 = 0
		if value.(bool) {
			val = 1
		}
		return val
	case string:
		sval := strings.TrimSpace(value.(string))
		if RegexpNumeric.MatchString(sval) {
			if strings.Contains(sval, ".") {
				val, _ := strconv.ParseFloat(sval, 64)
				if val < 0 || val > math.MaxUint64 {
					gbox.WARN("Data overflow during type conversion. value: %v", val)
				}
				return uint64(val)
			} else if strings.Contains(sval, "-") {
				val, _ := strconv.ParseInt(sval, 10, 64)
				if val < 0 {
					gbox.WARN("Data overflow during type conversion. value: %v", val)
				}
				return uint64(val)
			} else {
				val, _ := strconv.ParseUint(sval, 10, 64)
				if val < 0 || val > math.MaxUint64 {
					gbox.WARN("Data overflow during type conversion. value: %v", val)
				}
				return val
			}
		}
	case int:
		val := value.(int)
		if val < 0 {
			gbox.WARN("Data overflow during type conversion. value: %v", val)
		}
		return uint64(val)
	case int8:
		val := value.(int8)
		if val < 0 {
			gbox.WARN("Data overflow during type conversion. value: %v", val)
		}
		return uint64(val)
	case int16:
		val := value.(int16)
		if val < 0 {
			gbox.WARN("Data overflow during type conversion. value: %v", val)
		}
		return uint64(val)
	case int32:
		val := value.(int32)
		if val < 0 {
			gbox.WARN("Data overflow during type conversion. value: %v", val)
		}
		return uint64(val)
	case int64:
		val := value.(int64)
		if val < 0 {
			gbox.WARN("Data overflow during type conversion. value: %v", val)
		}
		return uint64(val)
	case uint:
		return uint64(value.(uint))
	case uint8:
		return uint64(value.(uint8))
	case uint16:
		return uint64(value.(uint16))
	case uint32:
		return uint64(value.(uint32))
	case uint64:
		return value.(uint64)
	case float32:
		val := value.(float32)
		if val < 0 || val > math.MaxUint64 {
			gbox.WARN("Data overflow during type conversion. value: %v", val)
		}
		return uint64(val)
	case float64:
		val := value.(float64)
		if val < 0 || val > math.MaxUint64 {
			gbox.WARN("Data overflow during type conversion. value: %v", val)
		}
		return uint64(val)
	}
	return 0
}

// ToFloat32 转换为 float32 类型
func ToFloat32(value any) float32 {
	if value == nil {
		return 0
	}
	switch value.(type) {
	case bool:
		var val float32 = 0
		if value.(bool) {
			val = 1
		}
		return val
	case string:
		sval := strings.TrimSpace(value.(string))
		if RegexpNumeric.MatchString(sval) {
			val, _ := strconv.ParseFloat(sval, 64)
			if val > math.MaxFloat32 || float64(val) < math.SmallestNonzeroFloat32 {
				gbox.WARN("Data overflow during type conversion. value: %v", val)
			}
			return float32(val)
		}
	case int:
		val := value.(int)
		if float64(val) > math.MaxFloat32 || float64(val) < math.SmallestNonzeroFloat32 {
			gbox.WARN("Data overflow during type conversion. value: %v", val)
		}
		return float32(val)
	case int8:
		return float32(value.(int8))
	case int16:
		return float32(value.(int16))
	case int32:
		return float32(value.(int32))
	case int64:
		val := value.(int64)
		if float64(val) > math.MaxFloat32 || float64(val) < math.SmallestNonzeroFloat32 {
			gbox.WARN("Data overflow during type conversion. value: %v", val)
		}
		return float32(val)
	case uint:
		val := value.(uint)
		if float64(val) > math.MaxFloat32 || float64(val) < math.SmallestNonzeroFloat32 {
			gbox.WARN("Data overflow during type conversion. value: %v", val)
		}
		return float32(val)
	case uint8:
		return float32(value.(uint8))
	case uint16:
		return float32(value.(uint16))
	case uint32:
		val := value.(uint32)
		if float64(val) > math.MaxFloat32 || float64(val) < math.SmallestNonzeroFloat32 {
			gbox.WARN("Data overflow during type conversion. value: %v", val)
		}
		return float32(val)
	case uint64:
		val := value.(uint64)
		if float64(val) > math.MaxFloat32 || float64(val) < math.SmallestNonzeroFloat32 {
			gbox.WARN("Data overflow during type conversion. value: %v", val)
		}
		return float32(val)
	case float32:
		return value.(float32)
	case float64:
		val := value.(float64)
		if val > math.MaxFloat32 || val < math.SmallestNonzeroFloat32 {
			gbox.WARN("Data overflow during type conversion. value: %v", val)
		}

		return float32(val)
	}
	return 0
}

// ToFloat64 转换为 float64 类型
func ToFloat64(value any) float64 {
	if value == nil {
		return 0
	}
	switch value.(type) {
	case bool:
		var val float64 = 0
		if value.(bool) {
			val = 1
		}
		return val
	case string:
		sval := strings.TrimSpace(value.(string))
		if RegexpNumeric.MatchString(sval) {
			val, _ := strconv.ParseFloat(sval, 64)
			return val
		}
	case int:
		val := value.(int)
		if float64(val) > math.MaxFloat64 {
			gbox.WARN("Data overflow during type conversion. value: %v", val)
		}
		return float64(val)
	case int8:
		return float64(value.(int8))
	case int16:
		return float64(value.(int16))
	case int32:
		return float64(value.(int32))
	case int64:
		val := value.(int64)
		if float64(val) > math.MaxFloat64 {
			gbox.WARN("Data overflow during type conversion. value: %v", val)
		}
		return float64(val)
	case uint:
		val := value.(uint)
		if float64(val) > math.MaxFloat64 {
			gbox.WARN("Data overflow during type conversion. value: %v", val)
		}
		return float64(val)
	case uint8:
		return float64(value.(uint8))
	case uint16:
		return float64(value.(uint16))
	case uint32:
		val := value.(uint32)
		if float64(val) > math.MaxFloat64 {
			gbox.WARN("Data overflow during type conversion. value: %v", val)
		}
		return float64(val)
	case uint64:
		val := value.(uint64)
		if float64(val) > math.MaxFloat64 {
			gbox.WARN("Data overflow during type conversion. value: %v", val)
		}
		return float64(val)
	case float32:
		return float64(value.(float32))
	case float64:
		return value.(float64)
	}
	return 0
}
