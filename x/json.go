// Copyright © 2021 - 2023 vity <vityme@icloud.com>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package x

import (
	"encoding/json"
	"errors"
	"strconv"
	"strings"
)

type jsonType int8

const (
	jsonNumber  jsonType = iota // JSON 数值类型，整形、浮点型
	jsonString                  // JSON 字符串类型
	jsonBoolean                 // JSON 布尔类型
	jsonArray                   // JSON 数组类型
	jsonObject                  // JSON 对象类型
	jsonNull                    // JSON 空对象
)

// JsonFromStringE 解析JSON字符串数据，并返回解析错误
func JsonFromStringE(jsonStr string) (*JsonNode, error) {
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

// JsonFromString 解析JSON字符串，非法字符以{}或[]返回
func JsonFromString(val string, array bool) *JsonNode {
	if val == "" {
		var node *JsonNode
		if array {
			node, _ = JsonFromStringE("[]")
		} else {
			node, _ = JsonFromStringE("{}")
		}
		return node
	} else {
		if node, err := JsonFromStringE(val); err != nil {
			if array {
				node, _ = JsonFromStringE("[]")
			} else {
				node, _ = JsonFromStringE("{}")
			}
			return node
		} else {
			return node
		}
	}
}

// JsonToString 转为JSON字符串
func JsonToString(value any) string {
	jsonVal, err := json.Marshal(value)
	if err != nil {
		return "{}"
	}
	return string(jsonVal)
}

// JsonNode JSON 结构
type JsonNode struct {
	rawString *string              // JSON 原始字符串，解析的时候存储
	object    map[string]*JsonNode // JSON 对象类型
	array     []*JsonNode          // JSON 数组类型
	value     any                  // JSON 字段值
	jType     jsonType             // JSON 字段值类型
}

// ToMap 转换为map对象  map[string]any
func (n *JsonNode) ToMap() map[string]any {
	if n == nil || n.rawString == nil || *n.rawString == "" {
		return MapEmpty[any]()
	}
	var m map[string]any
	_ = json.Unmarshal([]byte(*n.rawString), &m)
	return m
}

// ToSlice 转换为slice []map[string]any
func (n *JsonNode) ToSlice() []map[string]any {
	if n == nil || n.rawString == nil || *n.rawString == "" {
		return SliceEmpty[map[string]any]()
	}
	var s []map[string]any
	_ = json.Unmarshal([]byte(*n.rawString), &s)
	return s
}

// ToArray 转换为array []any
func (n *JsonNode) ToArray() []any {
	if n == nil || n.rawString == nil || *n.rawString == "" {
		return ArrayEmpty[any]()
	}
	var a []any
	_ = json.Unmarshal([]byte(*n.rawString), &a)
	return a
}

// IsEmpty 是否空JSON
func (n *JsonNode) IsEmpty() bool {
	return n == nil || n.rawString == nil || *n.rawString == "" || n.Size() == 0
}

// IsArray 是否JSONArray类型
func (n *JsonNode) IsArray() bool {
	return n != nil && n.rawString != nil && *n.rawString != "" && n.object == nil && n.array != nil
}

// IsObject 是否JSONObject 类型
func (n *JsonNode) IsObject() bool {
	return n != nil && n.rawString != nil && *n.rawString != "" && n.object != nil && n.array == nil
}

// Size 长度，对象类型：字段个数，数组类型：数组长度
func (n *JsonNode) Size() int {
	if n == nil || n.rawString == nil || *n.rawString == "" {
		return 0
	}
	if n.IsArray() {
		return len(n.array)
	} else if n.IsObject() {
		return len(n.object)
	}
	return 0
}

// Keys 获取JSON对象所有字段名称
func (n *JsonNode) Keys() []string {
	if n.IsObject() && !n.IsEmpty() {
		i := 0
		keys := make([]string, len(n.object))
		for k := range n.object {
			keys[i] = k
			i++
		}
		return keys
	}
	return []string{}
}

// ContainsKey 是否包含指定Key
func (n *JsonNode) ContainsKey(key string) bool {
	return SliceContains[string](n.Keys(), key)
}

// Name 获取JSON对象指定字段的值
func (n *JsonNode) Name(key string) *JsonNode {
	if n.IsObject() && !n.IsEmpty() {
		return n.object[key]
	}
	return nil
}

// Index 获取JSON数组指定索引的值
func (n *JsonNode) Index(index int) *JsonNode {
	if n.IsArray() && !n.IsEmpty() && index < n.Size() {
		return n.array[index]
	}
	return nil
}

// String 获取JSON字段指定的值内容，string 类型
func (n *JsonNode) String() string {
	if n == nil || n.rawString == nil {
		return ""
	}
	return *n.rawString
}

// Int64 获取JSON字段指定的值内容，int64类型
func (n *JsonNode) Int64() int64 {
	if n == nil || n.rawString == nil {
		return 0
	}
	switch n.jType {
	case jsonString:
		if val, err := strconv.ParseFloat(n.value.(string), 64); err == nil {
			return int64(val)
		} else {
			panic(errors.New("invalid number value"))
		}
	case jsonNumber:
		return int64(n.value.(float64))
	case jsonBoolean:
		if n.value.(bool) {
			return 1
		} else {
			return 0
		}
	}
	panic(errors.New("invalid number value"))
}

// Float64 获取JSON字段指定的值内容，float64类型
func (n *JsonNode) Float64() float64 {
	if n == nil || n.rawString == nil {
		return 0
	}
	switch n.jType {
	case jsonString:
		if val, err := strconv.ParseFloat(n.value.(string), 64); err == nil {
			return val
		} else {
			panic(errors.New("invalid number value"))
		}
	case jsonNumber:
		return n.value.(float64)
	case jsonBoolean:
		if n.value.(bool) {
			return 1
		} else {
			return 0
		}
	}
	panic(errors.New("invalid number value"))

}

// Boolean 获取JSON字段指定的值内容，bool 类型
func (n *JsonNode) Boolean() bool {
	if n == nil || n.rawString == nil {
		return false
	}
	switch n.jType {
	case jsonString:
		if strings.EqualFold(strings.ToLower(n.value.(string)), "true") {
			return true
		} else if strings.EqualFold(strings.ToLower(n.value.(string)), "false") {
			return false
		}
		panic(errors.New("invalid boolean value"))
	case jsonNumber:
		return int64(n.value.(float64)) > 0
	case jsonBoolean:
		return n.value.(bool)
	}
	panic(errors.New("invalid boolean value"))
}

// 从map解析为JSON对象
func (n *JsonNode) parseMap(val map[string]any) (*JsonNode, error) {
	node := &JsonNode{}
	node.object = make(map[string]*JsonNode)
	for k, v := range val {
		if n, err := n.convert(v); err != nil {
			return nil, err
		} else {
			node.object[k] = n
		}
	}
	return node, nil
}

// 从列表解析为JSON对象
func (n *JsonNode) parseSlice(val []any) (*JsonNode, error) {
	node := &JsonNode{}
	node.array = make([]*JsonNode, len(val), len(val))
	for i := 0; i < len(val); i++ {
		if n, err := n.convert(val[i]); err != nil {
			return nil, err
		} else {
			node.array[i] = n
		}

	}
	return node, nil
}

// 转换为JSON Value
func (n *JsonNode) convert(val any) (*JsonNode, error) {
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
			if now, err := n.parseSlice(val.([]any)); err != nil {
				return nil, err
			} else {
				raw := JsonToString(val)
				node.jType = jsonArray
				node.array = now.array
				node.rawString = &raw
			}
		case map[string]any:
			if now, err := n.parseMap(val.(map[string]any)); err != nil {
				return nil, err
			} else {
				raw := JsonToString(val)
				node.jType = jsonObject
				node.object = now.object
				node.rawString = &raw
			}
		}
	}
	return node, nil
}
