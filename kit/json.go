// Copyright © 2021 - 2022 vity <vityme@icloud.com>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package k

import (
	"encoding/json"
	"errors"
	"strconv"
	"strings"
)

// ToJSONString 转为JSON字符串
func ToJSONString(value any) string {
	jsonVal, err := json.Marshal(value)
	if err != nil {
		return "{}"
	}
	return string(jsonVal)
}

type jsonType int8

const (
	jsonNumber  jsonType = iota // JSON 数值类型，整形、浮点型
	jsonString                  // JSON 字符串类型
	jsonBoolean                 // JSON 布尔类型
	jsonArray                   // JSON 数组类型
	jsonObject                  // JSON 对象类型
	jsonNull                    // JSON 空对象
)

// JsonNode JSON 结构
type JsonNode struct {
	rawString *string              // JSON 原始字符串，解析的时候存储
	object    map[string]*JsonNode // JSON 对象类型
	array     []*JsonNode          // JSON 数组类型
	value     any                  // JSON 字段值
	jType     jsonType             // JSON 字段值类型
}

// ParseJSONString 解析JSON字符串数据
func ParseJSONString(jsonStr string) (*JsonNode, error) {
	jsonStr = strings.TrimSpace(jsonStr)

	node := new(JsonNode)
	node.rawString = &jsonStr
	if strings.HasPrefix(jsonStr, "{") && strings.HasSuffix(jsonStr, "}") {
		var jval map[string]any
		if err := json.Unmarshal([]byte(jsonStr), &jval); err != nil {
			return nil, err
		}
		if rst, err := node.parseMap(jval); err != nil {
			return nil, err
		} else {
			node.object = rst.object
			return node, nil
		}

	} else if strings.HasPrefix(jsonStr, "[") && strings.HasSuffix(jsonStr, "]") {
		var jval []any
		if err := json.Unmarshal([]byte(jsonStr), &jval); err != nil {
			return nil, err
		}
		if rst, err := node.parseSlice(jval); err != nil {
			return nil, err
		} else {
			node.array = rst.array
			return node, nil
		}
	} else {
		return nil, errors.New("invalid JSON string")
	}
}

// ParseJSONStringForce 强制解析JSON字符串，非法字符以{}或[]返回
func ParseJSONStringForce(val string, array bool) *JsonNode {
	if val == "" {
		var node *JsonNode
		if array {
			node, _ = ParseJSONString("[]")
		} else {
			node, _ = ParseJSONString("{}")
		}
		return node
	} else {
		if node, err := ParseJSONString(val); err != nil {
			if array {
				node, _ = ParseJSONString("[]")
			} else {
				node, _ = ParseJSONString("{}")
			}
			return node
		} else {
			return node
		}
	}
}

// ToMap 转换为Map对象
func (jn *JsonNode) ToMap() map[string]any {
	var m map[string]any
	_ = json.Unmarshal([]byte(*jn.rawString), &m)
	return m
}

// ToSlice 转换为Slice对象
func (jn *JsonNode) ToSlice() []any {
	var s []any
	_ = json.Unmarshal([]byte(*jn.rawString), &s)
	return s
}

// IsEmpty 是否空JSON
func (jn *JsonNode) IsEmpty() bool {
	return jn.Size() == 0
}

// IsArray 是否JSONArray类型
func (jn *JsonNode) IsArray() bool {
	return jn.object == nil && jn.array != nil
}

// IsObject 是否JSONObject 类型
func (jn *JsonNode) IsObject() bool {
	return jn.object != nil && jn.array == nil
}

// Size 长度，对象类型：字段个数，数组类型：数组长度
func (jn *JsonNode) Size() int {
	if jn.IsArray() {
		return len(jn.array)
	} else if jn.IsObject() {
		return len(jn.object)
	}
	return 0
}

// Keys 获取JSON对象所有字段名称
func (jn *JsonNode) Keys() []string {
	if jn.IsObject() && !jn.IsEmpty() {
		i := 0
		keys := make([]string, len(jn.object))
		for k := range jn.object {
			keys[i] = k
			i++
		}
		return keys
	}
	return []string{}
}

// Name 获取JSON对象指定字段的值
func (jn *JsonNode) Name(key string) *JsonNode {
	if jn.IsObject() && !jn.IsEmpty() {
		return jn.object[key]
	}
	return nil
}

