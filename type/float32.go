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
	"github.com/mvity/go-box/x"
	"math"
	"strconv"
	"strings"
)

// Float32 A nullable safe type, based on Golang basic type, supports JSON escape, Database field definition
type Float32 struct {
	value *float32
}

// NewFloat32 Generates a new object of type t.Float32 based on the specified value
// Warning: When value is passed in as uint(X) and float(X), be aware of data overflow and loss of precision
func NewFloat32(value any) Float32 {
	intv := Float32{}
	if value == nil {
		return intv
	}
	switch value.(type) {
	case bool:
		var val float32 = 0
		if value.(bool) {
			val = 1
		}
		intv.value = &val
	case string:
		sval := strings.TrimSpace(value.(string))
		if x.RegexpNumeric.MatchString(sval) {
			fval, _ := strconv.ParseFloat(value.(string), 64)
			if fval > math.MaxFloat32 || float64(fval) < math.SmallestNonzeroFloat32 {
				gbox.WARN("Data overflow during type conversion. value: %v", fval)
			}
			val := float32(fval)
			intv.value = &val
		}
	case int:
		ival := value.(int)
		if float64(ival) > math.MaxFloat32 || float64(ival) < math.SmallestNonzeroFloat32 {
			gbox.WARN("Data overflow during type conversion. value: %v", ival)
		}
		val := float32(ival)
		intv.value = &val
	case int8:
		ival := value.(int8)
		val := float32(ival)
		intv.value = &val
	case int16:
		ival := value.(int16)
		val := float32(ival)
		intv.value = &val
	case int32:
		ival := value.(int32)
		val := float32(ival)
		intv.value = &val
	case int64:
		ival := value.(int64)
		if float64(ival) > math.MaxFloat32 || float64(ival) < math.SmallestNonzeroFloat32 {
			gbox.WARN("Data overflow during type conversion. value: %v", ival)
		}
		val := float32(ival)
		intv.value = &val
	case uint:
		uval := value.(uint)
		if float64(uval) > math.MaxFloat32 || float64(uval) < math.SmallestNonzeroFloat32 {
			gbox.WARN("Data overflow during type conversion. value: %v", uval)
		}
		val := float32(uval)
		intv.value = &val
	case uint8:
		uval := value.(uint8)
		val := float32(uval)
		intv.value = &val
	case uint16:
		uval := value.(uint16)
		val := float32(uval)
		intv.value = &val
	case uint32:
		uval := value.(uint32)
		if float64(uval) > math.MaxFloat32 || float64(uval) < math.SmallestNonzeroFloat32 {
			gbox.WARN("Data overflow during type conversion. value: %v", uval)
		}
		val := float32(uval)
		intv.value = &val
	case uint64:
		uval := value.(uint64)
		if float64(uval) > math.MaxFloat32 || float64(uval) < math.SmallestNonzeroFloat32 {
			gbox.WARN("Data overflow during type conversion. value: %v", uval)
		}
		val := float32(uval)
		intv.value = &val
	case float32:
		val := value.(float32)
		intv.value = &val
	case float64:
		fval := value.(float64)
		if fval > math.MaxFloat32 || fval < math.SmallestNonzeroFloat32 {
			gbox.WARN("Data overflow during type conversion. value: %v", fval)
		}

		val := float32(fval)
		intv.value = &val
	}

	return intv
}

// String Returns the string value of this object, implements the Stringer interface.
func (i Float32) String() string {
	if i.IsNil() {
		return "NaN"
	}
	return fmt.Sprintf("%v", *i.value)
}

// Float32Value Returns the numeric value of this object
func (i Float32) Float32Value() float32 {
	if i.IsNil() {
		return 0
	}
	return *i.value
}

// IsNil Check if the object is empty
func (i Float32) IsNil() bool {
	return i.value == nil
}

// MarshalJSON implements the encoding json interface.
func (i Float32) MarshalJSON() ([]byte, error) {
	if i.IsNil() {
		return json.Marshal(nil)
	}
	sval := fmt.Sprintf("%0.7f", *i.value)
	fval, err := strconv.ParseFloat(sval, 64)
	if err != nil || fval > math.MaxFloat32 || fval < math.SmallestNonzeroFloat32 {
		fval = 0
	}
	return json.Marshal(fval)
}

// UnmarshalJSON implements the encoding json interface.
func (i *Float32) UnmarshalJSON(data []byte) error {
	sval := strings.Trim(string(data), "\"")
	if !x.RegexpNumeric.MatchString(sval) {
		i.value = nil
		return nil
	}
	var value any
	if err := json.Unmarshal([]byte(sval), &value); err != nil {
		return err
	}
	i.value = NewFloat32(value).value
	return nil
}

// Value implements the driver Valuer interface.
func (i Float32) Value() (driver.Value, error) {
	if i.IsNil() {
		return nil, nil
	}
	return *i.value, nil
}

// Scan implements the driver Scanner interface.
func (i *Float32) Scan(value any) error {
	i.value = NewFloat32(value).value
	return nil
}
