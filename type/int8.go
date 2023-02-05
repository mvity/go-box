// Copyright © 2021 - 2023 vity <vityme@icloud.com>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

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

// Int8 A nullable safe type, based on Golang basic type, supports JSON escape, Database field definition
type Int8 struct {
	value *int8
}

// NewInt8 Generates a new object of type t.Int8 based on the specified value
// Warning: When value is passed in as uint(X) and float(X), be aware of data overflow and loss of precision
func NewInt8(value any) Int8 {
	intv := Int8{}
	if value == nil {
		return intv
	}
	switch value.(type) {
	case bool:
		var val int8 = 0
		if value.(bool) {
			val = 1
		}
		intv.value = &val
	case string:
		sval := strings.TrimSpace(value.(string))
		if x.RegexpNumeric.MatchString(sval) {
			if strings.Contains(sval, ".") {
				fval, _ := strconv.ParseFloat(value.(string), 64)
				if fval > math.MaxInt8 {
					gbox.WARN("Data overflow during type conversion. value: %v", fval)
				}
				val := int8(fval)
				intv.value = &val
			} else {
				ival, _ := strconv.ParseInt(value.(string), 10, 64)
				if ival > math.MaxInt8 {
					gbox.WARN("Data overflow during type conversion. value: %v", ival)
				}
				val := int8(ival)
				intv.value = &val
			}
		}
	case int:
		ival := value.(int)
		if ival > math.MaxInt8 {
			gbox.WARN("Data overflow during type conversion. value: %v", ival)
		}
		val := int8(ival)
		intv.value = &val
	case int8:
		val := value.(int8)
		intv.value = &val
	case int16:
		ival := value.(int16)
		if ival > math.MaxInt8 {
			gbox.WARN("Data overflow during type conversion. value: %v", ival)
		}
		val := int8(ival)
		intv.value = &val
	case int32:
		ival := value.(int32)
		if ival > math.MaxInt8 {
			gbox.WARN("Data overflow during type conversion. value: %v", ival)
		}
		val := int8(ival)
		intv.value = &val
	case int64:
		ival := value.(int64)
		if ival > math.MaxInt8 {
			gbox.WARN("Data overflow during type conversion. value: %v", ival)
		}
		val := int8(ival)
		intv.value = &val
	case uint:
		uval := value.(uint)
		if uval > math.MaxInt8 {
			gbox.WARN("Data overflow during type conversion. value: %v", uval)
		}
		val := int8(uval)
		intv.value = &val
	case uint8:
		uval := value.(uint8)
		if uval > math.MaxInt8 {
			gbox.WARN("Data overflow during type conversion. value: %v", uval)
		}
		val := int8(uval)
		intv.value = &val
	case uint16:
		uval := value.(uint16)
		if uval > math.MaxInt8 {
			gbox.WARN("Data overflow during type conversion. value: %v", uval)
		}
		val := int8(uval)
		intv.value = &val
	case uint32:
		uval := value.(uint32)
		if uval > math.MaxInt8 {
			gbox.WARN("Data overflow during type conversion. value: %v", uval)
		}
		val := int8(uval)
		intv.value = &val
	case uint64:
		uval := value.(uint64)
		if uval > math.MaxInt8 {
			gbox.WARN("Data overflow during type conversion. value: %v", uval)
		}
		val := int8(uval)
		intv.value = &val
	case float32:
		fval := value.(float32)
		if fval > math.MaxInt8 {
			gbox.WARN("Data overflow during type conversion. value: %v", fval)
		}
		val := int8(fval)
		intv.value = &val
	case float64:
		fval := value.(float64)
		if fval > math.MaxInt8 {
			gbox.WARN("Data overflow during type conversion. value: %v", fval)
		}
		val := int8(fval)
		intv.value = &val
	}

	return intv
}

// String Returns the string value of this object, implements the Stringer interface.
func (i Int8) String() string {
	if i.IsNil() {
		return "NaN"
	}
	return fmt.Sprintf("%v", *i.value)
}

// Int8Value Returns the numeric value of this object
func (i Int8) Int8Value() int8 {
	if i.IsNil() {
		return 0
	}
	return *i.value
}

// IsNil Check if the object is empty
func (i Int8) IsNil() bool {
	return i.value == nil
}

// MarshalJSON implements the encoding json interface.
func (i Int8) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.value)
}

// UnmarshalJSON implements the encoding json interface.
func (i *Int8) UnmarshalJSON(data []byte) error {
	sval := strings.Trim(string(data), "\"")
	if !x.RegexpNumeric.MatchString(sval) {
		i.value = nil
		return nil
	}
	var value any
	if err := json.Unmarshal([]byte(sval), &value); err != nil {
		return err
	}
	i.value = NewInt8(value).value
	return nil
}

// Value implements the driver Valuer interface.
func (i Int8) Value() (driver.Value, error) {
	if i.IsNil() {
		return nil, nil
	}
	return *i.value, nil
}

// Scan implements the driver Scanner interface.
func (i *Int8) Scan(value any) error {
	i.value = NewInt8(value).value
	return nil
}
