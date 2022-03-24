// Copyright © 2021 - 2022 vity <vityme@icloud.com>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package t

import (
	"encoding/json"
	"math"
	"testing"
)

func TestNewUInt(t *testing.T) {
	var intValue1 = math.MaxInt - 1
	var intValue2 = math.MinInt + 1
	var int8Value1 int8 = math.MaxInt8 - 1
	var int8Value2 int8 = math.MinInt8 + 1
	var int16Value1 int16 = math.MaxInt16 - 1
	var int16Value2 int16 = math.MinInt16 + 1
	var int32Value1 int32 = math.MaxInt32 - 1
	var int32Value2 int32 = math.MinInt32 + 1
	var int64Value1 int64 = math.MaxInt64 - 1
	var int64Value2 int64 = math.MinInt64 + 1
	var uintValue uint = math.MaxUint - 1
	var uint8Value uint8 = math.MaxUint8 - 1
	var uint16Value uint16 = math.MaxUint16 - 1
	var uint32Value uint32 = math.MaxUint32 - 1
	var uint64Value uint64 = math.MaxUint64 - 1
	var float32Value1 float32 = math.MaxFloat32 - 1
	var float32Value2 float32 = math.SmallestNonzeroFloat32 + 1
	var float64Value1 = math.MaxFloat64 - 1
	var float64Value2 = math.SmallestNonzeroFloat64 + 1
	values := []interface{}{
		intValue1, intValue2, int8Value1, int8Value2, int16Value1, int16Value2,
		int32Value1, int32Value2, int64Value1, int64Value2,
		uintValue, uint8Value, uint16Value, uint32Value, uint64Value,
		float32Value1, float32Value2, float64Value1, float64Value2, true, false,
		"abcd", '2', "234", "-345", "456.789", "0.0.0.0", nil, func() {}, []string{}, struct{}{},
	}
	for _, value := range values {
		intv := NewUInt(value)
		t.Logf("valid, type: %T, ok? %v, value: %v, intv: %v ", value, !intv.IsNil(), value, intv.String())

	}
}

type testUIntObj struct {
	Value UInt `json:"value"`
}

func TestMarshalJSONUInt(t *testing.T) {
	var intValue1 = math.MaxInt - 1
	var intValue2 = math.MinInt + 1
	var int8Value1 int8 = math.MaxInt8 - 1
	var int8Value2 int8 = math.MinInt8 + 1
	var int16Value1 int16 = math.MaxInt16 - 1
	var int16Value2 int16 = math.MinInt16 + 1
	var int32Value1 int32 = math.MaxInt32 - 1
	var int32Value2 int32 = math.MinInt32 + 1
	var int64Value1 int64 = math.MaxInt64 - 1
	var int64Value2 int64 = math.MinInt64 + 1
	var uintValue uint = math.MaxUint - 1
	var uint8Value uint8 = math.MaxUint8 - 1
	var uint16Value uint16 = math.MaxUint16 - 1
	var uint32Value uint32 = math.MaxUint32 - 1
	var uint64Value uint64 = math.MaxUint64 - 1
	var float32Value1 float32 = math.MaxFloat32 - 1
	var float32Value2 float32 = math.SmallestNonzeroFloat32 + 1
	var float64Value1 = math.MaxFloat64 - 1
	var float64Value2 = math.SmallestNonzeroFloat64 + 1
	values := []interface{}{
		intValue1, intValue2, int8Value1, int8Value2, int16Value1, int16Value2,
		int32Value1, int32Value2, int64Value1, int64Value2,
		uintValue, uint8Value, uint16Value, uint32Value, uint64Value,
		float32Value1, float32Value2, float64Value1, float64Value2, true, false,
		"abcd", '2', "234", "-345", "456.789", "0.0.0.0", nil, func() {}, []string{}, struct{}{},
	}
	var objs []testUIntObj
	for _, value := range values {
		obj := testUIntObj{
			NewUInt(value),
		}
		objs = append(objs, obj)
	}
	for _, obj := range objs {
		if data, err := json.Marshal(obj); err != nil {
			t.Fatalf("error. value: %v, data: %s, err: %v", obj.Value, string(data), err)
		} else {
			t.Logf("valid. value: %v, json: %s, err: %v", obj.Value, string(data), err)
		}
	}
}

func TestUnmarshalJSONUInt(t *testing.T) {

	values := []string{
		"{\"value\":9223372036854775806}",
		"{\"value\":-9223372036854775807}",
		"{\"value\":126}",
		"{\"value\":-127}",
		"{\"value\":32766}",
		"{\"value\":-32767}",
		"{\"value\":2147483646}",
		"{\"value\":-2147483647}",
		"{\"value\":9223372036854775806}",
		"{\"value\":-9223372036854775807}",
		"{\"value\":18446744073709551614}",
		"{\"value\":254}",
		"{\"value\":65534}",
		"{\"value\":4294967294}",
		"{\"value\":18446744073709551614}",
		"{\"value\":3.4028235}",
		"{\"value\":-3.4028235}",
		"{\"value\":1.7976931348623157}",
		"{\"value\":-1.7976931348623157}",
		"{\"value\":1}",
		"{\"value\":0}",
		"{\"value\":\"abcd\"}",
		"{\"value\":\"234\"}",
		"{\"value\":\"-345\"}",
		"{\"value\":\"456.789\"}",
		"{\"value\":\"0.0.0.0\"}",
		"{\"value\":null}",
		"{\"value\":true}",
		"{\"value\":false}",
		"{\"value\":{}}",
		"{\"value\":[]}",
	}
	for _, value := range values {
		var result testUIntObj
		if err := json.Unmarshal([]byte(value), &result); err != nil {
			t.Fatalf("error. value: %v , err: %v ", value, err)
		} else {
			t.Logf("valid, value: %v > nil:%v, %v,%d", value, result.Value.IsNil(), result.Value.String(),
				result.Value.UIntValue())
		}
	}
}
