// Copyright © 2021 - 2022 vity <vityme@icloud.com>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package t

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	k "github.com/mvity/go-box/kit"
	"strings"
)

// Bool A nullable safe type, based on Golang basic type, supports JSON escape, Database field definition
type Bool struct {
	value *bool
}

var trueStrs = []string{"true", "1", "yes", "ok"}
var falseStrs = []string{"false", "0", "no"}

// NewBool Generates a new object of type t.Bool based on the specified value
// Warning: When value is passed in as uint(X) and float(X), be aware of data overflow and loss of precision
func NewBool(value interface{}) Bool {
	boolv := Bool{}
	if value == nil {
		return boolv
	}
	switch value.(type) {
	case bool:
		var val = value.(bool)
		boolv.value = &val
	case string:
		sval := strings.ToLower(strings.TrimSpace(value.(string)))
		if sval != "" {
			var val bool
			if k.InStringSlice(sval, trueStrs) {
				val = true
			} else if k.InStringSlice(sval, falseStrs) {
				val = false
			} else {
				break
			}
			boolv.value = &val
		}
	case int:
		ival := value.(int)
		val := ival > 0
		boolv.value = &val
	case int8:
		ival := value.(int8)
		val := ival > 0
		boolv.value = &val
	case int16:
		ival := value.(int16)
		val := ival > 0
		boolv.value = &val
	case int32:
		ival := value.(int32)
		val := ival > 0
		boolv.value = &val
	case int64:
		ival := value.(int64)
		val := ival > 0
		boolv.value = &val
	case uint:
		uval := value.(uint)
		val := uval > 0
		boolv.value = &val
	case uint8:
		uval := value.(uint8)
		val := uval > 0
		boolv.value = &val
	case uint16:
		uval := value.(uint16)
		val := uval > 0
		boolv.value = &val
	case uint32:
		uval := value.(uint32)
		val := uval > 0
		boolv.value = &val
	case uint64:
		uval := value.(uint64)
		val := uval > 0
		boolv.value = &val
	case float32:
		fval := value.(float32)
		val := fval > 0
		boolv.value = &val
	case float64:
		fval := value.(float64)
		val := fval > 0
		boolv.value = &val
	}

	return boolv
}

// String Returns the string value of this object, implements the Stringer interface.
func (b Bool) String() string {
	if b.IsNil() {
		return "NaN"
	}
	return fmt.Sprintf("%v", *b.value)
}

// BoolValue Returns the numeric value of this object
func (b Bool) BoolValue() bool {
	if b.IsNil() {
		return false
	}
	return *b.value
}

// IsNil Check if the object is empty
func (b Bool) IsNil() bool {
	return b.value == nil
}

// MarshalJSON implements the encoding json interface.
func (b Bool) MarshalJSON() ([]byte, error) {
	return json.Marshal(b.value)
}

// UnmarshalJSON implements the encoding json interface.
func (b *Bool) UnmarshalJSON(data []byte) error {
	sval := strings.Trim(string(data), "\"")
	var value interface{}
	if err := json.Unmarshal([]byte(sval), &value); err != nil {
		return err
	}
	if value == nil {
		b.value = nil
		return nil
	}
	intv := NewBool(value)
	b.value = intv.value
	return nil
}

// Value implements the driver Valuer interface.
func (b Bool) Value() (driver.Value, error) {
	if b.IsNil() {
		return nil, nil
	}
	return *b.value, nil
}

// Scan implements the driver Scanner interface.
func (b *Bool) Scan(value interface{}) error {
	//if value == nil {
	//	b.value = nil
	//	return nil
	//}
	//if val, ok := value.(bool); !ok {
	//	return errors.New(fmt.Sprint("Failed to unmarshal int value:", value))
	//} else {
	//	b.value = &val
	//}
	intv := NewBool(value)
	b.value = intv.value
	return nil
}
