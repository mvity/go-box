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

// Int A nullable safe type, based on Golang basic type, supports JSON escape, Database field definition
type Int struct {
	value *int
}

// NewInt Generates a new object of type t.Int based on the specified value
// Warning: When value is passed in as uint(X) and float(X), be aware of data overflow and loss of precision
func NewInt(value interface{}) Int {
	intv := Int{}
	if value == nil {
		return intv
	}
	switch value.(type) {
	case bool:
		var val = 0
		if value.(bool) {
			val = 1
		}
		intv.value = &val
	case string:
		sval := strings.TrimSpace(value.(string))
		if v.RegexpNumeric.MatchString(sval) {
			if strings.Contains(sval, ".") {
				fval, _ := strconv.ParseFloat(value.(string), 64)
				if fval > math.MaxInt {
					gbox.WARN("Data overflow during type conversion. value: %v", fval)
				}
				val := int(fval)
				intv.value = &val
			} else {
				ival, _ := strconv.ParseInt(value.(string), 10, 64)
				if ival > math.MaxInt {
					gbox.WARN("Data overflow during type conversion. value: %v", ival)
				}
				val := int(ival)
				intv.value = &val
			}
		}
	case int:
		val := value.(int)
		intv.value = &val
	case int8:
		ival := value.(int8)
		val := int(ival)
		intv.value = &val
	case int16:
		ival := value.(int16)
		val := int(ival)
		intv.value = &val
	case int32:
		ival := value.(int32)
		val := int(ival)
		intv.value = &val
	case int64:
		ival := value.(int64)
		if ival > math.MaxInt {
			gbox.WARN("Data overflow during type conversion. value: %v", ival)
		}
		val := int(ival)
		intv.value = &val
	case uint:
		uval := value.(uint)
		if uval > math.MaxInt {
			gbox.WARN("Data overflow during type conversion. value: %v", uval)
		}
		val := int(uval)
		intv.value = &val
	case uint8:
		uval := value.(uint8)
		val := int(uval)
		intv.value = &val
	case uint16:
		uval := value.(uint16)
		val := int(uval)
		intv.value = &val
	case uint32:
		uval := value.(uint32)
		val := int(uval)
		intv.value = &val
	case uint64:
		uval := value.(uint64)
		if uval > math.MaxInt {
			gbox.WARN("Data overflow during type conversion. value: %v", uval)
		}
		val := int(uval)
		intv.value = &val
	case float32:
		fval := value.(float32)
		if fval > math.MaxInt {
			gbox.WARN("Data overflow during type conversion. value: %v", fval)
		}
		val := int(fval)
		intv.value = &val
	case float64:
		fval := value.(float64)
		if fval > math.MaxInt {
			gbox.WARN("Data overflow during type conversion. value: %v", fval)
		}
		val := int(fval)
		intv.value = &val
	}

	return intv
}

// String Returns the string value of this object, implements the Stringer interface.
func (i Int) String() string {
	if i.IsNil() {
		return "NaN"
	}
	return fmt.Sprintf("%v", *i.value)
}

// IntValue Returns the numeric value of this object
func (i Int) IntValue() int {
	if i.IsNil() {
		return 0
	}
	return *i.value
}

// IsNil Check if the object is empty
func (i Int) IsNil() bool {
	return i.value == nil
}

// MarshalJSON implements the encoding json interface.
func (i Int) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.value)
}

// UnmarshalJSON implements the encoding json interface.
func (i *Int) UnmarshalJSON(data []byte) error {
	sval := strings.Trim(string(data), "\"")
	if !v.RegexpNumeric.MatchString(sval) {
		i.value = nil
		return nil
	}
	var value interface{}
	if err := json.Unmarshal([]byte(sval), &value); err != nil {
		return err
	}
	i.value = NewInt(value).value
	return nil
}

// Value implements the driver Valuer interface.
func (i Int) Value() (driver.Value, error) {
	if i.IsNil() {
		return nil, nil
	}
	return *i.value, nil
}

// Scan implements the driver Scanner interface.
func (i *Int) Scan(value interface{}) error {
	i.value = NewInt(value).value
	return nil
}
