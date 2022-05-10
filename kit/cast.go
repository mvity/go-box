// Copyright © 2021 - 2022 vity <vityme@icloud.com>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package k

import (
	gbox "github.com/mvity/go-box"
	v "github.com/mvity/go-box/validator"
	"math"
	"strconv"
	"strings"
)

var trueStrs = []string{"true", "1", "yes", "ok"}
var falseStrs = []string{"false", "0", "no"}

// CastToBool 转换为 bool 类型
func CastToBool(value any) bool {

	if value == nil {
		return false
	}
	switch value.(type) {
	case bool:
		return value.(bool)
	case string:
		sval := strings.ToLower(strings.TrimSpace(value.(string)))
		if sval != "" {
			var val bool
			if SliceContains(sval, trueStrs) {
				val = true
			} else if SliceContains(sval, falseStrs) {
				val = false
			} else {
				break
			}
			return val
		}
	case int:
		ival := value.(int)
		val := ival > 0
		return val
	case int8:
		ival := value.(int8)
		val := ival > 0
		return val
	case int16:
		ival := value.(int16)
		val := ival > 0
		return val
	case int32:
		ival := value.(int32)
		val := ival > 0
		return val
	case int64:
		ival := value.(int64)
		val := ival > 0
		return val
	case uint:
		uval := value.(uint)
		val := uval > 0
		return val
	case uint8:
		uval := value.(uint8)
		val := uval > 0
		return val
	case uint16:
		uval := value.(uint16)
		val := uval > 0
		return val
	case uint32:
		uval := value.(uint32)
		val := uval > 0
		return val
	case uint64:
		uval := value.(uint64)
		val := uval > 0
		return val
	case float32:
		fval := value.(float32)
		val := fval > 0
		return val
	case float64:
		fval := value.(float64)
		val := fval > 0
		return val
	}
	return false
}

// CastToString 转换为 string 类型
func CastToString(value any) string {
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

// CastToInt 转换为 int 类型
func CastToInt(value any) int {
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
		if v.RegexpNumeric.MatchString(sval) {
			if strings.Contains(sval, ".") {
				fval, _ := strconv.ParseFloat(value.(string), 64)
				if fval > math.MaxInt {
					gbox.WARN("Data overflow during type conversion. value: %v", fval)
				}
				val := int(fval)
				return val
			} else {
				ival, _ := strconv.ParseInt(value.(string), 10, 64)
				if ival > math.MaxInt {
					gbox.WARN("Data overflow during type conversion. value: %v", ival)
				}
				val := int(ival)
				return val
			}
		}
	case int:
		val := value.(int)
		return val
	case int8:
		ival := value.(int8)
		val := int(ival)
		return val
	case int16:
		ival := value.(int16)
		val := int(ival)
		return val
	case int32:
		ival := value.(int32)
		val := int(ival)
		return val
	case int64:
		ival := value.(int64)
		if ival > math.MaxInt {
			gbox.WARN("Data overflow during type conversion. value: %v", ival)
		}
		val := int(ival)
		return val
	case uint:
		uval := value.(uint)
		if uval > math.MaxInt {
			gbox.WARN("Data overflow during type conversion. value: %v", uval)
		}
		val := int(uval)
		return val
	case uint8:
		uval := value.(uint8)
		val := int(uval)
		return val
	case uint16:
		uval := value.(uint16)
		val := int(uval)
		return val
	case uint32:
		uval := value.(uint32)
		val := int(uval)
		return val
	case uint64:
		uval := value.(uint64)
		if uval > math.MaxInt {
			gbox.WARN("Data overflow during type conversion. value: %v", uval)
		}
		val := int(uval)
		return val
	case float32:
		fval := value.(float32)
		if fval > math.MaxInt {
			gbox.WARN("Data overflow during type conversion. value: %v", fval)
		}
		val := int(fval)
		return val
	case float64:
		fval := value.(float64)
		if fval > math.MaxInt {
			gbox.WARN("Data overflow during type conversion. value: %v", fval)
		}
		val := int(fval)
		return val
	}
	return 0
}

