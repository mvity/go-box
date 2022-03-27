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

// Int16 A nullable safe type, based on Golang basic type, supports JSON escape, Database field definition
type Int16 struct {
	value *int16
}

// NewInt16 Generates a new object of type t.Int16 based on the specified value
// Warning: When value is passed in as uint(X) and float(X), be aware of data overflow and loss of precision
func NewInt16(value interface{}) Int16 {
	intv := Int16{}
	if value == nil {
		return intv
	}
	switch value.(type) {
	case bool:
		var val int16 = 0
		if value.(bool) {
			val = 1
		}
		intv.value = &val
	case string:
		sval := strings.TrimSpace(value.(string))
		if v.RegexpNumeric.MatchString(sval) {
			if strings.Contains(sval, ".") {
				fval, _ := strconv.ParseFloat(value.(string), 64)
				if fval > math.MaxInt16 {
					gbox.WARN("Data overflow during type conversion. value: %v", fval)
				}
				val := int16(fval)
				intv.value = &val
			} else {
				ival, _ := strconv.ParseInt(value.(string), 10, 64)
				if ival > math.MaxInt16 {
					gbox.WARN("Data overflow during type conversion. value: %v", ival)
				}
				val := int16(ival)
				intv.value = &val
			}
		}
	case int:
		ival := value.(int)
		if ival > math.MaxInt16 {
			gbox.WARN("Data overflow during type conversion. value: %v", ival)
		}
		val := int16(ival)
		intv.value = &val
	case int8:
		ival := value.(int8)
		val := int16(ival)
		intv.value = &val
	case int16:
		val := value.(int16)
		intv.value = &val
	case int32:
		ival := value.(int32)
		if ival > math.MaxInt16 {
			gbox.WARN("Data overflow during type conversion. value: %v", ival)
		}
		val := int16(ival)
		intv.value = &val
	case int64:
		ival := value.(int64)
		if ival > math.MaxInt16 {
			gbox.WARN("Data overflow during type conversion. value: %v", ival)
		}
		val := int16(ival)
		intv.value = &val
	case uint:
		uval := value.(uint)
		if uval > math.MaxInt16 {
			gbox.WARN("Data overflow during type conversion. value: %v", uval)
		}
		val := int16(uval)
		intv.value = &val
	case uint8:
		uval := value.(uint8)
		val := int16(uval)
		intv.value = &val
	case uint16:
		uval := value.(uint16)
		if uval > math.MaxInt16 {
			gbox.WARN("Data overflow during type conversion. value: %v", uval)
		}
		val := int16(uval)
		intv.value = &val
	case uint32:
		uval := value.(uint32)
		if uval > math.MaxInt16 {
			gbox.WARN("Data overflow during type conversion. value: %v", uval)
		}
		val := int16(uval)
		intv.value = &val
	case uint64:
		uval := value.(uint64)
		if uval > math.MaxInt16 {
			gbox.WARN("Data overflow during type conversion. value: %v", uval)
		}
		val := int16(uval)
		intv.value = &val
	case float32:
		fval := value.(float32)
		if fval > math.MaxInt16 {
			gbox.WARN("Data overflow during type conversion. value: %v", fval)
		}
		val := int16(fval)
		intv.value = &val
	case float64:
		fval := value.(float64)
		if fval > math.MaxInt16 {
			gbox.WARN("Data overflow during type conversion. value: %v", fval)
		}
		val := int16(fval)
		intv.value = &val
	}

	return intv
}

// String Returns the string value of this object, implements the Stringer interface.
func (i Int16) String() string {
	if i.IsNil() {
		return "NaN"
	}
	return fmt.Sprintf("%v", *i.value)
}

// Int16Value Returns the numeric value of this object
func (i Int16) Int16Value() int16 {
	if i.IsNil() {
		return 0
	}
	return *i.value
}

// IsNil Check if the object is empty
func (i Int16) IsNil() bool {
	return i.value == nil
}

// MarshalJSON implements the encoding json interface.
func (i Int16) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.value)
}

// UnmarshalJSON implements the encoding json interface.
func (i *Int16) UnmarshalJSON(data []byte) error {
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
	i.value = NewInt16(value).value
	return nil
}

// Value implements the driver Valuer interface.
func (i Int16) Value() (driver.Value, error) {
	if i.IsNil() {
		return nil, nil
	}
	return *i.value, nil
}

// Scan implements the driver Scanner interface.
func (i *Int16) Scan(value interface{}) error {
	//if value == nil {
	//	i.value = nil
	//	return nil
	//}
	//if val, ok := value.(int16); !ok {
	//	return errors.New(fmt.Sprint("Failed to unmarshal int16 value:", value))
	//} else {
	//	i.value = &val
	//}
	i.value = NewInt16(value).value
	return nil
}
