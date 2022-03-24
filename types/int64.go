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

// Int64 A nullable safe type, based on Golang basic type, supports JSON escape, Database field definition
type Int64 struct {
	value *int64
}

// NewInt64 Generates a new object of type t.Int64 based on the specified value
// Warning: When value is passed in as uint(X) and float(X), be aware of data overflow and loss of precision
func NewInt64(value interface{}) Int64 {
	intv := Int64{}
	if value == nil {
		return intv
	}
	switch value.(type) {
	case bool:
		var val int64 = 0
		if value.(bool) {
			val = 1
		}
		intv.value = &val
	case string:
		sval := value.(string)
		if v.RegexpNumeric.MatchString(sval) {
			if strings.Contains(sval, ".") {
				fval, _ := strconv.ParseFloat(value.(string), 64)
				if fval > math.MaxInt64 {
					gbox.WARN("Data overflow during type conversion. value: %v", fval)
				}
				val := int64(fval)
				intv.value = &val
			} else {
				ival, _ := strconv.ParseInt(value.(string), 10, 64)
				if ival > math.MaxInt64 {
					gbox.WARN("Data overflow during type conversion. value: %v", ival)
				}
				val := ival
				intv.value = &val
			}
		}
	case int:
		ival := value.(int)
		val := int64(ival)
		intv.value = &val
	case int8:
		ival := value.(int8)
		val := int64(ival)
		intv.value = &val
	case int16:
		ival := value.(int16)
		val := int64(ival)
		intv.value = &val
	case int32:
		ival := value.(int32)
		val := int64(ival)
		intv.value = &val
	case int64:
		val := value.(int64)
		intv.value = &val
	case uint:
		uval := value.(uint)
		if uval > math.MaxInt64 {
			gbox.WARN("Data overflow during type conversion. value: %v", uval)
		}
		val := int64(uval)
		intv.value = &val
	case uint8:
		uval := value.(uint8)
		val := int64(uval)
		intv.value = &val
	case uint16:
		uval := value.(uint16)
		val := int64(uval)
		intv.value = &val
	case uint32:
		uval := value.(uint32)
		val := int64(uval)
		intv.value = &val
	case uint64:
		uval := value.(uint64)
		if uval > math.MaxInt64 {
			gbox.WARN("Data overflow during type conversion. value: %v", uval)
		}
		val := int64(uval)
		intv.value = &val
	case float32:
		fval := value.(float32)
		if fval > math.MaxInt64 {
			gbox.WARN("Data overflow during type conversion. value: %v", fval)
		}
		val := int64(fval)
		intv.value = &val
	case float64:
		fval := value.(float64)
		if fval > math.MaxInt64 {
			gbox.WARN("Data overflow during type conversion. value: %v", fval)
		}
		val := int64(fval)
		intv.value = &val
	}

	return intv
}

// String Returns the string value of this object, implements the Stringer interface.
func (i Int64) String() string {
	if i.IsNil() {
		return "NaN"
	}
	return fmt.Sprintf("%v", *i.value)
}

// Int64Value Returns the numeric value of this object
func (i Int64) Int64Value() int64 {
	if i.IsNil() {
		return 0
	}
	return *i.value
}

// IsNil Check if the object is empty
func (i Int64) IsNil() bool {
	return i.value == nil
}

// MarshalJSON implements the encoding json interface.
func (i Int64) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.value)
}

// UnmarshalJSON implements the encoding json interface.
func (i *Int64) UnmarshalJSON(data []byte) error {
	sval := strings.Trim(string(data), "\"")
	if !v.RegexpNumeric.MatchString(sval) {
		i.value = nil
		return nil
	}
	var value interface{}
	if err := json.Unmarshal([]byte(sval), &value); err != nil {
		return err
	}
	if value == nil {
		i.value = nil
		return nil
	}
	intv := NewInt64(value)
	i.value = intv.value
	return nil
}

// Value implements the driver Valuer interface.
func (i Int64) Value() (driver.Value, error) {
	if i.IsNil() {
		return nil, nil
	}
	return *i.value, nil
}

// Scan implements the driver Scanner interface.
func (i *Int64) Scan(value interface{}) error {
	if value == nil {
		i.value = nil
		return nil
	}
	if val, ok := value.(int64); !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal int64 value:", value))
	} else {
		i.value = &val
	}
	return nil
}