// CastToInt8 转换为 int8 类型
func CastToInt8(value any) int8 {
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
		if v.RegexpNumeric.MatchString(sval) {
			if strings.Contains(sval, ".") {
				fval, _ := strconv.ParseFloat(value.(string), 64)
				if fval > math.MaxInt8 {
					gbox.WARN("Data overflow during type conversion. value: %v", fval)
				}
				val := int8(fval)
				return val
			} else {
				ival, _ := strconv.ParseInt(value.(string), 10, 64)
				if ival > math.MaxInt8 {
					gbox.WARN("Data overflow during type conversion. value: %v", ival)
				}
				val := int8(ival)
				return val
			}
		}
	case int:
		ival := value.(int)
		if ival > math.MaxInt8 {
			gbox.WARN("Data overflow during type conversion. value: %v", ival)
		}
		val := int8(ival)
		return val
	case int8:
		val := value.(int8)
		return val
	case int16:
		ival := value.(int16)
		if ival > math.MaxInt8 {
			gbox.WARN("Data overflow during type conversion. value: %v", ival)
		}
		val := int8(ival)
		return val
	case int32:
		ival := value.(int32)
		if ival > math.MaxInt8 {
			gbox.WARN("Data overflow during type conversion. value: %v", ival)
		}
		val := int8(ival)
		return val
	case int64:
		ival := value.(int64)
		if ival > math.MaxInt8 {
			gbox.WARN("Data overflow during type conversion. value: %v", ival)
		}
		val := int8(ival)
		return val
	case uint:
		uval := value.(uint)
		if uval > math.MaxInt8 {
			gbox.WARN("Data overflow during type conversion. value: %v", uval)
		}
		val := int8(uval)
		return val
	case uint8:
		uval := value.(uint8)
		if uval > math.MaxInt8 {
			gbox.WARN("Data overflow during type conversion. value: %v", uval)
		}
		val := int8(uval)
		return val
	case uint16:
		uval := value.(uint16)
		if uval > math.MaxInt8 {
			gbox.WARN("Data overflow during type conversion. value: %v", uval)
		}
		val := int8(uval)
		return val
	case uint32:
		uval := value.(uint32)
		if uval > math.MaxInt8 {
			gbox.WARN("Data overflow during type conversion. value: %v", uval)
		}
		val := int8(uval)
		return val
	case uint64:
		uval := value.(uint64)
		if uval > math.MaxInt8 {
			gbox.WARN("Data overflow during type conversion. value: %v", uval)
		}
		val := int8(uval)
		return val
	case float32:
		fval := value.(float32)
		if fval > math.MaxInt8 {
			gbox.WARN("Data overflow during type conversion. value: %v", fval)
		}
		val := int8(fval)
		return val
	case float64:
		fval := value.(float64)
		if fval > math.MaxInt8 {
			gbox.WARN("Data overflow during type conversion. value: %v", fval)
		}
		val := int8(fval)
		return val
	}

	return 0
}

// CastToInt16 转换为 int16 类型
func CastToInt16(value any) int16 {
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
		if v.RegexpNumeric.MatchString(sval) {
			if strings.Contains(sval, ".") {
				fval, _ := strconv.ParseFloat(value.(string), 64)
				if fval > math.MaxInt16 {
					gbox.WARN("Data overflow during type conversion. value: %v", fval)
				}
				val := int16(fval)
				return val
			} else {
				ival, _ := strconv.ParseInt(value.(string), 10, 64)
				if ival > math.MaxInt16 {
					gbox.WARN("Data overflow during type conversion. value: %v", ival)
				}
				val := int16(ival)
				return val
			}
		}
	case int:
		ival := value.(int)
		if ival > math.MaxInt16 {
			gbox.WARN("Data overflow during type conversion. value: %v", ival)
		}
		val := int16(ival)
		return val
	case int8:
		ival := value.(int8)
		val := int16(ival)
		return val
	case int16:
		val := value.(int16)
		return val
	case int32:
		ival := value.(int32)
		if ival > math.MaxInt16 {
			gbox.WARN("Data overflow during type conversion. value: %v", ival)
		}
		val := int16(ival)
		return val
	case int64:
		ival := value.(int64)
		if ival > math.MaxInt16 {
			gbox.WARN("Data overflow during type conversion. value: %v", ival)
		}
		val := int16(ival)
		return val
	case uint:
		uval := value.(uint)
		if uval > math.MaxInt16 {
			gbox.WARN("Data overflow during type conversion. value: %v", uval)
		}
		val := int16(uval)
		return val
	case uint8:
		uval := value.(uint8)
		val := int16(uval)
		return val
	case uint16:
		uval := value.(uint16)
		if uval > math.MaxInt16 {
			gbox.WARN("Data overflow during type conversion. value: %v", uval)
		}
		val := int16(uval)
		return val
	case uint32:
		uval := value.(uint32)
		if uval > math.MaxInt16 {
			gbox.WARN("Data overflow during type conversion. value: %v", uval)
		}
		val := int16(uval)
		return val
	case uint64:
		uval := value.(uint64)
		if uval > math.MaxInt16 {
			gbox.WARN("Data overflow during type conversion. value: %v", uval)
		}
		val := int16(uval)
		return val
	case float32:
		fval := value.(float32)
		if fval > math.MaxInt16 {
			gbox.WARN("Data overflow during type conversion. value: %v", fval)
		}
		val := int16(fval)
		return val
	case float64:
		fval := value.(float64)
		if fval > math.MaxInt16 {
			gbox.WARN("Data overflow during type conversion. value: %v", fval)
		}
		val := int16(fval)
		return val
	}

	return 0
}

