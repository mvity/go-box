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
	"github.com/mvity/go-box/x"
	"strconv"
)

// Bool A nullable safe type, based on Golang basic type, supports JSON escape, Database field definition
type Bool struct {
	value *bool
}

var trueStrs = []string{"true", "1", "yes", "ok"}
var falseStrs = []string{"false", "0", "no"}

// NewBool Generates a new object of type t.Bool based on the specified value
// Warning: When value is passed in as uint(X) and float(X), be aware of data overflow and loss of precision
func NewBool(value any) Bool {
	boolv := Bool{}
	if value == nil {
		return boolv
	}
	switch value.(type) {
	case bool:
		var val = value.(bool)
		boolv.value = &val
	case string:
		//sval := strings.ToLower(strings.TrimSpace(value.(string)))
		val, _ := strconv.ParseBool(x.ToString(value))
		boolv.value = &val
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
	var value any
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	b.value = NewBool(value).value
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
func (b *Bool) Scan(value any) error {
	b.value = NewBool(value).value
	return nil
}
