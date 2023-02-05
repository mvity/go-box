// Copyright © 2021 - 2023 vity <vityme@icloud.com>.
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
	"github.com/mvity/go-box/x"
	"math"
	"strconv"
	"strings"
)

// Date A nullable safe type, based on Golang basic type, supports JSON escape, Database field definition
type Date struct {
	value *int32
}

const dayms = 24 * 60 * 60 * 1000

// NewDate Generates a new object of type t.Date based on the specified value
// Warning: When value is passed in as uint(X) and float(X), be aware of data overflow and loss of precision
func NewDate(value any) Date {
	intv := Date{}
	if value == nil {
		return intv
	}
	switch value.(type) {
	case string:
		sval := strings.TrimSpace(value.(string))
		if x.RegexpNumeric.MatchString(sval) && !strings.Contains(sval, ".") {
			ival, _ := strconv.ParseInt(value.(string), 10, 64)
			if ival < math.MaxInt32 {
				val := int32(ival)
				intv.value = &val
			}
		}
	case int:
		ival := value.(int)
		val := int32(ival)
		intv.value = &val
	case int8:
		ival := value.(int8)
		val := int32(ival)
		intv.value = &val
	case int16:
		ival := value.(int16)
		val := int32(ival)
		intv.value = &val
	case int32:
		val := value.(int32)
		intv.value = &val
	case uint:
		uval := value.(uint)
		if uval > math.MaxInt32 {
			gbox.WARN("Data overflow during type conversion. value: %v", uval)
		}
		val := int32(uval)
		intv.value = &val
	case uint8:
		uval := value.(uint8)
		val := int32(uval)
		intv.value = &val
	case uint16:
		uval := value.(uint16)
		val := int32(uval)
		intv.value = &val
	case uint32:
		uval := value.(uint32)
		if uval > math.MaxInt32 {
			gbox.WARN("Data overflow during type conversion. value: %v", uval)
		}
		val := int32(uval)
		intv.value = &val
	default:
		panic(errors.New("unsupported data, The specified data cannot be converted to an object of type Date"))
	}

	return intv
}

// String Returns the string value of this object, implements the Stringer interface.
func (i Date) String() string {
	if i.IsNil() {
		return "NaN"
	}
	return fmt.Sprintf("%v", *i.value)
}

// DateValue Returns the numeric value of this object
func (i Date) DateValue() int32 {
	if i.IsNil() {
		return 0
	}
	return *i.value
}

// IsNil Check if the object is empty
func (i Date) IsNil() bool {
	return i.value == nil
}

// MarshalJSON implements the encoding json interface.
func (i Date) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.value)
}

// UnmarshalJSON implements the encoding json interface.
func (i *Date) UnmarshalJSON(data []byte) error {
	sval := strings.Trim(string(data), "\"")
	if !x.RegexpNumeric.MatchString(sval) {
		i.value = nil
		return nil
	}
	var value any
	if err := json.Unmarshal([]byte(sval), &value); err != nil {
		return err
	}
	if value == nil {
		i.value = nil
		return nil
	}
	intv := NewDate(value)
	i.value = intv.value
	return nil
}

// Value implements the driver Valuer interface.
func (i Date) Value() (driver.Value, error) {
	if i.IsNil() {
		return nil, nil
	}
	return *i.value, nil
}

// Scan implements the driver Scanner interface.
func (i *Date) Scan(value any) error {
	if value == nil {
		i.value = nil
		return nil
	}
	if val, ok := value.(int32); !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal int value:", value))
	} else {
		i.value = &val
	}
	return nil
}