// CastToInt32 转换为 int32 类型
func CastToInt32(value any) int32 {
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
		if v.RegexpNumeric.MatchString(sval) {
			if strings.Contains(sval, ".") {
				fval, _ := strconv.ParseFloat(value.(string), 64)
				if fval > math.MaxInt32 {
					gbox.WARN("Data overflow during type conversion. value: %v", fval)
				}
				val := int32(fval)
				return val
			} else {
				ival, _ := strconv.ParseInt(value.(string), 10, 64)
				if ival > math.MaxInt32 {
					gbox.WARN("Data overflow during type conversion. value: %v", ival)
				}
				val := int32(ival)
				return val
			}
		}
	case int:
		ival := value.(int)
		if ival > math.MaxInt32 {
			gbox.WARN("Data overflow during type conversion. value: %v", ival)
		}
		val := int32(ival)
		return val
	case int8:
		ival := value.(int8)
		val := int32(ival)
		return val
	case int16:
		ival := value.(int16)
		val := int32(ival)
		return val
	case int32:
		val := value.(int32)
		return val
	case int64:
		ival := value.(int64)
		if ival > math.MaxInt32 {
			gbox.WARN("Data overflow during type conversion. value: %v", ival)
		}
		val := int32(ival)
		return val
	case uint:
		uval := value.(uint)
		if uval > math.MaxInt32 {
			gbox.WARN("Data overflow during type conversion. value: %v", uval)
		}
		val := int32(uval)
		return val
	case uint8:
		uval := value.(uint8)
		val := int32(uval)
		return val
	case uint16:
		uval := value.(uint16)
		val := int32(uval)
		return val
	case uint32:
		uval := value.(uint32)
		if uval > math.MaxInt32 {
			gbox.WARN("Data overflow during type conversion. value: %v", uval)
		}
		val := int32(uval)
		return val
	case uint64:
		uval := value.(uint64)
		if uval > math.MaxInt32 {
			gbox.WARN("Data overflow during type conversion. value: %v", uval)
		}
		val := int32(uval)
		return val
	case float32:
		fval := value.(float32)
		if fval > math.MaxInt32 {
			gbox.WARN("Data overflow during type conversion. value: %v", fval)
		}
		val := int32(fval)
		return val
	case float64:
		fval := value.(float64)
		if fval > math.MaxInt32 {
			gbox.WARN("Data overflow during type conversion. value: %v", fval)
		}
		val := int32(fval)
		return val
	}

	return 0
}

