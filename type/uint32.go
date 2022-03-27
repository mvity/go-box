// Copyright © 2021 - 2022 vity <vityme@icloud.com>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package t

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	gbox "github.com/mvity/go-box"
	v "github.com/mvity/go-box/validator"
	"math"
	"strconv"
	"strings"
)

// UInt32 A nullable safe type, based on Golang basic type, supports JSON escape, Database field definition
type UInt32 struct {
	value *uint32
}

// NewUInt32 Generates a new object of type t.UInt32 based on the specified value
// Warning: When value is passed in as uint(X) and float(X), be aware of data overflow and loss of precision
func NewUInt32(value interface{}) UInt32 {
	uintv := UInt32{}
	if value == nil {
		return uintv
	}
	switch value.(type) {
	case bool:
		var val uint32 = 0
		if value.(bool) {
			val = 1
		}
		uintv.value = &val
	case string:
		sval := strings.TrimSpace(value.(string))
		if v.RegexpNumeric.MatchString(sval) {
			if strings.Contains(sval, ".") {
				fval, _ := strconv.ParseFloat(value.(string), 64)
				if fval < 0 || fval > math.MaxUint32 {
					gbox.WARN("Data overflow during type conversion. value: %v", fval)
				}
				val := uint32(fval)
				uintv.value = &val
			} else if strings.Contains(sval, "-") {
				ival, _ := strconv.ParseInt(value.(string), 10, 64)
				if ival < 0 || ival > math.MaxUint32 {
					gbox.WARN("Data overflow during type conversion. value: %v", ival)
				}
				val := uint32(ival)
				uintv.value = &val
			} else {
				ival, _ := strconv.ParseUint(value.(string), 10, 64)
				if ival < 0 || ival > math.MaxUint32 {
					gbox.WARN("Data overflow during type conversion. value: %v", ival)
				}
				val := uint32(ival)
				uintv.value = &val
			}
		}
	case int:
		ival := value.(int)
		if ival < 0 || ival > math.MaxUint32 {
			gbox.WARN("Data overflow during type conversion. value: %v", ival)
		}
		val := uint32(ival)
		uintv.value = &val
	case int8:
		ival := value.(int8)
		if ival < 0 {
			gbox.WARN("Data overflow during type conversion. value: %v", ival)
		}
		val := uint32(ival)
		uintv.value = &val
	case int16:
		ival := value.(int16)
		val := uint32(ival)
		uintv.value = &val
	case int32:
		ival := value.(int32)
		val := uint32(ival)
		uintv.value = &val
	case int64:
		ival := value.(int64)
		if ival < 0 || ival > math.MaxUint32 {
			gbox.WARN("Data overflow during type conversion. value: %v", ival)
		}
		val := uint32(ival)
		uintv.value = &val
	case uint:
		uval := value.(uint)
		if uval > math.MaxUint32 {
			gbox.WARN("Data overflow during type conversion. value: %v", uval)
		}
		val := uint32(uval)
		uintv.value = &val
	case uint8:
		uval := value.(uint8)
		val := uint32(uval)
		uintv.value = &val
	case uint16:
		uval := value.(uint16)
		val := uint32(uval)
		uintv.value = &val
	case uint32:
		val := value.(uint32)
		uintv.value = &val
	case uint64:
		uval := value.(uint64)
		if uval > math.MaxUint32 {
			gbox.WARN("Data overflow during type conversion. value: %v", uval)
		}
		val := uint32(uval)
		uintv.value = &val
	case float32:
		fval := value.(float32)
		if fval < 0 || fval > math.MaxUint32 {
			gbox.WARN("Data overflow during type conversion. value: %v", fval)
		}
		val := uint32(fval)
		uintv.value = &val
	case float64:
		fval := value.(float64)
		if fval < 0 || fval > math.MaxUint32 {
			gbox.WARN("Data overflow during type conversion. value: %v", fval)
		}
		val := uint32(fval)
		uintv.value = &val
	}

	return uintv
}

// String Returns the string value of this object, implements the Stringer interface.
func (u UInt32) String() string {
	if u.IsNil() {
		return "NaN"
	}
	return fmt.Sprintf("%v", *u.value)
}

// UIntValue Returns the numeric value of this object
func (u UInt32) UIntValue() uint32 {
	if u.IsNil() {
		return 0
	}
	return *u.value
}

// IsNil Check if the object is empty
func (u UInt32) IsNil() bool {
	return u.value == nil
}

// MarshalJSON implements the encoding json interface.
func (u UInt32) MarshalJSON() ([]byte, error) {
	return json.Marshal(u.value)
}

// UnmarshalJSON implements the encoding json interface.
func (u *UInt32) UnmarshalJSON(data []byte) error {
	sval := strings.Trim(string(data), "\"")
	if !v.RegexpNumeric.MatchString(sval) {
		u.value = nil
		return nil
	}
	var value interface{}
	if err := json.Unmarshal([]byte(sval), &value); err != nil {
		return err
	}
	if value == nil {
		u.value = nil
		return nil
	}
	uintv := NewUInt32(value)
	u.value = uintv.value
	return nil
}

// Value implements the driver Valuer interface.
func (u UInt32) Value() (driver.Value, error) {
	if u.IsNil() {
		return nil, nil
	}
	return *u.value, nil
}

// Scan implements the driver Scanner interface.
func (u *UInt32) Scan(value interface{}) error {
	if value == nil {
		u.value = nil
		return nil
	}
	if val, ok := value.(uint32); !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal uint32 value:", value))
	} else {
		u.value = &val
	}
	return nil
}
