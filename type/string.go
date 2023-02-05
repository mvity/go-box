// Copyright © 2021 - 2023 vity <vityme@icloud.com>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package t

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"strings"
)

// String A nullable safe type, based on Golang basic type, supports JSON escape, Database field definition
type String struct {
	value *string
}

// NewString Generates a new object of type t.String based on the specified value
// Warning: When value is passed in as uint(X) and float(X), be aware of data overflow and loss of precision
func NewString(value any) String {
	strv := String{}
	if value == nil {
		return strv
	}
	switch value.(type) {
	case bool:
		val := fmt.Sprintf("%v", value.(bool))
		strv.value = &val
	case string:
		sval := strings.TrimSpace(value.(string))
		strv.value = &sval
	case int:
		val := fmt.Sprintf("%v", value.(int))
		strv.value = &val
	case int8:
		val := fmt.Sprintf("%v", value.(int8))
		strv.value = &val
	case int16:
		val := fmt.Sprintf("%v", value.(int16))
		strv.value = &val
	case int32:
		val := fmt.Sprintf("%v", value.(int32))
		strv.value = &val
	case int64:
		val := fmt.Sprintf("%v", value.(int64))
		strv.value = &val
	case uint:
		val := fmt.Sprintf("%v", value.(uint))
		strv.value = &val
	case uint8:
		val := fmt.Sprintf("%v", value.(uint8))
		strv.value = &val
	case uint16:
		val := fmt.Sprintf("%v", value.(uint16))
		strv.value = &val
	case uint32:
		val := fmt.Sprintf("%v", value.(uint32))
		strv.value = &val
	case uint64:
		val := fmt.Sprintf("%v", value.(uint64))
		strv.value = &val
	case float32:
		val := fmt.Sprintf("%v", value.(float32))
		strv.value = &val
	case float64:
		val := fmt.Sprintf("%v", value.(float64))
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
	var value any
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	s.value = NewString(value).value
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
func (s *String) Scan(value any) error {
	s.value = NewString(value).value
	return nil
}