// CastToInt64 转换为 int64 类型
func CastToInt64(value any) int64 {
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
		if v.RegexpNumeric.MatchString(sval) {
			if strings.Contains(sval, ".") {
				fval, _ := strconv.ParseFloat(value.(string), 64)
				if fval > math.MaxInt64 {
					gbox.WARN("Data overflow during type conversion. value: %v", fval)
				}
				val := int64(fval)
				return val
			} else {
				ival, _ := strconv.ParseInt(value.(string), 10, 64)
				if ival > math.MaxInt64 {
					gbox.WARN("Data overflow during type conversion. value: %v", ival)
				}
				val := ival
				return val
			}
		}
	case int:
		ival := value.(int)
		val := int64(ival)
		return val
	case int8:
		ival := value.(int8)
		val := int64(ival)
		return val
	case int16:
		ival := value.(int16)
		val := int64(ival)
		return val
	case int32:
		ival := value.(int32)
		val := int64(ival)
		return val
	case int64:
		val := value.(int64)
		return val
	case uint:
		uval := value.(uint)
		if uval > math.MaxInt64 {
			gbox.WARN("Data overflow during type conversion. value: %v", uval)
		}
		val := int64(uval)
		return val
	case uint8:
		uval := value.(uint8)
		val := int64(uval)
		return val
	case uint16:
		uval := value.(uint16)
		val := int64(uval)
		return val
	case uint32:
		uval := value.(uint32)
		val := int64(uval)
		return val
	case uint64:
		uval := value.(uint64)
		if uval > math.MaxInt64 {
			gbox.WARN("Data overflow during type conversion. value: %v", uval)
		}
		val := int64(uval)
		return val
	case float32:
		fval := value.(float32)
		if fval > math.MaxInt64 {
			gbox.WARN("Data overflow during type conversion. value: %v", fval)
		}
		val := int64(fval)
		return val
	case float64:
		fval := value.(float64)
		if fval > math.MaxInt64 {
			gbox.WARN("Data overflow during type conversion. value: %v", fval)
		}
		val := int64(fval)
		return val
	}

	return 0
}

// CastToUInt 转换为 uint 类型
func CastToUInt(value any) uint {
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
		if v.RegexpNumeric.MatchString(sval) {
			if strings.Contains(sval, ".") {
				fval, _ := strconv.ParseFloat(value.(string), 64)
				if fval < 0 || fval > math.MaxUint {
					gbox.WARN("Data overflow during type conversion. value: %v", fval)
				}
				val := uint(fval)
				return val
			} else if strings.Contains(sval, "-") {
				ival, _ := strconv.ParseInt(value.(string), 10, 64)
				if ival < 0 {
					gbox.WARN("Data overflow during type conversion. value: %v", ival)
				}
				val := uint(ival)
				return val
			} else {
				ival, _ := strconv.ParseUint(value.(string), 10, 64)
				if ival < 0 || ival > math.MaxUint {
					gbox.WARN("Data overflow during type conversion. value: %v", ival)
				}
				val := uint(ival)
				return val
			}
		}
	case int:
		ival := value.(int)
		if ival < 0 {
			gbox.WARN("Data overflow during type conversion. value: %v", ival)
		}
		val := uint(ival)
		return val
	case int8:
		ival := value.(int8)
		if ival < 0 {
			gbox.WARN("Data overflow during type conversion. value: %v", ival)
		}
		val := uint(ival)
		return val
	case int16:
		ival := value.(int16)
		if ival < 0 {
			gbox.WARN("Data overflow during type conversion. value: %v", ival)
		}
		val := uint(ival)
		return val
	case int32:
		ival := value.(int32)
		if ival < 0 {
			gbox.WARN("Data overflow during type conversion. value: %v", ival)
		}
		val := uint(ival)
		return val
	case int64:
		ival := value.(int64)
		if ival < 0 {
			gbox.WARN("Data overflow during type conversion. value: %v", ival)
		}
		val := uint(ival)
		return val
	case uint:
		val := value.(uint)
		return val
	case uint8:
		uval := value.(uint8)
		val := uint(uval)
		return val
	case uint16:
		uval := value.(uint16)
		val := uint(uval)
		return val
	case uint32:
		uval := value.(uint32)
		val := uint(uval)
		return val
	case uint64:
		uval := value.(uint64)
		if uval > math.MaxUint {
			gbox.WARN("Data overflow during type conversion. value: %v", uval)
		}
		val := uint(uval)
		return val
	case float32:
		fval := value.(float32)
		if fval < 0 || fval > math.MaxUint {
			gbox.WARN("Data overflow during type conversion. value: %v", fval)
		}
		val := uint(fval)
		return val
	case float64:
		fval := value.(float64)
		if fval < 0 || fval > math.MaxUint {
			gbox.WARN("Data overflow during type conversion. value: %v", fval)
		}
		val := uint(fval)
		return val
	}

	return 0
}