// Index 获取JSON数组指定索引的值
func (jn *JsonNode) Index(index int) *JsonNode {
	if jn.IsArray() && !jn.IsEmpty() && index < jn.Size() {
		return jn.array[index]
	}
	return nil
}

// String 获取JSON字段指定的值内容，string 类型
func (jn *JsonNode) String() string {
	if jn == nil && jn.rawString == nil {
		panic(errors.New("invalid string value"))
	}
	return *jn.rawString
}

// Int64 获取JSON字段指定的值内容，int64类型
func (jn *JsonNode) Int64() int64 {
	if jn == nil && jn.value == nil {
		panic(errors.New("invalid number value"))
	}
	switch jn.jType {
	case jsonString:
		if val, err := strconv.ParseFloat(jn.value.(string), 64); err == nil {
			return int64(val)
		} else {
			panic(errors.New("invalid number value"))
		}
	case jsonNumber:
		return int64(jn.value.(float64))
	case jsonBoolean:
		if jn.value.(bool) {
			return 1
		} else {
			return 0
		}
	}
	panic(errors.New("invalid number value"))
}

// Float64 获取JSON字段指定的值内容，float64类型
func (jn *JsonNode) Float64() float64 {
	if jn == nil && jn.value == nil {
		panic(errors.New("invalid number value"))
	}
	switch jn.jType {
	case jsonString:
		if val, err := strconv.ParseFloat(jn.value.(string), 64); err == nil {
			return val
		} else {
			panic(errors.New("invalid number value"))
		}
	case jsonNumber:
		return jn.value.(float64)
	case jsonBoolean:
		if jn.value.(bool) {
			return 1
		} else {
			return 0
		}
	}
	panic(errors.New("invalid number value"))

}

// Boolean 获取JSON字段指定的值内容，bool 类型
func (jn *JsonNode) Boolean() bool {
	if jn == nil && jn.value == nil {
		panic(errors.New("invalid boolean value"))
	}
	switch jn.jType {
	case jsonString:
		if strings.EqualFold(strings.ToLower(jn.value.(string)), "true") {
			return true
		} else if strings.EqualFold(strings.ToLower(jn.value.(string)), "false") {
			return false
		}
		panic(errors.New("invalid boolean value"))
	case jsonNumber:
		return int64(jn.value.(float64)) > 0
	case jsonBoolean:
		return jn.value.(bool)
	}
	panic(errors.New("invalid boolean value"))
}

// 从map解析为JSON对象
func (jn *JsonNode) parseMap(val map[string]any) (*JsonNode, error) {
	node := &JsonNode{}
	node.object = make(map[string]*JsonNode)
	for k, v := range val {
		if n, err := jn.convert(v); err != nil {
			return nil, err
		} else {
			node.object[k] = n
		}
	}
	return node, nil
}

// 从列表解析为JSON对象
func (jn *JsonNode) parseSlice(val []any) (*JsonNode, error) {
	node := &JsonNode{}
	node.array = make([]*JsonNode, len(val), len(val))
	for i := 0; i < len(val); i++ {
		if n, err := jn.convert(val[i]); err != nil {
			return nil, err
		} else {
			node.array[i] = n
		}

	}
	return node, nil
}

// 转换为JSON Value
func (jn *JsonNode) convert(val any) (*JsonNode, error) {
	//fmt.Printf("Convert -> %value , %T \n", val, val)
	node := &JsonNode{}
	if val == nil {
		node.jType = jsonNull
		node.value = nil
		node.rawString = nil
	} else {
		switch val.(type) {
		case string:
			raw := val.(string)
			node.jType = jsonString
			node.value = val
			node.rawString = &raw
		case float64:
			raw := strconv.FormatFloat(val.(float64), 'f', -1, 64)
			node.jType = jsonNumber
			node.value = val
			node.rawString = &raw
		case bool:
			raw := strconv.FormatBool(val.(bool))
			node.jType = jsonBoolean
			node.value = val
			node.rawString = &raw
		case []any:
			if now, err := jn.parseSlice(val.([]any)); err != nil {
				return nil, err
			} else {
				raw := ToJSONString(val)
				node.jType = jsonArray
				node.array = now.array
				node.rawString = &raw
			}
		case map[string]any:
			if now, err := jn.parseMap(val.(map[string]any)); err != nil {
				return nil, err
			} else {
				raw := ToJSONString(val)
				node.jType = jsonObject
				node.object = now.object
				node.rawString = &raw
			}
		}
	}
	return node, nil
}
