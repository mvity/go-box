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
	v "github.com/mvity/go-box/validator"
	"strings"
)

// String A nullable safe type, based on Golang basic type, supports JSON escape, Database field definition
type String struct {
	value *string
}

// NewString Generates a new object of type t.String based on the specified value
// Warning: When value is passed in as uint(X) and float(X), be aware of data overflow and loss of precision
func NewString(value interface{}) String {
	strv := String{}
	if value == nil {
		return strv
	}
	switch value.(type) {
	case bool:
		var bval = value.(bool)
		val := fmt.Sprintf("%v", bval)
		strv.value = &val
	case string:
		sval := strings.TrimSpace(value.(string))
		strv.value = &sval
	case int:
		ival := value.(int)
		val := fmt.Sprintf("%v", ival)
		strv.value = &val
	case int8:
		ival := value.(int8)
		val := fmt.Sprintf("%v", ival)
		strv.value = &val
	case int16:
		ival := value.(int16)
		val := fmt.Sprintf("%v", ival)
		strv.value = &val
	case int32:
		ival := value.(int32)
		val := fmt.Sprintf("%v", ival)
		strv.value = &val
	case int64:
		ival := value.(int64)
		val := fmt.Sprintf("%v", ival)
		strv.value = &val
	case uint:
		uval := value.(uint)
		val := fmt.Sprintf("%v", uval)
		strv.value = &val
	case uint8:
		uval := value.(uint8)
		val := fmt.Sprintf("%v", uval)
		strv.value = &val
	case uint16:
		uval := value.(uint16)
		val := fmt.Sprintf("%v", uval)
		strv.value = &val
	case uint32:
		uval := value.(uint32)
		val := fmt.Sprintf("%v", uval)
		strv.value = &val
	case uint64:
		uval := value.(uint64)
		val := fmt.Sprintf("%v", uval)
		strv.value = &val
	case float32:
		fval := value.(float32)
		val := fmt.Sprintf("%v", fval)
		strv.value = &val
	case float64:
		fval := value.(float64)
		val := fmt.Sprintf("%v", fval)
		strv.value = &val
	}

	return strv
}

// String Returns the string value of this object, implements the Stringer interface.
func (s String) String() string {
	if s.IsNil() {
		return "NaN"
	}
	return fmt.Sprintf("%v", *s.value)
}

// StringValue Returns the numeric value of this object
func (s String) StringValue() string {
	if s.IsNil() {
		return ""
	}
	return *s.value
}

// IsNil Check if the object is empty
func (s String) IsNil() bool {
	return s.value == nil
}

// MarshalJSON implements the encoding json interface.
func (s String) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.value)
}

// UnmarshalJSON implements the encoding json interface.
func (s *String) UnmarshalJSON(data []byte) error {
	sval := strings.Trim(string(data), "\"")
	if !v.RegexpNumeric.MatchString(sval) {
		s.value = nil
		return nil
	}
	var value interface{}
	if err := json.Unmarshal([]byte(sval), &value); err != nil {
		return err
	}
	if value == nil {
		s.value = nil
		return nil
	}
	intv := NewString(value)
	s.value = intv.value
	return nil
}

// Value implements the driver Valuer interface.
func (s String) Value() (driver.Value, error) {
	if s.IsNil() {
		return nil, nil
	}
	return *s.value, nil
}

// Scan implements the driver Scanner interface.
func (s *String) Scan(value interface{}) error {
	if value == nil {
		s.value = nil
		return nil
	}
	if val, ok := value.(string); !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal int value:", value))
	} else {
		s.value = &val
	}
	return nil
}