// CastToUInt8 转换为 uint8 类型
func CastToUInt8(value any) uint8 {
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
		if v.RegexpNumeric.MatchString(sval) {
			if strings.Contains(sval, ".") {
				fval, _ := strconv.ParseFloat(value.(string), 64)
				if fval < 0 || fval > math.MaxUint8 {
					gbox.WARN("Data overflow during type conversion. value: %v", fval)
				}
				val := uint8(fval)
				return val
			} else if strings.Contains(sval, "-") {
				ival, _ := strconv.ParseInt(value.(string), 10, 64)
				if ival < 0 || ival > math.MaxUint8 {
					gbox.WARN("Data overflow during type conversion. value: %v", ival)
				}
				val := uint8(ival)
				return val
			} else {
				ival, _ := strconv.ParseUint(value.(string), 10, 64)
				if ival < 0 || ival > math.MaxUint8 {
					gbox.WARN("Data overflow during type conversion. value: %v", ival)
				}
				val := uint8(ival)
				return val
			}
		}
	case int:
		ival := value.(int)
		if ival < 0 || ival > math.MaxUint8 {
			gbox.WARN("Data overflow during type conversion. value: %v", ival)
		}
		val := uint8(ival)
		return val
	case int8:
		ival := value.(int8)
		if ival < 0 {
			gbox.WARN("Data overflow during type conversion. value: %v", ival)
		}
		val := uint8(ival)
		return val
	case int16:
		ival := value.(int16)
		if ival < 0 || ival > math.MaxUint8 {
			gbox.WARN("Data overflow during type conversion. value: %v", ival)
		}
		val := uint8(ival)
		return val
	case int32:
		ival := value.(int32)
		if ival < 0 || ival > math.MaxUint8 {
			gbox.WARN("Data overflow during type conversion. value: %v", ival)
		}
		val := uint8(ival)
		return val
	case int64:
		ival := value.(int64)
		if ival < 0 || ival > math.MaxUint8 {
			gbox.WARN("Data overflow during type conversion. value: %v", ival)
		}
		val := uint8(ival)
		return val
	case uint:
		uval := value.(uint)
		if uval > math.MaxUint8 {
			gbox.WARN("Data overflow during type conversion. value: %v", uval)
		}
		val := uint8(uval)
		return val
	case uint8:
		val := value.(uint8)
		return val
	case uint16:
		uval := value.(uint16)
		if uval > math.MaxUint8 {
			gbox.WARN("Data overflow during type conversion. value: %v", uval)
		}
		val := uint8(uval)
		return val
	case uint32:
		uval := value.(uint32)
		if uval > math.MaxUint8 {
			gbox.WARN("Data overflow during type conversion. value: %v", uval)
		}
		val := uint8(uval)
		return val
	case uint64:
		uval := value.(uint64)
		if uval > math.MaxUint8 {
			gbox.WARN("Data overflow during type conversion. value: %v", uval)
		}
		val := uint8(uval)
		return val
	case float32:
		fval := value.(float32)
		if fval < 0 || fval > math.MaxUint8 {
			gbox.WARN("Data overflow during type conversion. value: %v", fval)
		}
		val := uint8(fval)
		return val
	case float64:
		fval := value.(float64)
		if fval < 0 || fval > math.MaxUint8 {
			gbox.WARN("Data overflow during type conversion. value: %v", fval)
		}
		val := uint8(fval)
		return val
	}

	return 0
}

