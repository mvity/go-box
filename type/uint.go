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

// UInt A nullable safe type, based on Golang basic type, supports JSON escape, Database field definition
type UInt struct {
	value *uint
}

// NewUInt Generates a new object of type t.UInt based on the specified value
// Warning: When value is passed in as uint(X) and float(X), be aware of data overflow and loss of precision
func NewUInt(value interface{}) UInt {
	uintv := UInt{}
	if value == nil {
		return uintv
	}
	switch value.(type) {
	case bool:
		var val uint = 0
		if value.(bool) {
			val = 1
		}
		uintv.value = &val
	case string:
		sval := strings.TrimSpace(value.(string))
		if v.RegexpNumeric.MatchString(sval) {
			if strings.Contains(sval, ".") {
				fval, _ := strconv.ParseFloat(value.(string), 64)
				if fval < 0 || fval > math.MaxUint {
					gbox.WARN("Data overflow during type conversion. value: %v", fval)
				}
				val := uint(fval)
				uintv.value = &val
			} else if strings.Contains(sval, "-") {
				ival, _ := strconv.ParseInt(value.(string), 10, 64)
				if ival < 0 {
					gbox.WARN("Data overflow during type conversion. value: %v", ival)
				}
				val := uint(ival)
				uintv.value = &val
			} else {
				ival, _ := strconv.ParseUint(value.(string), 10, 64)
				if ival < 0 || ival > math.MaxUint {
					gbox.WARN("Data overflow during type conversion. value: %v", ival)
				}
				val := uint(ival)
				uintv.value = &val
			}
		}
	case int:
		ival := value.(int)
		if ival < 0 {
			gbox.WARN("Data overflow during type conversion. value: %v", ival)
		}
		val := uint(ival)
		uintv.value = &val
	case int8:
		ival := value.(int8)
		if ival < 0 {
			gbox.WARN("Data overflow during type conversion. value: %v", ival)
		}
		val := uint(ival)
		uintv.value = &val
	case int16:
		ival := value.(int16)
		if ival < 0 {
			gbox.WARN("Data overflow during type conversion. value: %v", ival)
		}
		val := uint(ival)
		uintv.value = &val
	case int32:
		ival := value.(int32)
		if ival < 0 {
			gbox.WARN("Data overflow during type conversion. value: %v", ival)
		}
		val := uint(ival)
		uintv.value = &val
	case int64:
		ival := value.(int64)
		if ival < 0 {
			gbox.WARN("Data overflow during type conversion. value: %v", ival)
		}
		val := uint(ival)
		uintv.value = &val
	case uint:
		val := value.(uint)
		uintv.value = &val
	case uint8:
		uval := value.(uint8)
		val := uint(uval)
		uintv.value = &val
	case uint16:
		uval := value.(uint16)
		val := uint(uval)
		uintv.value = &val
	case uint32:
		uval := value.(uint32)
		val := uint(uval)
		uintv.value = &val
	case uint64:
		uval := value.(uint64)
		if uval > math.MaxUint {
			gbox.WARN("Data overflow during type conversion. value: %v", uval)
		}
		val := uint(uval)
		uintv.value = &val
	case float32:
		fval := value.(float32)
		if fval < 0 || fval > math.MaxUint {
			gbox.WARN("Data overflow during type conversion. value: %v", fval)
		}
		val := uint(fval)
		uintv.value = &val
	case float64:
		fval := value.(float64)
		if fval < 0 || fval > math.MaxUint {
			gbox.WARN("Data overflow during type conversion. value: %v", fval)
		}
		val := uint(fval)
		uintv.value = &val
	}

	return uintv
}

// String Returns the string value of this object, implements the Stringer interface.
func (u UInt) String() string {
	if u.IsNil() {
		return "NaN"
	}
	return fmt.Sprintf("%v", *u.value)
}

// UIntValue Returns the numeric value of this object
func (u UInt) UIntValue() uint {
	if u.IsNil() {
		return 0
	}
	return *u.value
}

// IsNil Check if the object is empty
func (u UInt) IsNil() bool {
	return u.value == nil
}

// MarshalJSON implements the encoding json interface.
func (u UInt) MarshalJSON() ([]byte, error) {
	return json.Marshal(u.value)
}

// UnmarshalJSON implements the encoding json interface.
func (u *UInt) UnmarshalJSON(data []byte) error {
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
	uintv := NewUInt(value)
	u.value = uintv.value
	return nil
}

// Value implements the driver Valuer interface.
func (u UInt) Value() (driver.Value, error) {
	if u.IsNil() {
		return nil, nil
	}
	return *u.value, nil
}

// Scan implements the driver Scanner interface.
func (u *UInt) Scan(value interface{}) error {
	if value == nil {
		u.value = nil
		return nil
	}
	if val, ok := value.(uint); !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal uint value:", value))
	} else {
		u.value = &val
	}
	return nil
}
