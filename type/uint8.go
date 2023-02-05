/*
 * Copyright © 2021 - 2023 vity <vityme@icloud.com>.
 *
 * Use of this source code is governed by an MIT-style
 * license that can be found in the LICENSE file.
 */

package t

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	gbox "github.com/mvity/go-box"
	"github.com/mvity/go-box/x"
	"math"
	"strconv"
	"strings"
)

// UInt8 A nullable safe type, based on Golang basic type, supports JSON escape, Database field definition
type UInt8 struct {
	value *uint8
}

// NewUInt8 Generates a new object of type t.UInt8 based on the specified value
// Warning: When value is passed in as uint(X) and float(X), be aware of data overflow and loss of precision
func NewUInt8(value any) UInt8 {
	uintv := UInt8{}
	if value == nil {
		return uintv
	}
	switch value.(type) {
	case bool:
		var val uint8 = 0
		if value.(bool) {
			val = 1
		}
		uintv.value = &val
	case string:
		sval := strings.TrimSpace(value.(string))
		if x.RegexpNumeric.MatchString(sval) {
			if strings.Contains(sval, ".") {
				fval, _ := strconv.ParseFloat(value.(string), 64)
				if fval < 0 || fval > math.MaxUint8 {
					gbox.WARN("Data overflow during type conversion. value: %v", fval)
				}
				val := uint8(fval)
				uintv.value = &val
			} else if strings.Contains(sval, "-") {
				ival, _ := strconv.ParseInt(value.(string), 10, 64)
				if ival < 0 || ival > math.MaxUint8 {
					gbox.WARN("Data overflow during type conversion. value: %v", ival)
				}
				val := uint8(ival)
				uintv.value = &val
			} else {
				ival, _ := strconv.ParseUint(value.(string), 10, 64)
				if ival < 0 || ival > math.MaxUint8 {
					gbox.WARN("Data overflow during type conversion. value: %v", ival)
				}
				val := uint8(ival)
				uintv.value = &val
			}
		}
	case int:
		ival := value.(int)
		if ival < 0 || ival > math.MaxUint8 {
			gbox.WARN("Data overflow during type conversion. value: %v", ival)
		}
		val := uint8(ival)
		uintv.value = &val
	case int8:
		ival := value.(int8)
		if ival < 0 {
			gbox.WARN("Data overflow during type conversion. value: %v", ival)
		}
		val := uint8(ival)
		uintv.value = &val
	case int16:
		ival := value.(int16)
		if ival < 0 || ival > math.MaxUint8 {
			gbox.WARN("Data overflow during type conversion. value: %v", ival)
		}
		val := uint8(ival)
		uintv.value = &val
	case int32:
		ival := value.(int32)
		if ival < 0 || ival > math.MaxUint8 {
			gbox.WARN("Data overflow during type conversion. value: %v", ival)
		}
		val := uint8(ival)
		uintv.value = &val
	case int64:
		ival := value.(int64)
		if ival < 0 || ival > math.MaxUint8 {
			gbox.WARN("Data overflow during type conversion. value: %v", ival)
		}
		val := uint8(ival)
		uintv.value = &val
	case uint:
		uval := value.(uint)
		if uval > math.MaxUint8 {
			gbox.WARN("Data overflow during type conversion. value: %v", uval)
		}
		val := uint8(uval)
		uintv.value = &val
	case uint8:
		val := value.(uint8)
		uintv.value = &val
	case uint16:
		uval := value.(uint16)
		if uval > math.MaxUint8 {
			gbox.WARN("Data overflow during type conversion. value: %v", uval)
		}
		val := uint8(uval)
		uintv.value = &val
	case uint32:
		uval := value.(uint32)
		if uval > math.MaxUint8 {
			gbox.WARN("Data overflow during type conversion. value: %v", uval)
		}
		val := uint8(uval)
		uintv.value = &val
	case uint64:
		uval := value.(uint64)
		if uval > math.MaxUint8 {
			gbox.WARN("Data overflow during type conversion. value: %v", uval)
		}
		val := uint8(uval)
		uintv.value = &val
	case float32:
		fval := value.(float32)
		if fval < 0 || fval > math.MaxUint8 {
			gbox.WARN("Data overflow during type conversion. value: %v", fval)
		}
		val := uint8(fval)
		uintv.value = &val
	case float64:
		fval := value.(float64)
		if fval < 0 || fval > math.MaxUint8 {
			gbox.WARN("Data overflow during type conversion. value: %v", fval)
		}
		val := uint8(fval)
		uintv.value = &val
	}

	return uintv
}

// String Returns the string value of this object, implements the Stringer interface.
func (u UInt8) String() string {
	if u.IsNil() {
		return "NaN"
	}
	return fmt.Sprintf("%v", *u.value)
}

// UIntValue Returns the numeric value of this object
func (u UInt8) UIntValue() uint8 {
	if u.IsNil() {
		return 0
	}
	return *u.value
}

// IsNil Check if the object is empty
func (u UInt8) IsNil() bool {
	return u.value == nil
}

// MarshalJSON implements the encoding json interface.
func (u UInt8) MarshalJSON() ([]byte, error) {
	return json.Marshal(u.value)
}

// UnmarshalJSON implements the encoding json interface.
func (u *UInt8) UnmarshalJSON(data []byte) error {
	sval := strings.Trim(string(data), "\"")
	if !x.RegexpNumeric.MatchString(sval) {
		u.value = nil
		return nil
	}
	var value any
	if err := json.Unmarshal([]byte(sval), &value); err != nil {
		return err
	}
	u.value = NewUInt8(value).value
	return nil
}

// Value implements the driver Valuer interface.
func (u UInt8) Value() (driver.Value, error) {
	if u.IsNil() {
		return nil, nil
	}
	return *u.value, nil
}

// Scan implements the driver Scanner interface.
func (u *UInt8) Scan(value any) error {
	u.value = NewUInt8(value).value
	return nil
}