// CastToUInt16 转换为 uint16 类型
func CastToUInt16(value any) uint16 {
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
		if v.RegexpNumeric.MatchString(sval) {
			if strings.Contains(sval, ".") {
				fval, _ := strconv.ParseFloat(value.(string), 64)
				if fval < 0 || fval > math.MaxUint16 {
					gbox.WARN("Data overflow during type conversion. value: %v", fval)
				}
				val := uint16(fval)
				return val
			} else if strings.Contains(sval, "-") {
				ival, _ := strconv.ParseInt(value.(string), 10, 64)
				if ival < 0 || ival > math.MaxUint16 {
					gbox.WARN("Data overflow during type conversion. value: %v", ival)
				}
				val := uint16(ival)
				return val
			} else {
				ival, _ := strconv.ParseUint(value.(string), 10, 64)
				if ival < 0 || ival > math.MaxUint16 {
					gbox.WARN("Data overflow during type conversion. value: %v", ival)
				}
				val := uint16(ival)
				return val
			}
		}
	case int:
		ival := value.(int)
		if ival < 0 || ival > math.MaxUint16 {
			gbox.WARN("Data overflow during type conversion. value: %v", ival)
		}
		val := uint16(ival)
		return val
	case int8:
		ival := value.(int8)
		if ival < 0 {
			gbox.WARN("Data overflow during type conversion. value: %v", ival)
		}
		val := uint16(ival)
		return val
	case int16:
		ival := value.(int16)
		val := uint16(ival)
		return val
	case int32:
		ival := value.(int32)
		if ival < 0 || ival > math.MaxUint16 {
			gbox.WARN("Data overflow during type conversion. value: %v", ival)
		}
		val := uint16(ival)
		return val
	case int64:
		ival := value.(int64)
		if ival < 0 || ival > math.MaxUint16 {
			gbox.WARN("Data overflow during type conversion. value: %v", ival)
		}
		val := uint16(ival)
		return val
	case uint:
		uval := value.(uint)
		if uval > math.MaxUint16 {
			gbox.WARN("Data overflow during type conversion. value: %v", uval)
		}
		val := uint16(uval)
		return val
	case uint8:
		uval := value.(uint8)
		val := uint16(uval)
		return val
	case uint16:
		val := value.(uint16)
		return val
	case uint32:
		uval := value.(uint32)
		if uval > math.MaxUint16 {
			gbox.WARN("Data overflow during type conversion. value: %v", uval)
		}
		val := uint16(uval)
		return val
	case uint64:
		uval := value.(uint64)
		if uval > math.MaxUint16 {
			gbox.WARN("Data overflow during type conversion. value: %v", uval)
		}
		val := uint16(uval)
		return val
	case float32:
		fval := value.(float32)
		if fval < 0 || fval > math.MaxUint16 {
			gbox.WARN("Data overflow during type conversion. value: %v", fval)
		}
		val := uint16(fval)
		return val
	case float64:
		fval := value.(float64)
		if fval < 0 || fval > math.MaxUint16 {
			gbox.WARN("Data overflow during type conversion. value: %v", fval)
		}
		val := uint16(fval)
		return val
	}

	return 0
}

// CastToUInt32 转换为 uint32 类型
func CastToUInt32(value any) uint32 {
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
		if v.RegexpNumeric.MatchString(sval) {
			if strings.Contains(sval, ".") {
				fval, _ := strconv.ParseFloat(value.(string), 64)
				if fval < 0 || fval > math.MaxUint32 {
					gbox.WARN("Data overflow during type conversion. value: %v", fval)
				}
				val := uint32(fval)
				return val
			} else if strings.Contains(sval, "-") {
				ival, _ := strconv.ParseInt(value.(string), 10, 64)
				if ival < 0 || ival > math.MaxUint32 {
					gbox.WARN("Data overflow during type conversion. value: %v", ival)
				}
				val := uint32(ival)
				return val
			} else {
				ival, _ := strconv.ParseUint(value.(string), 10, 64)
				if ival < 0 || ival > math.MaxUint32 {
					gbox.WARN("Data overflow during type conversion. value: %v", ival)
				}
				val := uint32(ival)
				return val
			}
		}
	case int:
		ival := value.(int)
		if ival < 0 || ival > math.MaxUint32 {
			gbox.WARN("Data overflow during type conversion. value: %v", ival)
		}
		val := uint32(ival)
		return val
	case int8:
		ival := value.(int8)
		if ival < 0 {
			gbox.WARN("Data overflow during type conversion. value: %v", ival)
		}
		val := uint32(ival)
		return val
	case int16:
		ival := value.(int16)
		val := uint32(ival)
		return val
	case int32:
		ival := value.(int32)
		val := uint32(ival)
		return val
	case int64:
		ival := value.(int64)
		if ival < 0 || ival > math.MaxUint32 {
			gbox.WARN("Data overflow during type conversion. value: %v", ival)
		}
		val := uint32(ival)
		return val
	case uint:
		uval := value.(uint)
		if uval > math.MaxUint32 {
			gbox.WARN("Data overflow during type conversion. value: %v", uval)
		}
		val := uint32(uval)
		return val
	case uint8:
		uval := value.(uint8)
		val := uint32(uval)
		return val
	case uint16:
		uval := value.(uint16)
		val := uint32(uval)
		return val
	case uint32:
		val := value.(uint32)
		return val
	case uint64:
		uval := value.(uint64)
		if uval > math.MaxUint32 {
			gbox.WARN("Data overflow during type conversion. value: %v", uval)
		}
		val := uint32(uval)
		return val
	case float32:
		fval := value.(float32)
		if fval < 0 || fval > math.MaxUint32 {
			gbox.WARN("Data overflow during type conversion. value: %v", fval)
		}
		val := uint32(fval)
		return val
	case float64:
		fval := value.(float64)
		if fval < 0 || fval > math.MaxUint32 {
			gbox.WARN("Data overflow during type conversion. value: %v", fval)
		}
		val := uint32(fval)
		return val
	}

	return 0
}

