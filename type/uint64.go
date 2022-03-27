// Copyright © 2021 - 2022 vity <vityme@icloud.com>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package t

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	gbox "github.com/mvity/go-box"
	v "github.com/mvity/go-box/validator"
	"math"
	"strconv"
	"strings"
)

// UInt64 A nullable safe type, based on Golang basic type, supports JSON escape, Database field definition
type UInt64 struct {
	value *uint64
}

// NewUInt64 Generates a new object of type t.UInt64 based on the specified value
// Warning: When value is passed in as uint(X) and float(X), be aware of data overflow and loss of precision
func NewUInt64(value interface{}) UInt64 {
	uintv := UInt64{}
	if value == nil {
		return uintv
	}
	switch value.(type) {
	case bool:
		var val uint64 = 0
		if value.(bool) {
			val = 1
		}
		uintv.value = &val
	case string:
		sval := strings.TrimSpace(value.(string))
		if v.RegexpNumeric.MatchString(sval) {
			if strings.Contains(sval, ".") {
				fval, _ := strconv.ParseFloat(value.(string), 64)
				if fval < 0 || fval > math.MaxUint64 {
					gbox.WARN("Data overflow during type conversion. value: %v", fval)
				}
				val := uint64(fval)
				uintv.value = &val
			} else if strings.Contains(sval, "-") {
				ival, _ := strconv.ParseInt(value.(string), 10, 64)
				if ival < 0 {
					gbox.WARN("Data overflow during type conversion. value: %v", ival)
				}
				val := uint64(ival)
				uintv.value = &val
			} else {
				ival, _ := strconv.ParseUint(value.(string), 10, 64)
				if ival < 0 || ival > math.MaxUint64 {
					gbox.WARN("Data overflow during type conversion. value: %v", ival)
				}
				val := uint64(ival)
				uintv.value = &val
			}
		}
	case int:
		ival := value.(int)
		if ival < 0 {
			gbox.WARN("Data overflow during type conversion. value: %v", ival)
		}
		val := uint64(ival)
		uintv.value = &val
	case int8:
		ival := value.(int8)
		if ival < 0 {
			gbox.WARN("Data overflow during type conversion. value: %v", ival)
		}
		val := uint64(ival)
		uintv.value = &val
	case int16:
		ival := value.(int16)
		if ival < 0 {
			gbox.WARN("Data overflow during type conversion. value: %v", ival)
		}
		val := uint64(ival)
		uintv.value = &val
	case int32:
		ival := value.(int32)
		if ival < 0 {
			gbox.WARN("Data overflow during type conversion. value: %v", ival)
		}
		val := uint64(ival)
		uintv.value = &val
	case int64:
		ival := value.(int64)
		if ival < 0 {
			gbox.WARN("Data overflow during type conversion. value: %v", ival)
		}
		val := uint64(ival)
		uintv.value = &val
	case uint:
		uval := value.(uint)
		val := uint64(uval)
		uintv.value = &val
	case uint8:
		uval := value.(uint8)
		val := uint64(uval)
		uintv.value = &val
	case uint16:
		uval := value.(uint16)
		val := uint64(uval)
		uintv.value = &val
	case uint32:
		uval := value.(uint32)
		val := uint64(uval)
		uintv.value = &val
	case uint64:
		val := value.(uint64)
		uintv.value = &val
	case float32:
		fval := value.(float32)
		if fval < 0 || fval > math.MaxUint64 {
			gbox.WARN("Data overflow during type conversion. value: %v", fval)
		}
		val := uint64(fval)
		uintv.value = &val
	case float64:
		fval := value.(float64)
		if fval < 0 || fval > math.MaxUint64 {
			gbox.WARN("Data overflow during type conversion. value: %v", fval)
		}
		val := uint64(fval)
		uintv.value = &val
	}

	return uintv
}

// String Returns the string value of this object, implements the Stringer interface.
func (u UInt64) String() string {
	if u.IsNil() {
		return "NaN"
	}
	return fmt.Sprintf("%v", *u.value)
}

// UIntValue Returns the numeric value of this object
func (u UInt64) UIntValue() uint64 {
	if u.IsNil() {
		return 0
	}
	return *u.value
}

// IsNil Check if the object is empty
func (u UInt64) IsNil() bool {
	return u.value == nil
}

// MarshalJSON implements the encoding json interface.
func (u UInt64) MarshalJSON() ([]byte, error) {
	return json.Marshal(u.value)
}

// UnmarshalJSON implements the encoding json interface.
func (u *UInt64) UnmarshalJSON(data []byte) error {
	sval := strings.Trim(string(data), "\"")
	if !v.RegexpNumeric.MatchString(sval) {
		u.value = nil
		return nil
	}
	var value interface{}
	if err := json.Unmarshal([]byte(sval), &value); err != nil {
		return err
	}
	u.value = NewUInt64(value).value
	return nil
}

// Value implements the driver Valuer interface.
func (u UInt64) Value() (driver.Value, error) {
	if u.IsNil() {
		return nil, nil
	}
	return *u.value, nil
}

// Scan implements the driver Scanner interface.
func (u *UInt64) Scan(value interface{}) error {
	u.value = NewUInt64(value).value
	return nil
}
