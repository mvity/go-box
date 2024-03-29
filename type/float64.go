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

// Float64 A nullable safe type, based on Golang basic type, supports JSON escape, Database field definition
type Float64 struct {
	value *float64
}

// NewFloat64 Generates a new object of type t.Float64 based on the specified value
// Warning: When value is passed in as uint(X) and float(X), be aware of data overflow and loss of precision
func NewFloat64(value any) Float64 {
	intv := Float64{}
	if value == nil {
		return intv
	}
	switch value.(type) {
	case bool:
		var val float64 = 0
		if value.(bool) {
			val = 1
		}
		intv.value = &val
	case string:
		sval := strings.TrimSpace(value.(string))
		if x.RegexpNumeric.MatchString(sval) {
			val, _ := strconv.ParseFloat(value.(string), 64)
			intv.value = &val
		}
	case int:
		ival := value.(int)
		if float64(ival) > math.MaxFloat64 {
			gbox.WARN("Data overflow during type conversion. value: %v", ival)
		}
		val := float64(ival)
		intv.value = &val
	case int8:
		ival := value.(int8)
		val := float64(ival)
		intv.value = &val
	case int16:
		ival := value.(int16)
		val := float64(ival)
		intv.value = &val
	case int32:
		ival := value.(int32)
		val := float64(ival)
		intv.value = &val
	case int64:
		ival := value.(int64)
		if float64(ival) > math.MaxFloat64 {
			gbox.WARN("Data overflow during type conversion. value: %v", ival)
		}
		val := float64(ival)
		intv.value = &val
	case uint:
		uval := value.(uint)
		if float64(uval) > math.MaxFloat64 {
			gbox.WARN("Data overflow during type conversion. value: %v", uval)
		}
		val := float64(uval)
		intv.value = &val
	case uint8:
		uval := value.(uint8)
		val := float64(uval)
		intv.value = &val
	case uint16:
		uval := value.(uint16)
		val := float64(uval)
		intv.value = &val
	case uint32:
		uval := value.(uint32)
		if float64(uval) > math.MaxFloat64 {
			gbox.WARN("Data overflow during type conversion. value: %v", uval)
		}
		val := float64(uval)
		intv.value = &val
	case uint64:
		uval := value.(uint64)
		if float64(uval) > math.MaxFloat64 {
			gbox.WARN("Data overflow during type conversion. value: %v", uval)
		}
		val := float64(uval)
		intv.value = &val
	case float32:
		fval := value.(float32)
		val := float64(fval)
		intv.value = &val
	case float64:
		val := value.(float64)
		intv.value = &val
	}

	return intv
}

// String Returns the string value of this object, implements the Stringer interface.
func (i Float64) String() string {
	if i.IsNil() {
		return "NaN"
	}
	return fmt.Sprintf("%v", *i.value)
}

// Float64Value Returns the numeric value of this object
func (i Float64) Float64Value() float64 {
	if i.IsNil() {
		return 0
	}
	return *i.value
}

// IsNil Check if the object is empty
func (i Float64) IsNil() bool {
	return i.value == nil
}

// MarshalJSON implements the encoding json interface.
func (i Float64) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.value)
}

// UnmarshalJSON implements the encoding json interface.
func (i *Float64) UnmarshalJSON(data []byte) error {
	sval := strings.Trim(string(data), "\"")
	if !x.RegexpNumeric.MatchString(sval) {
		i.value = nil
		return nil
	}
	var value any
	if err := json.Unmarshal([]byte(sval), &value); err != nil {
		return err
	}
	i.value = NewFloat64(value).value
	return nil
}

// Value implements the driver Valuer interface.
func (i Float64) Value() (driver.Value, error) {
	if i.IsNil() {
		return nil, nil
	}
	return *i.value, nil
}

// Scan implements the driver Scanner interface.
func (i *Float64) Scan(value any) error {
	i.value = NewFloat64(value).value
	return nil
}