// CastToUInt64 转换为 uint64 类型
func CastToUInt64(value any) uint64 {
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
		if v.RegexpNumeric.MatchString(sval) {
			if strings.Contains(sval, ".") {
				fval, _ := strconv.ParseFloat(value.(string), 64)
				if fval < 0 || fval > math.MaxUint64 {
					gbox.WARN("Data overflow during type conversion. value: %v", fval)
				}
				val := uint64(fval)
				return val
			} else if strings.Contains(sval, "-") {
				ival, _ := strconv.ParseInt(value.(string), 10, 64)
				if ival < 0 {
					gbox.WARN("Data overflow during type conversion. value: %v", ival)
				}
				val := uint64(ival)
				return val
			} else {
				ival, _ := strconv.ParseUint(value.(string), 10, 64)
				if ival < 0 || ival > math.MaxUint64 {
					gbox.WARN("Data overflow during type conversion. value: %v", ival)
				}
				val := uint64(ival)
				return val
			}
		}
	case int:
		ival := value.(int)
		if ival < 0 {
			gbox.WARN("Data overflow during type conversion. value: %v", ival)
		}
		val := uint64(ival)
		return val
	case int8:
		ival := value.(int8)
		if ival < 0 {
			gbox.WARN("Data overflow during type conversion. value: %v", ival)
		}
		val := uint64(ival)
		return val
	case int16:
		ival := value.(int16)
		if ival < 0 {
			gbox.WARN("Data overflow during type conversion. value: %v", ival)
		}
		val := uint64(ival)
		return val
	case int32:
		ival := value.(int32)
		if ival < 0 {
			gbox.WARN("Data overflow during type conversion. value: %v", ival)
		}
		val := uint64(ival)
		return val
	case int64:
		ival := value.(int64)
		if ival < 0 {
			gbox.WARN("Data overflow during type conversion. value: %v", ival)
		}
		val := uint64(ival)
		return val
	case uint:
		uval := value.(uint)
		val := uint64(uval)
		return val
	case uint8:
		uval := value.(uint8)
		val := uint64(uval)
		return val
	case uint16:
		uval := value.(uint16)
		val := uint64(uval)
		return val
	case uint32:
		uval := value.(uint32)
		val := uint64(uval)
		return val
	case uint64:
		val := value.(uint64)
		return val
	case float32:
		fval := value.(float32)
		if fval < 0 || fval > math.MaxUint64 {
			gbox.WARN("Data overflow during type conversion. value: %v", fval)
		}
		val := uint64(fval)
		return val
	case float64:
		fval := value.(float64)
		if fval < 0 || fval > math.MaxUint64 {
			gbox.WARN("Data overflow during type conversion. value: %v", fval)
		}
		val := uint64(fval)
		return val
	}

	return 0
}

// CastToFloat32 转换为 float32 类型
func CastToFloat32(value any) float32 {
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
		if v.RegexpNumeric.MatchString(sval) {
			fval, _ := strconv.ParseFloat(value.(string), 64)
			if fval > math.MaxFloat32 || float64(fval) < math.SmallestNonzeroFloat32 {
				gbox.WARN("Data overflow during type conversion. value: %v", fval)
			}
			val := float32(fval)
			return val
		}
	case int:
		ival := value.(int)
		if float64(ival) > math.MaxFloat32 || float64(ival) < math.SmallestNonzeroFloat32 {
			gbox.WARN("Data overflow during type conversion. value: %v", ival)
		}
		val := float32(ival)
		return val
	case int8:
		ival := value.(int8)
		val := float32(ival)
		return val
	case int16:
		ival := value.(int16)
		val := float32(ival)
		return val
	case int32:
		ival := value.(int32)
		val := float32(ival)
		return val
	case int64:
		ival := value.(int64)
		if float64(ival) > math.MaxFloat32 || float64(ival) < math.SmallestNonzeroFloat32 {
			gbox.WARN("Data overflow during type conversion. value: %v", ival)
		}
		val := float32(ival)
		return val
	case uint:
		uval := value.(uint)
		if float64(uval) > math.MaxFloat32 || float64(uval) < math.SmallestNonzeroFloat32 {
			gbox.WARN("Data overflow during type conversion. value: %v", uval)
		}
		val := float32(uval)
		return val
	case uint8:
		uval := value.(uint8)
		val := float32(uval)
		return val
	case uint16:
		uval := value.(uint16)
		val := float32(uval)
		return val
	case uint32:
		uval := value.(uint32)
		if float64(uval) > math.MaxFloat32 || float64(uval) < math.SmallestNonzeroFloat32 {
			gbox.WARN("Data overflow during type conversion. value: %v", uval)
		}
		val := float32(uval)
		return val
	case uint64:
		uval := value.(uint64)
		if float64(uval) > math.MaxFloat32 || float64(uval) < math.SmallestNonzeroFloat32 {
			gbox.WARN("Data overflow during type conversion. value: %v", uval)
		}
		val := float32(uval)
		return val
	case float32:
		val := value.(float32)
		return val
	case float64:
		fval := value.(float64)
		if fval > math.MaxFloat32 || fval < math.SmallestNonzeroFloat32 {
			gbox.WARN("Data overflow during type conversion. value: %v", fval)
		}

		val := float32(fval)
		return val
	}

	return 0
}

// CastToFloat64 转换为 float64 类型
func CastToFloat64(value any) float64 {
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
		if v.RegexpNumeric.MatchString(sval) {
			val, _ := strconv.ParseFloat(value.(string), 64)
			return val
		}
	case int:
		ival := value.(int)
		if float64(ival) > math.MaxFloat64 {
			gbox.WARN("Data overflow during type conversion. value: %v", ival)
		}
		val := float64(ival)
		return val
	case int8:
		ival := value.(int8)
		val := float64(ival)
		return val
	case int16:
		ival := value.(int16)
		val := float64(ival)
		return val
	case int32:
		ival := value.(int32)
		val := float64(ival)
		return val
	case int64:
		ival := value.(int64)
		if float64(ival) > math.MaxFloat64 {
			gbox.WARN("Data overflow during type conversion. value: %v", ival)
		}
		val := float64(ival)
		return val
	case uint:
		uval := value.(uint)
		if float64(uval) > math.MaxFloat64 {
			gbox.WARN("Data overflow during type conversion. value: %v", uval)
		}
		val := float64(uval)
		return val
	case uint8:
		uval := value.(uint8)
		val := float64(uval)
		return val
	case uint16:
		uval := value.(uint16)
		val := float64(uval)
		return val
	case uint32:
		uval := value.(uint32)
		if float64(uval) > math.MaxFloat64 {
			gbox.WARN("Data overflow during type conversion. value: %v", uval)
		}
		val := float64(uval)
		return val
	case uint64:
		uval := value.(uint64)
		if float64(uval) > math.MaxFloat64 {
			gbox.WARN("Data overflow during type conversion. value: %v", uval)
		}
		val := float64(uval)
		return val
	case float32:
		fval := value.(float32)
		val := float64(fval)
		return val
	case float64:
		val := value.(float64)
		return val
	}

	return 0
}
